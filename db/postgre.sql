/*
Navicat PGSQL Data Transfer

Source Server         : 本地测试
Source Server Version : 101600
Source Host           : localhost:5432
Source Database       : test
Source Schema         : public

Target Server Type    : PGSQL
Target Server Version : 101600
File Encoding         : 65001

Date: 2021-04-05 18:06:11
*/


-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dictionaries";
CREATE TABLE "public"."sys_dictionaries" (
"id" int4 NOT NULL,
"created_at" timestamp(6),
"updated_at" timestamp(6),
"deleted_at" timestamp(6),
"name" varchar(255) COLLATE "default",
"type" varchar(255) COLLATE "default",
"status" int4,
"desc" varchar(255) COLLATE "default"
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO "public"."sys_dictionaries" VALUES ('3', '2020-07-05 15:27:31', '2020-07-05 15:27:31', null, '数据库int类型', 'int', '1', 'int类型对应的数据库类型');
INSERT INTO "public"."sys_dictionaries" VALUES ('4', '2020-07-05 15:33:07', '2020-07-05 16:07:18', null, '数据库时间日期类型', '*time.Time', '1', '数据库时间日期类型');
INSERT INTO "public"."sys_dictionaries" VALUES ('5', '2020-07-05 15:34:23', '2020-07-05 15:52:45', null, '数据库浮点型', 'float64', '1', '数据库浮点型');
INSERT INTO "public"."sys_dictionaries" VALUES ('6', '2020-07-05 15:35:05', '2020-07-05 15:35:05', null, '数据库字符串', 'string', '1', '数据库字符串');
INSERT INTO "public"."sys_dictionaries" VALUES ('7', '2020-07-05 15:36:48', '2020-07-05 15:36:48', null, '数据库bool类型', 'bool', '1', '数据库bool类型');

-- ----------------------------
-- Alter Sequences Owned By
-- ----------------------------

-- ----------------------------
-- Primary Key structure for table sys_dictionaries
-- ----------------------------
ALTER TABLE "public"."sys_dictionaries" ADD PRIMARY KEY ("id");
/*
Navicat PGSQL Data Transfer

Source Server         : 本地测试
Source Server Version : 101600
Source Host           : localhost:5432
Source Database       : test
Source Schema         : public

Target Server Type    : PGSQL
Target Server Version : 101600
File Encoding         : 65001

Date: 2021-04-05 18:06:18
*/


-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dictionary_details";
CREATE TABLE "public"."sys_dictionary_details" (
"id" int4 NOT NULL,
"created_at" timestamp(6),
"updated_at" timestamp(6),
"deleted_at" timestamp(6),
"label" varchar(255) COLLATE "default",
"value" int4,
"status" int4,
"sort" int4,
"sys_dictionary_id" int4
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO "public"."sys_dictionary_details" VALUES ('1', '2020-07-05 15:31:41', '2020-07-05 15:31:41', null, 'int2', '1', '1', '1', '3');
INSERT INTO "public"."sys_dictionary_details" VALUES ('2', '2020-07-05 15:31:52', '2020-07-05 15:31:52', null, 'int4', '2', '1', '2', '3');
INSERT INTO "public"."sys_dictionary_details" VALUES ('3', '2020-07-05 15:34:35', '2020-07-05 15:34:35', null, 'float4', '1', '1', '1', '5');
INSERT INTO "public"."sys_dictionary_details" VALUES ('4', '2020-07-05 15:32:11', '2020-07-05 15:32:11', null, 'int8', '4', '1', '4', '3');
INSERT INTO "public"."sys_dictionary_details" VALUES ('5', '2020-07-05 15:34:35', '2020-07-05 15:34:35', null, 'float8', '1', '1', '1', '5');
INSERT INTO "public"."sys_dictionary_details" VALUES ('6', '2020-07-05 15:33:21', '2020-07-05 15:33:21', null, 'time', '1', '1', '1', '4');
INSERT INTO "public"."sys_dictionary_details" VALUES ('7', '2020-07-05 15:33:25', '2020-07-05 15:33:25', null, 'timetz', '2', '1', '2', '4');
INSERT INTO "public"."sys_dictionary_details" VALUES ('8', '2020-07-05 15:33:35', '2020-07-05 15:33:35', null, 'timestamptz', '3', '1', '3', '4');
INSERT INTO "public"."sys_dictionary_details" VALUES ('9', '2020-07-05 15:33:42', '2020-07-05 15:33:42', null, 'timestamp', '5', '1', '5', '4');
INSERT INTO "public"."sys_dictionary_details" VALUES ('10', '2020-07-05 15:34:35', '2020-07-05 15:34:35', null, 'money', '1', '1', '1', '5');
INSERT INTO "public"."sys_dictionary_details" VALUES ('11', '2020-07-05 15:34:35', '2020-07-05 15:34:35', null, 'numeric', '1', '1', '1', '5');
INSERT INTO "public"."sys_dictionary_details" VALUES ('12', '2020-07-05 15:34:41', '2020-07-05 15:34:41', null, 'decimal', '2', '1', '2', '5');
INSERT INTO "public"."sys_dictionary_details" VALUES ('13', '2020-07-05 15:37:45', '2020-07-05 15:37:45', null, 'bool', '0', '1', '0', '7');
INSERT INTO "public"."sys_dictionary_details" VALUES ('14', '2020-07-05 15:53:25', '2020-07-05 15:53:25', null, 'char', '0', '1', '0', '6');
INSERT INTO "public"."sys_dictionary_details" VALUES ('15', '2020-07-05 15:53:29', '2020-07-05 15:53:29', null, 'varchar', '1', '1', '1', '6');
INSERT INTO "public"."sys_dictionary_details" VALUES ('16', '2020-07-05 15:53:29', '2020-07-05 15:53:29', null, 'geometry', '1', '1', '1', '6');
INSERT INTO "public"."sys_dictionary_details" VALUES ('18', '2020-07-05 15:53:48', '2020-07-05 15:53:48', null, 'text', '4', '1', '4', '6');
INSERT INTO "public"."sys_dictionary_details" VALUES ('20', '2020-07-05 15:54:02', '2020-07-05 15:54:02', null, 'date', '6', '1', '6', '4');

-- ----------------------------
-- Alter Sequences Owned By
-- ----------------------------

-- ----------------------------
-- Primary Key structure for table sys_dictionary_details
-- ----------------------------
ALTER TABLE "public"."sys_dictionary_details" ADD PRIMARY KEY ("id");
