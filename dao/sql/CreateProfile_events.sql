CREATE TABLE create_profile_events (
    id INT PRIMARY KEY AUTO_INCREMENT,
    chain_name VARCHAR(255) NOT NULL,
    chain_id INT NOT NULL,
    to_address VARCHAR(255) NOT NULL,
    profile_id INT NOT NULL,
    handle VARCHAR(255) NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    metadata TEXT NOT NULL
);