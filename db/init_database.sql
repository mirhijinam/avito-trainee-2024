CREATE TABLE users (
    id          serial PRIMARY KEY
);

CREATE TABLE tags (
    id          serial PRIMARY KEY
);

CREATE TABLE users_tags (
    id              serial PRIMARY KEY,
    user_id         integer,
    tag_id          integer,

    CONSTRAINT user_tag_unique UNIQUE (user_id, tag_id)
);

CREATE TABLE features (
    id          serial PRIMARY KEY
);

CREATE TABLE banners (
    id                      serial PRIMARY KEY,            
    feature_id              integer NOT null,
    tag_ids                 integer[] NOT null,
    additional_info         jsonb,

    FOREIGN KEY (feature_id)            REFERENCES features(id),
    FOREIGN KEY (tag_id)                REFERENCES tags(id)
);

