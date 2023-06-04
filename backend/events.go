package backend

import (
	"database/sql"
	"log"
	"real-time-forum/db"
	"strconv"

	"github.com/gorilla/websocket"
)

func MakeEvent(creator int, name string, description string, time string, gid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	_, err := db.GetEventByNameInGroup(name, gid)
	log.Println("here")
	log.Println(err)
	log.Println("here2")
	if err != sql.ErrNoRows && err != nil {
		r.Errrr = "EVENT make DB error"
		conn.WriteJSON(r)
	} else {
		event := db.Event{
			Id:      0,
			Creator: creator,
			Name:    name,
			Text:    description,
			Time:    time,
			Gid:     gid,
		}
		err := db.CreateEvent(event)
		if err != nil {
			r.Errrr = "EVENT make ERR 1"
			conn.WriteJSON(r)
		}
		g, err := db.GetGroupById(gid)
		if err != nil {
			r.Errrr = "EVENT make ERR 2"
			conn.WriteJSON(r)
		}
		ugs, err := db.GetUserGroups(creator)
		if err != nil {
			r.Errrr = "EVENT make ERR 3"
			conn.WriteJSON(r)
		}
		ags, err := db.GetAllGroups()
		if err != nil {
			r.Errrr = "EVENT make ERR 4"
			conn.WriteJSON(r)
		}
		usr, err := db.GetGroupUsers(g.Id)
		if err != nil {
			r.Errrr = "EVENT make ERR 5"
			conn.WriteJSON(r)
		}
		gthrd, err := db.GetThreadsInGroup(g.Id)
		if err != nil {
			r.Errrr = "EVENT make ERR 6"
			conn.WriteJSON(r)
		}
		gevents, err := db.GetEventsByGroup(g.Id)
		if err != nil {
			r.Errrr = "EVENT make ERR 7"
			conn.WriteJSON(r)
		}
		in, _ := db.GetInvited(gid)
		out, _ := db.GetNotInvited(gid)
		y, _ := db.GetUserEvents(creator)
		allData := &GData{
			InGroups:   ugs,
			AllGroups:  ags,
			Group:      g,
			Members:    usr,
			Events:     gevents,
			Ustatus:    true,
			Threads:    gthrd,
			Invited:    in,
			Outsiders:  out,
			UserEvents: y,
		}
		reply := Reply{"group_" + strconv.Itoa(g.Id), allData}
		for _, i := range allData.Members {
			MakeNotifications("event", creator, i.UId, g.Id, g.Name, name, conn, allConnections)
		}
		for _, oneConnection := range allConnections {
			if contains(allData.Members, oneConnection.Uid) {
				oneConnection.Conn.WriteJSON(reply)
			}
		}
	}
}

func SetEventStatus(uid int, eid int, gid int, status string, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	x, _ := db.GetEventUserById(uid, eid)
	log.Println(x)
	if !x {
		log.Println("first time")
		err := db.AddUserEventConnection(uid, eid, status)
		if err != nil {
			r.Errrr = "EVENT set ERR"
			conn.WriteJSON(r)
		}
	} else {
		log.Println("second time")
		log.Println(status)
		err := db.UpdateUserEventConnection(uid, eid, status)
		if err != nil {
			r.Errrr = "EVENT set ERR"
			conn.WriteJSON(r)
		}
	}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "EVENT set ERR"
		conn.WriteJSON(r)
	}
	ugs, err := db.GetUserGroups(uid)
	if err != nil {
		r.Errrr = "EVENT set ERR"
		conn.WriteJSON(r)
	}
	ags, err := db.GetAllGroups()
	if err != nil {
		r.Errrr = "EVENT set ERR"
		conn.WriteJSON(r)
	}
	usr, err := db.GetGroupUsers(g.Id)
	if err != nil {
		r.Errrr = "EVENT set ERR"
		conn.WriteJSON(r)
	}
	gthrd, err := db.GetThreadsInGroup(g.Id)
	if err != nil {
		r.Errrr = "EVENT set ERR"
		conn.WriteJSON(r)
	}
	gevents, err := db.GetEventsByGroup(g.Id)
	if err != nil {
		r.Errrr = "EVENT set ERR"
		conn.WriteJSON(r)
	}
	y, _ := db.GetUserEvents(uid)
	in, _ := db.GetInvited(gid)
	out, _ := db.GetNotInvited(gid)
	allData := &GData{
		InGroups:   ugs,
		AllGroups:  ags,
		Group:      g,
		Members:    usr,
		Events:     gevents,
		Threads:    gthrd,
		Ustatus:    true,
		UserEvents: y,
		Invited:    in,
		Outsiders:  out,
	}
	reply := Reply{"group_" + strconv.Itoa(g.Id), allData}
	for _, oneConnection := range allConnections {
		if contains(allData.Members, oneConnection.Uid) {
			oneConnection.Conn.WriteJSON(reply)
		}
	}

}
