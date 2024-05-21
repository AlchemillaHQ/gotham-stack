const { addIconSelectors } = require('@iconify/tailwind');

/** @type {import('tailwindcss').Config} */

module.exports = {
  content: ["./public/**/*.js", "./templates/**/*.{html,templ,jet}"],
  theme: {
    extend: {},
  },
  plugins: [
    addIconSelectors(['mdi', 'mdi-light', 'flagpack']),
  ],
  darkMode: 'class',
}