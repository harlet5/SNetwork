-- +migrate Up
CREATE TABLE IF NOT EXISTS "Threads" (	
		"TId"				INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"TBody"			TEXT NOT NULL,
		"TUId"			INTEGER NOT NULL REFERENCES Users(UId),
		"TTime"			TEXT,
		"TGId"			INTEGER,
		"TPriv"			TEXT
);
-- +migrate Down
DROP TABLE "Threads";
