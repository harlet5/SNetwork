-- +migrate Up
CREATE TABLE IF NOT EXISTS "threadpics" (
		"tid" INTEGER REFERENCES Threads(TId),
		"name" TEXT
	  );

-- +migrate Down
DROP TABLE "threadpics";
