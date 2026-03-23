import { useState, type ChangeEvent } from "react";
import "./App.css";
import { Box, Button, Stack, TextField } from "@mui/material";

const pc = new RTCPeerConnection({
  iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
});

const channel = pc.createDataChannel("data");
channel.onopen = () => {
  console.log("channel open");
};

channel.onmessage = (e) => {
  console.log("Recieved: ", e.data);
};
function App() {
  const [offer, setOffer] = useState<string>("");
  const [answer, setAnswer] = useState<string>("");
  const [answerInput, setAnswerInput] = useState<string>("");
  const [remoteAnswer, setRemoteAnswer] = useState<string>("");

  async function getOffer() {
    setTimeout(() => {
      setOffer(JSON.stringify(pc.localDescription));
    }, 2000);

    const offerDesc = await pc.createOffer();
    await pc.setLocalDescription(offerDesc);
  }

  async function getAnswer() {
    await pc.setRemoteDescription(JSON.parse(answerInput));
    const answer = await pc.createAnswer();
    await pc.setLocalDescription(answer);

    setTimeout(() => {
      setAnswer(JSON.stringify(pc.localDescription));
    });
  }

  async function setRemote() {
    await pc.setRemoteDescription(JSON.parse(remoteAnswer));
  }
  return (
    <Box>
      <Stack>
        <TextField sx={{ bgcolor: "white" }} multiline value={offer} />
        <Button variant="contained" onClick={getOffer} disabled={!!offer}>
          Generate Offer
        </Button>
      </Stack>
      <Stack mt={5}>
        <TextField
          sx={{ bgcolor: "white" }}
          value={answerInput}
          onChange={(e: ChangeEvent<HTMLInputElement>) => {
            setAnswerInput(e.target.value);
          }}
        />
        <Button variant="contained" onClick={getAnswer}>
          Generate Answer
        </Button>
        <TextField sx={{ bgcolor: "white" }} value={answer} />
      </Stack>
      <Stack mt={5}>
        <TextField
          sx={{ bgcolor: "white" }}
          value={remoteAnswer}
          onChange={(e: ChangeEvent<HTMLInputElement>) => {
            setRemoteAnswer(e.target.value);
          }}
        />
        <Button variant="contained" onClick={setRemote}>
          Set Answer
        </Button>
      </Stack>
    </Box>
  );
}

export default App;
