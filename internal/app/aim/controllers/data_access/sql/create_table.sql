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

CREATE TABLE IF NOT EXISTS buttons (
	app_id TEXT NOT NULL,
	
	PRIMARY KEY (app_id),
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS switches (
	app_id TEXT NOT NULL,

	PRIMARY KEY (app_id),
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS thermostats (
	app_id TEXT NOT NULL,
	scale REAL NOT NULL,
	maximum_heating_temp INTEGER NOT NULL,
	minimum_heating_temp INTEGER NOT NULL,
	maximum_cooling_temp INTEGER NOT NULL,
	minimum_cooling_temp INTEGER NOT NULL,

	PRIMARY KEY (app_id),
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS commands (
	command_id TEXT PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	app_id TEXT NOT NULL,
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
	UNIQUE (name, app_id)
);

CREATE TABLE IF NOT EXISTS irdata (
	command_id TEXT PRIMARY KEY NOT NULL,
	irdata BLOB,
	FOREIGN KEY (command_id) REFERENCES commands (command_id) ON DELETE CASCADE
);