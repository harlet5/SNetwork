-- +migrate Up
CREATE TABLE IF NOT EXISTS "Cats" (
		"CId"       INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"CName"     TEXT NOT NULL
);
-- +migrate Down
DROP TABLE "Cats";
