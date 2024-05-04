INSERT INTO users (id, first_name, last_name, email, password)
VALUES (
    '123123',
    'John',
    'Maclane',
    'john@maclane.com',
    '12312'
  )
RETURNING *;

INSERT INTO organization (id, name, slug, user_id)
VALUES (
    '151515',
    'John',
    'john',
    '123123'
  )
RETURNING *;