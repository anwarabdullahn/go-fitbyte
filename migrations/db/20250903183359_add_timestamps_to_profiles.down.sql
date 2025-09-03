-- Remove timestamp columns from profiles table
DROP INDEX IF EXISTS idx_profiles_deleted_at;

ALTER TABLE "profiles" 
DROP COLUMN IF EXISTS "deleted_at",
DROP COLUMN IF EXISTS "updated_at",
DROP COLUMN IF EXISTS "created_at";