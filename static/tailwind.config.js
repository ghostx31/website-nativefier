/** @type {import('tailwindcss').Config} */

module.exports = {
  content: ['./main.css', './dist/*.html'],
  theme: {
    extend: {
      fontFamily: {
        fantasque: ['Fantasque Sans Mono', 'monospace'],
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@catppuccin/tailwindcss')({
      defaultFlavour: ''
    })
  ],
}
