# t_goods
CREATE TABLE `t_goods`  (
    `pk_id` varchar(255) NOT NULL,
    `goods_name` varchar(255) NULL,
    `price` bigint NULL,
    `sale_num` int NULL,
    `stock` int NULL,
    `is_seckill` tinyint(2) NULL,
    `seckill_time` bigint NULL,
    `version` int NULL,
    `update_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`pk_id`)
);
# t_userinfo
CREATE TABLE `t_user_info`  (
    `pk_id` varchar(255) NOT NULL,
    `user_name` varchar(255) NULL,
    `password` varchar(255) NULL,
    `token` varchar(255) NULL,
    `create_time` datetime NULL,
    PRIMARY KEY (`pk_id`)
);
CREATE TABLE `t_user_address`  (
  `pk_id` varchar(255) NOT NULL,
  `fk_user_id` varchar(255) NULL,
  `province` varchar(255) NULL,
  `city` varchar(255) NULL,
  `detail` varchar(255) NULL,
  `create_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`pk_id`)
);
CREATE TABLE `t_user_fund`  (
  `pk_id` varchar(255) NOT NULL,
  `fk_user_id` varchar(255) NULL,
  `balance` bigint NULL,
  `version` int(11) NULL,
  `update_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`pk_id`)
);

CREATE TABLE `t_user_fund_record`  (
  `pk_id` varchar(255) NOT NULL,
  `fk_user_id` varchar(255) NULL,
  `amount` bigint NULL,
  `type` int(11) NULL,
  `source` int(11) NULL,
  `create_time` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`pk_id`)
);

CREATE TABLE `t_order`  (
  `pk_id` varchar(255) NOT NULL,
  `fk_good_id` varchar(255) NULL,
  `fk_user_id` varchar(255) NULL,
  `goods_name` varchar(255) NULL,
  `price` bigint(10) NULL,
  `num` int(11) NULL,
  `total_price` bigint(10) NULL,
  `user_name` varchar(255) NULL,
  `delivery_address` varchar(255) NULL,
  `pay_time` datetime NULL,
  `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`pk_id`)
);