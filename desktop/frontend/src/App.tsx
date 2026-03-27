import { EventsOn } from "../wailsjs/runtime/runtime";
import { SendSignal } from "../wailsjs/go/main/App";
import { ChangeEvent, useEffect, useState } from "react";
import { Box, Stack } from "@mui/system";
import { Button, Divider, TextField, Typography } from "@mui/material";
import { Message } from "./types/message";

function App() {
  const [msg, setMsg] = useState("");
  const [textValue, setTextValue] = useState("");

  useEffect(() => {
    const unsubscribe = EventsOn("signal", (message: string) => {
      setMsg(message);
    });
    return unsubscribe;
  }, []);

  function CreateSession() {
    const payload: Message = { type: "create_session" };
    SendSignal(JSON.stringify(payload));
  }

  function JoinSession() {
    const payload: Message = {
      type: "join_session",
      data: { session_id: textValue },
    };
    SendSignal(JSON.stringify(payload));
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
        <Button fullWidth variant="contained" onClick={CreateSession}>
          Skapa session
        </Button>
        <Divider />
        <TextField
          value={textValue}
          onChange={(e: ChangeEvent<HTMLInputElement>) => {
            setTextValue(e.target.value);
          }}
        />
        <Button fullWidth variant="contained" onClick={JoinSession}>
          Anslut till session
        </Button>
        <Typography>{msg}</Typography>
      </Stack>
    </Box>
  );
}

export default App;
