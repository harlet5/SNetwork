-- +migrate Up
CREATE TABLE IF NOT EXISTS "userevents" (
		"uid" INTEGER REFERENCES users(UId),
		"eid" INTEGER REFERENCES events(id) ON DELETE CASCADE,
		"status" text
	  );
-- +migrate Down
DROP TABLE "userevents";
