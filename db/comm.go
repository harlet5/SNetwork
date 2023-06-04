package db

import "log"

type Comm struct {
	CoId    int
	CoTId   int
	CoUId   int
	CoUName string
	CoBody  string
	CoTime  string
	CoPics  []string
}

func CreateComm(tid int, uid int, body string, time string) (int, error) {
	q := "INSERT INTO Comms(cTId, cUId, cBody, cTime) VALUES (?,?,?,?) RETURNING cId"
	cid := -1
	err := DB.QueryRow(q, tid, uid, body, time).Scan(&cid)
	DbErrHandler(false, "Chat create | execute", err)
	return cid, err
}

func GetThreadComms(id string) []Comm {
	var comms []Comm
	rows, err := DB.Query("SELECT * FROM Comms WHERE cTId = " + id)
	DbErrHandler(false, "Comm get | query", err)
	defer rows.Close()
	for rows.Next() {
		var comm Comm
		if err := rows.Scan(&comm.CoId, &comm.CoTId, &comm.CoUId, &comm.CoBody, &comm.CoTime); err != nil {
			DbErrHandler(false, "Comm get | scan", err)
		}
		comm.CoUName, _ = GetUname(comm.CoUId)
		comm.CoPics, _ = GetCommPics(comm.CoId)
		comms = append(comms, comm)
	}
	return comms
}

func CreateCommPicConnection(cid int, name string) error {
	insertSQL := `INSERT INTO commpics(cid, name) VALUES (?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		return err
	}
	_, err = statement.Exec(cid, name)
	if err != nil {
		return err
	}
	return nil
}

func GetCommPics(cid int) ([]string, error) {
	var pics []string
	var pic string
	rows, err := DB.Query("SELECT name FROM commpics WHERE cid = ?", cid)
	if err != nil {
		log.Println(err)
		return pics, err
	}
	defer rows.Close()
	for rows.Next() {
		switch err = rows.Scan(&pic); err {
		case nil:
			pics = append(pics, pic)
		default:
			log.Println(err)
			return pics, err
		}
	}
	return pics, nil
}
