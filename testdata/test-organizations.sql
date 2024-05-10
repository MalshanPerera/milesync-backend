
INSERT INTO organization (id, user_id, name, slug)
VALUES (
    '151515',
    '123123',
    'John',
    'john'
)
RETURNING *;