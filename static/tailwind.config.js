/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./main.css", "./dist/*.html"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@catppuccin/tailwindcss')({
      defaultFlavour: 'mocha'
    })
  ],
}
