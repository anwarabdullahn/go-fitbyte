-- Add timestamp columns to activities table
ALTER TABLE "activities" 
ADD COLUMN "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN "deleted_at" timestamp NULL;

-- Create index for deleted_at to optimize soft delete queries
CREATE INDEX idx_activities_deleted_at ON "activities" ("deleted_at");