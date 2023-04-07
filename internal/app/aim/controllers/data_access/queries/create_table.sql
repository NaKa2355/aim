PRAGMA foreign_keys=true;

CREATE TABLE IF NOT EXISTS appliances (
	app_id TEXT PRIMARY KEY NOT NULL, 
	name TEXT NOT NULL UNIQUE, 
	device_id TEXT NOT NULL
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
	FOREIGN KEY (app_id) REFERENCES appliances (app_id) ON DELETE CASCADE
);


DROP VIEW IF EXISTS appliances_sti;

CREATE VIEW appliances_sti AS 
SELECT  apps.app_id, apps.app_type, a.name, a.device_id
FROM
(
	SELECT 0 AS app_type, app_id FROM customs
	UNION
	SELECT 1 AS app_type, app_id FROM buttons
	UNION
	SELECT 2 AS app_type, app_id FROM toggles
	UNION
	SELECT 3 AS app_type, app_id FROM thermostats
) AS apps
LEFT JOIN appliances a ON apps.app_id = a.app_id;


