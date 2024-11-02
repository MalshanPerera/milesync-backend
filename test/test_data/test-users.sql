INSERT INTO users (id, first_name, last_name, email, password)
VALUES 
    ('123123', 'John', 'Maclane', 'john@maclane.com', '12312'),
    ('456456', 'Jane', 'Doe', 'jane@doe.com', '45645'),
    ('789789', 'Alice', 'Smith', 'alice@smith.com', '78978')
RETURNING *;
