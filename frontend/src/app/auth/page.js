"use client";

import { useState, useEffect, useRef } from "react";
import Link from "next/link";

export default function AuthPage() {
  const [currentTab, setCurrentTab] = useState("signIn");
  const [currentStep, setCurrentStep] = useState({
    signIn: 1,
    signUp: 1,
  });
  const [signInTimer, setSignInTimer] = useState(60);
  const [signUpTimer, setSignUpTimer] = useState(60);
  const [isTimerActive, setIsTimerActive] = useState({
    signIn: false,
    signUp: false,
  });

  const [formData, setFormData] = useState({
    // Sign In
    signInIdentifier: "",
    signInOtp: ["", "", "", "", "", ""],
    // Sign Up
    firstName: "",
    lastName: "",
    signUpEmail: "",
    signUpPhone: "",
    signUpPassword: "",
    signUpOtp: ["", "", "", "", "", ""],
  });

  const signInTimerRef = useRef(null);
  const signUpTimerRef = useRef(null);

  // Timer effect
  useEffect(() => {
    if (isTimerActive.signIn && signInTimer > 0) {
      signInTimerRef.current = setInterval(() => {
        setSignInTimer((prev) => {
          if (prev <= 1) {
            setIsTimerActive((prev) => ({ ...prev, signIn: false }));
            return 60;
          }
          return prev - 1;
        });
      }, 1000);
    }

    return () => {
      if (signInTimerRef.current) clearInterval(signInTimerRef.current);
    };
  }, [isTimerActive.signIn, signInTimer]);

  useEffect(() => {
    if (isTimerActive.signUp && signUpTimer > 0) {
      signUpTimerRef.current = setInterval(() => {
        setSignUpTimer((prev) => {
          if (prev <= 1) {
            setIsTimerActive((prev) => ({ ...prev, signUp: false }));
            return 60;
          }
          return prev - 1;
        });
      }, 1000);
    }

    return () => {
      if (signUpTimerRef.current) clearInterval(signUpTimerRef.current);
    };
  }, [isTimerActive.signUp, signUpTimer]);

  const switchTab = (tab) => {
    setCurrentTab(tab);
    setCurrentStep({
      signIn: 1,
      signUp: 1,
    });
    setFormData({
      signInIdentifier: "",
      signInOtp: ["", "", "", "", "", ""],
      firstName: "",
      lastName: "",
      signUpEmail: "",
      signUpPhone: "",
      signUpPassword: "",
      signUpOtp: ["", "", "", "", "", ""],
    });
    setIsTimerActive({ signIn: false, signUp: false });
  };

  const handleInputChange = (e) => {
    const { id, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [id]: value,
    }));
  };

  const handleOtpChange = (index, value, type) => {
    if (!/^\d*$/.test(value)) return;

    const otpField = type === "signIn" ? "signInOtp" : "signUpOtp";
    const newOtp = [...formData[otpField]];
    newOtp[index] = value;

    setFormData((prev) => ({
      ...prev,
      [otpField]: newOtp,
    }));

    // Auto-focus next input
    if (value && index < 5) {
      const nextInput = document.querySelector(
        `.${type}-otp[data-index="${index + 1}"]`,
      );
      if (nextInput) nextInput.focus();
    }
  };

  const handleOtpKeyDown = (e, index, type) => {
    if (e.key === "Backspace" && !e.target.value && index > 0) {
      const prevInput = document.querySelector(
        `.${type}-otp[data-index="${index - 1}"]`,
      );
      if (prevInput) prevInput.focus();
    }
  };

  const handleOtpPaste = (e, type) => {
    e.preventDefault();
    const pasteData = e.clipboardData.getData("text").slice(0, 6);
    if (!/^\d+$/.test(pasteData)) return;

    const otpField = type === "signIn" ? "signInOtp" : "signUpOtp";
    const newOtp = pasteData.split("").concat(Array(6).fill("")).slice(0, 6);

    setFormData((prev) => ({
      ...prev,
      [otpField]: newOtp,
    }));

    const lastInput = document.querySelector(
      `.${type}-otp[data-index="${Math.min(pasteData.length - 1, 5)}"]`,
    );
    if (lastInput) lastInput.focus();
  };

  const startTimer = (type) => {
    setIsTimerActive((prev) => ({ ...prev, [type]: true }));
    if (type === "signIn") {
      setSignInTimer(60);
    } else {
      setSignUpTimer(60);
    }
  };

  const handleSignInStep1Submit = (e) => {
    e.preventDefault();
    if (!formData.signInIdentifier.trim()) {
      alert("Please enter your email or phone number");
      return;
    }
    setCurrentStep((prev) => ({ ...prev, signIn: 2 }));
    startTimer("signIn");
    alert("Verification code sent to " + formData.signInIdentifier);

    setTimeout(() => {
      const firstInput = document.querySelector('.signin-otp[data-index="0"]');
      if (firstInput) firstInput.focus();
    }, 100);
  };

  const handleSignInOtpSubmit = (e) => {
    e.preventDefault();
    const otp = formData.signInOtp.join("");
    if (otp.length === 6) {
      alert("Sign in successful! Redirecting to dashboard...");
      console.log("Signed in with OTP:", otp);
    } else {
      alert("Please enter the complete 6-digit code");
    }
  };

  const handleSignUpStep1Submit = (e) => {
    e.preventDefault();

    if (
      !formData.firstName.trim() ||
      !formData.lastName.trim() ||
      !formData.signUpEmail.trim() ||
      !formData.signUpPhone.trim()
    ) {
      alert("Please fill in all required fields");
      return;
    }

    if (formData.signUpPassword.length < 8) {
      alert("Password must be at least 8 characters long");
      return;
    }

    setCurrentStep((prev) => ({ ...prev, signUp: 2 }));
    startTimer("signUp");
    alert("Verification code sent to " + formData.signUpEmail);

    setTimeout(() => {
      const firstInput = document.querySelector('.signup-otp[data-index="0"]');
      if (firstInput) firstInput.focus();
    }, 100);
  };

  const handleSignUpOtpSubmit = (e) => {
    e.preventDefault();
    const otp = formData.signUpOtp.join("");
    if (otp.length === 6) {
      setCurrentStep((prev) => ({ ...prev, signUp: 3 }));
      console.log("Account created with OTP:", otp);
    } else {
      alert("Please enter the complete 6-digit code");
    }
  };

  return (
    <div className="auth-container min-h-screen relative">
      {/* Floating Background Shapes */}
      <div className="floating-shapes">
        <div className="shape"></div>
        <div className="shape"></div>
        <div className="shape"></div>
        <div className="shape"></div>
      </div>

      {/* Main Container */}
      <div className="min-h-screen flex items-center justify-center p-4 relative z-10">
        <div className="w-full max-w-md">
          {/* Header */}
          <div className="text-center mb-8 animate-fade-in">
            <h1 className="text-4xl font-bold text-white mb-2">Oplan</h1>
            <p className="text-white text-opacity-80">
              Welcome to your event management platform
            </p>
          </div>

          {/* Auth Card */}
          <div className="bg-white rounded-2xl shadow-2xl p-8 animate-bounce-in">
            {/* Toggle Buttons */}
            <div className="flex bg-gray-100 rounded-xl p-1 mb-8">
              <button
                onClick={() => switchTab("signIn")}
                className={`flex-1 py-3 px-4 rounded-lg font-medium transition-all duration-300 ${
                  currentTab === "signIn"
                    ? "bg-primary text-white"
                    : "text-gray-600 hover:text-gray-900"
                }`}
              >
                Sign In
              </button>
              <button
                onClick={() => switchTab("signUp")}
                className={`flex-1 py-3 px-4 rounded-lg font-medium transition-all duration-300 ${
                  currentTab === "signUp"
                    ? "bg-primary text-white"
                    : "text-gray-600 hover:text-gray-900"
                }`}
              >
                Sign Up
              </button>
            </div>

            {/* Sign In Form */}
            {currentTab === "signIn" && (
              <div>
                {/* Step 1: Email/Phone */}
                {currentStep.signIn === 1 && (
                  <div className="step-container active">
                    <h2 className="text-2xl font-bold text-gray-900 mb-2">
                      Welcome back!
                    </h2>
                    <p className="text-gray-600 mb-6">
                      Enter your email or phone number to continue
                    </p>

                    <form onSubmit={handleSignInStep1Submit}>
                      <div className="mb-6">
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                          Email or Phone Number
                        </label>
                        <input
                          type="text"
                          id="signInIdentifier"
                          required
                          value={formData.signInIdentifier}
                          onChange={handleInputChange}
                          className="placeholder w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                          placeholder="Enter email or phone number"
                        />
                      </div>

                      <button
                        type="submit"
                        className="w-full bg-primary text-white py-3 rounded-lg hover:bg-indigo-700 transition-colors font-medium mb-4"
                      >
                        Continue
                      </button>
                    </form>

                    <div className="text-center">
                      <p className="text-sm text-gray-600">
                        Don&apos;t have an account?{" "}
                        <button
                          onClick={() => switchTab("signUp")}
                          className="text-primary hover:text-indigo-700 font-medium"
                        >
                          Sign up
                        </button>
                      </p>
                    </div>
                  </div>
                )}

                {/* Step 2: OTP Verification */}
                {currentStep.signIn === 2 && (
                  <div className="step-container active">
                    <div className="text-center mb-6">
                      <div className="w-16 h-16 bg-primary bg-opacity-10 rounded-full flex items-center justify-center mx-auto mb-4">
                        <svg
                          className="w-8 h-8 text-primary"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M12 15v2m-6 4h12a2 2 0 002-2v-1a2 2 0 00-2-2H6a2 2 0 00-2 2v1a2 2 0 002 2zM12 9a3 3 0 100-6 3 3 0 000 6z"
                          ></path>
                        </svg>
                      </div>
                      <h2 className="text-2xl font-bold text-gray-900 mb-2">
                        Verify your identity
                      </h2>
                      <p className="text-gray-600 mb-6">
                        We&apos;ve sent a 6-digit code to
                        <br />
                        <span className="font-semibold text-gray-900">
                          {formData.signInIdentifier}
                        </span>
                      </p>
                    </div>

                    <form onSubmit={handleSignInOtpSubmit}>
                      <div className="mb-6">
                        <label className="block text-sm font-medium text-gray-700 mb-3 text-center">
                          Enter verification code
                        </label>
                        <div className="flex justify-center space-x-3 mb-4">
                          {[0, 1, 2, 3, 4, 5].map((index) => (
                            <input
                              key={index}
                              type="text"
                              maxLength={1}
                              value={formData.signInOtp[index]}
                              onChange={(e) =>
                                handleOtpChange(index, e.target.value, "signIn")
                              }
                              onKeyDown={(e) =>
                                handleOtpKeyDown(e, index, "signIn")
                              }
                              onPaste={(e) => handleOtpPaste(e, "signIn")}
                              className="otp-input signin-otp"
                              data-index={index}
                            />
                          ))}
                        </div>
                      </div>

                      <button
                        type="submit"
                        className="w-full bg-primary text-white py-3 rounded-lg hover:bg-indigo-700 transition-colors font-medium mb-4"
                      >
                        Verify & Sign In
                      </button>

                      <div className="text-center">
                        <p className="text-sm text-gray-600 mb-2">
                          Didn&apos;t receive the code?
                        </p>
                        <button
                          type="button"
                          onClick={() => {
                            if (!isTimerActive.signIn) {
                              startTimer("signIn");
                              alert("New verification code sent!");
                            }
                          }}
                          disabled={isTimerActive.signIn}
                          className="text-primary hover:text-indigo-700 font-medium text-sm disabled:opacity-50"
                        >
                          {isTimerActive.signIn
                            ? `Resend Code (${signInTimer}s)`
                            : "Resend Code"}
                        </button>
                      </div>

                      <button
                        type="button"
                        onClick={() =>
                          setCurrentStep((prev) => ({ ...prev, signIn: 1 }))
                        }
                        className="w-full mt-4 text-gray-600 hover:text-gray-900 text-sm"
                      >
                        ← Back to email/phone
                      </button>
                    </form>
                  </div>
                )}
              </div>
            )}

            {/* Sign Up Form */}
            {currentTab === "signUp" && (
              <div>
                {/* Step 1: Basic Info */}
                {currentStep.signUp === 1 && (
                  <div className="step-container active">
                    <h2 className="text-2xl font-bold text-gray-900 mb-2">
                      Create your account
                    </h2>
                    <p className="text-gray-600 mb-6">
                      Join thousands of event organizers and attendees
                    </p>

                    <form onSubmit={handleSignUpStep1Submit}>
                      <div className="grid grid-cols-2 gap-4 mb-4">
                        <div>
                          <label className="block text-sm font-medium text-gray-700 mb-2">
                            First Name
                          </label>
                          <input
                            type="text"
                            id="firstName"
                            required
                            value={formData.firstName}
                            onChange={handleInputChange}
                            className="placeholder w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                            placeholder="John"
                          />
                        </div>
                        <div>
                          <label className="block text-sm font-medium text-gray-700 mb-2">
                            Last Name
                          </label>
                          <input
                            type="text"
                            id="lastName"
                            required
                            value={formData.lastName}
                            onChange={handleInputChange}
                            className="placeholder w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                            placeholder="Doe"
                          />
                        </div>
                      </div>

                      <div className="mb-4">
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                          Email Address
                        </label>
                        <input
                          type="email"
                          id="signUpEmail"
                          required
                          value={formData.signUpEmail}
                          onChange={handleInputChange}
                          className="placeholder w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                          placeholder="john.doe@example.com"
                        />
                      </div>

                      <div className="mb-4">
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                          Phone Number
                        </label>
                        <input
                          type="tel"
                          id="signUpPhone"
                          required
                          value={formData.signUpPhone}
                          onChange={handleInputChange}
                          className="placeholder w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                          placeholder="+254 7XX XXX XXX"
                        />
                      </div>

                      <div className="mb-6">
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                          Password
                        </label>
                        <input
                          type="password"
                          id="signUpPassword"
                          required
                          value={formData.signUpPassword}
                          onChange={handleInputChange}
                          className="placeholder w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary focus:border-primary transition-colors"
                          placeholder="Create a strong password"
                        />
                        <p className="text-xs text-gray-500 mt-1">
                          Minimum 8 characters with letters and numbers
                        </p>
                      </div>

                      <button
                        type="submit"
                        className="w-full bg-primary text-white py-3 rounded-lg hover:bg-indigo-700 transition-colors font-medium mb-4"
                      >
                        Continue
                      </button>
                    </form>

                    <div className="text-center">
                      <p className="text-sm text-gray-600">
                        Already have an account?{" "}
                        <button
                          onClick={() => switchTab("signIn")}
                          className="text-primary hover:text-indigo-700 font-medium"
                        >
                          Sign in
                        </button>
                      </p>
                    </div>
                  </div>
                )}

                {/* Step 2: OTP Verification */}
                {currentStep.signUp === 2 && (
                  <div className="step-container active">
                    <div className="text-center mb-6">
                      <div className="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-4">
                        <svg
                          className="w-8 h-8 text-green-600"
                          fill="none"
                          stroke="currentColor"
                          viewBox="0 0 24 24"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                          ></path>
                        </svg>
                      </div>
                      <h2 className="text-2xl font-bold text-gray-900 mb-2">
                        Verify your email
                      </h2>
                      <p className="text-gray-600 mb-6">
                        We&apos;ve sent a 6-digit code to
                        <br />
                        <span className="font-semibold text-gray-900">
                          {formData.signUpEmail}
                        </span>
                      </p>
                    </div>

                    <form onSubmit={handleSignUpOtpSubmit}>
                      <div className="mb-6">
                        <label className="block text-sm font-medium text-gray-700 mb-3 text-center">
                          Enter verification code
                        </label>
                        <div className="flex justify-center space-x-3 mb-4">
                          {[0, 1, 2, 3, 4, 5].map((index) => (
                            <input
                              key={index}
                              type="text"
                              maxLength={1}
                              value={formData.signUpOtp[index]}
                              onChange={(e) =>
                                handleOtpChange(index, e.target.value, "signUp")
                              }
                              onKeyDown={(e) =>
                                handleOtpKeyDown(e, index, "signUp")
                              }
                              onPaste={(e) => handleOtpPaste(e, "signUp")}
                              className="otp-input signup-otp"
                              data-index={index}
                            />
                          ))}
                        </div>
                      </div>

                      <button
                        type="submit"
                        className="w-full bg-primary text-white py-3 rounded-lg hover:bg-indigo-700 transition-colors font-medium mb-4"
                      >
                        Verify & Create Account
                      </button>

                      <div className="text-center">
                        <p className="text-sm text-gray-600 mb-2">
                          Didn&apos;t receive the code?
                        </p>
                        <button
                          type="button"
                          onClick={() => {
                            if (!isTimerActive.signUp) {
                              startTimer("signUp");
                              alert(
                                "New verification code sent to your email!",
                              );
                            }
                          }}
                          disabled={isTimerActive.signUp}
                          className="text-primary hover:text-indigo-700 font-medium text-sm disabled:opacity-50"
                        >
                          {isTimerActive.signUp
                            ? `Resend Code (${signUpTimer}s)`
                            : "Resend Code"}
                        </button>
                      </div>

                      <button
                        type="button"
                        onClick={() =>
                          setCurrentStep((prev) => ({ ...prev, signUp: 1 }))
                        }
                        className="w-full mt-4 text-gray-600 hover:text-gray-900 text-sm"
                      >
                        ← Back to account details
                      </button>
                    </form>
                  </div>
                )}

                {/* Step 3: Success */}
                {currentStep.signUp === 3 && (
                  <div className="step-container active">
                    <div className="text-center">
                      <div className="w-20 h-20 bg-green-100 rounded-full flex items-center justify-center mx-auto mb-6">
                        <svg
                          className="w-10 h-10 text-green-600"
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
                      </div>
                      <h2 className="text-2xl font-bold text-gray-900 mb-2">
                        Welcome to Oplan!
                      </h2>
                      <p className="text-gray-600 mb-8">
                        Your account has been created successfully.
                        <br />
                        You can now start discovering and creating amazing
                        events.
                      </p>

                      <Link href="/dashboard">
                        <button className="w-full bg-primary text-white py-3 rounded-lg hover:bg-indigo-700 transition-colors font-medium mb-3">
                          Go to Dashboard
                        </button>
                      </Link>

                      <Link href="/events">
                        <button className="w-full border border-gray-300 text-gray-700 py-3 rounded-lg hover:bg-gray-50 transition-colors font-medium">
                          Browse Events
                        </button>
                      </Link>
                    </div>
                  </div>
                )}
              </div>
            )}
          </div>

          {/* Terms & Privacy */}
          <div className="text-center mt-6 text-white text-opacity-80 text-sm">
            <p>
              By continuing, you agree to our{" "}
              <Link
                href="/terms"
                className="text-white hover:text-opacity-100 underline"
              >
                Terms of Service
              </Link>{" "}
              and{" "}
              <Link
                href="/privacy"
                className="text-white hover:text-opacity-100 underline"
              >
                Privacy Policy
              </Link>
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
