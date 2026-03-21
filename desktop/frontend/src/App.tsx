import { ThemeProvider } from "@emotion/react";
import { CssBaseline } from "@mui/material";
import { Box } from "@mui/system";
import { default_theme } from "./themes/default";
import Home from "./views/Home";

function App() {
  return (
    <ThemeProvider theme={default_theme}>
      <CssBaseline />
      <Box sx={{ height: "100vh", width: "100vw" }}>
        <Home />
      </Box>
    </ThemeProvider>
  );
}

export default App;
