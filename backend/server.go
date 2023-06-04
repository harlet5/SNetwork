package backend

import (
	"real-time-forum/db"

	"github.com/gorilla/websocket"
)

type Data struct {
	Act     HarletIsStoopid
	Session map[string]interface{}
}

type HarletIsStoopid struct {
	Page string
	Data map[string]interface{}
}

type EData struct {
	Errrr string
}

type CoData struct {
	Sid    string
	Udata  UData
	Unseen int
}

type USocket struct {
	Conn *websocket.Conn
	Uid  int
}

type Reply struct {
	Act  string
	Data interface{}
}

func UpdateConnections(id int, oper string, conn *websocket.Conn, cons []USocket) []USocket {
	if !containscon(cons, conn) {
		cons = append(cons, USocket{Conn: conn, Uid: id})
	} else if oper == "login" {
		for i, con := range cons {
			if con.Conn == conn {
				cons[i].Uid = id
				return cons
			}
		}
	} else {
		for i, con := range cons {
			if con.Uid == id {
				cons[i].Uid = 0
				return cons
			}
		}
	}
	return cons
}

func DeleteConn(conn *websocket.Conn, cons []USocket) []USocket {
	if containscon(cons, conn) {
		for i, v := range cons {
			if v.Conn == conn {
				ret := make([]USocket, 0)
				ret = append(ret, cons[:i]...)
				return append(ret, cons[i+1:]...)
			}
		}
	}
	return cons
}

func CheckSess(uid int, sid string) bool {
	if sid == "" {
		return true
	}
	db.DelExpired(uid, sid)
	x, _ := db.GetSpesSess(sid, uid)
	if !x {
		return false
	}
	db.UpdateSession(sid, uid)
	return true
}

func containscon(s []USocket, conn *websocket.Conn) bool {
	for _, v := range s {
		if v.Conn == conn {
			return true
		}
	}
	return false
}
