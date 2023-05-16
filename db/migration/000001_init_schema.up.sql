CREATE TABLE "todos" (
  "id" bigserial PRIMARY KEY,  
  "title" varchar NOT NULL,
  "note" varchar NOT NULL,
  "due_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);