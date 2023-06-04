-- +migrate Up
CREATE TABLE IF NOT EXISTS "ThreadsInCats" (
		"TCtId"     INTEGER NOT NULL REFERENCES Threads("TId"),
		"TCcId"     INTEGER NOT NULL REFERENCES Cats("CId")
);
-- +migrate Down
DROP TABLE "ThreadsInCats";
