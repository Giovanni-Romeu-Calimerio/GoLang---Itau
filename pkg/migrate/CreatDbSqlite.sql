CREATE TABLE products (
    ID    INTEGER             PRIMARY KEY,
    Name  TEXT                NOT NULL,
    Code  INTEGER             NOT NULL,
    Price INTEGER             NOT NULL,
    CreatedAt [CURRENT_TIMESTAMP],
    UpdatedAt [CURRENT_TIMESTAMP]
);

