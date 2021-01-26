CREATE TABLE `tweets` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int,
  `content` text NOT NULL,
  `parent_tweet` int,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`parent_tweet`) REFERENCES `tweets` (`id`) ON DELETE CASCADE
);
