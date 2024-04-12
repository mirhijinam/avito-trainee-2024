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
    additional_info jsonb,
    FOREIGN KEY (feature_id) REFERENCES features(id)
);

-- many-to-many between banners and tags
CREATE TABLE banner_tags (
    banner_id integer NOT NULL,
    tag_id    integer NOT NULL,
    PRIMARY KEY (banner_id, tag_id),
    FOREIGN KEY (banner_id) REFERENCES banners(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE RESTRICT
);
