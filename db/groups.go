package db

import (
	"database/sql"
	"log"
)

type Group struct {
	Id      int
	Creator int
	Name    string
	Text    string
}

type Invite struct {
	Sid int
	Rid int
	Gid int
}

type Req struct {
	Sid int
	Gid int
}

func CreateGroup(group Group) error {
	insertSQL := `INSERT INTO groups(creator, name, description) VALUES (?, ?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(group.Creator, group.Name, group.Text)
	if err != nil {
		return err
	}
	t, _ := GetGroupByName(group.Name)
	AddUserGroupConnection(group.Creator, t.Id)
	return nil
}

func GetGroupById(gid int) (Group, error) {
	var group Group
	row := DB.QueryRow("SELECT * FROM groups WHERE id = ?", gid)
	switch err := row.Scan(&group.Id, &group.Creator, &group.Name, &group.Text); err {
	case nil:
		return group, nil
	case sql.ErrNoRows:
		return group, err
	default:
		return group, err
	}
}

func GetAllGroups() ([]Group, error) {
	var group Group
	var groups []Group
	rows, err := DB.Query("SELECT * FROM groups")
	if err != nil {
		return groups, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&group.Id, &group.Creator, &group.Name, &group.Text); err {
		case nil:
			groups = append(groups, group)
		default:
			return groups, err
		}
	}
	return groups, nil
}

func GetGroupsByCreator(creator int) ([]Group, error) {
	var group Group
	var groups []Group
	rows, err := DB.Query("SELECT * FROM groups WHERE creator = ?", creator)
	if err != nil {
		return groups, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&group.Id, &group.Creator, &group.Name, &group.Text); err {
		case nil:
			groups = append(groups, group)
		default:
			return groups, err
		}
	}
	return groups, nil
}

func GetGroupByName(name string) (Group, error) {
	var group Group
	row := DB.QueryRow("SELECT * FROM groups WHERE name = ?", name)
	switch err := row.Scan(&group.Id, &group.Creator, &group.Name, &group.Text); err {
	case nil:
		return group, nil
	case sql.ErrNoRows:
		return group, err
	default:
		return group, err
	}
}

func AddUserGroupConnection(uid int, gid int) error {
	insertSQL := `INSERT INTO usergroups(uid, gid) VALUES (?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(uid, gid)
	if err != nil {
		return err
	}
	return nil
}

func GetUserGroups(uid int) ([]Group, error) {
	var ug Group
	var ugs []Group
	rows, err := DB.Query("SELECT * FROM groups WHERE id IN ( SELECT gid FROM usergroups WHERE uid = ?)", uid)
	if err != nil {
		return ugs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&ug.Id, &ug.Creator, &ug.Name, &ug.Text); err {
		case nil:
			ugs = append(ugs, ug)
		default:
			return ugs, err
		}
	}
	return ugs, nil
}

func GetGroupUsers(gid int) ([]User, error) {
	var user User
	var ugs []User
	rows, err := DB.Query("SELECT Users.UId, UFirst, ULast, UEmail, UAge, UGender, UName, UTime, UPic, UNick, UText FROM users WHERE Users.UId IN ( SELECT usergroups.uid FROM usergroups WHERE gid = ?)", gid)
	if err != nil {
		log.Println(err)
		return ugs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UAge, &user.UGender, &user.UEmail, &user.UName, &user.UTime, &user.UPic, &user.UNick, &user.UText); err {
		case nil:
			ugs = append(ugs, user)
		default:
			return ugs, err
		}
	}
	return ugs, nil
}

func RemoveUserFromGroup(uid int, gid int) error {
	_, err := DB.Exec("DELETE FROM usergroups WHERE uid = ? AND gid = ?", uid, gid)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}

