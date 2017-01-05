## Database(gradmin)

### 管理员用户表

表的结构 `gra_user`
    
    CREATE TABLE `gra_user` (
      `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户自增ID',
      `username` varchar(32) NOT NULL COMMENT '用户名',
      `password` char(32) NOT NULL COMMENT '用户密码，32位',
      `salt` char(8) NOT NULL COMMENT '用户的随机盐值，8位',
      `add_time` date NOT NULL COMMENT '创建时间',
      `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '用户状态，0禁用;1正常',
      `last_ip` varchar(15) NOT NULL COMMENT '最后一次登录IP',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    


### Redis配置表

表的结构 `gra_redis`

    
    CREATE TABLE `gra_redis` (
      `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
      `mark` varchar(64) NOT NULL COMMENT '链接实例的名称',
      `network` tinyint(4) NOT NULL DEFAULT '0' COMMENT '链接类型, 0 TCP; 1 unix domain sockets',
      `address` varchar(255) NOT NULL COMMENT '链接地址',
      `connect_timeout` tinyint(3) UNSIGNED NOT NULL DEFAULT '10' COMMENT '链接超时时间，单位s,默认为10s',
      `read_timeout` tinyint(3) UNSIGNED NOT NULL DEFAULT '5' COMMENT '读超时时间，单位s,默认5s',
      `write_timeout` tinyint(3) UNSIGNED NOT NULL DEFAULT '5' COMMENT '写超时时间，单位s,默认5s',
      `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '实例状态，1开启,0禁用',
      `add_time` date NOT NULL COMMENT '添加时间',
      `update_time` date NOT NULL COMMENT '修改时间',
      `editor_id` int(10) UNSIGNED NOT NULL COMMENT '编辑者ID',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    
