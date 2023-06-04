package db

import "log"

type Notification struct {
	Sender  int
	Reciver int
	Status  string
	Ntype   string
	Sname   string
	Rname   string
	Gid     int
	Gname   string
	Ename   string
}

func CreateNotification(nfc Notification) error {
	x, _ := GetSpecificNotification(nfc)
	log.Println("hmmm")
	if x {
		log.Println("hmmm2")
		return nil
	}
	insertSQL := `INSERT INTO notifications(type, sender, reciver, status, gid, gname, ename) VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(nfc.Ntype, nfc.Sender, nfc.Reciver, nfc.Status, nfc.Gid, nfc.Gname, nfc.Ename)
	if err != nil {
		return err
	}
	return nil
}

func GetNotifications(id int) ([]Notification, error) {
	var nfcs []Notification
	var nfc Notification
	rows, err := DB.Query("SELECT type, sender, reciver, status, gid, gname, ename FROM notifications WHERE sender = ? OR reciver = ? ORDER BY nid DESC", id, id)
	ename := []string{}
	if err != nil {
		return nfcs, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&nfc.Ntype, &nfc.Sender, &nfc.Reciver, &nfc.Status, &nfc.Gid, &nfc.Gname, &nfc.Ename); err {
		case nil:
			if nfc.Ename != "" {
				if !containsevent(ename, nfc.Ename) {
					ename = append(ename, nfc.Ename)
					nfc.Sname, _ = GetUname(nfc.Sender)
					nfc.Rname, _ = GetUname(nfc.Reciver)
					nfcs = append(nfcs, nfc)
				}
			} else {
				nfc.Sname, _ = GetUname(nfc.Sender)
				nfc.Rname, _ = GetUname(nfc.Reciver)
				nfcs = append(nfcs, nfc)
			}
		default:
			return nfcs, err
		}
	}
	return nfcs, nil
}

func UpdateNfcStatus(nfc Notification) error {
	_, err := DB.Exec("UPDATE notifications SET status = ? WHERE sender = ? AND reciver = ? AND type = ? AND gid = ? AND gname = ?", nfc.Status, nfc.Sender, nfc.Reciver, nfc.Ntype, nfc.Gid, nfc.Gname)
	if err != nil {
		DbErrHandler(false, "IMAGE", err)
		return err
	}
	return nil
}

func containsevent(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetSpecificNotification(nfc Notification) (bool, error) {
	rows, err := DB.Query("SELECT * FROM notifications WHERE type = ? AND sender = ? AND reciver = ? AND status = ? AND gid = ? AND gname = ? AND ename = ?", nfc.Ntype, nfc.Sender, nfc.Reciver, nfc.Status, nfc.Gid, nfc.Gname, nfc.Ename)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
