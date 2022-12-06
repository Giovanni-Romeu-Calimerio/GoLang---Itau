CREATE TABLE products (
    id    INTEGER             PRIMARY KEY,
    name  TEXT                NOT NULL,
    code  INTEGER             NOT NULL,
    price INTEGER             NOT NULL,
    datac [CURRENT_TIMESTAMP],
    datau [CURRENT_TIMESTAMP]
);

