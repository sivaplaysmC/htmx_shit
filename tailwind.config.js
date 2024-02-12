/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: "jit",
  content: ["./**/**/*.{html,css,templ}"],
  plugins: [
    require("daisyui")
  ],
  daisyui: {
    themes: ["nord", "dark", "light", "cupcake"]
  }
}




