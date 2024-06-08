CREATE TABLE IF NOT EXISTS application (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    type VARCHAR(50) NOT NULL,
    bank VARCHAR(50) NOT NULL,
    opsyears INTEGER NOT NULL CHECK (opsyears > 0),
    ssm BOOLEAN NOT NULL,
    prevgateway BOOLEAN NOT NULL,
    prodtype VARCHAR(255) NOT NULL,
    storetype VARCHAR(255) NOT NULL,
    inventory BOOLEAN NOT NULL,
    reference VARCHAR(50),
    socmedia VARCHAR(50),
    litigation BOOLEAN NOT NULL,
    score INTEGER NOT NULL CHECK (score >= 0 AND score <= 10),
    created TIMESTAMP DEFAULT NOW(),
    CONSTRAINT name_bank_check CHECK (name = bank)
);