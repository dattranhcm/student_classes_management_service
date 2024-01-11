CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    user_type VARCHAR(20) NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT valid_user_type CHECK (user_type IN ('Student', 'Teacher', 'Administrator'))
);


CREATE TABLE classes (
    class_id SERIAL PRIMARY KEY,
    class_name VARCHAR(100) NOT NULL,
    teacher_id INT,
    day_of_week VARCHAR(50) NOT NULL, -- Monday, Tuesday, etc.
    start_time VARCHAR(100) NOT NULL,
    end_time VARCHAR(100) NOT NULL,
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT class_teacher
        FOREIGN KEY (teacher_id) REFERENCES users (user_id) ON DELETE CASCADE
);

CREATE TABLE students_classes (
    student_id INT,
    class_id INT,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (student_id, class_id),
    CONSTRAINT class_fk
        FOREIGN KEY (class_id) REFERENCES classes (class_id) ON DELETE CASCADE,
    CONSTRAINT student_fk
        FOREIGN KEY (student_id) REFERENCES users (user_id) ON DELETE CASCADE
)