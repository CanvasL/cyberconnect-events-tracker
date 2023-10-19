CREATE TABLE
    create_profile_events (
        id INT PRIMARY KEY AUTO_INCREMENT,
        chain_name VARCHAR(255) NOT NULL,
        chain_id INT NOT NULL,
        block_number INT NOT NULL,
        tx_hash VARCHAR(66) NOT NULL,
        `to` VARCHAR(255) NOT NULL,
        profile_id VARCHAR(255) NOT NULL,
        handle VARCHAR(255) NOT NULL,
        avatar VARCHAR(255),
        metadata TEXT
    );