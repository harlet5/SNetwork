package db

import (
	"database/sql"
	"log"
)

type Event struct {
	Id             int
	Creator        int
	UName          string
	Name           string
	Text           string
	Time           string
	Gid            int
	Yes            []User
	No             []User
	YesCount       int
	NoCount        int
	UndecidedCount int
}

type UserEvent struct {
	Uid    int
	Eid    int
	Status string
}

func CreateEvent(event Event) error {
	insertSQL := `INSERT INTO events(creator, title, description, time, gid) VALUES (?, ?, ?, ?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(event.Creator, event.Name, event.Text, event.Time, event.Gid)
	if err != nil {
		return err
	}
	return nil
}

func GetEventById(eid int) (Event, error) {
	var event Event
	row := DB.QueryRow("SELECT * FROM events WHERE id = ?", eid)
	switch err := row.Scan(&event.Id, &event.Creator, &event.Name, &event.Text, &event.Time, &event.Gid); err {
	case nil:
		y, _ := GetEventUsers(eid, "yes")
		n, _ := GetEventUsers(eid, "no")
		event.Yes = y
		event.No = n
		event.UName, _ = GetUname(event.Creator)
		return event, nil
	case sql.ErrNoRows:
		return event, err
	default:
		return event, err
	}
}

func GetEventsByGroup(gid int) ([]Event, error) {
	var event Event
	var events []Event
	rows, err := DB.Query("SELECT * FROM events WHERE gid = ?", gid)
	tmp, _ := GetGroupUsers(gid)
	if err != nil {
		log.Println(err)
		log.Println("why1")
		return events, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&event.Id, &event.Creator, &event.Name, &event.Text, &event.Time, &event.Gid); err {
		case nil:
			y, _ := GetEventUsers(event.Id, "yes")
			n, _ := GetEventUsers(event.Id, "no")
			event.YesCount = len(y)
			event.NoCount = len(n)
			event.UndecidedCount = len(tmp) - len(y) - len(n)
			event.Yes = y
			event.No = n
			event.UName, _ = GetUname(event.Creator)
			events = append(events, event)
		default:
			log.Println(err)
			log.Println("why2")
			return events, err
		}
	}
	return events, nil
}

func GetEventByNameInGroup(name string, gid int) (Event, error) {
	var event Event
	row := DB.QueryRow("SELECT * FROM events WHERE title = ? AND gid = ?", name, gid)
	switch err := row.Scan(&event.Id, &event.Creator, &event.Name, &event.Text, &event.Time, &event.Gid); err {
	case nil:
		y, _ := GetEventUsers(event.Id, "yes")
		n, _ := GetEventUsers(event.Id, "no")
		event.Yes = y
		event.No = n
		event.UName, _ = GetUname(event.Creator)
		return event, nil
	case sql.ErrNoRows:
		return event, err
	default:
		return event, err
	}
}

func AddUserEventConnection(uid int, eid int, status string) error {
	insertSQL := `INSERT INTO userevents(uid, eid, status) VALUES (?, ?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(uid, eid, status)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserEventConnection(uid int, eid int, status string) error {
	sql := `UPDATE userevents SET status = ? WHERE uid = ? AND eid = ?`
	stmt, err := DB.Prepare(sql)
	if err != nil {
		log.Println("sql", err)
		return err
	}
	_, err = stmt.Exec(status, uid, eid)
	if err != nil {
		log.Println("sql", err)
		return err
	}
	return nil
}

func DeleteEvent(eid int) error {
	insertSQL := `DELETE * FROM events WHERE eid = ?`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(eid)
	if err != nil {
		return err
	}
	return nil
}

func GetEventUsers(eid int, status string) ([]User, error) {
	var user User
	var users []User
	rows, err := DB.Query(`SELECT Users.UId, UFirst, ULast, UEmail, UAge, UGender, UName, UTime, UPic, UNick, UText FROM Users INNER JOIN userevents WHERE eid = ? AND userevents.uid = Users.UId AND status = ?`, eid, status)
	if err != nil {
		log.Println(err)
		log.Println("wh31")
		return users, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UAge, &user.UGender, &user.UEmail, &user.UName, &user.UTime, &user.UPic, &user.UNick, &user.UText); err {
		case nil:
			users = append(users, user)
		default:
			log.Println(err)
			return users, err
		}
	}
	return users, nil
}

func GetEventUserById(uid int, eid int) (bool, error) {
	rows, err := DB.Query("SELECT * FROM userevents WHERE uid = ? AND eid = ?", uid, eid)
	if err != nil {
		log.Println(err)
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		log.Println("here")
		return true, nil
	}
	return false, nil
}

func GetUserEvents(uid int) ([]UserEvent, error) {
	var ue UserEvent
	var ues []UserEvent
	rows, err := DB.Query("SELECT * FROM userevents WHERE uid = ? ", uid)
	if err != nil {
		log.Println(err)
		log.Println("wh32")
		return ues, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&ue.Uid, &ue.Eid, &ue.Status); err {
		case nil:
			ues = append(ues, ue)
		default:
			log.Println(err)
			return ues, err
		}
	}
	return ues, nil
}
