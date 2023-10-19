CREATE TABLE
    collect_paid_mw_set_events (
        id INT PRIMARY KEY AUTO_INCREMENT,
        chain_name VARCHAR(255) NOT NULL,
        chain_id INT NOT NULL,
        block_number INT NOT NULL,
        tx_hash VARCHAR(66) NOT NULL,
        `namespace` VARCHAR(255) NOT NULL,
        profile_id INT NOT NULL,
        essence_id INT NOT NULL,
        total_supply INT NOT NULL,
        amount INT NOT NULL,
        recipient VARCHAR(255) NOT NULL,
        currency VARCHAR(255) NOT NULL,
        subscribe_required TINYINT(1) NOT NULL
    );