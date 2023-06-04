package backend

import (
	"log"
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
	log.Println("frick")
	if !containscon(cons, conn) {
		log.Println("frick1")
		cons = append(cons, USocket{Conn: conn, Uid: id})
	} else if oper == "login" {
		for i, con := range cons {
			if con.Conn == conn {
				log.Println("frick2")
				log.Println("why are you not here1")
				log.Println(con.Uid)
				cons[i].Uid = id
				return cons
			}
		}
	} else {
		for i, con := range cons {
			if con.Uid == id {
				log.Println("frick3")
				log.Println("why are you not here2")
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
				log.Println("should")
				ret := make([]USocket, 0)
				ret = append(ret, cons[:i]...)
				return append(ret, cons[i+1:]...)
			}
		}
	}
	return cons
}

func CheckSess(uid int, sid string) bool {
	log.Println("checking")
	if sid == "" {
		log.Println("checking")
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
			log.Println("why tho")
			return true
		}
	}
	return false
}
