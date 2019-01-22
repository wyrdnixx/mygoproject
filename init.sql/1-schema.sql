
DROP TABLE IF EXISTS testtable;

CREATE TABLE testtable
(
    id    SERIAL PRIMARY KEY NOT NULL,         
    name   varchar(40) NOT NULL CHECK (name <> '')
    );