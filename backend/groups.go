package backend

import (
	"database/sql"
	"log"
	"real-time-forum/db"
	"strconv"

	"github.com/gorilla/websocket"
)

type GData struct {
	InGroups   []db.Group
	AllGroups  []db.Group
	Group      db.Group
	Events     []db.Event
	Members    []db.User
	Invited    []db.User
	Outsiders  []db.User
	Threads    []db.Thread
	Creator    db.User
	Ustatus    bool
	UserEvents []db.UserEvent
	Ureqstatus bool
}

type GsData struct {
	InGroups  []db.Group
	AllGroups []db.Group
}

func MakeGroup(creator int, name string, description string, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	_, err := db.GetGroupByName(name)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "DB error"
		conn.WriteJSON(r)
	} else {
		if err != sql.ErrNoRows {
			r.Errrr = "Group exists"
			conn.WriteJSON(r)
			return
		}
		group := db.Group{
			Id:      0,
			Creator: creator,
			Name:    name,
			Text:    description,
		}
		err := db.CreateGroup(group)
		if err != nil {
			r.Errrr = "Group make ERR"
			conn.WriteJSON(r)
		}
		g, err := db.GetGroupByName(group.Name)
		if err != nil {
			r.Errrr = "Group make ERR"
			conn.WriteJSON(r)
		}
		ugs, err := db.GetUserGroups(g.Creator)
		if err != nil {
			r.Errrr = "Group make ERR"
			conn.WriteJSON(r)
		}
		ags, err := db.GetAllGroups()
		if err != nil {
			r.Errrr = "Group make ERR"
			conn.WriteJSON(r)
		}
		usr, err := db.GetGroupUsers(g.Id)
		if err != nil {
			r.Errrr = "Group make ERR"
			conn.WriteJSON(r)
		}
		u, err := db.GetUser(creator)
		if err != nil {
			r.Errrr = "Group make ERR"
			conn.WriteJSON(r)
		}
		allData := &GData{
			InGroups:   ugs,
			AllGroups:  ags,
			Group:      g,
			Members:    usr,
			Events:     []db.Event{},
			Threads:    []db.Thread{},
			Invited:    []db.User{},
			Outsiders:  []db.User{},
			Creator:    u,
			UserEvents: []db.UserEvent{},
			Ustatus:    true,
			Ureqstatus: false,
		}
		reply := Reply{"groups", allData}
		for _, oneConnection := range allConnections {
			oneConnection.Conn.WriteJSON(reply)
		}
	}
}

func ShowGroups(userId int, conn *websocket.Conn) {
	r := EData{}
	groups, err := db.GetAllGroups()
	if err != nil {
		r.Errrr = "Groups show ERR"
		conn.WriteJSON(r)
	}
	ugroups, err := db.GetUserGroups(userId)
	if err != nil {
		r.Errrr = "Groups show ERR"
		conn.WriteJSON(r)
	}
	allData := &GsData{
		InGroups:  ugroups,
		AllGroups: groups,
	}
	reply := Reply{"groups", allData}
	conn.WriteJSON(reply)
}

