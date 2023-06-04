package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"real-time-forum/backend"
	"real-time-forum/db"
	"strconv"

	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

var upgrader = websocket.Upgrader{}
var spooderMan []backend.USocket

func main() {
	db.DbUtil()
	data, _ := os.Stat("./db/forum.db")
	if data == nil {
		db.FillTables()
	}
	db.ClearSess()
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("frontend"))))
	fmt.Println("http://localhost:8000")
	http.HandleFunc("/ws", scoreWithWebSocket)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func scoreWithWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	var conn, _ = upgrader.Upgrade(w, r, nil)
	go func(conn *websocket.Conn) {
		for {
			mType, msg, err := conn.ReadMessage()
			if err != nil || mType == 8 {
				log.Println("hereeeeee")
				spooderMan = backend.DeleteConn(conn, spooderMan)
				conn.Close()
				return
			}
			log.Println("spooderman: ", spooderMan)
			in := backend.Data{}
			out := backend.EData{}
			json.Unmarshal(msg, &in)
			log.Println("there")
			log.Println(in)
			if !backend.CheckSess(int(in.Session["userid"].(float64)), in.Session["sessionid"].(string)) {
				log.Println("checked logout")
				log.Println(int(in.Session["userid"].(float64)), in.Session["sessionid"].(string))
				spooderMan = backend.DeleteConn(conn, spooderMan)
				backend.Logout(int(in.Session["userid"].(float64)), conn, spooderMan)
			} else {
				switch in.Act.Page {
				case "front":
					backend.ShowProfile(int(in.Act.Data["cur"].(float64)), int(in.Act.Data["uid"].(float64)), conn)
				case "chatroom":
					backend.SendChatResponce(int(in.Act.Data["uid"].(float64)), conn)
				case "onechat":
					ouid, _ := strconv.Atoi(in.Act.Data["ouid"].(string))
					backend.ShowChat(int(in.Act.Data["uid"].(float64)), ouid, int(in.Act.Data["count"].(float64)), conn, spooderMan)
				case "newchat":
					ouid, _ := strconv.Atoi(in.Act.Data["ouid"].(string))
					backend.PostChat(int(in.Act.Data["uid"].(float64)), ouid, in.Act.Data["body"].(string), in.Act.Data["time"].(string), int(in.Act.Data["count"].(float64)), conn, spooderMan)
				case "threads":
					backend.ShowThreads(in.Act.Data["uname"].(string), conn, spooderMan)
				case "catthreads":
					backend.ShowCat(in.Act.Data["uname"].(string), in.Act.Data["cname"].(string), conn, spooderMan)
				case "onethread":
					backend.ShowThread(in.Act.Data["tid"].(string), spooderMan)
				case "newcomm":
					tid, err := strconv.Atoi(in.Act.Data["tid"].(string))
					if err != nil {
						out.Errrr = "Error: TId"
						conn.WriteJSON(out)
					}
					tmp := in.Act.Data["images"].([]interface{})
					img := []string{}
					for _, j := range tmp {
						img = append(img, j.(string))
					}
					backend.PostComm(tid, int(in.Act.Data["uid"].(float64)), in.Act.Data["body"].(string), in.Act.Data["time"].(string), in.Act.Data["tid"].(string), img, conn, spooderMan)
				case "newpost":
					tmp := in.Act.Data["categories"].([]interface{})
					cat := []string{}
					for _, j := range tmp {
						cat = append(cat, j.(string))
					}
					tmp = in.Act.Data["whosees"].([]interface{})
					who := []string{}
					for _, j := range tmp {
						who = append(who, j.(string))
					}
					tmp = in.Act.Data["images"].([]interface{})
					img := []string{}
					for _, j := range tmp {
						img = append(img, j.(string))
					}
					backend.PostThread(cat, int(in.Act.Data["uid"].(float64)), in.Act.Data["content"].(string), in.Act.Data["privacy"].(string), who, img, in.Act.Data["time"].(string), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "login":
					spooderMan = backend.LogIn(in.Act.Data["uname"].(string), in.Act.Data["upass"].(string), conn, spooderMan)
				case "logout":
					backend.Logout(int(in.Act.Data["cid"].(float64)), conn, spooderMan)
				case "signup":
					log.Println("here")
					spooderMan = backend.SignUp(in.Act.Data["ufname"].(string), in.Act.Data["ulname"].(string), in.Act.Data["uage"].(string), in.Act.Data["ugender"].(string), in.Act.Data["uname"].(string), in.Act.Data["upass"].(string), in.Act.Data["umail"].(string), in.Act.Data["upic"].(string), in.Act.Data["unickname"].(string), in.Act.Data["uaboutme"].(string), conn, spooderMan)
				case "groups":
					backend.ShowGroups(int(in.Act.Data["uid"].(float64)), conn)
				case "onegroup":
					backend.ShowOneGroup(int(in.Act.Data["uid"].(float64)), int(in.Act.Data["gid"].(float64)), int(in.Act.Data["count"].(float64)), conn, spooderMan)
				case "newgroup":
					backend.MakeGroup(int(in.Act.Data["uid"].(float64)), in.Act.Data["gname"].(string), in.Act.Data["gtext"].(string), conn, spooderMan)
				case "leavegroup":
					backend.LeaveGroup(int(in.Act.Data["uid"].(float64)), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "grpinv":
					backend.SendInvite(int(in.Act.Data["sid"].(float64)), int(in.Act.Data["rid"].(float64)), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "grpreq":
					backend.RequestJoin(int(in.Act.Data["sid"].(float64)), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "grpinvacc":
					backend.AcceptInv(int(in.Act.Data["rid"].(float64)), int(in.Act.Data["sid"].(float64)), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "grpreqacc":
					backend.AcceptJoin(int(in.Act.Data["sid"].(float64)), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "grpinvdec":
					backend.DeclineInv(int(in.Act.Data["rid"].(float64)), int(in.Act.Data["sid"].(float64)), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "grpreqdec":
					backend.DeclineJoin(int(in.Act.Data["sid"].(float64)), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "flwreq":
					backend.MakeFollowRequest(int(in.Act.Data["sid"].(float64)), int(in.Act.Data["rid"].(float64)), conn, spooderMan)
				case "flwacc":
					backend.AcceptFollowReq(int(in.Act.Data["sid"].(float64)), int(in.Act.Data["rid"].(float64)), conn, spooderMan)
				case "flwdec":
					backend.DeclineFollowReq(int(in.Act.Data["sid"].(float64)), int(in.Act.Data["rid"].(float64)), conn, spooderMan)
				case "flwrem":
					backend.RemoveFollower(int(in.Act.Data["sid"].(float64)), int(in.Act.Data["rid"].(float64)), conn, spooderMan)
				case "notifications":
					backend.ShowNotifications(int(in.Act.Data["uid"].(float64)), conn, spooderMan)
				case "newgroupchat":
					backend.PostGroupChat(int(in.Act.Data["uid"].(float64)), int(in.Act.Data["ouid"].(float64)), in.Act.Data["body"].(string), in.Act.Data["time"].(string), int(in.Act.Data["count"].(float64)), conn, spooderMan)
				case "makeevent":
					backend.MakeEvent(int(in.Act.Data["uid"].(float64)), in.Act.Data["name"].(string), in.Act.Data["desc"].(string), in.Act.Data["time"].(string), int(in.Act.Data["gid"].(float64)), conn, spooderMan)
				case "setevent":
					backend.SetEventStatus(int(in.Act.Data["uid"].(float64)), int(in.Act.Data["eid"].(float64)), int(in.Act.Data["gid"].(float64)), in.Act.Data["status"].(string), conn, spooderMan)
				case "setprivacy":
					backend.SetPrivateProfile(int(in.Act.Data["uid"].(float64)), conn)
				case "setconnid":
					spooderMan = backend.UpdateConnections(int(in.Act.Data["uid"].(float64)), "login", conn, spooderMan)
					if int(in.Act.Data["uid"].(float64)) > 0 {
						c, _ := db.GetUnseenNr(int(in.Act.Data["uid"].(float64)))
						conn.WriteJSON(backend.Reply{Act: "unread", Data: c})
					}
				}
			}
		}
	}(conn)
}

func init() {
	if err := db.InitDb(); err != nil {
		log.Fatal(err)
	}
	migratePtr := flag.Bool("migrate", false, "a bool")

	// populatePtr := flag.Bool("populate", false, "a bool")
	flag.Parse()

	if *migratePtr {
		if err := db.Migrate(); err != nil {
			log.Fatal(err)
		}
    db.DbUtil()
    db.FillTables()
	}

	// if *populatePtr {
	// 	if err := db.Populate(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
}
