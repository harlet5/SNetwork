-- +migrate Up
CREATE TABLE IF NOT EXISTS "Users" (
	  "UId"				INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"UFirst"			TEXT NOT NULL,
		"ULast"			TEXT NOT NULL,
		"UAge"			TEXT NOT NULL,
		"UGender"			TEXT NOT NULL,
		"UEmail"			TEXT NOT NULL,
		"UName"			TEXT NOT NULL,
		"UPass"			TEXT NOT NULL,
		"UTime"			TEXT,
		"UPic"		TEXT,
		"UNick"			TEXT,
		"UText"			TEXT,
		"UPriv"			TEXT	
	);
-- +migrate Down
DROP TABLE "Users";
