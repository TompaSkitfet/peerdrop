import { Box, Button, TextField } from "@mui/material";
import { Peer } from "../webrtc/peer";
import { useEffect, useState } from "react";
import { SignalingClient } from "../webrtc/signaling";

const signaling = new SignalingClient("ws://localhost:8080/ws")

export default function Home() {
  const [peer] = useState(new Peer())
  const [myPeerId, setMyPeerId] = useState<string>()
  const [targetPeerId, setTargetPeerId] = useState<string>()
  const [id, setId] = useState<string>("")

  useEffect(() => {
    signaling.onMessage(async (msg) => {
      if (msg.type === "signal") {
        const { from, data } = msg.data

        switch (data.type) {
          case "offer":
            const answer = await peer.createAnswer(data)
            signaling.send({
              type: "signal",
              data: {
                to: from,
                data: answer
              }
            })
            break

          case "answer":
            await peer.setAnswer(data)
            break

          case "ice_candidate":
            await peer.addIceCandidate(data.candidate)
            break
        }
      }
      else {
        switch (msg.type) {
          case "session_created":
            setMyPeerId(msg.data.peerId)
            console.log("session created", msg.data.sessionId)
            break

          case "session_joined":
            setMyPeerId(msg.data.peerId)
            setTargetPeerId(msg.data.peers[0])
            break

          case "peer_joined":
            setTargetPeerId(msg.data.peerId)

            // 🔥 HÄR ska du starta WebRTC
            peer.createDataChannel()
            const offer = await peer.createOffer()

            signaling.send({
              type: "signal",
              data: {
                to: msg.data.peerId,
                data: offer
              }
            })
            break
        }
      }
    })


    peer.onIceCandidate = (candidate) => {
      signaling.send({
        type: "signal",
        data: {
          to: targetPeerId,
          data: {
            type: "ice_candidate",
            candidate
          }
        }
      })
    }
  }, [])


  const createSession = () => {
    signaling.send({
      type: "create_session"
    })
  }

  const joinSession = () => {
    signaling.send({
      type: "join_session",
      data: { sessionId: id }
    })
  }

  return <Box sx={{ height: "100%", width: "100%", display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center", gap: 2 }}>
    <Button sx={{ width: "20%" }} variant="contained" onClick={createSession}>Create Session</Button>
    <TextField value={id} onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
      setId(event.target.value)
    }} />
    < Button sx={{ width: "20%" }
    } variant="contained" onClick={joinSession}>Join Session</Button>
  </Box>
}
