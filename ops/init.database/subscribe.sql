CREATE TABLE `subscribe` (
    `id`            varchar(26) NOT NULL,
    `connection_id` varchar(26) NOT NULL,
    `topic`         varchar(26) NOT NULL,
    `expiry_at`     datetime NOT NULL,
    `create_at`     datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `index_connection_id_topic` (`connection_id`, `topic`),
    KEY `index_connection_id` (`connection_id`),
    KEY `index_topic` (`topic`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
