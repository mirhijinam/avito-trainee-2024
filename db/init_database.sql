CREATE TABLE tags (
    id serial PRIMARY KEY
);

-- one-to-many between tags and users
CREATE TABLE users (
    id serial PRIMARY KEY,
    tag_id integer,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE SET NULL
);

CREATE TABLE features (
    id serial PRIMARY KEY
);

CREATE TABLE banners (
    id              serial PRIMARY KEY,
    feature_id      integer NOT NULL,
    content         jsonb,
    is_active       boolean,
    created_at      timestamp,
    updated_at      timestamp,
    FOREIGN KEY (feature_id) REFERENCES features(id)
);

-- many-to-many between banners and tags
CREATE TABLE banners_tags (
    banner_id integer NOT NULL,
    tag_id    integer NOT NULL,
    PRIMARY KEY (banner_id, tag_id),
    FOREIGN KEY (banner_id) REFERENCES banners(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE RESTRICT
);

INSERT INTO tags (id)
SELECT 1
WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 1);

INSERT INTO tags (id)
SELECT 2
WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 2);

INSERT INTO tags (id)
SELECT 3
WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 3);

INSERT INTO tags (id)
SELECT 4
WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 4);

INSERT INTO tags (id)
SELECT 5
WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 5);

INSERT INTO tags (id)
SELECT 6
WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 6);

INSERT INTO tags (id)
SELECT 7
WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 7);

INSERT INTO tags (id)
SELECT 8
WHERE NOT EXISTS (SELECT 1 FROM tags WHERE id = 8);


INSERT INTO features (id)
SELECT 101
WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 101);

INSERT INTO features (id)
SELECT 102
WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 102);

INSERT INTO features (id)
SELECT 103
WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 103);

INSERT INTO features (id)
SELECT 104
WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 104);

INSERT INTO features (id)
SELECT 105
WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 105);

INSERT INTO features (id)
SELECT 123
WHERE NOT EXISTS (SELECT 1 FROM features WHERE id = 123);
