-- +migrate Up
CREATE TABLE IF NOT EXISTS "groupchats" (
		"id"			INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"uid"			INTEGER NOT NULL,
		"gid"			INTEGER NOT NULL,
		"body" 			TEXT NOT NULL,
		"time"			TEXT NOT NULL
	);

-- +migrate Down
drop table "groupchats";
