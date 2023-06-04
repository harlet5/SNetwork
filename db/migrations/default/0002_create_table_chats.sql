-- +migrate Up
CREATE TABLE IF NOT EXISTS "Chats" (
		"ChId"          INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"ChUIdSender"   INTEGER NOT NULL,
		"ChUIdReciver"  INTEGER NOT NULL,
		"ChBody"        TEXT NOT NULL,
		"ChTime"        TEXT NOT NULL,
		"ChStatus"      TEXT
);
-- +migrate Down
DROP TABLE "Chats";
