-- +migrate Up
CREATE TABLE IF NOT EXISTS "Sess" (
		"SId"       TEXT NOT NULL PRIMARY KEY,
		"SUId"      INTEGER NOT NULL REFERENCES Users("UId") ON DELETE CASCADE UNIQUE ON CONFLICT REPLACE,
		"STime"     TIMESTAMP DEFAULT (strftime('%s', 'now'))
);
-- +migrate Down
DROP TABLE "Sess";
