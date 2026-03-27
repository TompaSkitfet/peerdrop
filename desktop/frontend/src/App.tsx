import { EventsOn } from "../wailsjs/runtime/runtime";
import { SendSignal } from "../wailsjs/go/main/App";
import { ChangeEvent, useEffect, useState } from "react";
import { Box, Stack } from "@mui/system";
import { Button, Paper, TextField, Typography } from "@mui/material";

function App() {
  const [msg, setMsg] = useState("");
  const [textValue, setTextValue] = useState("");

  useEffect(() => {
    const unsubscribe = EventsOn("signal", (message: string) => {
      setMsg(message);
    });
    return unsubscribe();
  }, []);

  function SendMessage() {
    SendSignal(textValue);
  }

  return (
    <Box
      sx={{
        m: 2,
        height: "100vh",
        width: "100vw",
        display: "flex",
        alignItems: "center",
        flexDirection: "column",
      }}
    >
      <Stack sx={{ alignItems: "center", gap: 2 }}>
        <Typography>Websocket Test</Typography>
        <TextField
          value={textValue}
          onChange={(e: ChangeEvent<HTMLInputElement>) => {
            setTextValue(e.target.value);
          }}
        />
        <Button fullWidth variant="contained" onClick={SendMessage}>
          Skicka
        </Button>
        <Typography>{msg}</Typography>
      </Stack>
    </Box>
  );
}

export default App;
