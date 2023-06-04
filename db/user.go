package db

import "log"

type User struct {
	UId     int
	UFirst  string
	ULast   string
	UAge    string
	UGender string
	UEmail  string
	UName   string
	UPass   string
	UTime   string
	UPic    string
	UNick   string
	UText   string
	UPriv   string
}

type Other struct {
	OId     int
	OName   string
	OActive bool
	Unread  int
	OProf   string
}

func GetUsers(id int) []Other {
	var users []Other
	rows, err := DB.Query("SELECT UId, UName FROM Users;")
	if err != nil {
		DbErrHandler(false, "Users get | query", err)
	}
	for rows.Next() {
		var user Other
		if err := rows.Scan(&user.OId, &user.OName); err != nil {
			DbErrHandler(false, "Users get | scan", err)
		}
		user.Unread, _ = GetUnseenNrPerChat(id, user.OId)
		user.OProf, _ = GetProfPic(user.OId)
		user.OActive = GetActive(user.OId)
		users = append(users, user)
	}
	return users
}

func GetUserByUname(uname string) (User, error) {
	user := User{}
	rows, err := DB.Query(`SELECT * FROM Users where UName = ?`, uname)
	if err != nil {
		DbErrHandler(false, "User uname | query", err)
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UAge, &user.UGender, &user.UEmail, &user.UName, &user.UPass, &user.UTime, &user.UPic, &user.UNick, &user.UText, &user.UPriv); err != nil {
			DbErrHandler(false, "User uname | scan", err)
			return user, err
		}
	}
	return user, nil
}

func GetUser(uid int) (User, error) {
	user := User{}
	log.Println(uid)
	rows, err := DB.Query(`SELECT * FROM Users where UId = ?`, uid)
	if err != nil {
		DbErrHandler(false, "User get | query", err)
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UAge, &user.UGender, &user.UEmail, &user.UName, &user.UPass, &user.UTime, &user.UPic, &user.UNick, &user.UText, &user.UPriv); err != nil {
			DbErrHandler(false, "User get | scan", err)
			return user, err
		}
	}
	return user, nil
}

func GetPass(uname string) (string, error) {
	rows, err := DB.Query(`SELECT UPass FROM Users WHERE UName = ?`, uname)
	if err != nil {
		DbErrHandler(false, "User pass | query", err)
		return "", err
	}
	var pass string
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&pass); err != nil {
			DbErrHandler(false, "User pass | query", err)
		}
	}
	return pass, nil
}

func GetPassByemail(uemail string) (string, error) {
	rows, err := DB.Query(`SELECT UPass FROM Users WHERE UEmail = ?`, uemail)
	if err != nil {
		DbErrHandler(false, "User pass | query", err)
		return "", err
	}
	var pass string
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&pass); err != nil {
			DbErrHandler(false, "User pass | query", err)
		}
	}
	return pass, nil
}

func GetUId(uname string) (int, error) {
	rows, err := DB.Query(`SELECT UId FROM Users WHERE UName = ?`, uname)
	if err != nil {
		DbErrHandler(false, "User uid | query", err)
		return 0, err
	}
	var id int
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			DbErrHandler(false, "User uid | scan", err)
		}
	}
	return id, nil
}

func GetUIdByEmail(email string) (int, error) {
	rows, err := DB.Query(`SELECT UId FROM Users WHERE UEmail = ?`, email)
	if err != nil {
		DbErrHandler(false, "User uid | query", err)
		return 0, err
	}
	var id int
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			DbErrHandler(false, "User uid | scan", err)
		}
	}
	return id, nil
}

func GetUname(uid int) (string, error) {
	rows, err := DB.Query(`SELECT UName FROM Users WHERE UId = ?`, uid)
	if err != nil {
		DbErrHandler(false, "User uname | query", err)
		return "", err
	}
	var name string
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			DbErrHandler(false, "User uname | query", err)
		}
	}
	return name, nil
}

func CreateUser(user User) error {
	stm, err := DB.Prepare(`INSERT INTO Users (UFirst, ULast, UEmail, UAge, UGender, UName, UPass, UTime, UPic, UNick, UText, UPriv) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`)
	log.Println(err)
	if err != nil {
		DbErrHandler(false, "User create | prepare", err)
		return err
	}
	_, err = stm.Exec(user.UFirst, user.ULast, user.UEmail, user.UAge, user.UGender, user.UName, user.UPass, user.UTime, user.UPic, user.UNick, user.UText, user.UPriv)
	if err != nil {
		DbErrHandler(false, "User create | execute", err)
	}
	return nil
}

func UpdProfPic(pic string, id int) error {
	_, err := DB.Exec("UPDATE Users SET UPic = ? where UId = ?", pic, id)
	if err != nil {
		DbErrHandler(false, "IMAGE", err)
		return err
	}
	return nil
}

func UpdPrivacy(id int, priv string) error {
	_, err := DB.Exec("UPDATE Users SET UPriv = ? where UId = ?", priv, id)
	if err != nil {
		DbErrHandler(false, "PRIV UPDATE", err)
		return err
	}
	return nil
}

func GetProfPic(uid int) (string, error) {
	rows, err := DB.Query(`SELECT UPic FROM Users WHERE UId = ?`, uid)
	if err != nil {
		DbErrHandler(false, "User pic | query", err)
		return "default.png", err
	}
	var pic string
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&pic); err != nil {
			DbErrHandler(false, "User pic | scan", err)
		}
	}
	return pic, nil
}
