-- Create the Record table
CREATE TABLE IF NOT EXISTS dynamic_dns (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    hostname VARCHAR(255) NOT NULL,
    enabled BOOLEAN NOT NULL
);

-- Create the DNSLog table
CREATE TABLE IF NOT EXISTS dns_log (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL,
    ip VARCHAR(45) NOT NULL
);