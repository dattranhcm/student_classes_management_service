CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL, -- Use appropriate hashing for passwords
    user_type VARCHAR(20) NOT NULL, -- Student, Teacher, Administrator
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT valid_user_type CHECK (user_type IN ('Student', 'Teacher', 'Administrator'))
);

CREATE TABLE classes (
    class_id SERIAL PRIMARY KEY,
    class_name VARCHAR(100) NOT NULL,
    teacher_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    created_at timestamp,
    updated_at timestamp
    -- Add other class-related attributes as needed
);

CREATE TABLE students_classes (
    student_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    class_id INT REFERENCES classes(class_id) ON DELETE CASCADE,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (student_id, class_id)
);

CREATE TABLE schedules (
    schedule_id SERIAL PRIMARY KEY,
    class_id INT REFERENCES classes(class_id) ON DELETE CASCADE,
    day_of_week VARCHAR(10) NOT NULL, -- Monday, Tuesday, etc.
    start_time VARCHAR(100) NOT NULL,
    end_time VARCHAR(100) NOT NULL,
    created_at timestamp,
    updated_at timestamp
    -- Add other schedule-related attributes as needed
);