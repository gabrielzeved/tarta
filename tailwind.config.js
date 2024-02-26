/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{templ,go,html}"], // This is where your HTML templates / JSX files are located
  theme: {
    colors: {
      primary: {
        100: "#CCFFFB",
        200: "#99FFFE",
        300: "#66F4FF",
        400: "#3FE3FF",
        500: "#00C8FF",
        600: "#009BDB",
        700: "#0075B7",
        800: "#005393",
        900: "#003C7A",
      },
      secondary: {
        100: "#CAD9FD",
        200: "#97B2FB",
        300: "#6286F5",
        400: "#3A63EB",
        500: "#002EDE",
        600: "#0023BE",
        700: "#001A9F",
        800: "#001280",
        900: "#000C6A",
      },
      neutral: {
        100: "#FCFCFC",
        200: "#E7E4E3",
        300: "#B8B4B3",
        400: "#726E6E",
        500: "#151414",
        600: "#120E0E",
        700: "#0F0A0A",
        800: "#0C0607",
        900: "#0A0306",
      },
      success: {
        900: "#209E53",
        500: "#25B860",
        100: "#2AD16D",
      },
      warning: {
        900: "#E58200",
        500: "#FF9000",
        100: "#FFA633",
      },
      error: {
        900: "#DB3D3D",
        500: "#F64545",
        100: "#FF5C5C",
      },
    },
    extend: {
      fontFamily: {},
    },
  },
  plugins: [],
};
