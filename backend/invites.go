package backend

import (
	"log"
	"real-time-forum/db"
	"strconv"

	"github.com/gorilla/websocket"
)

type GInv struct {
	User  db.User
	Group db.Group
}

type GNew struct {
	Users []db.User
	Group db.Group
}

type GInvs struct {
	Invites []db.Invite
}

type GReqs struct {
	Reqs []db.Req
}

func SendInvite(sender int, reciver int, gid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	err := db.CreateGroupInvite(sender, reciver, gid)
	if err != nil {
		r.Errrr = "INV send ERR1"
		log.Println(err)
		conn.WriteJSON(r)
	}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "INV send ERR2"
		log.Println(err)
		conn.WriteJSON(r)
	}
	sen, err := db.GetUser(sender)
	if err != nil {
		r.Errrr = "INV send ERR3"
		log.Println(err)
		conn.WriteJSON(r)
	}
	rec, err := db.GetUser(reciver)
	if err != nil {
		r.Errrr = "INV send ERR4"
		log.Println(err)
		conn.WriteJSON(r)
	}
	db.GetGroupById(gid)
	dataSender := GInv{
		User:  rec,
		Group: g,
	}
	dataReciver := GInv{
		User:  sen,
		Group: g,
	}
	MakeNotifications("invite", sen.UId, rec.UId, g.Id, g.Name, "", conn, allConnections)
	reply1 := Reply{"invrec_" + strconv.Itoa(rec.UId), dataReciver}
	reply2 := Reply{"invsend_" + strconv.Itoa(sen.UId), dataSender}
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == rec.UId {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == sen.UId {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

func DeclineInv(rid int, sid int, gid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	err := db.DeleteGroupInv(rid, gid)
	if err != nil {
		r.Errrr = "INV decline ERR1"
		log.Println(err)
		conn.WriteJSON(r)
	}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "INV decline ERR2"
		log.Println(err)
		conn.WriteJSON(r)
	}
	sen, err := db.GetUser(sid)
	if err != nil {
		r.Errrr = "INV decline ERR3"
		log.Println(err)
		conn.WriteJSON(r)
	}
	rec, err := db.GetUser(rid)
	if err != nil {
		r.Errrr = "INV decline ERR4"
		log.Println(err)
		conn.WriteJSON(r)
	}
	dataSender := GInv{
		User:  rec,
		Group: g,
	}
	dataReciver := GInv{
		User:  sen,
		Group: g,
	}
	ChangeNotificationStatus("invite", sen.UId, rec.UId, "declined", g.Id, g.Name, conn, allConnections)
	reply1 := Reply{"decinvrec_" + strconv.Itoa(rec.UId), dataReciver}
	reply2 := Reply{"decinvsend_" + strconv.Itoa(sen.UId), dataSender}
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == rec.UId {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == sen.UId {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

func AcceptInv(rid int, sid int, gid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	err := db.DeleteGroupInv(rid, gid)
	if err != nil {
		r.Errrr = "INV accept ERR1"
		log.Println(err)
		conn.WriteJSON(r)
	}
	err = db.AddUserGroupConnection(rid, gid)
	if err != nil {
		r.Errrr = "INV accept ERR2"
		log.Println(err)
		conn.WriteJSON(r)
	}
	u, err := db.GetGroupUsers(gid)
	if err != nil {
		r.Errrr = "INV accept ERR3"
		log.Println(err)
		conn.WriteJSON(r)
	}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "INV accept ERR4"
		log.Println(err)
		conn.WriteJSON(r)
	}
	sen, err := db.GetUser(sid)
	if err != nil {
		r.Errrr = "INV accept ERR5"
		log.Println(err)
		conn.WriteJSON(r)
	}
	rec, err := db.GetUser(rid)
	if err != nil {
		r.Errrr = "INV accept ERR6"
		log.Println(err)
		conn.WriteJSON(r)
	}
	dataSender := GInv{
		User:  rec,
		Group: g,
	}
	dataReciver := GInv{
		User:  sen,
		Group: g,
	}
	dataMembers := GNew{
		Users: u,
		Group: g,
	}
	reply1 := Reply{"accinvrec_" + strconv.Itoa(rec.UId), dataReciver}
	reply2 := Reply{"accinvsend_" + strconv.Itoa(sen.UId), dataSender}
	reply3 := Reply{"updatemembers_" + strconv.Itoa(gid), dataMembers}
	ChangeNotificationStatus("invite", sen.UId, rec.UId, "accepted", g.Id, g.Name, conn, allConnections)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == rec.UId {
			oneConnection.Conn.WriteJSON(reply1)
		} else if oneConnection.Uid == sen.UId {
			oneConnection.Conn.WriteJSON(reply2)
		}
		oneConnection.Conn.WriteJSON(reply3)
	}
}

func RequestJoin(sender int, gid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	err := db.CreateJoinReq(sender, gid)
	if err != nil {
		r.Errrr = "Join req ERR1"
		log.Println(err)
		conn.WriteJSON(r)
	}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "Join req  ERR2"
		log.Println(err)
		conn.WriteJSON(r)
	}
	sen, err := db.GetUser(sender)
	if err != nil {
		r.Errrr = "Join req  ERR3"
		log.Println(err)
		conn.WriteJSON(r)
	}
	rec, err := db.GetUser(g.Creator)
	if err != nil {
		r.Errrr = "Join req  ERR4"
		log.Println(err)
		conn.WriteJSON(r)
	}
	dataSender := GInv{
		User:  rec,
		Group: g,
	}
	dataReciver := GInv{
		User:  sen,
		Group: g,
	}
	reply1 := Reply{"joinrez_" + strconv.Itoa(rec.UId), dataReciver}
	reply2 := Reply{"joinsend_" + strconv.Itoa(sen.UId), dataSender}
	MakeNotifications("joinreq", sen.UId, rec.UId, g.Id, g.Name, "", conn, allConnections)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == rec.UId {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == sen.UId {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

func DeclineJoin(sender int, gid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	err := db.DeleteJoinRequest(sender, gid)
	if err != nil {
		r.Errrr = "Join decline ERR1"
		log.Println(err)
		conn.WriteJSON(r)
	}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "Join decline ERR2"
		log.Println(err)
		conn.WriteJSON(r)
	}
	sen, err := db.GetUser(sender)
	if err != nil {
		r.Errrr = "Join decline ERR3"
		log.Println(err)
		conn.WriteJSON(r)
	}
	rec, err := db.GetUser(g.Creator)
	if err != nil {
		r.Errrr = "Join decline ERR4"
		log.Println(err)
		conn.WriteJSON(r)
	}
	dataSender := GInv{
		User:  rec,
		Group: g,
	}
	dataReciver := GInv{
		User:  sen,
		Group: g,
	}
	reply1 := Reply{"decjoinrec_" + strconv.Itoa(rec.UId), dataReciver}
	reply2 := Reply{"decjoinsend_" + strconv.Itoa(sen.UId), dataSender}
	ChangeNotificationStatus("joinreq", sen.UId, rec.UId, "declined", g.Id, g.Name, conn, allConnections)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == rec.UId {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == sen.UId {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

func AcceptJoin(sender int, gid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	err := db.DeleteJoinRequest(sender, gid)
	if err != nil {
		r.Errrr = "Join accept ERR1"
		log.Println(err)
		conn.WriteJSON(r)
	}
	err = db.AddUserGroupConnection(sender, gid)
	if err != nil {
		r.Errrr = "Join accept ERR2"
		log.Println(err)
		conn.WriteJSON(r)
	}
	u, err := db.GetGroupUsers(gid)
	if err != nil {
		r.Errrr = "Join accept ERR3"
		log.Println(err)
		conn.WriteJSON(r)
	}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "Join accept ERR4"
		log.Println(err)
		conn.WriteJSON(r)
	}
	sen, err := db.GetUser(sender)
	if err != nil {
		r.Errrr = "Join accept ERR5"
		log.Println(err)
		conn.WriteJSON(r)
	}
	rec, err := db.GetUser(g.Creator)
	if err != nil {
		r.Errrr = "Join accept ERR6"
		log.Println(err)
		conn.WriteJSON(r)
	}
	dataSender := GInv{
		User:  rec,
		Group: g,
	}
	dataReciver := GInv{
		User:  sen,
		Group: g,
	}
	dataMembers := GNew{
		Users: u,
		Group: g,
	}
	reply1 := Reply{"accjoinrec_" + strconv.Itoa(rec.UId), dataReciver}
	reply2 := Reply{"accjoinsend_" + strconv.Itoa(sen.UId), dataSender}
	reply3 := Reply{"updatemembers_" + strconv.Itoa(gid), dataMembers}
	ChangeNotificationStatus("joinreq", sen.UId, rec.UId, "accepted", g.Id, g.Name, conn, allConnections)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == rec.UId {
			oneConnection.Conn.WriteJSON(reply1)
		} else if oneConnection.Uid == sen.UId {
			oneConnection.Conn.WriteJSON(reply2)
		}
		oneConnection.Conn.WriteJSON(reply3)
	}
}

func SeeInvites(reciver int, conn *websocket.Conn) {
	r := EData{}
	invs, err := db.GetUserInvites(reciver)
	if err != nil {
		r.Errrr = "Inv see ERR"
		log.Println(err)
		conn.WriteJSON(r)
	}
	allData := &GInvs{
		Invites: invs,
	}
	reply := Reply{"Invites", allData}
	conn.WriteJSON(reply)
}

func SeeRequests(gid int, conn *websocket.Conn) {
	r := EData{}
	reqs, err := db.GetJoinRequests(gid)
	if err != nil {
		r.Errrr = "Req see ERR"
		log.Println(err)
		conn.WriteJSON(r)
	}
	allData := &GReqs{
		Reqs: reqs,
	}
	reply := Reply{"Reqs", allData}
	conn.WriteJSON(reply)
}
