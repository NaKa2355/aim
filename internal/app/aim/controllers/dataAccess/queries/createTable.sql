PRAGMA foreign_keys=true;

CREATE TABLE IF NOT EXISTS remotes (
	remote_id TEXT PRIMARY KEY NOT NULL, 
	name TEXT NOT NULL UNIQUE, 
	device_id TEXT NOT NULL,
	tag TEXT
);

CREATE TABLE IF NOT EXISTS buttons (
	button_id TEXT NOT NULL,
	remote_id TEXT NOT NULL,
	name TEXT NOT NULL,
	tag TEXT,
	irdata BLOB NOT NULL,
	FOREIGN KEY (remote_id) REFERENCES remotes (remote_id) ON DELETE CASCADE,
	PRIMARY KEY (button_id, remote_id)
	UNIQUE (name, remote_id)
);
