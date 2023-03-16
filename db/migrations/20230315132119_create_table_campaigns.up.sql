CREATE TABLE `campaigns` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `user_id` int DEFAULT NULL,
                             `name` varchar(255) DEFAULT NULL,
                             `summary` varchar(255) DEFAULT NULL,
                             `description` text,
                             `perks` text,
                             `backer_count` int DEFAULT NULL,
                             `goal_amount` int DEFAULT NULL,
                             `current_amount` int DEFAULT NULL,
                             `slug` varchar(255) DEFAULT NULL,
                             `created_at` datetime DEFAULT NULL,
                             `updated_at` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB;