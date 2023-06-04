-- +migrate Up
CREATE TABLE IF NOT EXISTS "grouprequests" (
		"sid" INTEGER REFERENCES Users(UId),
		"gid" INTEGER REFERENCES groups(id)
	  );

-- +migrate Down
DROP TABLE "grouprequests";
