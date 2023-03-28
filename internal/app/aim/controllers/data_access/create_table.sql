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
	FOREIGN KEY(app_type) REFERENCES appliance_types(type_id) 
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

CREATE TABLE IF NOT EXISTS customs (
	app_id TEXT PRIMARY KEY NOT NULL,
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS buttons (
	app_id TEXT PRIMARY KEY NOT NULL,
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS toggles (
	app_id TEXT PRIMARY KEY NOT NULL,
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS thermostats (
	app_id TEXT PRIMARY KEY NOT NULL,
	scale REAL NOT NULL,
	minimum_heating_temp INT NOT NULL,
	maximum_heating_temp INT NOT NULL,
	minimum_cooling_temp INT NOT NULL,
	maximum_cooling_temp INT NOT NULL,
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
);