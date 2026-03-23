import { Box, Button, TextField, Typography } from "@mui/material";
import "./App.css";
import { useEffect, useState, type ChangeEvent } from "react";

const socket = new WebSocket("ws://localhost:8080/ws");
function App() {
  const [value, setValue] = useState<string>("");

  useEffect(() => {
    socket.onopen = () => {
      console.log("connected");
    };
  }, [socket]);

  const onSend = () => {
    socket.send(value);
  };

  return (
    <Box sx={{ mt: 2 }}>
      <Typography>Test</Typography>
      <TextField
        value={value}
        onChange={(e: ChangeEvent<HTMLInputElement>) => {
          setValue(e.target.value);
        }}
      />
      <Button onClick={onSend} variant="contained">
        Send
      </Button>
    </Box>
  );
}

export default App;
