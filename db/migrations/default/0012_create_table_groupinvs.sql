-- +migrate Up
CREATE TABLE IF NOT EXISTS "groupinvs" (
		"sid" INTEGER REFERENCES Users(UId),
		"rid" INTEGER REFERENCES Users(UId),
		"gid" INTEGER REFERENCES groups(id)
	  );
-- +migrate Down
DROP TABLE "groupinvs";
