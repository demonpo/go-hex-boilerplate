-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "name" text NULL,
  "email" text NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "uni_users_email" to table: "users"
CREATE UNIQUE INDEX "uni_users_email" ON "public"."users" ("email");
