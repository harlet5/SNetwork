-- +migrate Up
CREATE TABLE IF NOT EXISTS "notifications" (
		"nid" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"type" TEXT,
		"sender" INTEGER,
		"reciver" INTEGER,
		"status" TEXT,
		"gid" INTEGER,
		"gname"	TEXT,
		"ename" TEXT
	  );

-- +migrate Down
drop table "notifications";
