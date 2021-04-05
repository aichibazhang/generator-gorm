
/*
Navicat MySQL Data Transfer

Source Server         : 本地测试
Source Server Version : 50728
Source Host           : localhost:3306
Source Database       : go-model

Target Server Type    : MYSQL
Target Server Version : 50728
File Encoding         : 65001

Date: 2021-04-05 18:07:50
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `label` varchar(255) DEFAULT NULL COMMENT '展示值',
  `value` int(11) DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` int(11) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` int(11) DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES ('12', '2020-07-05 15:31:41', '2020-07-05 15:31:41', null, 'smallint', '1', '1', '1', '3');
INSERT INTO `sys_dictionary_details` VALUES ('13', '2020-07-05 15:31:52', '2020-07-05 15:31:52', null, 'mediumint', '2', '1', '2', '3');
INSERT INTO `sys_dictionary_details` VALUES ('14', '2020-07-05 15:32:04', '2020-07-05 15:32:04', null, 'int', '3', '1', '3', '3');
INSERT INTO `sys_dictionary_details` VALUES ('15', '2020-07-05 15:32:11', '2020-07-05 15:32:11', null, 'bigint', '4', '1', '4', '3');
INSERT INTO `sys_dictionary_details` VALUES ('19', '2020-07-05 15:33:16', '2020-07-05 15:33:16', null, 'data', '0', '1', '0', '4');
INSERT INTO `sys_dictionary_details` VALUES ('20', '2020-07-05 15:33:21', '2020-07-05 15:33:21', null, 'time', '1', '1', '1', '4');
INSERT INTO `sys_dictionary_details` VALUES ('21', '2020-07-05 15:33:25', '2020-07-05 15:33:25', null, 'year', '2', '1', '2', '4');
INSERT INTO `sys_dictionary_details` VALUES ('22', '2020-07-05 15:33:35', '2020-07-05 15:33:35', null, 'datetime', '3', '1', '3', '4');
INSERT INTO `sys_dictionary_details` VALUES ('23', '2020-07-05 15:33:42', '2020-07-05 15:33:42', null, 'timestamp', '5', '1', '5', '4');
INSERT INTO `sys_dictionary_details` VALUES ('24', '2020-07-05 15:34:30', '2020-07-05 15:34:30', null, 'float', '0', '1', '0', '5');
INSERT INTO `sys_dictionary_details` VALUES ('25', '2020-07-05 15:34:35', '2020-07-05 15:34:35', null, 'double', '1', '1', '1', '5');
INSERT INTO `sys_dictionary_details` VALUES ('26', '2020-07-05 15:34:41', '2020-07-05 15:34:41', null, 'decimal', '2', '1', '2', '5');
INSERT INTO `sys_dictionary_details` VALUES ('27', '2020-07-05 15:37:45', '2020-07-05 15:37:45', null, 'tinyint', '0', '1', '0', '7');
INSERT INTO `sys_dictionary_details` VALUES ('28', '2020-07-05 15:53:25', '2020-07-05 15:53:25', null, 'char', '0', '1', '0', '6');
INSERT INTO `sys_dictionary_details` VALUES ('29', '2020-07-05 15:53:29', '2020-07-05 15:53:29', null, 'varchar', '1', '1', '1', '6');
INSERT INTO `sys_dictionary_details` VALUES ('30', '2020-07-05 15:53:35', '2020-07-05 15:53:35', null, 'tinyblob', '2', '1', '2', '6');
INSERT INTO `sys_dictionary_details` VALUES ('31', '2020-07-05 15:53:40', '2020-07-05 15:53:40', null, 'tinytext', '3', '1', '3', '6');
INSERT INTO `sys_dictionary_details` VALUES ('32', '2020-07-05 15:53:48', '2020-07-05 15:53:48', null, 'text', '4', '1', '4', '6');
INSERT INTO `sys_dictionary_details` VALUES ('33', '2020-07-05 15:53:55', '2020-07-05 15:53:55', null, 'blob', '5', '1', '5', '6');
INSERT INTO `sys_dictionary_details` VALUES ('34', '2020-07-05 15:54:02', '2020-07-05 15:54:02', null, 'date', '6', '1', '6', '4');
INSERT INTO `sys_dictionary_details` VALUES ('35', '2020-07-05 15:54:09', '2020-07-05 15:54:09', null, 'mediumtext', '7', '1', '7', '6');
INSERT INTO `sys_dictionary_details` VALUES ('36', '2020-07-05 15:54:16', '2020-07-05 15:54:16', null, 'longblob', '8', '1', '8', '6');
INSERT INTO `sys_dictionary_details` VALUES ('37', '2020-07-05 15:54:24', '2020-07-05 15:54:24', null, 'longtext', '9', '1', '9', '6');
/*
Navicat MySQL Data Transfer

Source Server         : 本地测试
Source Server Version : 50728
Source Host           : localhost:3306
Source Database       : go-model

Target Server Type    : MYSQL
Target Server Version : 50728
File Encoding         : 65001

Date: 2021-04-05 18:07:43
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(255) DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES ('3', '2020-07-05 15:27:31', '2020-07-05 15:27:31', null, '数据库int类型', 'int', '1', 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES ('4', '2020-07-05 15:33:07', '2020-07-05 16:07:18', null, '数据库时间日期类型', '*time.Time', '1', '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES ('5', '2020-07-05 15:34:23', '2020-07-05 15:52:45', null, '数据库浮点型', 'float64', '1', '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES ('6', '2020-07-05 15:35:05', '2020-07-05 15:35:05', null, '数据库字符串', 'string', '1', '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES ('7', '2020-07-05 15:36:48', '2020-07-05 15:36:48', null, '数据库bool类型', 'bool', '1', '数据库bool类型');
