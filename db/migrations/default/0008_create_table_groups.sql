-- +migrate Up
CREATE TABLE IF NOT EXISTS "groups" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"creator" INTEGER,
		"name" text,
		"description" text
	  );
-- +migrate Down
DROP TABLE "groups";
