-- Create boards table
CREATE TABLE boards (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW()
);

-- Create columns table
CREATE TABLE columns (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         board_id INT REFERENCES boards(id) ON DELETE CASCADE,
                         position INT NOT NULL,  -- Position defines the order of the columns in the board
                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW()
);

-- Create tasks table
CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       column_id INT REFERENCES columns(id) ON DELETE CASCADE,
                       position INT NOT NULL,  -- Position defines the order of tasks in the column
                       due_date DATE,
                       created_at TIMESTAMP DEFAULT NOW(),
                       updated_at TIMESTAMP DEFAULT NOW()
);
