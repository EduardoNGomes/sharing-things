-- Create "users" table
CREATE TABLE "users" (
  "uuid" uuid NOT NULL,
  "internal_id" bigserial NOT NULL,
  "name" character varying(50) NOT NULL,
  "email" character varying(50) NOT NULL,
  "password" character varying(100) NOT NULL,
  "created_at" timestamptz NULL DEFAULT now(),
  "updated_at" timestamptz NULL DEFAULT now(),
  PRIMARY KEY ("uuid")
);
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
-- Create index "idx_users_internal_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_internal_id" ON "users" ("internal_id");
