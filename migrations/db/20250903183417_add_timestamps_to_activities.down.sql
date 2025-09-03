-- Remove timestamp columns from activities table
DROP INDEX IF EXISTS idx_activities_deleted_at;

ALTER TABLE "activities" 
DROP COLUMN IF EXISTS "deleted_at",
DROP COLUMN IF EXISTS "updated_at",
DROP COLUMN IF EXISTS "created_at";