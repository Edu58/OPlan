"use client";

import { useState, useEffect } from "react";

export default function CreateEvent() {
  const [currentStep, setCurrentStep] = useState(1);
  const [selectedPricing, setSelectedPricing] = useState(null);
  const [eventType, setEventType] = useState("");
  const [registrationFee, setRegistrationFee] = useState("free");
  const [formData, setFormData] = useState({
    eventName: "",
    eventCategory: "",
    eventDescription: "",
    startDateTime: "",
    endDateTime: "",
    venueName: "",
    venueAddress: "",
    venueCity: "",
    venueCountry: "",
    virtualPlatform: "",
    visibility: "public",
    registrationPrice: "",
    eventRules: "",
    additionalInfo: "",
  });

  const totalSteps = 4;
  const stepTexts = [
    "Event Details",
    "Location & Capacity",
    "Rules & Settings",
    "Review & Payment",
  ];

  const handleInputChange = (e) => {
    const { id, name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [id || name]: value,
    }));
  };

  const selectPricing = (type) => {
    setSelectedPricing(type);
  };

  const togglePriceInput = (value) => {
    setRegistrationFee(value);
  };

  const validateStep = (step) => {
    const requiredFields = {
      1: [
        "eventName",
        "eventCategory",
        "eventType",
        "eventDescription",
        "startDateTime",
        "endDateTime",
      ],
      2: [],
      3: [],
      4: [],
    };

    if (step === 2) {
      if (!selectedPricing) {
        alert("Please select a pricing tier for your event.");
        return false;
      }

      if (eventType === "in-person" || eventType === "hybrid") {
        const venueFields = ["venueName", "venueAddress"];
        for (let field of venueFields) {
          if (!formData[field]?.trim()) {
            alert("Please fill in all venue information.");
            return false;
          }
        }
      }
      return true;
    }

    const fields = requiredFields[step] || [];
    for (let field of fields) {
      if (!formData[field]?.trim()) {
        alert("Please fill in all required fields.");
        return false;
      }
    }

    if (step === 1) {
      const startDate = new Date(formData.startDateTime);
      const endDate = new Date(formData.endDateTime);

      if (startDate >= endDate) {
        alert("End date must be after start date.");
        return false;
      }

      if (startDate < new Date()) {
        alert("Event start date cannot be in the past.");
        return false;
      }
    }

    return true;
  };

  const handleNext = () => {
    if (validateStep(currentStep)) {
      if (currentStep < totalSteps) {
        setCurrentStep(currentStep + 1);
      }
    }
  };

  const handlePrev = () => {
    if (currentStep > 1) {
      setCurrentStep(currentStep - 1);
    }
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    if (!selectedPricing) {
      alert("Please select a pricing tier.");
      return;
    }

    const fileInput = document.querySelector('input[type="file"]');
    if (!fileInput?.files.length) {
      alert("Please upload payment proof to complete your event creation.");
      return;
    }

    alert(
      "Event creation submitted successfully! You will receive a confirmation email once payment is verified.",
    );
    console.log("Event submitted with pricing:", selectedPricing);
  };

  const getPaymentAmount = () => {
    if (!selectedPricing) return 0;
    return selectedPricing === "small" ? 2000 : 5000;
  };

  const showLocationFields =
    eventType === "in-person" || eventType === "hybrid";
  const showVirtualFields = eventType === "virtual" || eventType === "hybrid";

  return (
    <>
      {/* Create Event Header */}
      <section className="gradient-bg pt-24 pb-12 text-white">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
          <div className="animate-fade-in">
            <h1 className="text-4xl md:text-5xl font-bold mb-4">
              Create Your
              <span className="block text-yellow-300">Amazing Event</span>
            </h1>
            <p className="text-xl opacity-90 mb-8">
              Bring people together with professional event management tools
            </p>

            {/* Progress Indicator */}
            <div className="flex items-center justify-center space-x-4 mb-8">
              {[1, 2, 3, 4].map((step, index) => (
                <div key={step} className="flex items-center">
                  <div
                    className={`w-10 h-10 rounded-full flex items-center justify-center font-bold text-sm transition-all ${
                      step < currentStep
                        ? "bg-green-500"
                        : step === currentStep
                          ? "bg-white text-primary"
                          : "bg-white bg-opacity-20 text-white"
                    }`}
                  >
                    {step}
                  </div>
                  {index < 3 && (
                    <div className="w-16 h-1 bg-white bg-opacity-20 rounded"></div>
                  )}
                </div>
              ))}
            </div>

            <div className="text-center">
              <span className="current-step-text text-lg font-semibold">
                {stepTexts[currentStep - 1]}
              </span>
            </div>
          </div>
        </div>
      </section>

      {/* Main Form Container */}
      <section className="py-12">
        <div className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="bg-white rounded-2xl shadow-xl p-8">
            <form id="createEventForm" onSubmit={handleSubmit}>
              {/* Step 1: Event Details */}
              <div
                id="step1"
                className={`step-container ${currentStep === 1 ? "active" : ""}`}
                style={{ display: currentStep === 1 ? "block" : "none" }}
              >
                <h2 className="text-2xl font-bold text-gray-900 mb-6">
                  Tell us about your event
                </h2>

                <div className="space-y-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Event Name *
                    </label>
                    <input
                      type="text"
                      id="eventName"
                      required
                      value={formData.eventName}
                      onChange={handleInputChange}
                      className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                      placeholder="Enter your event name"
                    />
                  </div>

                  <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div>
                      <label className="block text-sm font-medium text-gray-700 mb-2">
                        Category *
                      </label>
                      <select
                        id="eventCategory"
                        required
                        value={formData.eventCategory}
                        onChange={handleInputChange}
                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                      >
                        <option value="">Select category</option>
                        <option value="technology">Technology</option>
                        <option value="business">Business</option>
                        <option value="arts">Arts & Culture</option>
                        <option value="sports">Sports</option>
                        <option value="education">Education</option>
                        <option value="networking">Networking</option>
                        <option value="workshop">Workshop</option>
                        <option value="conference">Conference</option>
                        <option value="other">Other</option>
                      </select>
                    </div>

                    <div>
                      <label className="block text-sm font-medium text-gray-700 mb-2">
                        Event Type *
                      </label>
                      <select
                        id="eventType"
                        required
                        value={eventType}
                        onChange={(e) => {
                          setEventType(e.target.value);
                          handleInputChange(e);
                        }}
                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                      >
                        <option value="">Select type</option>
                        <option value="in-person">In-Person</option>
                        <option value="virtual">Virtual</option>
                        <option value="hybrid">Hybrid</option>
                      </select>
                    </div>
                  </div>

                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Event Description *
                    </label>
                    <textarea
                      id="eventDescription"
                      required
                      rows={4}
                      value={formData.eventDescription}
                      onChange={handleInputChange}
                      className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                      placeholder="Describe your event, what attendees can expect, key highlights..."
                    ></textarea>
                  </div>

                  <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div>
                      <label className="block text-sm font-medium text-gray-700 mb-2">
                        Start Date & Time *
                      </label>
                      <input
                        type="datetime-local"
                        id="startDateTime"
                        required
                        value={formData.startDateTime}
                        onChange={handleInputChange}
                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                      />
                    </div>

                    <div>
                      <label className="block text-sm font-medium text-gray-700 mb-2">
                        End Date & Time *
                      </label>
                      <input
                        type="datetime-local"
                        id="endDateTime"
                        required
                        value={formData.endDateTime}
                        onChange={handleInputChange}
                        className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                      />
                    </div>
                  </div>
                </div>
              </div>

              {/* Step 2: Location & Capacity */}
              <div
                id="step2"
                className={`step-container ${currentStep === 2 ? "active" : ""}`}
                style={{ display: currentStep === 2 ? "block" : "none" }}
              >
                <h2 className="text-2xl font-bold text-gray-900 mb-6">
                  Where will your event take place?
                </h2>

                <div className="space-y-6">
                  {showLocationFields && (
                    <div id="locationFields">
                      <div className="mb-6">
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                          Venue Name *
                        </label>
                        <input
                          type="text"
                          id="venueName"
                          value={formData.venueName}
                          onChange={handleInputChange}
                          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                          placeholder="e.g., Nairobi Convention Center"
                        />
                      </div>

                      <div className="mb-6">
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                          Full Address *
                        </label>
                        <input
                          type="text"
                          id="venueAddress"
                          value={formData.venueAddress}
                          onChange={handleInputChange}
                          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                          placeholder="Street address, City, Country"
                        />
                      </div>

                      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                        <div>
                          <label className="block text-sm font-medium text-gray-700 mb-2">
                            City
                          </label>
                          <input
                            type="text"
                            id="venueCity"
                            value={formData.venueCity}
                            onChange={handleInputChange}
                            className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                            placeholder="Nairobi"
                          />
                        </div>
                        <div>
                          <label className="block text-sm font-medium text-gray-700 mb-2">
                            Country
                          </label>
                          <input
                            type="text"
                            id="venueCountry"
                            value={formData.venueCountry}
                            onChange={handleInputChange}
                            className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                            placeholder="Kenya"
                          />
                        </div>
                      </div>
                    </div>
                  )}

                  {showVirtualFields && (
                    <div id="virtualFields">
                      <div className="bg-blue-50 border border-blue-200 rounded-lg p-6">
                        <h3 className="text-lg font-semibold text-blue-900 mb-4">
                          Virtual Event Details
                        </h3>
                        <div>
                          <label className="block text-sm font-medium text-blue-700 mb-2">
                            Platform
                          </label>
                          <select
                            id="virtualPlatform"
                            value={formData.virtualPlatform}
                            onChange={handleInputChange}
                            className="w-full px-4 py-3 border border-blue-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                          >
                            <option value="">Select platform</option>
                            <option value="zoom">Zoom</option>
                            <option value="teams">Microsoft Teams</option>
                            <option value="meet">Google Meet</option>
                            <option value="webex">Cisco Webex</option>
                            <option value="other">Other</option>
                          </select>
                        </div>
                      </div>
                    </div>
                  )}

                  <div className="bg-yellow-50 border border-yellow-200 rounded-lg p-6">
                    <h3 className="text-lg font-semibold text-gray-900 mb-4">
                      Expected Attendance & Pricing
                    </h3>
                    <p className="text-sm text-gray-600 mb-4">
                      Choose your expected attendance to determine the creation
                      fee. This is a one-time payment to publish your event.
                    </p>

                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                      {/* Small Event */}
                      <div
                        className={`pricing-card border-2 rounded-xl p-6 bg-white cursor-pointer transition-all hover:transform hover:-translate-y-1 hover:shadow-xl ${
                          selectedPricing === "small"
                            ? "border-primary shadow-lg"
                            : "border-gray-200"
                        }`}
                        onClick={() => selectPricing("small")}
                      >
                        <div className="flex items-center justify-between mb-4">
                          <div>
                            <h4 className="text-lg font-bold text-gray-900">
                              Small Event
                            </h4>
                            <p className="text-sm text-gray-600">
                              Up to 199 attendees
                            </p>
                          </div>
                          <div
                            className={`w-6 h-6 border-2 rounded-full flex items-center justify-center ${
                              selectedPricing === "small"
                                ? "border-primary bg-primary"
                                : "border-gray-300"
                            }`}
                          >
                            {selectedPricing === "small" && (
                              <div className="w-2 h-2 bg-white rounded-full"></div>
                            )}
                          </div>
                        </div>
                        <div className="text-center">
                          <div className="text-3xl font-bold text-green-600 mb-2">
                            KSh 2,000
                          </div>
                          <p className="text-xs text-gray-500">
                            One-time creation fee
                          </p>
                        </div>
                        <ul className="mt-4 space-y-2 text-sm text-gray-600">
                          <li className="flex items-center space-x-2">
                            <svg
                              className="w-4 h-4 text-green-500"
                              fill="none"
                              stroke="currentColor"
                              viewBox="0 0 24 24"
                            >
                              <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth={2}
                                d="M5 13l4 4L19 7"
                              ></path>
                            </svg>
                            <span>Full event management</span>
                          </li>
                          <li className="flex items-center space-x-2">
                            <svg
                              className="w-4 h-4 text-green-500"
                              fill="none"
                              stroke="currentColor"
                              viewBox="0 0 24 24"
                            >
                              <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth={2}
                                d="M5 13l4 4L19 7"
                              ></path>
                            </svg>
                            <span>Registration system</span>
                          </li>
                          <li className="flex items-center space-x-2">
                            <svg
                              className="w-4 h-4 text-green-500"
                              fill="none"
                              stroke="currentColor"
                              viewBox="0 0 24 24"
                            >
                              <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth={2}
                                d="M5 13l4 4L19 7"
                              ></path>
                            </svg>
                            <span>Analytics & insights</span>
                          </li>
                        </ul>
                      </div>

                      {/* Large Event */}
                      <div
                        className={`pricing-card border-2 rounded-xl p-6 bg-white cursor-pointer transition-all hover:transform hover:-translate-y-1 hover:shadow-xl ${
                          selectedPricing === "large"
                            ? "border-primary shadow-lg"
                            : "border-gray-200"
                        }`}
                        onClick={() => selectPricing("large")}
                      >
                        <div className="flex items-center justify-between mb-4">
                          <div>
                            <h4 className="text-lg font-bold text-gray-900">
                              Large Event
                            </h4>
                            <p className="text-sm text-gray-600">
                              200+ attendees
                            </p>
                          </div>
                          <div
                            className={`w-6 h-6 border-2 rounded-full flex items-center justify-center ${
                              selectedPricing === "large"
                                ? "border-primary bg-primary"
                                : "border-gray-300"
                            }`}
                          >
                            {selectedPricing === "large" && (
                              <div className="w-2 h-2 bg-white rounded-full"></div>
                            )}
                          </div>
                        </div>
                        <div className="text-center">
                          <div className="text-3xl font-bold text-blue-600 mb-2">
                            KSh 5,000
                          </div>
                          <p className="text-xs text-gray-500">
                            One-time creation fee
                          </p>
                        </div>
                        <ul className="mt-4 space-y-2 text-sm text-gray-600">
                          <li className="flex items-center space-x-2">
                            <svg
                              className="w-4 h-4 text-green-500"
                              fill="none"
                              stroke="currentColor"
                              viewBox="0 0 24 24"
                            >
                              <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth={2}
                                d="M5 13l4 4L19 7"
                              ></path>
                            </svg>
                            <span>Everything in Small Event</span>
                          </li>
                          <li className="flex items-center space-x-2">
                            <svg
                              className="w-4 h-4 text-green-500"
                              fill="none"
                              stroke="currentColor"
                              viewBox="0 0 24 24"
                            >
                              <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth={2}
                                d="M5 13l4 4L19 7"
                              ></path>
                            </svg>
                            <span>Priority support</span>
                          </li>
                          <li className="flex items-center space-x-2">
                            <svg
                              className="w-4 h-4 text-green-500"
                              fill="none"
                              stroke="currentColor"
                              viewBox="0 0 24 24"
                            >
                              <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth={2}
                                d="M5 13l4 4L19 7"
                              ></path>
                            </svg>
                            <span>Advanced analytics</span>
                          </li>
                          <li className="flex items-center space-x-2">
                            <svg
                              className="w-4 h-4 text-green-500"
                              fill="none"
                              stroke="currentColor"
                              viewBox="0 0 24 24"
                            >
                              <path
                                strokeLinecap="round"
                                strokeLinejoin="round"
                                strokeWidth={2}
                                d="M5 13l4 4L19 7"
                              ></path>
                            </svg>
                            <span>Custom branding options</span>
                          </li>
                        </ul>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              {/* Step 3: Event Rules & Settings */}
              <div
                id="step3"
                className={`step-container ${currentStep === 3 ? "active" : ""}`}
                style={{ display: currentStep === 3 ? "block" : "none" }}
              >
                <h2 className="text-2xl font-bold text-gray-900 mb-6">
                  Event Rules & Settings
                </h2>

                <div className="space-y-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Event Visibility
                    </label>
                    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                      <label className="flex items-center p-4 border-2 border-gray-200 rounded-lg cursor-pointer hover:border-primary transition-colors">
                        <input
                          type="radio"
                          name="visibility"
                          value="public"
                          className="mr-3"
                          checked={formData.visibility === "public"}
                          onChange={handleInputChange}
                        />
                        <div>
                          <div className="font-medium text-gray-900">
                            Public
                          </div>
                          <div className="text-sm text-gray-600">
                            Anyone can find and register
                          </div>
                        </div>
                      </label>
                      <label className="flex items-center p-4 border-2 border-gray-200 rounded-lg cursor-pointer hover:border-primary transition-colors">
                        <input
                          type="radio"
                          name="visibility"
                          value="private"
                          className="mr-3"
                          checked={formData.visibility === "private"}
                          onChange={handleInputChange}
                        />
                        <div>
                          <div className="font-medium text-gray-900">
                            Private
                          </div>
                          <div className="text-sm text-gray-600">
                            Invitation only
                          </div>
                        </div>
                      </label>
                    </div>
                  </div>

                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Registration Fee
                    </label>
                    <div className="flex items-center space-x-4">
                      <label className="flex items-center">
                        <input
                          type="radio"
                          name="registrationFee"
                          value="free"
                          className="mr-2"
                          checked={registrationFee === "free"}
                          onChange={(e) => togglePriceInput(e.target.value)}
                        />
                        <span>Free Event</span>
                      </label>
                      <label className="flex items-center">
                        <input
                          type="radio"
                          name="registrationFee"
                          value="paid"
                          className="mr-2"
                          checked={registrationFee === "paid"}
                          onChange={(e) => togglePriceInput(e.target.value)}
                        />
                        <span>Paid Event</span>
                      </label>
                    </div>
                    {registrationFee === "paid" && (
                      <div id="priceInputDiv" className="mt-4">
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                          Registration Price (KSh)
                        </label>
                        <input
                          type="number"
                          id="registrationPrice"
                          min="0"
                          value={formData.registrationPrice}
                          onChange={handleInputChange}
                          className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                          placeholder="0"
                        />
                      </div>
                    )}
                  </div>

                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Event Rules & Guidelines
                    </label>
                    <textarea
                      id="eventRules"
                      rows={5}
                      value={formData.eventRules}
                      onChange={handleInputChange}
                      className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                      placeholder="Enter any rules, guidelines, or requirements for attendees...
Example:
- Professional attire required
- Valid ID needed for entry
- No outside food or drinks
- Bring your own laptop for workshops"
                    ></textarea>
                  </div>

                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      Additional Information
                    </label>
                    <textarea
                      id="additionalInfo"
                      rows={3}
                      value={formData.additionalInfo}
                      onChange={handleInputChange}
                      className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                      placeholder="Any additional details, contact information, or special notes..."
                    ></textarea>
                  </div>
                </div>
              </div>

              {/* Step 4: Payment & Confirmation */}
              <div
                id="step4"
                className={`step-container ${currentStep === 4 ? "active" : ""}`}
                style={{ display: currentStep === 4 ? "block" : "none" }}
              >
                <h2 className="text-2xl font-bold text-gray-900 mb-6">
                  Review & Payment
                </h2>

                {/* Event Summary */}
                <div className="bg-gray-50 rounded-xl p-6 mb-8">
                  <h3 className="text-lg font-semibold text-gray-900 mb-4">
                    Event Summary
                  </h3>
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
                    <div>
                      <span className="text-gray-600">Event Name:</span>
                      <span className="font-medium text-gray-900 ml-2">
                        {formData.eventName || "-"}
                      </span>
                    </div>
                    <div>
                      <span className="text-gray-600">Category:</span>
                      <span className="font-medium text-gray-900 ml-2">
                        {formData.eventCategory || "-"}
                      </span>
                    </div>
                    <div>
                      <span className="text-gray-600">Date:</span>
                      <span className="font-medium text-gray-900 ml-2">
                        {formData.startDateTime
                          ? new Date(formData.startDateTime).toLocaleString()
                          : "-"}
                      </span>
                    </div>
                    <div>
                      <span className="text-gray-600">Type:</span>
                      <span className="font-medium text-gray-900 ml-2">
                        {eventType || "-"}
                      </span>
                    </div>
                    <div>
                      <span className="text-gray-600">Expected Attendees:</span>
                      <span className="font-medium text-gray-900 ml-2">
                        {selectedPricing === "small"
                          ? "Up to 199"
                          : selectedPricing === "large"
                            ? "200+"
                            : "-"}
                      </span>
                    </div>
                    <div>
                      <span className="text-gray-600">Registration:</span>
                      <span className="font-medium text-gray-900 ml-2">
                        {registrationFee === "free"
                          ? "Free"
                          : `KSh ${formData.registrationPrice || 0}`}
                      </span>
                    </div>
                  </div>
                </div>

                {/* Payment Information */}
                <div className="border-2 border-primary border-opacity-20 rounded-xl p-6 mb-8">
                  <h3 className="text-lg font-semibold text-gray-900 mb-4">
                    Payment Details
                  </h3>
                  <div className="flex justify-between items-center mb-4">
                    <span className="text-gray-600">Event Creation Fee:</span>
                    <span className="text-2xl font-bold text-primary">
                      KSh {getPaymentAmount().toLocaleString()}
                    </span>
                  </div>
                  <p className="text-sm text-gray-600 mb-6">
                    This is a <strong>one-time fee</strong> to create and
                    publish your event on the Oplan platform. Once paid,
                    you&apos;ll have full access to manage registrations, send
                    updates, and track analytics.
                  </p>

                  {/* Payment Options */}
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                    {/* M-Pesa Payment */}
                    <div className="border border-green-200 rounded-lg p-4 bg-green-50">
                      <div className="flex items-center space-x-3 mb-3">
                        <div className="w-10 h-10 bg-green-600 rounded-lg flex items-center justify-center">
                          <svg
                            className="w-5 h-5 text-white"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                          >
                            <path
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              strokeWidth={2}
                              d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z"
                            ></path>
                          </svg>
                        </div>
                        <h4 className="font-semibold text-gray-900">
                          M-Pesa Payment
                        </h4>
                      </div>
                      <div className="space-y-1 text-sm text-gray-700">
                        <p>
                          <strong>Paybill:</strong> 247247
                        </p>
                        <p>
                          <strong>Account:</strong> EVENTCREATE
                        </p>
                        <p>
                          <strong>Amount:</strong> KSh{" "}
                          {getPaymentAmount().toLocaleString()}
                        </p>
                      </div>
                    </div>

                    {/* Bank Transfer */}
                    <div className="border border-blue-200 rounded-lg p-4 bg-blue-50">
                      <div className="flex items-center space-x-3 mb-3">
                        <div className="w-10 h-10 bg-blue-600 rounded-lg flex items-center justify-center">
                          <svg
                            className="w-5 h-5 text-white"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                          >
                            <path
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              strokeWidth={2}
                              d="M3 10h18M7 15h1m4 0h1m-7 4h12a3 3 0 003-3V8a3 3 0 00-3-3H6a3 3 0 00-3 3v8a3 3 0 003 3z"
                            ></path>
                          </svg>
                        </div>
                        <h4 className="font-semibold text-gray-900">
                          Bank Transfer
                        </h4>
                      </div>
                      <div className="space-y-1 text-sm text-gray-700">
                        <p>
                          <strong>Bank:</strong> KCB Bank
                        </p>
                        <p>
                          <strong>Account:</strong> 1234567890
                        </p>
                        <p>
                          <strong>Amount:</strong> KSh{" "}
                          {getPaymentAmount().toLocaleString()}
                        </p>
                      </div>
                    </div>
                  </div>

                  {/* Payment Proof Upload */}
                  <div className="mt-6 p-4 bg-yellow-50 border border-yellow-200 rounded-lg">
                    <h4 className="font-semibold text-yellow-800 mb-2">
                      Upload Payment Proof
                    </h4>
                    <p className="text-sm text-yellow-700 mb-3">
                      Please upload a screenshot or photo of your payment
                      confirmation
                    </p>
                    <input
                      type="file"
                      accept="image/*,.pdf"
                      required
                      className="w-full px-4 py-3 border border-yellow-300 rounded-lg focus:ring-2 focus:ring-yellow-500 focus:border-yellow-500 bg-white"
                    />
                  </div>
                </div>
              </div>

              {/* Navigation Buttons */}
              <div className="flex justify-between mt-12 pt-8 border-t border-gray-200">
                <button
                  type="button"
                  id="prevBtn"
                  onClick={handlePrev}
                  className="px-8 py-3 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition-colors font-medium"
                  style={{ display: currentStep === 1 ? "none" : "block" }}
                >
                  ← Previous
                </button>
                <div className="flex space-x-4">
                  <button
                    type="button"
                    id="nextBtn"
                    onClick={handleNext}
                    className="px-8 py-3 bg-primary text-white rounded-lg hover:bg-indigo-700 transition-colors font-medium"
                    style={{
                      display: currentStep === totalSteps ? "none" : "block",
                    }}
                  >
                    Next Step →
                  </button>
                  <button
                    type="submit"
                    id="submitBtn"
                    className="px-8 py-3 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors font-medium"
                    style={{
                      display: currentStep === totalSteps ? "block" : "none",
                    }}
                  >
                    Create Event
                  </button>
                </div>
              </div>
            </form>
          </div>

          {/* Trust & Security */}
          <div className="mt-8 text-center">
            <div className="flex justify-center items-center space-x-8 text-gray-500">
              <div className="flex items-center space-x-2">
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"
                  ></path>
                </svg>
                <span className="text-sm">Secure Payment</span>
              </div>
              <div className="flex items-center space-x-2">
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192L5.636 18.364M12 12h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                  ></path>
                </svg>
                <span className="text-sm">24/7 Support</span>
              </div>
              <div className="flex items-center space-x-2">
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M13 10V3L4 14h7v7l9-11h-7z"
                  ></path>
                </svg>
                <span className="text-sm">Instant Setup</span>
              </div>
            </div>
          </div>
        </div>
      </section>
    </>
  );
}
