-- +migrate Up
CREATE TABLE IF NOT EXISTS "events" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"creator" INTEGER,
		"title" text,
		"description" text,
		"time" DATE,
		"gid" INTEGER
	  );
-- +migrate Down
DROP TABLE "events";
