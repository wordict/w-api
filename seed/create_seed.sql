CREATE TABLE IF NOT EXISTS users(
   id VARCHAR NOT NULL PRIMARY KEY,
   email VARCHAR NOT NULL UNIQUE,
   password VARCHAR NOT NULL,
   created_at TIMESTAMP DEFAULT now(),
   updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS vocabulary(
    id                      VARCHAR NOT NULL PRIMARY KEY,
    word                    VARCHAR NOT NULL,
    user_id                 VARCHAR NOT NULL references users ON DELETE CASCADE ON UPDATE CASCADE,
    from_language           VARCHAR NOT NULL,
    to_language             VARCHAR NOT NULL,
    word_translation        VARCHAR,
    created_at              TIMESTAMP DEFAULT now(),
    updated_at              TIMESTAMP DEFAULT now()
);