-- modify "todolists" table
ALTER TABLE `todolists` ADD COLUMN `delete_flag` bigint NOT NULL DEFAULT 0;
