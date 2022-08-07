CREATE TABLE it_users (
    user_id VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL PRIMARY KEY,
    dtime_created TIMESTAMP,
    dtime_last_login TIMESTAMP,
    status VARCHAR(50) NOT NULL,
    password VARCHAR(255)
);
