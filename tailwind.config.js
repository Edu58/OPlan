/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./internal/frontend/**/*.templ", "./internal/frontend/**/*.html"],
  theme: {
    extend: {
      colors: {
        primary: "#6366f1",
        secondary: "#8b5cf6",
      },
      animation: {
        "fade-in": "fadeIn 0.6s ease-in-out",
        "slide-up": "slideUp 0.4s ease-out",
        "bounce-in": "bounceIn 0.6s ease-out",
      },
    },
  },
  plugins: [],
};
