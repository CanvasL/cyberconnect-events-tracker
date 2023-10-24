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
        avatar TEXT,
        metadata TEXT
    );

CREATE TABLE
    collect_paid_mw_set_events (
        id INT PRIMARY KEY AUTO_INCREMENT,
        chain_name VARCHAR(255) NOT NULL,
        chain_id INT NOT NULL,
        block_number INT NOT NULL,
        tx_hash VARCHAR(66) NOT NULL,
        `namespace` VARCHAR(255) NOT NULL,
        profile_id VARCHAR(255) NOT NULL,
        essence_id VARCHAR(255) NOT NULL,
        total_supply VARCHAR(255) NOT NULL,
        amount VARCHAR(255) NOT NULL,
        recipient VARCHAR(255) NOT NULL,
        currency VARCHAR(255) NOT NULL,
        subscribe_required TINYINT(1) NOT NULL
    );