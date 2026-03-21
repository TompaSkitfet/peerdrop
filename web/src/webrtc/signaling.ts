
export class SignalingClient {
  private ws: WebSocket;

  constructor(url: string) {
    this.ws = new WebSocket(url)
  }


  onMessage(callback: (data: any) => void) {
    this.ws.onmessage = (event) => {
      const msg = JSON.parse(event.data)
      callback(msg)
    }
  }

  send(data: any) {
    this.ws.send(JSON.stringify(data))
  }
}
