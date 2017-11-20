SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `w_user_list`
-- ----------------------------
DROP TABLE IF EXISTS `w_user_list`;
CREATE TABLE `w_user_list` (
  `uid` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` char(50) NOT NULL DEFAULT '' COMMENT '用户名，登陆可用',
  `password` char(50) NOT NULL COMMENT '经salt加密过的密码',
  `nickname` varchar(25) DEFAULT NULL,
  `salt` int(6) NOT NULL,
  `state` tinyint(1) NOT NULL COMMENT '用户状态',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像地址',
  `user_agent` varchar(255) NOT NULL COMMENT '用户登录的user_agent',
  `time_create` datetime DEFAULT NULL COMMENT '注册时间',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='W_后台登录用户表';

-- ----------------------------
-- Records of w_user_list
-- ----------------------------
INSERT INTO `w_user_list` VALUES ('1', 'moon', '1f0c428a6bfe454bb1115805cc0cdaaa', '月儿', '211314', '0', 'assets/img/avatar/Sora.jpg', 'Mozilla', null);
