import { Box, CssBaseline } from '@mui/material'
import { ThemeProvider } from '@mui/material/styles'
import { default_theme } from './themes/default'
import Home from './pages/Home'

function App() {

  return (
    <ThemeProvider theme={default_theme}>
      <Box sx={{ height: "100vh", width: "100vw" }}>
        <CssBaseline />
        <Home />
      </Box>
    </ThemeProvider>
  )
}

export default App
