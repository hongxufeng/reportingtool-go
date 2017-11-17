

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `w_user_list`
-- ----------------------------
DROP TABLE IF EXISTS `w_user_list`;
CREATE TABLE `w_user_list` (
  `uid` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户自增uid',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名，登陆可用',
  `password` varchar(255) NOT NULL COMMENT '经salt加密过的密码',
  `salt` smallint(6) NOT NULL,
  `state` tinyint(4) NOT NULL DEFAULT '1' COMMENT '用户状态',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像地址',
  `user_agent` varchar(255) DEFAULT NULL COMMENT '用户登录的user_agent',
  `time_create` datetime DEFAULT NULL COMMENT '注册时间',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='W_后台登录用户表';

-- ----------------------------
-- Records of w_user_list
-- ----------------------------
