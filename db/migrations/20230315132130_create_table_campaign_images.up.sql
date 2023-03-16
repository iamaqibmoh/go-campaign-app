CREATE TABLE `campaign_images` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `campaign_id` int DEFAULT NULL,
                                   `file_name` varchar(255) DEFAULT NULL,
                                   `is_primary` tinyint DEFAULT NULL,
                                   `created_at` datetime DEFAULT NULL,
                                   `updated_at` datetime DEFAULT NULL,
                                   PRIMARY KEY (`id`)
) ENGINE=InnoDB;