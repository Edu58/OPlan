"use client";

import { useState } from "react";
import Link from "next/link";

export default function DiscoverEvents() {
  const [searchQuery, setSearchQuery] = useState("");
  const [selectedCategory, setSelectedCategory] = useState("All Categories");
  const [activeFilter, setActiveFilter] = useState("all");
  const [sortBy, setSortBy] = useState("Date (Nearest)");

  const filters = [
    { id: "all", label: "All Events" },
    { id: "week", label: "This Week" },
    { id: "free", label: "Free Events" },
    { id: "virtual", label: "Virtual" },
    { id: "inperson", label: "In-Person" },
  ];

  const categories = [
    "All Categories",
    "Technology",
    "Business",
    "Arts & Culture",
    "Sports",
    "Education",
  ];

  const sortOptions = [
    "Date (Nearest)",
    "Popularity",
    "Price (Low to High)",
    "Price (High to Low)",
    "A-Z",
  ];

  const events = [
    {
      id: 1,
      title: "Tech Conference 2025",
      category: "Technology",
      categoryColor: "bg-blue-100 text-blue-800",
      gradientFrom: "from-blue-500",
      gradientTo: "to-purple-600",
      icon: "shield",
      initials: "TC",
      visibility: "Public",
      visibilityColor: "bg-green-500",
      attending: "245 attending",
      description:
        "Join industry leaders for a day of innovation, networking, and insights into the future of technology. Featuring keynote speakers from top tech companies.",
      date: "March 15, 2025 • 9:00 AM - 6:00 PM",
      location: "Nairobi Convention Center, Nairobi",
      locationType: "in-person",
      price: "KSh 2,500",
      isFavorited: false,
      rules: [
        "Professional attire required",
        "No outside food or drinks",
        "ID required for entry",
      ],
      buttonText: "Register Now",
    },
    {
      id: 2,
      title: "Digital Marketing Masterclass",
      category: "Business",
      categoryColor: "bg-purple-100 text-purple-800",
      gradientFrom: "from-purple-500",
      gradientTo: "to-pink-600",
      icon: "lightning",
      initials: "MW",
      visibility: "Private",
      visibilityColor: "bg-yellow-500",
      attending: "50 spots left",
      description:
        "Master the art of digital marketing with hands-on workshops covering SEO, social media, content marketing, and paid advertising strategies.",
      date: "March 22, 2025 • 10:00 AM - 4:00 PM",
      location: "Virtual Event (Zoom)",
      locationType: "virtual",
      price: "KSh 1,800",
      isFavorited: false,
      rules: [
        "Laptop required for exercises",
        "Camera must be on during sessions",
        "Invitation link sent 24h before",
      ],
      buttonText: "Request Invitation",
    },
    {
      id: 3,
      title: "Startup Founders Meetup",
      category: "Networking",
      categoryColor: "bg-green-100 text-green-800",
      gradientFrom: "from-green-500",
      gradientTo: "to-teal-600",
      icon: "star",
      initials: "SF",
      visibility: "Public",
      visibilityColor: "bg-green-500",
      attending: "89 attending",
      description:
        "Connect with fellow entrepreneurs, share experiences, and build meaningful relationships in Kenya's thriving startup ecosystem.",
      date: "April 5, 2025 • 6:00 PM - 9:00 PM",
      location: "iHub Nairobi, Ngong Road",
      locationType: "in-person",
      price: "Free",
      isFavorited: true,
      rules: [
        "Bring business cards",
        "Networking mindset required",
        "Light refreshments provided",
      ],
      buttonText: "Register Now",
    },
  ];

  const [eventsList, setEventsList] = useState(events);

  const toggleFavorite = (eventId) => {
    setEventsList((prev) =>
      prev.map((event) =>
        event.id === eventId
          ? { ...event, isFavorited: !event.isFavorited }
          : event,
      ),
    );
  };

  const renderIcon = (iconType) => {
    switch (iconType) {
      case "shield":
        return (
          <path d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
        );
      case "lightning":
        return <path d="M13 10V3L4 14h7v7l9-11h-7z" />;
      case "star":
        return (
          <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z" />
        );
      default:
        return null;
    }
  };

  return (
    <>
      {/* Hero Section */}
      <section className="gradient-bg pt-24 pb-16">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center text-white">
            <h1 className="text-5xl md:text-6xl font-bold mb-6 animate-fade-in">
              Discover Amazing
              <span className="block text-yellow-300">Events Near You</span>
            </h1>
            <p className="text-xl md:text-2xl mb-10 opacity-90 max-w-3xl mx-auto animate-slide-up">
              Connect with your community through extraordinary experiences.
              From tech meetups to cultural festivals.
            </p>

            {/* Search Bar */}
            <div className="max-w-4xl mx-auto bg-white rounded-2xl p-2 shadow-2xl animate-slide-up">
              <div className="flex flex-col md:flex-row gap-2">
                <div className="flex-1 relative">
                  <input
                    type="text"
                    placeholder="Search events..."
                    value={searchQuery}
                    onChange={(e) => setSearchQuery(e.target.value)}
                    className="w-full px-6 py-4 rounded-xl border-0 focus:ring-2 focus:ring-primary text-gray-900 placeholder-gray-500"
                  />
                  <svg
                    className="absolute right-4 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                    ></path>
                  </svg>
                </div>
                <div className="flex gap-2">
                  <select
                    value={selectedCategory}
                    onChange={(e) => setSelectedCategory(e.target.value)}
                    className="px-6 py-4 rounded-xl border-0 focus:ring-2 focus:ring-primary text-gray-900 bg-gray-50"
                  >
                    {categories.map((category) => (
                      <option key={category} value={category}>
                        {category}
                      </option>
                    ))}
                  </select>
                  <button className="bg-primary text-white px-8 py-4 rounded-xl hover:bg-indigo-700 transition-colors font-medium">
                    Search
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Filters and Sort */}
      <section className="py-8 bg-white border-b border-gray-200 sticky top-16 z-40">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-wrap items-center justify-between gap-4">
            <div className="flex flex-wrap items-center gap-3">
              <span className="text-gray-700 font-medium">Filter by:</span>
              {filters.map((filter) => (
                <button
                  key={filter.id}
                  onClick={() => setActiveFilter(filter.id)}
                  className={`px-4 py-2 rounded-full border-2 font-medium transition-all ${
                    activeFilter === filter.id
                      ? "border-primary bg-primary text-white"
                      : "border-gray-200 text-gray-600 hover:border-primary hover:text-primary"
                  }`}
                >
                  {filter.label}
                </button>
              ))}
            </div>
            <div className="flex items-center gap-3">
              <span className="text-gray-700 font-medium">Sort by:</span>
              <select
                value={sortBy}
                onChange={(e) => setSortBy(e.target.value)}
                className="px-4 py-2 rounded-lg border border-gray-200 focus:ring-2 focus:ring-primary focus:border-primary"
              >
                {sortOptions.map((option) => (
                  <option key={option} value={option}>
                    {option}
                  </option>
                ))}
              </select>
            </div>
          </div>
        </div>
      </section>

      {/* Events Grid */}
      <section className="py-12">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="mb-8">
            <h2 className="text-3xl font-bold text-gray-900 mb-2">
              Upcoming Events
            </h2>
            <p className="text-gray-600">Found 24 events in Nairobi</p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {eventsList.map((event) => (
              <div
                key={event.id}
                className="bg-white rounded-2xl shadow-lg overflow-hidden card-hover transition-all hover:transform hover:-translate-y-2 hover:shadow-2xl"
              >
                <div className="relative">
                  <div
                    className={`h-48 bg-linear-to-r ${event.gradientFrom} ${event.gradientTo} flex items-center justify-center`}
                  >
                    <div className="text-center text-white">
                      <div className="w-16 h-16 bg-white bg-opacity-20 rounded-full flex items-center justify-center mx-auto mb-3">
                        <svg
                          className="w-8 h-8"
                          fill="currentColor"
                          viewBox="0 0 24 24"
                        >
                          {renderIcon(event.icon)}
                        </svg>
                      </div>
                      <h3 className="text-lg font-bold">{event.initials}</h3>
                    </div>
                  </div>
                  <div className="absolute top-4 left-4">
                    <span
                      className={`${event.visibilityColor} text-white px-3 py-1 rounded-full text-sm font-medium`}
                    >
                      {event.visibility}
                    </span>
                  </div>
                  <div className="absolute top-4 right-4">
                    <button
                      onClick={() => toggleFavorite(event.id)}
                      className={`w-10 h-10 bg-white bg-opacity-90 rounded-full flex items-center justify-center transition-colors ${
                        event.isFavorited
                          ? "text-red-500"
                          : "text-gray-600 hover:text-red-500"
                      }`}
                    >
                      <svg
                        className="w-5 h-5"
                        fill={event.isFavorited ? "currentColor" : "none"}
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"
                        ></path>
                      </svg>
                    </button>
                  </div>
                </div>
                <div className="p-6">
                  <div className="flex items-center justify-between mb-3">
                    <span
                      className={`${event.categoryColor} px-3 py-1 rounded-full text-sm font-medium`}
                    >
                      {event.category}
                    </span>
                    <div className="flex items-center text-gray-500 text-sm">
                      <svg
                        className="w-4 h-4 mr-1"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                        ></path>
                      </svg>
                      {event.attending}
                    </div>
                  </div>
                  <h3 className="text-xl font-bold text-gray-900 mb-2">
                    {event.title}
                  </h3>
                  <p className="text-gray-600 mb-4 line-clamp-2">
                    {event.description}
                  </p>

                  <div className="space-y-3 mb-4">
                    <div className="flex items-center text-gray-600">
                      <svg
                        className="w-4 h-4 mr-3"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                        ></path>
                      </svg>
                      <span className="text-sm">{event.date}</span>
                    </div>
                    <div className="flex items-center text-gray-600">
                      {event.locationType === "virtual" ? (
                        <svg
                          className="w-4 h-4 mr-3"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M9.75 17L9 20l-1 1h8l-1-1-.75-3M3 13h18M5 17h14a2 2 0 002-2V5a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                          ></path>
                        </svg>
                      ) : (
                        <svg
                          className="w-4 h-4 mr-3"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
                          ></path>
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
                          ></path>
                        </svg>
                      )}
                      <span className="text-sm">{event.location}</span>
                    </div>
                    <div className="flex items-center text-gray-600">
                      <svg
                        className="w-4 h-4 mr-3"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"
                        ></path>
                      </svg>
                      <span className="text-sm font-semibold text-green-600">
                        {event.price}
                      </span>
                    </div>
                  </div>

                  <div className="mb-4">
                    <p className="text-xs text-gray-500 mb-2">Event Rules:</p>
                    <ul className="text-xs text-gray-600 space-y-1">
                      {event.rules.map((rule, index) => (
                        <li key={index}>• {rule}</li>
                      ))}
                    </ul>
                  </div>

                  <button className="w-full bg-primary text-white py-3 rounded-xl hover:bg-indigo-700 transition-colors font-medium">
                    {event.buttonText}
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>
    </>
  );
}
