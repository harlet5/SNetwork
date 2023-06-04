package db

import (
	"database/sql"
	"errors"
	"time"
)

func GetActive(uid int) bool {
	rows, err := DB.Query(`SELECT SUId, STime FROM Sess WHERE SUId = ?`, uid)
	if !errors.Is(err, sql.ErrNoRows) {
		DbErrHandler(false, "Sess actives | query", err)
	}
	var id int
	var end int64
	for rows.Next() {
		if err := rows.Scan(&id, &end); err != nil {
			DbErrHandler(false, "Sess actives | scan", err)
		}
	}
	if id != 0 {
		if time.Now().Unix()-900 < end {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func DelSess(user_id int) { // add sid
	sqlStatement := `DELETE FROM Sess WHERE SUId = $1`
	_, err := DB.Exec(sqlStatement, user_id)
	DbErrHandler(false, "Sess del | execute", err)
}

func GetSess(cookie string) (int, error) {
	rows, err := DB.Query(`SELECT SUId FROM Sess WHERE SId = ?`, cookie)
	if err != nil {
		DbErrHandler(false, "Sess get | query", err)
		return 0, err
	}
	var id int
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			DbErrHandler(false, "Sess get | scan", err)
		}
	}
	return id, nil
}

func CreateSess(token string, userId int, created int64) error {
	stm, err := DB.Prepare(`INSERT INTO Sess (SId, SUId, STime) values ($1, $2, strftime('%s', 'now'))`)
	if err != nil {
		DbErrHandler(false, "Sess create | prepare", err)
		return err
	}
	_, err = stm.Exec(token, userId)
	DbErrHandler(false, "Sess create | execute", err)
	return nil
}

func UpdateSession(token string, userid int) error {
	stm, err := DB.Prepare(`UPDATE Sess SET STime = (strftime('%s', 'now')) WHERE SId = ? AND SUId = ?`)
	if err != nil {
		DbErrHandler(false, "Sess create | prepare", err)
		return err
	}
	_, err = stm.Exec(token, userid)
	DbErrHandler(false, "Sess create | execute", err)
	return nil
}

func GetSpesSess(token string, userid int) (bool, error) {
	rows, err := DB.Query(`SELECT * FROM Sess WHERE SUId = ? AND SId = ?`, userid, token)
	if err != nil {
		DbErrHandler(false, "Sess get | query", err)
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}
func DelExpired(user_id int, sid string) { // add sid
	sqlStatement := `DELETE FROM Sess WHERE SUId = ? AND SId = ? AND STime <  (strftime('%s', 'now') - 900)`
	_, err := DB.Exec(sqlStatement, user_id, sid)
	DbErrHandler(false, "Sess del | execute", err)
}

func ClearSess() {
	stm := "DROP TABLE Sess"
	_, _ = DB.Exec(stm)
	stm = `CREATE TABLE IF NOT EXISTS Sess (
		SId 			TEXT NOT NULL PRIMARY KEY,
		SUId  			INTEGER NOT NULL REFERENCES Users(UId) ON DELETE CASCADE UNIQUE ON CONFLICT REPLACE,
		STime	 		TEXT NOT NULL
	);`
	_, _ = DB.Exec(stm)
}
