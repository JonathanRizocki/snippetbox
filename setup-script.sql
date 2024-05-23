-- UUID_generate_v4 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Sanity check it was created
SELECT * FROM pg_extension WHERE extname = 'uuid-ossp';

-- Create a `snippets` table.
CREATE TABLE snippets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT timezone('utc', current_timestamp),
    expires TIMESTAMPTZ NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);



-- Add some dummy records
INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    timezone('utc', current_timestamp),
    current_timestamp + interval '365 days'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    timezone('utc', current_timestamp),
    current_timestamp + interval '365 days'
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    timezone('utc', current_timestamp),
    current_timestamp + interval '7 days'
);
