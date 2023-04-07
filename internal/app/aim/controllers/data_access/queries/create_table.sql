PRAGMA foreign_keys=true;

CREATE TABLE IF NOT EXISTS appliances (
	app_id TEXT PRIMARY KEY NOT NULL, 
	name TEXT NOT NULL UNIQUE, 
	device_id TEXT NOT NULL,
	app_type INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS commands (
	com_id TEXT NOT NULL,
	app_id TEXT NOT NULL,
	name TEXT NOT NULL,
	irdata BLOB NOT NULL,
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE,
	PRIMARY KEY (com_id, app_id)
	UNIQUE (name, app_id)
);

