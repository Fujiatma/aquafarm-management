CREATE TABLE farms
(
    id         VARCHAR(36) PRIMARY KEY,
    farm_name  VARCHAR,
    user_id    VARCHAR(36),
    is_deleted BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE ponds
(
    id         VARCHAR(36) PRIMARY KEY,
    pond_name  VARCHAR,
    farm_id    VARCHAR(36),
    is_deleted BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE statistics
(
    id         VARCHAR(36) PRIMARY KEY,
    endpoint   VARCHAR,
    user_id    VARCHAR(36),
    count      INT,
    call_at    TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE users
(
    id         VARCHAR(36) PRIMARY KEY,
    user_name  VARCHAR NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);