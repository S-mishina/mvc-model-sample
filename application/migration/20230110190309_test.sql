-- create "todolists" table
CREATE TABLE `todolists` (`id` bigint NOT NULL AUTO_INCREMENT, `title` varchar(255) NOT NULL, `body` varchar(255) NOT NULL, `priority` bigint NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
