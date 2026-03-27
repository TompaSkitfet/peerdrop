import { createTheme } from "@mui/material/styles";

// Background: rgba(20, 30, 45, 1) — matches style.css / wails config
const background = "#141e2d";

const theme = createTheme({
  palette: {
    mode: "dark",
    background: {
      default: background,
      paper: "#1c2a3d",
    },
    primary: {
      main: "#4da8da",
      light: "#7dc4e8",
      dark: "#2b7aad",
      contrastText: "#ffffff",
    },
    secondary: {
      main: "#6c8ebf",
      light: "#94afd4",
      dark: "#47628f",
      contrastText: "#ffffff",
    },
    text: {
      primary: "#e8edf4",
      secondary: "#9badc4",
      disabled: "#556070",
    },
    divider: "rgba(255, 255, 255, 0.08)",
    action: {
      hover: "rgba(77, 168, 218, 0.08)",
      selected: "rgba(77, 168, 218, 0.14)",
      disabled: "rgba(255, 255, 255, 0.26)",
      disabledBackground: "rgba(255, 255, 255, 0.06)",
    },
  },
  typography: {
    fontFamily: [
      "Nunito",
      "-apple-system",
      "BlinkMacSystemFont",
      '"Segoe UI"',
      "Roboto",
      "Oxygen",
      "Ubuntu",
      "Cantarell",
      '"Fira Sans"',
      '"Droid Sans"',
      '"Helvetica Neue"',
      "sans-serif",
    ].join(","),
    h1: { fontWeight: 700 },
    h2: { fontWeight: 700 },
    h3: { fontWeight: 600 },
    h4: { fontWeight: 600 },
    h5: { fontWeight: 600 },
    h6: { fontWeight: 600 },
    button: { fontWeight: 600, textTransform: "none" },
  },
  shape: {
    borderRadius: 10,
  },
  components: {
    MuiCssBaseline: {
      styleOverrides: {
        html: {
          backgroundColor: background,
        },
        body: {
          backgroundColor: background,
        },
      },
    },
    MuiPaper: {
      styleOverrides: {
        root: {
          backgroundImage: "none",
          backgroundColor: "#1c2a3d",
        },
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: 8,
        },
        containedPrimary: {
          boxShadow: "0 2px 8px rgba(77, 168, 218, 0.25)",
          "&:hover": {
            boxShadow: "0 4px 14px rgba(77, 168, 218, 0.35)",
          },
        },
      },
    },
    MuiTextField: {
      defaultProps: {
        variant: "outlined",
      },
    },
    MuiOutlinedInput: {
      styleOverrides: {
        root: {
          "& .MuiOutlinedInput-notchedOutline": {
            borderColor: "rgba(255, 255, 255, 0.15)",
          },
          "&:hover .MuiOutlinedInput-notchedOutline": {
            borderColor: "rgba(77, 168, 218, 0.5)",
          },
        },
      },
    },
    MuiTooltip: {
      styleOverrides: {
        tooltip: {
          backgroundColor: "#253348",
          fontSize: "0.75rem",
        },
      },
    },
  },
});

export default theme;
