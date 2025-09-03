-- Remove timestamp columns from users table
DROP INDEX IF EXISTS idx_users_deleted_at;

ALTER TABLE "users" 
DROP COLUMN IF EXISTS "deleted_at",
DROP COLUMN IF EXISTS "updated_at",
DROP COLUMN IF EXISTS "created_at";