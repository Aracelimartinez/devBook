/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{vue,js,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "color-primary": "#FF8900",
        "color-secondary": "#eeeeee",
        "color-primary-bg": "#FCFAF7",
        "color-text-header": "#292828",
        "color-text-body": "#373434",
        "color-text-details": "#666666",
      }
    },
    fontFamily: {
      Roboto: ["Roboto", "sans-serif"],
    },
    container: {
      padding: "2rem",
      center: true,
    }
  },
  plugins: [],
}
