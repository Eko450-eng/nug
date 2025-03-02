CREATE TABLE IF NOT EXISTS task (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    description TEXT,

    project_id INTEGER,
    prio INTEGER DEFAULT 0,

    time TEXT DEFAULT (datetime('now')),
    deletedtime TEXT,               

    modified TEXT,         
    completed INTEGER DEFAULT 0,     
    deleted INTEGER DEFAULT 0,       

    FOREIGN KEY (project_id) REFERENCES project(id)
);

-- Tags
CREATE TABLE IF NOT EXISTS tags (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL UNIQUE,

    modified TEXT NOT NULL,       
    deleted INTEGER DEFAULT 0,
    deletedtime TEXT        
);

-- Company
CREATE TABLE IF NOT EXISTS company (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL UNIQUE,

    modified TEXT,         
    deleted INTEGER DEFAULT 0,       
    deletedtime TEXT
);

-- Person
CREATE TABLE IF NOT EXISTS person (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL UNIQUE,

    modified TEXT,         
    deleted INTEGER DEFAULT 0,       
    deletedtime TEXT
);

-- Links
CREATE TABLE IF NOT EXISTS links (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL UNIQUE,

    modified TEXT,         
    deleted INTEGER DEFAULT 0,       
    deletedtime TEXT               
);

CREATE TABLE IF NOT EXISTS project (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL UNIQUE,
    budget INTEGER,  -- Added missing comma

    modified TEXT,
    deleted INTEGER DEFAULT 0,       
    deletedtime TEXT    
);

CREATE TABLE IF NOT EXISTS task_project (
    task_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    PRIMARY KEY (task_id, project_id),
    FOREIGN KEY (task_id) REFERENCES task(id),
    FOREIGN KEY (project_id) REFERENCES project(id)
);

-- Relations
CREATE TABLE IF NOT EXISTS task_tags (
    task_id INTEGER NOT NULL,
    tag_id INTEGER NOT NULL,
    PRIMARY KEY (task_id, tag_id),
    FOREIGN KEY (task_id) REFERENCES task(id),  -- Fixed reference: was tasks(id)
    FOREIGN KEY (tag_id) REFERENCES tags(id)
);

-- Times
CREATE TABLE IF NOT EXISTS start_times (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    task_id INTEGER NOT NULL,
    time TEXT NOT NULL,
    FOREIGN KEY (task_id) REFERENCES task(id)  -- Fixed reference: was tasks(id)
);

CREATE TABLE IF NOT EXISTS end_times (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    task_id INTEGER NOT NULL,
    time TEXT NOT NULL,
    FOREIGN KEY (task_id) REFERENCES task(id)  -- Fixed reference: was tasks(id)
);

CREATE TABLE IF NOT EXISTS company_project (
    company_id INTEGER NOT NULL,
    project_id INTEGER NOT NULL,
    PRIMARY KEY (company_id, project_id),
    FOREIGN KEY (company_id) REFERENCES company(id),
    FOREIGN KEY (project_id) REFERENCES project(id)
);

CREATE TABLE IF NOT EXISTS company_person (
    company_id INTEGER NOT NULL,
    person_id INTEGER NOT NULL,
    PRIMARY KEY (company_id, person_id),
    FOREIGN KEY (company_id) REFERENCES company(id),
    FOREIGN KEY (person_id) REFERENCES person(id)
);
