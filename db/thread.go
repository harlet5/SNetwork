package db

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type Thread struct {
	TId    int
	TUId   int
	TUName string
	TProf  string
	TBody  string
	TCats  []Cat
	TTime  string
	TGid   int
	TPics  []string
	TPriv  string
	TWho   []string
}

func GetThread(id string) Thread {
	var thread Thread
	rows, err := DB.Query("SELECT * FROM Threads WHERE TId = " + id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Thread get | query", err)
		} else {
			DbErrHandler(true, "Thread get | query", err)
		}
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&thread.TId, &thread.TBody, &thread.TUId, &thread.TTime, &thread.TGid, &thread.TPriv); err != nil {
			DbErrHandler(true, "Thread get | scan", err)
		}
		u, err := GetUser(thread.TUId)
		if err != nil {
			log.Println(err)
			return thread
		}
		thread.TUName = u.UName
		thread.TProf = u.UPic
		thread.TCats, err = GetCatByThread(id)
		thread.TPics, err = GetThreadPics(thread.TId)
		thread.TWho, _ = GetThreadUsers(thread.TUId)
	}
	return thread
}

func GetLastThread() (int, error) {
	rows, err := DB.Query("SELECT TId FROM Threads WHERE TGId = ? ORDER BY TId DESC LIMIT 1", -1)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Thread newest | query", err)
		}
		return 0, err
	}
	var tid int
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&tid); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				DbErrHandler(false, "Thread newest | scan", err)
			}
			return 0, err
		}
	}
	return tid, nil
}

func GetThreads(uname string) ([]Thread, error) {
	var threads []Thread
	u, _ := GetUserByUname(uname)
	rows, err := DB.Query("SELECT * FROM Threads WHERE TGId = ? AND TUId = ? OR TPriv = 'Public' OR (TPriv = 'Private' AND ? IN (SELECT follower FROM followers WHERE followee = TUId)) OR (TPriv = 'Almost private' AND TId IN (SELECT tid FROM threadviewers))", -1, u.UId, u.UId)
	if err != nil {
		log.Println(err)
		if !errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Threads get | query", err)
		}
		return threads, err
	}
	defer rows.Close()
	for rows.Next() {
		var thread Thread
		if err := rows.Scan(&thread.TId, &thread.TBody, &thread.TUId, &thread.TTime, &thread.TGid, &thread.TPriv); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				DbErrHandler(false, "Threads get | scan", err)
			}
			log.Println(err)
			return threads, err
		}
		topic_id := strconv.Itoa(thread.TId)
		u, err := GetUser(thread.TUId)
		if err != nil {
			log.Println(err)
			return threads, err
		}
		thread.TUName = u.UName
		thread.TProf = u.UPic
		thread.TCats, err = GetCatByThread(topic_id)
		if err != nil {
			log.Println(err)
			return threads, err
		}
		thread.TPics, err = GetThreadPics(thread.TId)
		if err != nil {
			log.Println(err)
			return threads, err
		}
		thread.TWho, err = GetThreadUsers(thread.TUId)
		if err != nil {
			log.Println(err)
			return threads, err
		}
		threads = append(threads, thread)
	}
	return threads, nil
}

func CreateThread(t Thread) error {
	query := "INSERT INTO Threads (TBody, TUId, TTime, TGId, TPriv) VALUES (?,?,?,?,?) RETURNING TId"
	if err := DB.QueryRow(query, t.TBody, t.TUId, t.TTime, t.TGid, t.TPriv).Scan(&t.TId); err != nil {
		log.Println(err)
		DbErrHandler(false, "Thread create | query", err)
		return err
	}
	return nil
}

