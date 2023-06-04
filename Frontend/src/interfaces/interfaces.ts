export interface follower {
  UAge: string;
  UEmail: string;
  UFirst: string;
  UGender: string;
  UId: number;
  ULast: string;
  UName: string;
  UNick: string;
  UPass: string;
  UPic: string;
  UText: string;
  UTime: string;
}

export interface group {
  Creator: number;
  Id: number;
  Name: string;
  Text: string;
}
export interface user {
  UAge: string;
  UEmail: string;
  UFirst: string;
  UGender: string;
  UId: number;
  ULast: string;
  UName: string;
  UNick: string;
  UPass: string;
  UPic: string;
  UPriv: string;
  UText: string;
  UTime: string;
}

export interface event {
  Creator: number;
  Gid: number;
  Id: number;
  Name: string;
  No: user[];
  NoCount: number;
  Text: string;
  Time: string;
  UName: string;
  UndecidedCount: number;
  Yes: user[];
}

export interface Chat {
  OId: number;
  OName: string;
  OActive: boolean;
  OProf: string;
  Unread: number;
}

export interface Message {
  ChBody: string;
  ChId: number;
  ChOId: number;
  ChOName: string;
  ChStatus: string;
  ChTime: string;
  ChUId: number;
  ChUName: string;
  Unread: number;
}

export interface gchatmsg {
  ChBody: string;
  ChGId: number;
  ChId: number;
  ChTime: string;
  ChUId: number;
  ChUName: string;
}

export interface thread {

}

export interface profileinf {
  uFirst: string,
  uLast: string,
  uAge: string,
  uGender: string,
  uEmail: string,
  uName: string,
  uTime: string,
  uPic: string,
  uText: string,
  uNick: string,
  uPriv: boolean,
  fStatus: boolean,
  uThreads: {},
}
