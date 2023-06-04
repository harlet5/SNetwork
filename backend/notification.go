package backend

import (
	"database/sql"
	"log"
	"real-time-forum/db"

	"github.com/gorilla/websocket"
)

type NData struct {
	Nfc []db.Notification
}

func ShowNotifications(uid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	ndata, err := db.GetNotifications(uid)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		r.Errrr = "Follow Notif error"
		conn.WriteJSON(r)
	}
	allData := &NData{
		Nfc: ndata,
	}
	reply := Reply{"Notifications", allData}
	conn.WriteJSON(reply)
}

func MakeNotifications(typ string, sender int, reciver int, gid int, gname string, ename string, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	nfc := db.Notification{
		Ntype:   typ,
		Sender:  sender,
		Reciver: reciver,
		Status:  "unseen",
		Gid:     gid,
		Gname:   gname,
		Ename:   ename,
	}
	x, err := db.GetSpecificNotification(nfc)
	log.Println("wut")
	log.Println(x)
	log.Println(err)
	if !x {
		log.Println("here")
		err := db.CreateNotification(nfc)
		if err != nil {
			r.Errrr = "Notification make ERR"
			conn.WriteJSON(r)
		}
		log.Println("hereagane")
		log.Println(sender)
		log.Println(reciver)
		reply1 := Reply{"NfcSent", nfc}
		reply2 := Reply{"NfcRecived", nfc}
		reply4 := Reply{"NfcNew", nfc}
		for _, oneConnection := range allConnections {
			log.Println(oneConnection.Uid)
			if oneConnection.Uid == sender {
				oneConnection.Conn.WriteJSON(reply1)
				oneConnection.Conn.WriteJSON(reply4)
			}
			if oneConnection.Uid == reciver {
				oneConnection.Conn.WriteJSON(reply2)
				oneConnection.Conn.WriteJSON(reply4)
			}
		}
	}
}

func ChangeNotificationStatus(typ string, sender int, reciver int, status string, gid int, gname string, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	nfc := db.Notification{
		Ntype:   typ,
		Sender:  sender,
		Reciver: reciver,
		Status:  status,
		Gid:     gid,
		Gname:   gname,
	}
	err := db.UpdateNfcStatus(nfc)
	if err != nil {
		r.Errrr = "chenge notif status ERR"
		conn.WriteJSON(r)
	}
	reply1 := Reply{"NfcSentChange", nfc}
	reply2 := Reply{"NfcRecivedChange", nfc}
	reply4 := Reply{"NfcNew", nfc}
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == sender {
			oneConnection.Conn.WriteJSON(reply1)
			oneConnection.Conn.WriteJSON(reply4)
		}
		if oneConnection.Uid == reciver {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}
