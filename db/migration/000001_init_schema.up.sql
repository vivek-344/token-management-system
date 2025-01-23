CREATE TABLE "token" (
  "token_id" varchar(12) NOT NULL,
  "usage_count" integer NOT NULL,
  "last_updated" timestamp NOT NULL
);

CREATE INDEX ON "token" ("usage_count");

CREATE INDEX ON "token" ("last_updated");