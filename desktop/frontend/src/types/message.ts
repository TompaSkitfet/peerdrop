export type Message = {
  type: MessageType;
  data?: Data;
};

type MessageType = "create_session" | "join_session" | "sdp" | "ice-candidate";

type Data = JoinSessionData | SDPData | ICECandidateDate;

export type JoinSessionData = {
  session_id: string;
};

export type SDPData = {
  type: string;
  sdp: string;
};

export type ICECandidateDate = {
  candidate: string;
  sdpMid: string;
  sdpMLineUndex: number;
};
