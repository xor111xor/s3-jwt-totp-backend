CREATE DATABASE store;

\c store;

CREATE TABLE users (uuid_user UUID PRIMARY KEY, mail VARCHAR(50), hash VARCHAR(100), date_reg TIMESTAMP, otp_cache VARCHAR(50));

CREATE TABLE files(uuid_file UUID PRIMARY KEY, file_name VARCHAR(50), upload_date TIMESTAMP, size BIGINT, bucket_name VARCHAR(50), uuid_user UUID);
