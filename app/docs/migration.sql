CREATE TABLE `users` (
  `id` varchar(16) NOT NULL COMMENT '主キー',
  `email` varchar(255) NOT NULL COMMENT 'メールアドレス',
  `password` varchar(255) NOT NULL COMMENT 'ハッシュ化されたパスワード',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
CREATE TABLE `permissions` (
  `resource_id` varchar(16) NOT NULL COMMENT 'リソースID',
  `role_id` varchar(16) NOT NULL COMMENT 'ロールID'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
CREATE TABLE `resources` (
  `id` varchar(16) NOT NULL COMMENT '主キー',
  `path` varchar(1024) NOT NULL COMMENT 'パス'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
CREATE TABLE `roles` (
  `id` varchar(16) NOT NULL COMMENT '主キー',
  `name` varchar(255) NOT NULL COMMENT '名前'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
CREATE TABLE `user_roles` (
  `user_id` varchar(16) NOT NULL COMMENT 'ユーザーID',
  `role_id` varchar(16) NOT NULL COMMENT 'ロールID'
) ENGINE = InnoDB DEFAULT CHARSET = latin1;
insert into resources (id, path)
values ('a', 'aaa'),
  ('b', 'bbb'),
  ('c', 'ccc');
insert into roles (id, name)
values ('x', 'xxx'),
  ('y', 'yyy'),
  ('z', 'zzz');
insert into `permissions` (resource_id, role_id)
values ('a', 'x'),
  ('b', 'y'),
  ('c', 'z'),
  ('a', 'y'),
  ('a', 'z');
insert into user_roles (user_id, role_id)
values ('1', 'x'),
  ('2', 'y'),
  ('3', 'z');