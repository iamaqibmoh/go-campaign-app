CREATE TABLE `transactions` (
                                `id` int NOT NULL AUTO_INCREMENT,
                                `campaign_id` int DEFAULT NULL,
                                `user_id` int DEFAULT NULL,
                                `amount` int DEFAULT NULL,
                                `status` varchar(255) DEFAULT NULL,
                                `code` varchar(255) DEFAULT NULL,
                                `created_at` datetime DEFAULT NULL,
                                `update_at` datetime DEFAULT NULL,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB;