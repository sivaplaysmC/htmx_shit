/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/**/*.{html,jss,css}"],
  theme: {
    extend: {},
  },
  plugins: [
    require("daisyui")
  ],
  daisyui: {
    themes: ["nord", "dim"],
  },
}

