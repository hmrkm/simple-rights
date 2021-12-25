CREATE TABLE `users` (
  `id` varchar(26) NOT NULL COMMENT '主キー',
  `email` varchar(255) NOT NULL COMMENT 'メールアドレス',
  `password` varchar(255) NOT NULL COMMENT 'ハッシュ化されたパスワード',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
CREATE TABLE `permissions` (
  `resource_id` varchar(26) NOT NULL COMMENT 'リソースID',
  `role_id` varchar(26) NOT NULL COMMENT 'ロールID'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
CREATE TABLE `resources` (
  `id` varchar(26) NOT NULL COMMENT '主キー',
  `path` varchar(1024) NOT NULL COMMENT 'パス'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
CREATE TABLE `roles` (
  `id` varchar(26) NOT NULL COMMENT '主キー',
  `name` varchar(255) NOT NULL COMMENT '名前'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
CREATE TABLE `user_roles` (
  `user_id` varchar(26) NOT NULL COMMENT 'ユーザーID',
  `role_id` varchar(26) NOT NULL COMMENT 'ロールID'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
