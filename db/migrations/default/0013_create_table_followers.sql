-- +migrate Up
CREATE TABLE IF NOT EXISTS "followers" (
		"follower" INTEGER REFERENCES Users(UId),
		"followee" INTEGER REFERENCES Users(UId)
	  );
-- +migrate Down
DROP TABLE "followers";
