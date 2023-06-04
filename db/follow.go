package db

import (
	"database/sql"
	"log"
)

type Follow struct {
	Sid int
	Rid int
}

func CreateFollower(follow Follow) error {
	x, _ := GetSpecificFollow(follow.Sid, follow.Rid)
	if x {
		return nil
	}
	insertSQL := `INSERT INTO followers(follower, followee) VALUES (?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec(follow.Sid, follow.Rid)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteFollower(follow Follow) error {
	_, err := DB.Exec("DELETE FROM followers WHERE follower = ? AND followee = ?", follow.Sid, follow.Rid)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}

func GetFollowers(rid int) ([]User, error) {
	var users []User
	rows, err := DB.Query("SELECT UId, UFirst, ULast, UEmail, UAge, UGender, UName, UTime, UPic, UNick, UText FROM Users WHERE Users.UId IN ( SELECT follower FROM followers WHERE followee = ?);", rid)
	DbErrHandler(false, "Followers get | query", err)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UEmail, &user.UAge, &user.UGender, &user.UName, &user.UTime, &user.UPic, &user.UNick, &user.UText); err != nil {
			log.Println(err)
			DbErrHandler(false, "Followers get | scan", err)
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetFollowed(sid int) ([]User, error) {
	var users []User
	rows, err := DB.Query("SELECT UId, UFirst, ULast, UEmail, UAge, UGender, UName, UTime, UPic, UNick, UText FROM Users WHERE Users.UId IN ( SELECT followee FROM followers WHERE follower = ?);", sid)
	DbErrHandler(false, "Followed get | query", err)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UEmail, &user.UAge, &user.UGender, &user.UName, &user.UTime, &user.UPic, &user.UNick, &user.UText); err != nil {
			log.Println(err)
			DbErrHandler(false, "Followed get | scan", err)
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateFollowRequest(follow Follow) error {
	x, _ := GetSpecificFollowRequest(follow.Sid, follow.Rid)
	if x {
		return nil
	}
	insertSQL := `INSERT INTO followrequests(follower, followee) VALUES (?, ?)`
	statement, err := DB.Prepare(insertSQL)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec(follow.Sid, follow.Rid)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetSpecificFollowRequest(sid int, rid int) (bool, error) {
	rows, err := DB.Query("SELECT * FROM followrequests WHERE follower = ? AND followee = ?", sid, rid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func GetSpecificFollow(sid int, rid int) (bool, error) {
	rows, err := DB.Query("SELECT * FROM followers WHERE follower = ? AND followee = ?", sid, rid)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func AcceptRequest(follow Follow) error {
	err := CreateFollower(follow)
	switch err {
	case nil:
		err = DeleteRequest(follow)
		switch err {
		case nil:
			return nil
		default:
			return err
		}
	default:
		return err
	}
}

func GetFollowRequestsRecived(rid int) ([]User, error) {
	var users []User
	rows, err := DB.Query("SELECT UId, UFirst, ULast, UEmail, UAge, UGender, UName, UTime, UPic, UNick, UText FROM Users WHERE Users.UId IN ( SELECT follower FROM followrequests WHERE followee = ?);", rid)
	DbErrHandler(false, "Follower GET get | query", err)
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UEmail, &user.UAge, &user.UGender, &user.UName, &user.UTime, &user.UPic, &user.UNick, &user.UText); err != nil {
			DbErrHandler(false, "Follower GET get | scan", err)
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetFollowRequestsSent(sid int) ([]User, error) {
	var users []User
	rows, err := DB.Query("SELECT UId, UFirst, ULast, UEmail, UAge, UGender, UName, UTime, UPic, UNick, UText FROM Users WHERE Users.UId IN ( SELECT follower FROM followrequests WHERE followee = ?);", sid)
	DbErrHandler(false, "Follower SENT get | query", err)
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UId, &user.UFirst, &user.ULast, &user.UEmail, &user.UAge, &user.UGender, &user.UName, &user.UTime, &user.UPic, &user.UNick, &user.UText); err != nil {
			DbErrHandler(false, "Followers SENT get | scan", err)
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}

func DeleteRequest(follow Follow) error {
	_, err := DB.Exec("DELETE FROM followrequests WHERE follower = ? AND followee = ?", follow.Sid, follow.Rid)
	switch err {
	case nil:
		return nil
	default:
		return err
	}
}

func CheckFollowStatus(sid int, rid int) (bool, error) {
	log.Println(sid, rid)
	row, err := DB.Query("SELECT * FROM followers WHERE follower = ? AND followee = ?", sid, rid)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	DbErrHandler(false, "Follower SENT get | query", err)
	defer row.Close()
	for row.Next() {
		log.Println(row)
		return true, nil
	}
	return false, nil
}
