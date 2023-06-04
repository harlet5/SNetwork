package db

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func DbUtil() {
	db, err := sql.Open("sqlite3", "./db/forum.db")
	if err != nil {
		panic(err.Error())
	}
	DB = db
}

func Init(db *sql.DB, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	requests := strings.SplitAfter(string(file), ";")
	for _, request := range requests {
		db.Exec(request)
	}
}

func CreateTables() {
	Tables := []string{`CREATE TABLE IF NOT EXISTS Cats (
		CId				INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		CName			TEXT NOT NULL
	);`,
		`CREATE TABLE IF NOT EXISTS ThreadsInCats (
		TCtId			INTEGER NOT NULL REFERENCES Threads(TId),
		TCcId			INTEGER NOT NULL REFERENCES Cats(CId)
	);`,
		`CREATE TABLE IF NOT EXISTS Threads (
		TId				INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		TBody			TEXT NOT NULL,
		TUId			INTEGER NOT NULL REFERENCES Users(UId),
		TTime			TEXT,
		TGId			INTEGER,
		TPriv			TEXT
	);`,
		`CREATE TABLE IF NOT EXISTS Comms (
		cId				INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		cTId			INTEGER NOT NULL REFERENCES Threads(TId),
		cUId			INTEGER NOT NULL REFERENCES Users(UId),
		cBody			TEXT NOT NULL,
		cTime			TEXT
	);`,
		`CREATE TABLE IF NOT EXISTS Users (
		UId				INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		UFirst			TEXT NOT NULL,
		ULast			TEXT NOT NULL,
		UAge			TEXT NOT NULL,
		UGender			TEXT NOT NULL,
		UEmail			TEXT NOT NULL,
		UName			TEXT NOT NULL,
		UPass			TEXT NOT NULL,
		UTime			TEXT,
		UPic			TEXT,
		UNick			TEXT,
		UText			TEXT,
		UPriv			TEXT
	);`,
		`CREATE TABLE IF NOT EXISTS Sess (
		SId 			TEXT NOT NULL PRIMARY KEY,
		SUId  			INTEGER NOT NULL REFERENCES Users(UId) ON DELETE CASCADE UNIQUE ON CONFLICT REPLACE,
		STime	 		TIMESTAMP DEFAULT (strftime('%s', 'now'))
	);`,
		`CREATE TABLE IF NOT EXISTS Chats (
		ChId			INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		ChUIdSender		INTEGER NOT NULL,
		ChUIdReciver	INTEGER NOT NULL,
		ChBody 			TEXT NOT NULL,
		ChTime			TEXT NOT NULL,
		ChStatus   		TEXT
	);`,
		`CREATE TABLE IF NOT EXISTS "groups" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"creator" INTEGER,
		"name" text,
		"description" text
	  );`,
		`CREATE TABLE IF NOT EXISTS "usergroups" (
		"uid" INTEGER REFERENCES Users(UId),
		"gid" INTEGER REFERENCES groups(id)
	  );`,
		`CREATE TABLE IF NOT EXISTS "events" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"creator" INTEGER,
		"title" text,
		"description" text,
		"time" DATE,
		"gid" INTEGER
	  );`,
		`CREATE TABLE IF NOT EXISTS "userevents" (
		"uid" INTEGER REFERENCES users(UId),
		"eid" INTEGER REFERENCES events(id) ON DELETE CASCADE,
		"status" text
	  );`,
		`CREATE TABLE IF NOT EXISTS "groupinvs" (
		"sid" INTEGER REFERENCES Users(UId),
		"rid" INTEGER REFERENCES Users(UId),
		"gid" INTEGER REFERENCES groups(id)
	  );`,
		`CREATE TABLE IF NOT EXISTS "followers" (
		"follower" INTEGER REFERENCES Users(UId),
		"followee" INTEGER REFERENCES Users(UId)
	  );`,
		`CREATE TABLE IF NOT EXISTS "followrequests" (
		"follower" INTEGER REFERENCES Users(UId),
		"followee" INTEGER REFERENCES Users(UId)
	  );`,
		`CREATE TABLE IF NOT EXISTS "allowedviewers" (
		"pid" INTEGER REFERENCES Threads(TId),
		"uid" INTEGER REFERENCES Users(UId)
	  );`,
		`CREATE TABLE IF NOT EXISTS "grouprequests" (
		"sid" INTEGER REFERENCES Users(UId),
		"gid" INTEGER REFERENCES groups(id)
	  );`,
		`CREATE TABLE IF NOT EXISTS "threadpics" (
		"tid" INTEGER REFERENCES Threads(TId),
		"name" TEXT
	  );`,
		`CREATE TABLE IF NOT EXISTS "threadviewers" (
		"tid" INTEGER REFERENCES Threads(TId),
		"name" TEXT
	  );`,
		`CREATE TABLE IF NOT EXISTS "commpics" (
		"cid" INTEGER REFERENCES Comms(cId),
		"name" TEXT
	  );`,
		`CREATE TABLE IF NOT EXISTS "notifications" (
		"nid" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"type" TEXT,
		"sender" INTEGER,
		"reciver" INTEGER,
		"status" TEXT,
		"gid" INTEGER,
		"gname"	TEXT,
		"ename" TEXT
	  );`,
		`CREATE TABLE IF NOT EXISTS "groupchats" (
		"id"			INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"uid"			INTEGER NOT NULL,
		"gid"			INTEGER NOT NULL,
		"body" 			TEXT NOT NULL,
		"time"			TEXT NOT NULL
	);`,
	}
	for _, table := range Tables {
		statement, err := DB.Prepare(table)
		if err != nil {
			log.Fatal(err.Error())
		}
		_, err = statement.Exec()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func FillTables() {
	//CreateTables()
	CreateCat("teema#1")
	CreateCat("teema#2")
	CreateCat("teema#3")
	CreateUser(User{0, "Dummy", "Uno", "264", "Male", "dumb@gmail.com", "Uno", "52A7Jn0Go", "10.03.2022", "default.png", "", "", "false"})
	CreateUser(User{0, "Dummy", "Dos", "19", "Male", "dumb2@gmail.com", "Dos", "9uEo8Zq0oOq6", "04.05.1988", "default.png", "", "", "false"})
	CreateUser(User{0, "Dummy", "Tres", "34", "Male", "dumb3@gmail.com", "Tres", "KwB2njMld", "25.11.2005", "default.png", "", "", "false"})
	epass, _ := bcrypt.GenerateFromPassword([]byte("name"), bcrypt.DefaultCost)
	user := User{UId: 0, UFirst: "name", ULast: "name", UAge: "20", UGender: "Male", UEmail: "name@gmail.com", UName: "name", UPass: string(epass), UTime: "10.03.2024", UPic: "default.png", UNick: "name", UText: "name", UPriv: "false"}
	CreateUser(user)
	epass, _ = bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
	user = User{UId: 0, UFirst: "first", ULast: "last", UAge: "20", UGender: "Female", UEmail: "123@gmail.com", UName: "123", UPass: string(epass), UTime: "10.03.2024", UPic: "default.png", UNick: "123", UText: "123", UPriv: "false"}
	CreateUser(user)
	CreateThread(Thread{0, 1, "Uno", "default.png", "Hello there", []Cat{}, "18.9.2021", -1, []string{}, "Public", []string{}})
	CreateThread(Thread{0, 2, "Dos", "default.png", "New phone", []Cat{}, "18.9.2021", -1, []string{}, "Public", []string{}})
	CreateThread(Thread{0, 3, "Tres", "default.png", "Du", []Cat{}, "18.9.2021", -1, []string{}, "Public", []string{}})
	CreateComm(1, 1, "ez", "16.04.2022")
	CreateComm(2, 2, "clap", "17.03.2022")
	CreateComm(3, 3, "why", "18.02.2022")
	CreateComm(1, 3, "tho", "19.01.2022")
	LinkThreadToCats(1, 1)
	LinkThreadToCats(1, 2)
	LinkThreadToCats(2, 1)
	LinkThreadToCats(3, 2)
	CreateChat(1, 2, "hi", "22.12.2022")
	CreateChat(1, 3, "hi", "22.12.2022")
	CreateChat(3, 2, "hi", "22.12.2022")
	CreateChat(2, 1, "why", "22.12.2022")
	CreateChat(2, 1, "tho", "22.12.2022")
}

func DbErrHandler(kind bool, why string, err error) {
	if err != nil && kind == true {
		log.Println(why)
		log.Fatal(err)
	} else if err != nil && kind == false {
		log.Println(why)
	}
}