func ShowOneGroup(userId int, gid int, count int, conn *websocket.Conn, allConnections []USocket) { //either name or id for group might redo
	r := EData{}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "Group show ERR"
		conn.WriteJSON(r)
	}
	ugs, err := db.GetUserGroups(userId)
	if err != nil {
		r.Errrr = "Group show ERR"
		conn.WriteJSON(r)
	}
	ags, err := db.GetAllGroups()
	if err != nil {
		r.Errrr = "Group show ERR"
		conn.WriteJSON(r)
	}
	usr, err := db.GetGroupUsers(g.Id)
	if err != nil {
		r.Errrr = "Group show ERR"
		conn.WriteJSON(r)
	}
	u, err := db.GetUser(g.Creator)
	if err != nil {
		r.Errrr = "Group show ERR"
		conn.WriteJSON(r)
	}
	u.UPass = ""
	log.Println(g.Id)
	e, err := db.GetEventsByGroup(g.Id)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "Group show ERR 1"
		conn.WriteJSON(r)
	}
	t, err := db.GetThreadsInGroup(g.Id)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "Group show ERR 2"
		conn.WriteJSON(r)
	}
	invd, err := db.GetInvited(g.Id)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "Group show ERR 2"
		conn.WriteJSON(r)
	}
	other, err := db.GetNotInvited(g.Id)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "Group show ERR 2"
		conn.WriteJSON(r)
	}
	uevent, err := db.GetUserEvents(userId)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "Group show ERR uevent"
		conn.WriteJSON(r)
	}
	stat := false
	for _, i := range ugs {
		if i.Id == gid {
			stat = true
		}
	}
	rstat, err := db.GetJoinRequest(userId, gid)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "Group show ERR uevent"
		conn.WriteJSON(r)
	}
	allData := &GData{}
	if stat {
		allData = &GData{
			InGroups:   ugs,
			AllGroups:  ags,
			Group:      g,
			Members:    usr,
			Events:     e,
			Threads:    t,
			Creator:    u,
			Invited:    invd,
			Outsiders:  other,
			Ustatus:    stat,
			UserEvents: uevent,
			Ureqstatus: false,
		}
	} else {
		allData = &GData{
			InGroups:   ugs,
			AllGroups:  ags,
			Group:      g,
			Members:    []db.User{},
			Events:     []db.Event{},
			Threads:    []db.Thread{},
			Creator:    u,
			Invited:    []db.User{},
			Outsiders:  []db.User{},
			Ustatus:    stat,
			UserEvents: []db.UserEvent{},
			Ureqstatus: rstat,
		}
	}
	reply := Reply{"group_" + strconv.Itoa(g.Id), allData}
	conn.WriteJSON(reply)
	ShowGroupChat(userId, gid, count, conn, allConnections)
}

func LeaveGroup(uid int, gid int, conn *websocket.Conn, allConnections []USocket) {
	r := EData{}
	g, err := db.GetGroupById(gid)
	if err != nil {
		r.Errrr = "Group leave ERR"
		conn.WriteJSON(r)
	}
	err = db.RemoveUserFromGroup(uid, g.Id)
	if err != nil {
		r.Errrr = "Group leave ERR"
		conn.WriteJSON(r)
	}
	ugs, err := db.GetUserGroups(g.Creator)
	if err != nil {
		r.Errrr = "Group leave ERR"
		conn.WriteJSON(r)
	}
	ags, err := db.GetAllGroups()
	if err != nil {
		r.Errrr = "Group leave ERR"
		conn.WriteJSON(r)
	}
	usr, err := db.GetGroupUsers(g.Id)
	if err != nil {
		r.Errrr = "Group leave ERR"
		conn.WriteJSON(r)
	}
	u, err := db.GetUser(g.Id)
	if err != nil {
		r.Errrr = "Group leave ERR"
		conn.WriteJSON(r)
	}
	invd, err := db.GetInvited(g.Id)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "Group show ERR 2"
		conn.WriteJSON(r)
	}
	other, err := db.GetNotInvited(g.Id)
	if err != nil && err != sql.ErrNoRows {
		r.Errrr = "Group show ERR 2"
		conn.WriteJSON(r)
	}
	allData := &GData{
		InGroups:  ugs,
		AllGroups: ags,
		Group:     g,
		Members:   usr,
		Events:    []db.Event{},
		Threads:   []db.Thread{},
		Creator:   u,
		Invited:   invd,
		Outsiders: other,
	}
	reply := Reply{"leftgroup_" + strconv.Itoa(g.Id), allData}
	for _, oneConnection := range allConnections {
		if contains(allData.Members, oneConnection.Uid) {
			oneConnection.Conn.WriteJSON(reply)
		}
	}
	conn.WriteJSON(Reply{"Groups", allData})
}
