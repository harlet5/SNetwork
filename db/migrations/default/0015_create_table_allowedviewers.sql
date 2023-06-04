-- +migrate Up
CREATE TABLE IF NOT EXISTS "allowedviewers" (
		"pid" INTEGER REFERENCES Threads(TId),
		"uid" INTEGER REFERENCES Users(UId)
	  );

-- +migrate Down
DROP TABLE "allowedviewrs";
