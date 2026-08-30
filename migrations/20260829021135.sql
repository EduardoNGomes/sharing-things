-- Create sequence for serial column "internal_id"
CREATE SEQUENCE IF NOT EXISTS "public"."users_internal_id_seq" OWNED BY "public"."users"."internal_id";
-- Modify "users" table
ALTER TABLE "public"."users" ALTER COLUMN "internal_id" SET DEFAULT nextval('"public"."users_internal_id_seq"'), ALTER COLUMN "internal_id" SET NOT NULL, ALTER COLUMN "name" TYPE character varying(50), ALTER COLUMN "name" SET NOT NULL, ALTER COLUMN "email" TYPE character varying(50), ALTER COLUMN "email" SET NOT NULL, ALTER COLUMN "created_at" SET DEFAULT now(), ALTER COLUMN "updated_at" SET DEFAULT now(), ADD COLUMN "password" character varying(50) NOT NULL;
-- Create index "idx_users_email" to table: "users"
CREATE UNIQUE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create index "idx_users_internal_id" to table: "users"
CREATE UNIQUE INDEX "idx_users_internal_id" ON "public"."users" ("internal_id");
