# PeerDrop Learning Roadmap (Detailed v5)

Denna roadmap bryter ner alla tekniker som behövs för att bygga PeerDrop
(P2P file sharing via WebRTC).

---

# 🗂️ Projektstruktur

## 📦 Projekt 1: Go Signaling Server

- 1.1 WebSocket Chat
- 1.2 Sessions
- 1.3 Peer Targeting

## 📦 Projekt 2: React Client

- 2 Frontend WebSocket Client

## 📦 Projekt 3: WebRTC Sandbox

- 3 WebRTC manuellt

## 📦 Projekt 4: Full App

- 4 WebRTC via signaling
- 5 DataChannel
- 6 File Transfer

## 📦 Projekt 5 (Valfri)

- 7 Desktop

---

# 🧱 1.1 WebSocket Chat

**Språk:** Go\
**Paket:** github.com/gorilla/websocket

### API / Metoder

- websocket.Upgrader
- conn.ReadMessage()
- conn.WriteMessage()
- http.HandleFunc

### Structs

- map\[\*websocket.Conn\]bool

### Vad du lär dig

- WebSocket lifecycle
- Broadcast
- Goroutines

---

# 🧩 1.2 Sessions

**Språk:** Go\
**Paket:** standard library

### API / Metoder

- createSession()
- joinSession()
- leaveSession()

### Structs

- map\[string\]\*Session
- map\[string\]\*Client

### Vad du lär dig

- State management
- Concurrency (mutex)

---

# 🎯 1.3 Peer Targeting

**Språk:** Go\
**Paket:** gorilla/websocket

### API / Metoder

- sendToPeer()
- lookup via map

### Structs

```go
type Client struct {
    ID string
    Conn *websocket.Conn
}
```

### Vad du lär dig

- Routing
- Identifiering

---

# 🌐 2 Frontend WebSocket Client

**Språk:** TypeScript (React)\
**Paket:** Vite, React

### API / Metoder

- new WebSocket()
- socket.send()
- socket.onmessage
- socket.onopen

### Structs

- JS objects (JSON messages)

### Vad du lär dig

- Realtidskommunikation
- State sync

---

# 🔌 3 WebRTC (Manuell)

**Språk:** TypeScript

### API / Metoder

- RTCPeerConnection
- createOffer()
- createAnswer()
- setLocalDescription()
- setRemoteDescription()
- onicecandidate

### Structs

- RTCSessionDescription
- RTCIceCandidate

### Vad du lär dig

- SDP
- ICE

---

# 🔁 4 WebRTC via Signaling

**Språk:** Go + TypeScript

### API / Metoder

- WebSocket messaging
- RTCPeerConnection methods

### Structs

- JSON signaling messages

### Vad du lär dig

- Integration
- Automatisering

---

# 📡 5 DataChannel

**Språk:** TypeScript

### API / Metoder

- createDataChannel()
- dc.send()
- dc.onmessage

### Structs

- RTCDataChannel

### Vad du lär dig

- P2P data

---

# 📦 6 File Transfer

**Språk:** TypeScript

### API / Metoder

- FileReader
- file.slice()
- readAsArrayBuffer()

### Structs

- Blob
- ArrayBuffer

### Vad du lär dig

- Chunking
- Large files

---

# 🖥️ 7 Desktop

**Språk:** Go + TS

### Paket

- Wails / Electron

### API / Metoder

- IPC (Wails bridge / Electron IPC)

### Structs

- App state

### Vad du lär dig

- Packaging

---

# 🧠 Mental modell

WebSocket → signaling\
WebRTC → connection\
DataChannel → data
