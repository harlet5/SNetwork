package db

import (
	"database/sql"
	"log"
)

type Chat struct {
	ChId     int
	ChUId    int
	ChOId    int
	ChUName  string
	ChOName  string
	ChBody   string
	ChTime   string
	ChStatus string
	Unread   int
}

type Gchat struct {
	ChId    int
	ChUId   int
	ChGId   int
	ChUName string
	ChBody  string
	ChTime  string
}

func CreateChat(uid int, oid int, body string, time string) error {
	if uid < 1 {
		return nil
	}
	statment, err := DB.Prepare(`INSERT INTO Chats (ChUIdSender, ChUIdReciver, ChBody, ChTime, ChStatus) values ($1, $2, $3, $4, $5)`)
	if err != nil {
		DbErrHandler(false, "Chat create | prepare", err)
		return err
	}
	_, err = statment.Exec(uid, oid, body, time, "false")
	if err != nil {
		DbErrHandler(false, "Chat create | execute", err)
		return err
	}
	return nil
}

func GetChat(uid int, oid int, mnum int) ([]Chat, error) {
	rows, err := DB.Query(`SELECT * FROM Chats WHERE (ChUIdSender = ? AND ChUIdReciver = ? ) OR (ChUIdSender = ? AND ChUIdReciver = ?) ORDER BY ChId DESC LIMIT 10 OFFSET ?`, uid, oid, oid, uid, mnum*10)
	if err != nil {
		DbErrHandler(false, "Chat get | query", err)
		return nil, err
	}
	var chats []Chat
	defer rows.Close()
	for rows.Next() {
		var chat Chat
		if err := rows.Scan(&chat.ChId, &chat.ChUId, &chat.ChOId, &chat.ChBody, &chat.ChTime, &chat.ChStatus); err != nil {
			DbErrHandler(false, "Chat get | scan", err)
		}
		chat.ChUName, _ = GetUname(chat.ChUId)
		chat.ChOName, _ = GetUname(chat.ChOId)
		chats = append(chats, chat)
	}
	SetSeen(uid, oid)

	return reverse(chats), nil
}

func SortChats(id int) []Chat {
	var chats []Chat
	rows, err := DB.Query("SELECT DISTINCT ChUIdSender, ChUIdReciver FROM Chats WHERE( ChUIdSender = ? OR ChUIdReciver = ? )  ORDER BY ChId DESC", id, id)
	if err != nil {
		DbErrHandler(false, "Chat sort | query", err)
	}
	defer rows.Close()
	for rows.Next() {
		var chat Chat
		if err := rows.Scan(&chat.ChUId, &chat.ChOId); err != nil {
			DbErrHandler(false, "Chat sort | scan", err)
		}

		chats = append(chats, chat)
	}
	return chats
}

func CreateGroupChat(uid int, oid int, body string, time string) error {
	statment, err := DB.Prepare(`INSERT INTO groupchats (uid, gid, body, time) values ($1, $2, $3, $4)`)
	if err != nil {
		DbErrHandler(false, "Chat create | prepare", err)
		return err
	}
	_, err = statment.Exec(uid, oid, body, time)
	if err != nil {
		DbErrHandler(false, "Chat create | execute", err)
		return err
	}
	return nil
}

func GetGroupChat(gid int, mnum int) ([]Gchat, error) {
	rows, err := DB.Query(`SELECT * FROM groupchats WHERE gid = ? ORDER BY id DESC LIMIT 10 OFFSET ?`, gid, mnum*10)
	if err != nil {
		DbErrHandler(false, "Chat get | query", err)
		return nil, err
	}
	var chats []Gchat
	defer rows.Close()
	for rows.Next() {
		var chat Gchat
		if err := rows.Scan(&chat.ChId, &chat.ChUId, &chat.ChGId, &chat.ChBody, &chat.ChTime); err != nil {
			DbErrHandler(false, "Chat get | scan", err)
		}
		chat.ChUName, _ = GetUname(chat.ChUId)
		chats = append(chats, chat)
	}
	return revgchat(chats), nil
}

func GetUnseenNr(uid int) (int, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM Chats WHERE ChUIdReciver = ? AND ChStatus = ?", uid, "false").Scan(&count)
	log.Println(err)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func GetUnseenNrPerChat(rid int, sid int) (int, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM Chats WHERE ChUIdSender = ? AND ChUIdReciver = ? AND ChStatus = ?", sid, rid, "false").Scan(&count)
	log.Println(err)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

func SetSeen(rid int, sid int) error {
	stm, err := DB.Prepare(`UPDATE Chats SET ChStatus = ? WHERE ChUIdSender = ? AND ChUIdReciver = ? AND ChStatus = ?`)
	log.Println(err)
	if err != nil {
		DbErrHandler(false, "Chat set | prepare", err)
		return err
	}
	_, err = stm.Exec("true", sid, rid, "false")
	DbErrHandler(false, "Chat set | execute", err)
	return nil
}

func reverse(arr []Chat) []Chat {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func revgchat(arr []Gchat) []Gchat {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func DelWrong() { // add sid
	sqlStatement := `DELETE FROM Chats WHERE ChUIdSender = ? OR ChUIdReciver = ?`
	_, err := DB.Exec(sqlStatement, -1, -1)
	DbErrHandler(false, "Sess del | execute", err)
}
