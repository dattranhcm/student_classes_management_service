CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL, -- Use appropriate hashing for passwords
    user_type VARCHAR(20) NOT NULL, -- Student, Teacher, Administrator
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT valid_user_type CHECK (user_type IN ('Student', 'Teacher', 'Administrator'))
)