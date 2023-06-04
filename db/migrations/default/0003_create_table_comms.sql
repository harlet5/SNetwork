-- +migrate Up
CREATE TABLE IF NOT EXISTS "Comms" (
		"cId"     INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"cTId"      INTEGER NOT NULL REFERENCES Threads("TId"),
		"cUId"      INTEGER NOT NULL REFERENCES Users("UId"),
		"cBody"     TEXT NOT NULL,
		"cTime"	TEXT
);
-- +migrate Down
DROP TABLE "Comms";
