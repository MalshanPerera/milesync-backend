INSERT INTO projects (id, user_id, organization_id, name, key_prefix, type) 
VALUES 
    ('111111', '123123', '151515', 'Project 1', 'P1', 'global'),
    ('222222', '123123', '151515', 'Project 2', 'P2', 'project'),
    ('333333', '123123', '151515', 'Project 3', 'P3', 'project')
RETURNING *;

