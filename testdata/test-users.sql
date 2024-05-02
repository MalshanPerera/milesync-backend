INSERT INTO users (first_name, last_name, email, password)
VALUES ('John', 'Maclane', 'john@maclane.com', '12312')
RETURNING *;