-- +migrate Up
CREATE TABLE IF NOT EXISTS "usergroups" (
		"uid" INTEGER REFERENCES Users(UId),
		"gid" INTEGER REFERENCES groups(id)
	  );
-- +migrate Down
DROP TABLE "usergroups";
