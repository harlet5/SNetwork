-- +migrate Up
CREATE TABLE IF NOT EXISTS "threadviewers" (
		"tid" INTEGER REFERENCES Threads(TId),
		"name" TEXT
	  );

-- +migrate Down
DROP TABLE "threadviewers";
