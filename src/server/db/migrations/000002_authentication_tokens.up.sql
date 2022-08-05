CREATE TABLE authentication_tokens (
           token_id VARCHAR(50) NOT NULL PRIMARY KEY,
           user_id VARCHAR(50),
           auth_token VARCHAR(255),
           generated_at TIMESTAMP,
           expires_at   TIMESTAMP
);
