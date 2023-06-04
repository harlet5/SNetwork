package backend

import (
	"real-time-forum/db"
	"sort"
	"strconv"

	"github.com/gorilla/websocket"
)

type CData struct {
	Chats  []db.Chat
	Others []db.Other
	Unseen int
}

func SendChatResponce(logedId int, conn *websocket.Conn) {
	newUsers := GetOthers(logedId)
	reply := Reply{"chatroom", newUsers}
	conn.WriteJSON(reply)
}

func GetOthers(logedId int) []db.Other {
	users := db.GetUsers(logedId)
	latestUsers := db.SortChats(logedId)
	latestID := []int{}
	newUsers := []db.Other{}
	for i := 0; i <= len(latestUsers)-1; i++ {
		if i < len(latestUsers) {
			if logedId == latestUsers[i].ChUId {
				latestID = append(latestID, latestUsers[i].ChOId)
			}
			if logedId == latestUsers[i].ChOId {
				latestID = append(latestID, latestUsers[i].ChUId)
			}
		}
	}
	for i := 0; i < len(latestID); i++ {
		if i < len(latestID)-1 {
			if latestID[i] == latestID[i+1] {
				latestID = SplitId(latestID, i)
			}
		}
	}
	for i := 0; i < len(latestID); i++ {
		for x := range users {
			if x < len(users) {
				if users[x].OId == latestID[i] {
					newUsers = append(newUsers, users[x])
					users = SplitOther(users, x)
				}
			}
		}
	}
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].OName < users[j].OName
	})
	for x := range users {
		users[x].Unread, _ = db.GetUnseenNrPerChat(logedId, users[x].OId)
		users[x].OProf, _ = db.GetProfPic(users[x].OId)
		newUsers = append(newUsers, users[x])
	}
	return newUsers
}

func SplitOther(s []db.Other, index int) []db.Other {
	return append(s[:index], s[index+1:]...)
}
func SplitId(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func ShowChat(fromId int, toId int, amount int, conn *websocket.Conn, allConnections []USocket) {
	response := EData{}
	data, err := db.GetChat(fromId, toId, amount)
	if err != nil {
		response.Errrr = "SHOW Chat ERR"
		conn.WriteJSON(response)
	}
	friendsFrom := GetOthers(fromId)
	friendsTo := GetOthers(toId)
	ufrom, _ := db.GetUnseenNr(fromId)
	uto, _ := db.GetUnseenNr(toId)
	allDataFrom := CData{
		Chats:  data,
		Others: friendsFrom,
		Unseen: ufrom,
	}
	allDataTo := CData{
		Chats:  data,
		Others: friendsTo,
		Unseen: uto,
	}
	toIdString := strconv.Itoa(toId)
	reply1 := Reply{"user_" + toIdString, allDataFrom}
	reply2 := Reply{"user_" + toIdString, allDataTo}
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == fromId {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == toId {
			oneConnection.Conn.WriteJSON(reply2)
		}
	}
}

func PostChat(fromId int, toId int, content string, date string, amount int, conn *websocket.Conn, allConnections []USocket) {
	response := EData{}
	err := db.CreateChat(fromId, toId, content, date)
	if err != nil {
		response.Errrr = "POST CREATE Chat ERR"
		conn.WriteJSON(response)
	}
	friends := GetOthers(fromId)
	friendsTo := GetOthers(toId)
	data, err := db.GetChat(fromId, toId, amount)
	rec, _ := db.GetUnseenNr(toId)
	sen, _ := db.GetUnseenNr(fromId)
	allData := CData{
		Chats:  data,
		Others: friends,
		Unseen: sen,
	}
	allData2 := CData{
		Chats:  data,
		Others: friendsTo,
		Unseen: rec,
	}
	if err != nil {
		response.Errrr = "POST GET Chat ERR"
		conn.WriteJSON(response)
	}
	toString := strconv.Itoa(toId)
	fromString := strconv.Itoa(fromId)
	reply4 := Reply{"newchat_" + toString, fromId}
	reply1 := Reply{"user_" + toString, allData}
	reply2 := Reply{"user_" + fromString, allData2}
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == fromId {
			oneConnection.Conn.WriteJSON(reply1)
		}
		if oneConnection.Uid == toId {
			oneConnection.Conn.WriteJSON(reply2)
			oneConnection.Conn.WriteJSON(reply4)
		}
	}
}

func ShowGroupChat(uid int, gid int, amount int, conn *websocket.Conn, allConnections []USocket) {
	response := EData{}
	data, err := db.GetGroupChat(gid, amount)
	if err != nil {
		response.Errrr = "SHOW GROUP Chat ERR"
		conn.WriteJSON(response)
	}
	toIdString := strconv.Itoa(gid)
	reply1 := Reply{"groupchat_" + toIdString, data}
	conn.WriteJSON(reply1)
	/*us, _ := db.GetGroupUsers(gid)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == uid {
			oneConnection.Conn.WriteJSON(reply1)
		} else if contains(us, oneConnection.Uid) {
			oneConnection.Conn.WriteJSON(reply1)
		}
	}*/
}

func PostGroupChat(uid int, gid int, content string, date string, amount int, conn *websocket.Conn, allConnections []USocket) {
	response := EData{}
	err := db.CreateGroupChat(uid, gid, content, date)
	if err != nil {
		response.Errrr = "POST CREATE GROUP Chat ERR"
		conn.WriteJSON(response)
	}
	data, err := db.GetGroupChat(gid, amount)
	if err != nil {
		response.Errrr = "POST GET GROUP Chat ERR"
		conn.WriteJSON(response)
	}
	toString := strconv.Itoa(gid)
	reply4 := Reply{"newgroupchat_" + toString, gid}
	reply1 := Reply{"groupchat_" + toString, data}
	us, _ := db.GetGroupUsers(gid)
	for _, oneConnection := range allConnections {
		if oneConnection.Uid == uid {
			oneConnection.Conn.WriteJSON(reply1)
		} else if contains(us, oneConnection.Uid) {
			oneConnection.Conn.WriteJSON(reply1)
			oneConnection.Conn.WriteJSON(reply4)
		}
	}
}

func contains(s []db.User, id int) bool {
	for _, v := range s {
		if v.UId == id {
			return true
		}
	}
	return false
}
