-- Add timestamp columns to users table
ALTER TABLE "users" 
ADD COLUMN "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN "updated_at" timestamp DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN "deleted_at" timestamp NULL;

-- Create index for deleted_at to optimize soft delete queries
CREATE INDEX idx_users_deleted_at ON "users" ("deleted_at");