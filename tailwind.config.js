/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "internal/port/htmx/template/**/*.html",
    "internal/port/htmx/template/*.html",
  ],
  theme: {
    extend: {
      keyframes: {
        "fade-in": {
          "0%": { opacity: "0" },
          "100%": { opacity: "1" },
        },
        "fade-out": {
          "0%": { opacity: "1" },
          "100%": { opacity: "0" },
        },
      },
      animation: {
        "fade-in": "fade-in 2s ease-in-out",
        "fade-out": "fade-out 2s ease-in-out",
      },
    },
  },
  plugins: [],
};
