CREATE TYPE attribute_types AS ENUM (
  'datetime',
  'text',
  'number'
);

CREATE TABLE IF NOT EXISTS profession (
    id uuid primary key,
    name varchar(255) not null
);

CREATE TABLE attribute (
  id uuid PRIMARY KEY,
  name varchar,
  type attribute_types
);

CREATE TABLE IF NOT EXISTS position (
    id uuid primary key,
    name varchar,
    profession_id uuid  references profession(id),
    company_id uuid not null
);

CREATE TABLE IF NOT EXISTS position_attributes (
    id uuid primary key,
    attribute_id uuid  references attribute(id),
    position_id uuid  references position(id),
    value varchar
);

CREATE TABLE company (
  id uuid PRIMARY KEY,
  name varchar
);

ALTER TABLE position ADD FOREIGN KEY (profession_id) REFERENCES profession (id);

ALTER TABLE position ADD FOREIGN KEY (company_id) REFERENCES company (id);

ALTER TABLE position_attributes ADD FOREIGN KEY (attribute_id) REFERENCES attribute (id);

ALTER TABLE position_attributes ADD FOREIGN KEY (position_id) REFERENCES position (id);
