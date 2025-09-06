-- Create user_files table
CREATE TABLE user_files (
    id SERIAL PRIMARY KEY,
    uri VARCHAR(255) NOT NULL,
    user_id INTEGER UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraint (one-to-one relationship)
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create indexes
CREATE UNIQUE INDEX idx_user_files_user_id ON user_files(user_id);