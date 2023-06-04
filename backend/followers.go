package backend

import (
	"real-time-forum/db"
	"strconv"

	"github.com/gorilla/websocket"
)

type FReqData struct {
	Sender  db.User
	Reciver db.User
}

type FUpdate struct {
	Sender       db.User
	Reciver      db.User
	SenderFData  FData
	ReciverFdata FData
}

type FReqsData struct {
	Sent    []db.User
	Recived []db.User
}

type FData struct {
	Followed  []db.User
	Followers []db.User
}

func MakeFollowRequest(sid int, rid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	follow := db.Follow{
		Sid: sid,
		Rid: rid,
	}
	err := db.CreateFollowRequest(follow)
	if err != nil {
		r.Errrr = "Follow REQ ERR"
		conn.WriteJSON(r)
	}
	u1, err := db.GetUser(sid)
	if err != nil {
		r.Errrr = "Follow REQ ERR"
		conn.WriteJSON(r)
	}
	u2, err := db.GetUser(rid)
	if err != nil {
		r.Errrr = "Follow REQ ERR"
		conn.WriteJSON(r)
	}
	allData := &FReqData{
		Sender:  u1,
		Reciver: u2,
	}
	reply1 := Reply{"followreqsent_" + strconv.Itoa(u1.UId), allData}
	reply2 := Reply{"followreqrecived_" + strconv.Itoa(u2.UId), allData}
	MakeNotifications("follow", sid, rid, -1, "", "", conn, allConnections)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == sid {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == rid {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

func AcceptFollowReq(sid int, rid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	follow := db.Follow{
		Sid: sid,
		Rid: rid,
	}
	err := db.AcceptRequest(follow)
	if err != nil {
		r.Errrr = "Follow ACCEPT ERR"
		conn.WriteJSON(r)
	}
	u1, err := db.GetUser(sid)
	if err != nil {
		r.Errrr = "Follow ACCEPT ERR"
		conn.WriteJSON(r)
	}
	u2, err := db.GetUser(rid)
	if err != nil {
		r.Errrr = "Follow ACCEPT ERR"
		conn.WriteJSON(r)
	}
	ChangeNotificationStatus("follow", sid, rid, "accepted", -1, "", conn, allConnections)
	db.CreateFollower(follow)
	f1, _ := ShowFollow(u1.UId)
	f2, _ := ShowFollow(u2.UId)
	allData := &FUpdate{
		Sender:       u1,
		Reciver:      u2,
		SenderFData:  f1,
		ReciverFdata: f2,
	}
	reply1 := Reply{"followsendaccpted_" + strconv.Itoa(u1.UId), allData}
	reply2 := Reply{"followrecaccepted_" + strconv.Itoa(u2.UId), allData}
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == sid {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == rid {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

func DeclineFollowReq(sid int, rid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	follow := db.Follow{
		Sid: sid,
		Rid: rid,
	}
	err := db.DeleteRequest(follow)
	if err != nil {
		r.Errrr = "Follow DECLINE ERR"
		conn.WriteJSON(r)
	}
	u1, err := db.GetUser(sid)
	if err != nil {
		r.Errrr = "Follow DECLINE ERR"
		conn.WriteJSON(r)
	}
	u2, err := db.GetUser(rid)
	if err != nil {
		r.Errrr = "Follow DECLINE ERR"
		conn.WriteJSON(r)
	}
	allData := &FReqData{
		Sender:  u1,
		Reciver: u2,
	}
	reply1 := Reply{"followsenddeclined_" + strconv.Itoa(u1.UId), allData}
	reply2 := Reply{"followrecdeclined_" + strconv.Itoa(u2.UId), allData}
	ChangeNotificationStatus("follow", sid, rid, "declined", -1, "", conn, allConnections)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == sid {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == rid {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

func RemoveFollower(sid int, rid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	follow := db.Follow{
		Sid: sid,
		Rid: rid,
	}
	err := db.DeleteRequest(follow)
	if err != nil {
		r.Errrr = "Follow REMOVE ERR"
		conn.WriteJSON(r)
	}
	u1, err := db.GetUser(sid)
	if err != nil {
		r.Errrr = "Follow REMOVE ERR"
		conn.WriteJSON(r)
	}
	u2, err := db.GetUser(rid)
	if err != nil {
		r.Errrr = "Follow REMOVE ERR"
		conn.WriteJSON(r)
	}
	allData := &FReqData{
		Sender:  u1,
		Reciver: u2,
	}
	reply1 := Reply{"followsendremoved_" + strconv.Itoa(u1.UId), allData}
	reply2 := Reply{"followrecremoved_" + strconv.Itoa(u2.UId), allData}
	MakeNotifications("follow", sid, rid, -1, "", "", conn, allConnections)
	db.DeleteFollower(follow)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == sid {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == rid {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

// For NOTIFICATIONS LATER
func ShowFollow(uid int) (FData, error) {
	fdata := FData{}
	followed, err := db.GetFollowed(uid)
	if err != nil {
		return fdata, err
	}
	followers, err := db.GetFollowers(uid)
	if err != nil {
		return fdata, err
	}
	fdata.Followed = followed
	fdata.Followers = followers
	return fdata, nil
}
func ShowFollowReq(uid int) (FData, error) {
	fdata := FData{}
	followed, err := db.GetFollowRequestsSent(uid)
	if err != nil {
		return fdata, err
	}
	followers, err := db.GetFollowRequestsRecived(uid)
	if err != nil {
		return fdata, err
	}
	fdata.Followed = followed
	fdata.Followers = followers
	return fdata, nil
}
