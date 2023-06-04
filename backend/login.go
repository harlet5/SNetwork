package backend

import (
	"log"
	"real-time-forum/db"
	"time"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UData struct {
	UName      string
	UId        int
	UEmail     string
	UTime      string
	UThreads   []db.Thread
	UFirst     string
	ULast      string
	UAge       string
	UGender    string
	UPic       string
	UNick      string
	UText      string
	Follow     FData
	FollowReq  FData
	Fstatus    bool
	UPriv      bool
	FReqStatus bool
}

func LogIn(uname string, upass string, conn *websocket.Conn, spooderman []USocket) []USocket {
	log.Println(spooderman)
	r := EData{}
	log.Println(uname)
	uid, err := db.GetUId(uname)
	if uid == 0 {
		uid, err = db.GetUIdByEmail(uname)
		log.Println(uid)
	}
	if err != nil {
		r.Errrr = "DB error for email/usernam"
		conn.WriteJSON(r)
	} else if uid == 0 {
		r.Errrr = "Username/email doesn't exist"
		conn.WriteJSON(r)
	} else {
		dbpass, err := db.GetPass(uname)
		if err != nil {
			r.Errrr = "DB error for password"
			conn.WriteJSON(r)
		}
		if dbpass == "" {
			dbpass, err = db.GetPassByemail(uname)
			if err != nil {
				r.Errrr = "DB error for password"
				conn.WriteJSON(r)
			}
		}
		err = bcrypt.CompareHashAndPassword([]byte(dbpass), []byte(upass))
		if err == nil {
			sid := uuid.NewV4()
			err := db.CreateSess(sid.String(), uid, time.Now().Unix())
			if err != nil {
				r.Errrr = "DB error for session"
				conn.WriteJSON(r)
			} else {
				user, err := db.GetUser(uid)
				if err != nil {
					r.Errrr = "DB error for user"
					conn.WriteJSON(r)
				}
				posts, err2 := db.GetThreadsByUser(user.UId)
				if err2 != nil {
					r.Errrr = "DB error for threads"
					conn.WriteJSON(r)
				}
				follow, err := ShowFollow(uid)
				if err != nil {
					r.Errrr = "FOLLOW error"
					conn.WriteJSON(r)
				}
				req, err := ShowFollowReq(uid)
				if err != nil {
					r.Errrr = "FOLLOWREQ error"
					conn.WriteJSON(r)
				}
				upriv := false
				switch user.UPriv {
				case "true":
					upriv = true
				default:
					upriv = false
				}
				spooderman = UpdateConnections(user.UId, "login", conn, spooderman)
				c, _ := db.GetUnseenNr(user.UId)
				conn.WriteJSON(Reply{"login", CoData{sid.String(), UData{user.UName, user.UId, user.UEmail, user.UTime, posts, user.UFirst, user.ULast, user.UAge, user.UGender, user.UPic, user.UNick, user.UText, follow, req, false, upriv, false}, c}})
				for _, con := range spooderman {
					if con.Uid != user.UId {
						con.Conn.WriteJSON(Reply{"added_u", user.UId})
					}
				}
			}
		} else {
			r.Errrr = "Password is incorrect"
			conn.WriteJSON(r)
		}
	}
	log.Println(spooderman)
	return spooderman
}

func Logout(id int, conn *websocket.Conn, spooderman []USocket) {
	db.DelSess(id)
	spooderman = UpdateConnections(id, "logout", conn, spooderman)
	conn.WriteJSON(Reply{"logout", id})
	for _, con := range spooderman {
		if con.Uid != id {
			con.Conn.WriteJSON(Reply{"loggedout", id})
		}
	}
}

func SignUp(ufirst string, ulast string, uage string, ugender string, uname string, upass string, uemail string, upic string, unick string, utext string, conn *websocket.Conn, spooderman []USocket) []USocket {
	log.Println(spooderman)
	r := EData{}
	uid, err := db.GetUId(uname)
	uid2, errr := db.GetUIdByEmail(uemail)
	if uid != 0 {
		r.Errrr = "Username used"
		conn.WriteJSON(r)
	} else if err != nil || errr != nil {
		r.Errrr = "DB error"
		conn.WriteJSON(r)
	} else if uid2 != 0 {
		r.Errrr = "Email used"
		conn.WriteJSON(r)
	} else {
		ctime := time.Now().Format("January 2, 2006")
		epass, err := bcrypt.GenerateFromPassword([]byte(upass), bcrypt.DefaultCost)
		if err != nil {
			r.Errrr = "Encryption error"
			conn.WriteJSON(r)
		}
		user := db.User{UId: 0, UFirst: ufirst, ULast: ulast, UAge: uage, UGender: ugender, UEmail: uemail, UName: uname, UPass: string(epass), UTime: ctime, UPic: "", UNick: unick, UText: utext, UPriv: "false"}
		err = db.CreateUser(user)
		if err != nil {
			r.Errrr = "Creation error"
			conn.WriteJSON(r)
		} else {
			if err != nil {
				r.Errrr = "Session error"
				conn.WriteJSON(r)
			} else {
				log.Println("hi")
				u, _ := db.GetUserByUname(uname)
				if upic != "" {
					ProfPic(upic, u.UId, "profile_", "", conn)
				} else {
					db.UpdProfPic("default.png", u.UId)
				}
				u, _ = db.GetUserByUname(uname)
				spooderman = UpdateConnections(0, "new", conn, spooderman)
				conn.WriteJSON(Reply{"signup", CoData{"", UData{u.UName, u.UId, u.UEmail, u.UTime, []db.Thread{}, u.UFirst, u.ULast, u.UAge, u.UGender, u.UPic, user.UNick, user.UText, FData{}, FData{}, false, false, false}, 0}})
			}
		}
	}
	log.Println(spooderman)
	return spooderman
}
