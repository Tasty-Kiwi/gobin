CREATE TABLE
  IF NOT EXISTS bins (
    uuid TEXT PRIMARY KEY NOT NULL,
    content TEXT NOT NULL,
    creation_date INTEGER NOT NULL
  )