func GetThreadsByCat(cid int) ([]Thread, error) {
	var threads []Thread
	rows, err := DB.Query(`SELECT TId, TBody, TUId, TTime, TGId, TPriv WHERE TGId = ? AND (TId IN (SELECT tid FROM threadviewers WHERE name = ?) OR TId NOT IN (SELECT tid FROM threadviewers)) FROM Threads INNER JOIN ThreadsInCats WHERE TCcId = ? AND TCtId = TId`, -1, cid)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Threads cat | query", err)
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var thread Thread
		if err := rows.Scan(&thread.TId, &thread.TBody, &thread.TUId, &thread.TTime, &thread.TGid, &thread.TPriv); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				DbErrHandler(false, "Threads cat | scan", err)
			}
			return nil, err
		}
		tid := strconv.Itoa(thread.TId)
		u, err := GetUser(thread.TUId)
		if err != nil {
			log.Println(err)
			return threads, err
		}
		thread.TUName = u.UName
		thread.TProf = u.UPic
		thread.TCats, err = GetCatByThread(tid)
		thread.TPics, err = GetThreadPics(thread.TId)
		thread.TWho, _ = GetThreadUsers(thread.TUId)
		if err != nil {
			return threads, err
		}
		threads = append(threads, thread)
	}
	return threads, nil
}

func GetThreadsByUser(id int) ([]Thread, error) {
	var threads []Thread
	rows, err := DB.Query("SELECT * FROM Threads where TUId =?", id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Threads user | query", err)
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var thread Thread
		if err := rows.Scan(&thread.TId, &thread.TBody, &thread.TUId, &thread.TTime, &thread.TGid, &thread.TPriv); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				DbErrHandler(false, "Threads user | scan", err)
			}
			return nil, err
		}
		tid := strconv.Itoa(thread.TId)
		u, err := GetUser(thread.TUId)
		if err != nil {
			log.Println(err)
			return threads, err
		}
		thread.TUName = u.UName
		thread.TProf = u.UPic
		thread.TCats, err = GetCatByThread(tid)
		thread.TPics, err = GetThreadPics(thread.TId)
		thread.TWho, _ = GetThreadUsers(thread.TUId)
		threads = append(threads, thread)
	}
	return threads, nil
}

func GetThreadsInGroup(gid int) ([]Thread, error) {
	var threads []Thread
	var thread Thread
	rows, err := DB.Query("SELECT * FROM Threads WHERE TGId = ?", gid)
	if err != nil {
		return threads, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&thread.TId, &thread.TBody, &thread.TUId, &thread.TTime, &thread.TGid, &thread.TPriv); err {
		case nil:
			tid := strconv.Itoa(thread.TId)
			u, err := GetUser(thread.TUId)
			if err != nil {
				log.Println(err)
				return threads, err
			}
			thread.TUName = u.UName
			thread.TProf = u.UPic
			thread.TCats, err = GetCatByThread(tid)
			thread.TPics, err = GetThreadPics(thread.TId)
			thread.TWho, _ = GetThreadUsers(thread.TUId)
			threads = append(threads, thread)
		default:
			return threads, err
		}
	}
	return threads, nil
}

func GetLastGroupThread(gid int) (int, error) {
	rows, err := DB.Query("SELECT TId FROM Threads WHERE TGId = ? ORDER BY TId DESC LIMIT 1", gid)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			DbErrHandler(false, "Thread newest | query", err)
		}
		return 0, err
	}
	var tid int
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&tid); err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				DbErrHandler(false, "Thread newest | scan", err)
			}
			return 0, err
		}
	}
	return tid, nil
}

func CreateThreadPicConnection(tid int, name string) error {
	insertSQL := `INSERT INTO threadpics(tid, name) VALUES (?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(tid, name)
	if err != nil {
		return err
	}
	return nil
}

func GetThreadPics(tid int) ([]string, error) {
	var pics []string
	var pic string
	rows, err := DB.Query("SELECT name FROM threadpics WHERE tid = ?", tid)
	if err != nil {
		return pics, err
	}
	defer rows.Close()
	r := 1
	for rows.Next() {
		switch err = rows.Scan(&pic); err {
		case nil:
			log.Println(r)
			r++
			pics = append(pics, pic)
		default:
			return pics, err
		}
	}
	return pics, nil
}

func CreateThreadUserConnection(tid int, name string) error {
	insertSQL := `INSERT INTO threadviewers(tid, name) VALUES (?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(tid, name)
	if err != nil {
		return err
	}
	return nil
}

func GetThreadUsers(tid int) ([]string, error) {
	var pics []string
	var pic string
	rows, err := DB.Query("SELECT name FROM threadviewers WHERE tid = ?", tid)
	if err != nil {
		return pics, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&pic); err {
		case nil:
			pics = append(pics, pic)
		default:
			return pics, err
		}
	}
	return pics, nil
}
