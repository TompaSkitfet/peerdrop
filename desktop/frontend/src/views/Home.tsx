import { Button } from "@mui/material";
import { Box } from "@mui/system";
import { useEffect, useState } from "react";
import { Peer } from "../webrtc/peer";
import { SignalingClient } from "../webrtc/signaling";

const signaling = new SignalingClient("ws://localhost:8080/ws")


export default function Home() {
  const [peer] = useState(new Peer())

  useEffect(() => {
    signaling.onMessage(async (msg) => {
      switch (msg.type) {
        case "offer":
          const answer = await peer.createAnswer(msg.payload)
          signaling.send({ type: "answer", payload: answer })
          break
        case "answer":
          await peer.setAnswer(msg.payload)
          break
        case "ice":
          await peer.addIceCandidate(msg.payload)
          break
      }
    })

    peer.onIceCandidate = (candidate) => {
      signaling.send({ type: "ice", payload: candidate })
    }
  }, [])

  const createSession = async () => {
    peer.createDataChannel()

    const offer = await peer.createOffer()

    signaling.send({
      type: "offer",
      payload: offer
    })
  }




  return (
    <Box sx={{ height: "100%", width: "100%", display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center", gap: 2 }}>
      <Button variant="contained" onClick={createSession}>Create Session</Button>
      <Button variant="contained">Join Session</Button>
    </Box>
  );
}
