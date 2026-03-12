# PeerDrop

PeerDrop is a small experimental project for **temporary peer-to-peer collaboration sessions**.

It uses a lightweight **Go signaling server** with WebSockets so clients can discover each other and establish direct P2P connections (e.g. WebRTC).

## Features

- WebSocket signaling server
- Create temporary sessions
- Join existing sessions
- JSON-based messages
- Designed for simple peer discovery

## Run

Requires Go.

```bash
git clone https://github.com/TompaSkitfet/peerdrop
cd peerdrop
go run main.go
