/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./public/**/*.js", "./templates/**/*.{html,templ,jet}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
  darkMode: 'media',
}