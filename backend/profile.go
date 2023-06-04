package backend

import (
	"log"
	"real-time-forum/db"

	"github.com/gorilla/websocket"
)

func ShowProfile(userId int, viewedId int, conn *websocket.Conn) {
	r := EData{}
	posts, err2 := db.GetThreadsByUser(viewedId)
	user, err := db.GetUser(viewedId)
	if err2 != nil {
		r.Errrr = "Profile ERR"
		conn.WriteJSON(r)
	}
	if err != nil {
		r.Errrr = "Profile ERR"
		conn.WriteJSON(r)
	}
	follow, err := ShowFollow(viewedId)
	if err != nil {
		r.Errrr = "FOLLOW error"
		conn.WriteJSON(r)
	}
	req, err := ShowFollowReq(viewedId)
	if err != nil {
		r.Errrr = "FOLLOWREQ error"
		conn.WriteJSON(r)
	}
	fstatus, err := db.CheckFollowStatus(userId, viewedId)
	log.Println(fstatus)
	if err != nil {
		r.Errrr = "FOLLOWREQ error"
		conn.WriteJSON(r)
	}
	if userId == viewedId {
		fstatus = true
	}
	rstat, err := db.GetSpecificFollowRequest(userId, viewedId)
	if err != nil {
		r.Errrr = "FOLLOWREQ error"
		conn.WriteJSON(r)
	}
	userInfo := UData{}
	switch user.UPriv {
	case "true":
		switch fstatus {
		case true:
			userInfo = UData{user.UName, user.UId, user.UEmail, user.UTime, posts, user.UFirst, user.ULast, user.UAge, user.UGender, user.UPic, user.UNick, user.UText, follow, req, fstatus, true, rstat}
		default:
			userInfo = UData{user.UName, user.UId, "", "", []db.Thread{}, "", "", "", "", user.UPic, "", "", FData{}, FData{}, fstatus, true, rstat}
		}
	default:
		userInfo = UData{user.UName, user.UId, user.UEmail, user.UTime, posts, user.UFirst, user.ULast, user.UAge, user.UGender, user.UPic, user.UNick, user.UText, follow, req, fstatus, false, rstat}
	}
	reply := Reply{"front", userInfo}
	conn.WriteJSON(reply)
}

func SetPrivateProfile(userId int, conn *websocket.Conn) {
	r := EData{}
	posts, err2 := db.GetThreadsByUser(userId)
	user, err := db.GetUser(userId)
	if err2 != nil {
		r.Errrr = "Profile ERR"
		conn.WriteJSON(r)
	}
	if err != nil {
		r.Errrr = "Profile ERR"
		conn.WriteJSON(r)
	}
	follow, err := ShowFollow(userId)
	if err != nil {
		r.Errrr = "FOLLOW error"
		conn.WriteJSON(r)
	}
	req, err := ShowFollowReq(userId)
	if err != nil {
		r.Errrr = "FOLLOWREQ error"
		conn.WriteJSON(r)
	}
	userInfo := UData{}
	switch user.UPriv {
	case "false":
		log.Println(user.UPriv)
		err = db.UpdPrivacy(userId, "true")
		if err != nil {
			r.Errrr = "Profile ERR"
			conn.WriteJSON(r)
		}
		user, err := db.GetUser(userId)
		if err != nil {
			r.Errrr = "Profile ERR"
			conn.WriteJSON(r)
		}
		log.Println(user.UPriv)
		userInfo = UData{user.UName, user.UId, user.UEmail, user.UTime, posts, user.UFirst, user.ULast, user.UAge, user.UGender, user.UPic, user.UNick, user.UText, follow, req, true, true, false}
	default:
		log.Println(user.UPriv)
		err = db.UpdPrivacy(userId, "false")
		if err != nil {
			r.Errrr = "Profile ERR"
			conn.WriteJSON(r)
		}
		user, err := db.GetUser(userId)
		if err != nil {
			r.Errrr = "Profile ERR"
			conn.WriteJSON(r)
		}
		log.Println(user.UPriv)
		userInfo = UData{user.UName, user.UId, user.UEmail, user.UTime, posts, user.UFirst, user.ULast, user.UAge, user.UGender, user.UPic, user.UNick, user.UText, follow, req, true, false, false}
	}
	reply := Reply{"front", userInfo}
	conn.WriteJSON(reply)
}
