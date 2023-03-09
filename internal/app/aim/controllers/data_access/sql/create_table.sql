PRAGMA foreign_keys=true;

CREATE TABLE IF NOT EXISTS appliance_types (
	type_id INTEGER PRIMARY KEY NOT NULL, 
	name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS appliances (
	app_id TEXT PRIMARY KEY NOT NULL, 
	name TEXT NOT NULL UNIQUE, 
	app_type INTEGER NOT NULL,
	device_id TEXT NOT NULL,
	opt TEXT,
	FOREIGN KEY(app_type) REFERENCES appliance_types(type_id) 
);

CREATE TABLE IF NOT EXISTS commands (
	com_id TEXT PRIMARY KEY NOT NULL,
	app_id TEXT NOT NULL,
	name TEXT NOT NULL,
	irdata BLOB NULL,
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
	--UNIQUE (name, app_id)
)