func CreateGroupInvite(uid int, rid int, gid int) error {
	x, _ := GetSpecificInvite(rid, gid)
	if x {
		return nil
	}
	insertSQL := `INSERT INTO groupinvs(sid, rid, gid) VALUES(?,?,?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec(uid, rid, gid)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetUserInvites(rid int) ([]Invite, error) {
	var inv Invite
	var invs []Invite
	rows, err := DB.Query("SELECT * FROM groupinvs WHERE rid = ?", rid)
	if err != nil {
		return invs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&inv.Sid, &inv.Rid, &inv.Gid); err {
		case nil:
			invs = append(invs, inv)
		default:
			return invs, err
		}
	}
	return invs, nil
}

/*
	func DeleteSpecificInv(sid int, rid int, gid int) error {
		_, err := DB.Exec("DELETE FROM groupinvs WHERE sid = ? AND rid = ? AND gid = ?", sid, rid, gid)
		switch err {
		case nil:
			return nil
		default:
			return err
		}
	}
*/
func DeleteGroupInv(rid int, gid int) error {
	_, err := DB.Exec("DELETE FROM groupinvs WHERE rid = ? AND gid = ?", rid, gid)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}

func CreateJoinReq(uid int, gid int) error {
	x, _ := GetSpecificRequest(uid, gid)
	if x {
		return nil
	}
	insertSQL := `INSERT INTO grouprequests(sid, gid) VALUES (?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(uid, gid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteJoinRequest(sid int, gid int) error {
	_, err := DB.Exec("DELETE FROM grouprequests WHERE sid = ? AND gid = ?", sid, gid)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}

func GetJoinRequests(gid int) ([]Req, error) {
	var req Req
	var reqs []Req
	rows, err := DB.Query("SELECT * FROM grouprequests WHERE gid = ?", gid)
	if err != nil {
		return reqs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&req.Sid, &req.Gid); err {
		case nil:
			reqs = append(reqs, req)
		default:
			return reqs, err
		}
	}
	return reqs, nil
}
func GetJoinRequest(uid int, gid int) (bool, error) {
	rows, err := DB.Query("SELECT * FROM grouprequests WHERE gid = ? AND sid = ?", gid, uid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return true, nil
}

func GetInvited(gid int) ([]User, error) {
	var user User
	var ugs []User
	rows, err := DB.Query("SELECT Users.UId, UFirst, ULast, UEmail, UAge, UGender, UName, UTime, UPic, UNick, UText FROM users WHERE Users.UId IN ( SELECT rid FROM groupinvs WHERE gid = ?)", gid)
	if err != nil {
		log.Println(err)
		return ugs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UAge, &user.UGender, &user.UEmail, &user.UName, &user.UTime, &user.UPic, &user.UNick, &user.UText); err {
		case nil:
			ugs = append(ugs, user)
		default:
			return ugs, err
		}
	}
	return ugs, nil
}

func GetNotInvited(gid int) ([]User, error) {
	var user User
	var ugs []User
	rows, err := DB.Query("SELECT Users.UId, UFirst, ULast, UEmail, UAge, UGender, UName, UTime, UPic, UNick, UText FROM users WHERE Users.UId NOT IN ( SELECT rid FROM groupinvs WHERE gid = ? AND rid) AND Users.UId NOT IN ( SELECT usergroups.uid FROM usergroups WHERE gid = ?)", gid, gid)
	if err != nil {
		log.Println(err)
		return ugs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UAge, &user.UGender, &user.UEmail, &user.UName, &user.UTime, &user.UPic, &user.UNick, &user.UText); err {
		case nil:
			ugs = append(ugs, user)
		default:
			return ugs, err
		}
	}
	return ugs, nil
}

func GetSpecificInvite(rid int, gid int) (bool, error) {
	rows, err := DB.Query("SELECT * FROM groupinvs WHERE rid = ? AND gid = ?", rid, gid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func GetSpecificRequest(sid int, gid int) (bool, error) {
	rows, err := DB.Query("SELECT * FROM grouprequests WHERE sid = ? AND gid = ?", sid, gid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
