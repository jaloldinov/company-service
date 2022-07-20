CREATE TYPE "attribute_types" AS ENUM (
  'datetime',
  'text',
  'number'
);

CREATE TABLE "profession" (
  "id" uuid PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "position" (
  "id" uuid,
  "name" varchar,
  "profession_id" uuid,
  "company_id" uuid,
  PRIMARY KEY ("id", "profession_id", "company_id")
);

CREATE TABLE "attribute" (
  "id" uuid PRIMARY KEY,
  "name" varchar,
  "type" attribute_types
);

CREATE TABLE "position_attributes" (
  "id" uuid,
  "attribute_id" uuid,
  "position_id" uuid,
  "value" varchar,
  PRIMARY KEY ("id", "attribute_id", "position_id")
);

CREATE TABLE "company" (
  "id" uuid PRIMARY KEY,
  "name" varchar
);

ALTER TABLE "position" ADD FOREIGN KEY ("profession_id") REFERENCES "profession" ("id");

ALTER TABLE "position" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "position_attributes" ADD FOREIGN KEY ("attribute_id") REFERENCES "attribute" ("id");

ALTER TABLE "position_attributes" ADD FOREIGN KEY ("position_id") REFERENCES "position" ("id");
