CREATE TABLE seller
(
    id          serial PRIMARY KEY,
    url         text                     NOT NULL DEFAULT '' UNIQUE,
    key         text                     NOT NULL DEFAULT '',
    Name        text                     NOT NULL DEFAULT '',
    tel         text                     NOT NULL DEFAULT '',
    prefecture  text                     NOT NULL DEFAULT '',
    address     text                     NOT NULL DEFAULT '',
    longitude   numeric                  NOT NULL,
    latitude    numeric                  NOT NULL,
    opened_at   timestamp with time zone NOT NULL DEFAULT 'NOW()',
    meta        jsonb,
    exported    bool                     NOT NULL DEFAULT FALSE,
    exported_at timestamp with time zone,
    disabled    bool                     NOT NULL DEFAULT FALSE,
    created_at  timestamp with time zone NOT NULL DEFAULT 'NOW()',
    updated_at  timestamp with time zone NOT NULL DEFAULT 'NOW()'
);

CREATE INDEX idx_seller_prefecture ON seller (prefecture);
CREATE INDEX idx_seller_open_at ON seller (opened_at);

CREATE TABLE domain
(
    id               serial PRIMARY KEY,
    key              text                     NOT NULL DEFAULT '' UNIQUE,
    url              text                     NOT NULL,
    allow_domains    varchar[],
    latest_visit_url text                     NOT NULL DEFAULT '',
    created_at       timestamp with time zone NOT NULL DEFAULT 'NOW()',
    updated_at       timestamp with time zone NOT NULL DEFAULT 'NOW()'
);

CREATE UNIQUE INDEX idx_domain_key ON domain (key);