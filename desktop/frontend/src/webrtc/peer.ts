export class Peer {
  pc: RTCPeerConnection
  dataChannel?: RTCDataChannel

  constructor() {
    this.pc = new RTCPeerConnection({
      iceServers: [
        { urls: "stun:stun.l.google.com:19302" }
      ]
    })

    this.pc.onicecandidate = (event) => {
      if (event.candidate) {
        this.onIceCandidate?.(event.candidate)
      }
    }

    this.pc.ondatachannel = (event) => {
      this.dataChannel = event.channel
      this.setupDataChannel()
    }
  }

  onIceCandidate?: (candidate: RTCIceCandidate) => void
  onMessage?: (msg: string) => void

  createDataChannel() {
    this.dataChannel = this.pc.createDataChannel("data")
    this.setupDataChannel()
  }

  private setupDataChannel() {
    if (!this.dataChannel) return;

    this.dataChannel.onmessage = (e) => {
      this.onMessage?.(e.data)
    }
  }

  async createOffer() {
    const offer = await this.pc.createOffer()
    await this.pc.setLocalDescription(offer)
    return offer
  }

  async createAnswer(offer: RTCSessionDescriptionInit) {
    await this.pc.setRemoteDescription(offer)

    const answer = await this.pc.createAnswer()
    await this.pc.setLocalDescription(answer)

    return answer
  }

  async setAnswer(answer: RTCSessionDescriptionInit) {
    await this.pc.setRemoteDescription(answer)
  }

  async addIceCandidate(candidate: RTCIceCandidateInit) {
    await this.pc.addIceCandidate(candidate)
  }
}
