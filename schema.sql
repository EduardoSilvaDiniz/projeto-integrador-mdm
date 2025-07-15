CREATE TABLE associated (
  number_card INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  group_id INTEGER NOT NULL,
  FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE
);

CREATE TABLE presence (
  number_card INTEGER NOT NULL,
  meeting_id INTEGER NOT NULL,
  is_presence BOOLEAN NOT NULL,
  PRIMARY KEY (number_card, meeting_id),
  FOREIGN KEY (number_card) REFERENCES associated (number_card) ON DELETE CASCADE,
  FOREIGN KEY (meeting_id) REFERENCES meeting (id) ON DELETE CASCADE
);

CREATE TABLE meeting (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  group_id INTEGER NOT NULL,
  address TEXT NOT NULL,
  date DATE NOT NULL,
  FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE
);

CREATE TABLE groups (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  hours DATE NOT NULL
);

CREATE TABLE payment (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  number_card INTEGER NOT NULL,
  ref_month TEXT NOT NULL,
  payment_date DATE NOT NULL DEFAULT CURRENT_DATE,
  FOREIGN KEY (number_card) REFERENCES associated (number_card) ON DELETE CASCADE
);
