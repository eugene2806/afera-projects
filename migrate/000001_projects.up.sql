CREATE TABLE IF NOT EXISTS projects (
             guid UUID PRIMARY KEY NOT NULL,
             alias VARCHAR(255) NOT NULL,
             name VARCHAR(255) NOT NULL,
             info TEXT,
             created_at TIMESTAMP NOT NULL,
             updated_at TIMESTAMP NOT NULL,
             deleted_at TIMESTAMP
);

INSERT INTO projects (guid, alias, name, info, created_at, updated_at, deleted_at) VALUES
            ('11111111-1111-1111-1111-111111111111', 'project_alpha', 'Alpha Project', 'The first project in the system.', '2024-01-01 10:00:00', '2024-01-02 12:00:00', NULL),
            ('22222222-2222-2222-2222-222222222222', 'project_beta', 'Beta Project', NULL, '2024-02-01 11:00:00', '2024-02-02 14:00:00', NULL),
            ('33333333-3333-3333-3333-333333333333', 'project_gamma', 'Gamma Project', 'This project has a detailed description.', '2024-03-01 09:00:00', '2024-03-02 15:00:00', '2024-05-01 10:00:00'),
            ('44444444-4444-4444-4444-444444444444', 'project_delta', 'Delta Project', NULL, '2024-04-01 08:00:00', '2024-04-02 16:00:00', NULL),
            ('55555555-5555-5555-5555-555555555555', 'project_epsilon', 'Epsilon Project', 'Epsilon is an experimental project.', '2024-05-01 07:00:00', '2024-05-02 17:00:00', '2024-06-01 12:00:00'),
            ('66666666-6666-6666-6666-666666666666', 'project_zeta', 'Zeta Project', NULL, '2024-06-01 06:00:00', '2024-06-02 18:00:00', NULL),
            ('77777777-7777-7777-7777-777777777777', 'project_eta', 'Eta Project', 'Another project with a focus on development.', '2024-07-01 05:00:00', '2024-07-02 19:00:00', NULL),
            ('88888888-8888-8888-8888-888888888888', 'project_theta', 'Theta Project', NULL, '2024-08-01 04:00:00', '2024-08-02 20:00:00', NULL),
            ('99999999-9999-9999-9999-999999999999', 'project_iota', 'Iota Project', 'Iota is in the ideation phase.', '2024-09-01 03:00:00', '2024-09-02 21:00:00', NULL),
            ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'project_kappa', 'Kappa Project', NULL, '2024-10-01 02:00:00', '2024-10-02 22:00:00', '2024-11-01 15:00:00'),
            ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'project_jopa', 'jopa Project', NULL, '2024-10-01 02:00:00', '2024-10-02 22:00:00', '2024-11-01 15:00:00');