CREATE TABLE "books" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "author" varchar NOT NULL,
  "desc" varchar NOT NULL,
  "isbn" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "books" ("id");

CREATE INDEX ON "books" ("title");

CREATE INDEX ON "books" ("author");
