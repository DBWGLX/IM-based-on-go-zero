CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  `account` varchar(64) NOT NULL,
  `password` varchar(128) NOT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `is_delete` tinyint NOT NULL DEFAULT 0,
  `create_time` timestamp DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  UNIQUE KEY `uk_account` (`account`),
  KEY `idx_phone` (`phone`),              -- phone 普通索引
  KEY `idx_email` (`email`),              -- email 普通索引
  KEY `idx_create_time` (`create_time`)   -- create_time 普通索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
