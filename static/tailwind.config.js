/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./main.css', './dist/*.html'],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@catppuccin/tailwindcss')({
      defaultFlavour: ''
    })
  ],
}
