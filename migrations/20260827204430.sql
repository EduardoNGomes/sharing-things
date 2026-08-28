-- Create "users" table
CREATE TABLE "public"."users" (
  "uuid" uuid NOT NULL,
  "internal_id" bigint NULL,
  "name" text NULL,
  "email" text NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  PRIMARY KEY ("uuid")
);
