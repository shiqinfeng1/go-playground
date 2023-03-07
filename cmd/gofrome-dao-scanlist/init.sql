
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;


DROP TABLE IF EXISTS `test_cert_info`;
CREATE TABLE `test_cert_info`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `version` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
  `identity` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户定义的证书标识',
  `cert_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证书类型',
  `issuer_identity` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '颁发机构标识',
  `subject` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证书主体信息',
  `validity_notbefore` datetime NULL DEFAULT NULL COMMENT '证书有效期开始',
  `validity_notafter` datetime NULL DEFAULT NULL COMMENT '证书有效期结束',
  `customer_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名称',
  `customer_idcard` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户身份证号',
  `customer_id` bigint UNSIGNED NOT NULL COMMENT '客户id',
  `serial_number` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证书序列号',
  `status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '证书状态',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `test_pci_serial_identity`(`identity` ASC) USING BTREE,
  UNIQUE INDEX `test_pci_serial_number`(`serial_number` ASC) USING BTREE,
  INDEX `test_pci_customer_id`(`customer_id` ASC) USING BTREE,
  INDEX `test_pci_customer_name`(`customer_name` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试证书信息表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of test_cert_info
-- ----------------------------
-- 申请账号id=1，个人证书identity11， 持有者customer_name1。下面有一个过期证书，一个吊销证书， 一个可用证书
INSERT INTO `test_cert_info` VALUES (1, 1, 'identity11', '个人', 'issuer_identity1', 'subject1', '2022-11-02 18:33:54', '2028-11-02 18:33:54','customer_name1','customer_idcard1',1,'serial_number11','已注销', '2022-11-02 18:33:54', '2022-11-02 18:33:54');
INSERT INTO `test_cert_info` VALUES (2, 1, 'identity11', '个人', 'issuer_identity1', 'subject1', '2022-11-02 18:33:54', '2028-11-02 18:33:54','customer_name1','customer_idcard1',1,'serial_number12','已过期', '2022-11-03 18:33:54', '2022-11-02 18:33:54');
INSERT INTO `test_cert_info` VALUES (3, 1, 'identity11', '个人', 'issuer_identity1', 'subject1', '2023-11-02 18:33:54', '2024-11-02 18:33:54','customer_name1','customer_idcard1',1,'serial_number13','可用', '2022-11-04 18:33:54', '2022-11-02 18:33:54');
-- 申请账号id=2，个人证书identity21，持有者customer_name2。下面有一个可用证书
INSERT INTO `test_cert_info` VALUES (4, 1, 'identity21', '个人', 'issuer_identity2', 'subject2', '2023-11-02 18:33:54', '2024-11-02 18:33:54','customer_name2','customer_idcard2',2,'serial_number21','可用', '2022-11-02 18:33:54', '2022-11-02 18:33:54');
-- 申请账号id=3，个人证书identity31，持有者customer_name3。下面有一个可用证书
INSERT INTO `test_cert_info` VALUES (5, 1, 'identity31', '个人', 'issuer_identity2', 'subject2', '2023-11-02 18:33:54', '2024-11-02 18:33:54','customer_name3','customer_idcard3',3,'serial_number31','可用', '2022-11-02 18:33:54', '2022-11-02 18:33:54');

-- 申请账号id=1，企业证书identity12，持有者customer_name5。下面有一个可用证书
INSERT INTO `test_cert_info` VALUES (6, 1, 'identity12', '企业', 'issuer_identity3', 'subject3', '2023-11-02 18:33:54', '2024-11-02 18:33:54','customer_name5','customer_idcard5',1,'serial_number14','可用', '2022-11-02 18:33:54', '2022-11-02 18:33:54');
-- 申请账号id=1，企业证书identity13，持有者customer_name5。下面有一个可用证书
INSERT INTO `test_cert_info` VALUES (7, 1, 'identity13', '域名', 'issuer_identity3', 'subject3', '2023-11-02 18:33:54', '2024-11-02 18:33:54','customer_name5','customer_idcard5',1,'serial_number15','可用', '2022-11-02 18:33:54', '2022-11-02 18:33:54');

-- ----------------------------
-- Records of test_cert_deviceId
-- ----------------------------
DROP TABLE IF EXISTS `test_cert_deviceId`;
CREATE TABLE `test_cert_deviceId`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `customer_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',
  `identity` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户定义的证书标识',
  `deviceId` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '设备号',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `test_cert_deviceId_identity`(`identity` ASC) USING BTREE,
  INDEX `test_cert_deviceId_deviceId`(`deviceId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '测试证书设备绑定表' ROW_FORMAT = COMPACT;

-- 申请账号id=1，个人证书identity11，下载到设备deviceId1，deviceId2，deviceId3
INSERT INTO `test_cert_deviceId` VALUES (1, 1,'identity11', 'deviceId1', '2022-11-02 18:33:54', '2022-11-02 18:33:54');
INSERT INTO `test_cert_deviceId` VALUES (2, 1,'identity11', 'deviceId2', '2022-11-02 18:33:54', '2022-11-02 18:33:54');
INSERT INTO `test_cert_deviceId` VALUES (3, 1,'identity11', 'deviceId3', '2022-11-02 18:33:54', '2022-11-02 18:33:54');
-- 申请账号id=2，个人证书identity21，下载到设备deviceId1
INSERT INTO `test_cert_deviceId` VALUES (4, 2,'identity21', 'deviceId1', '2022-11-02 18:33:54', '2022-11-02 18:33:54');
-- 申请账号id=1，企业证书identity12，下载到设备deviceId1
INSERT INTO `test_cert_deviceId` VALUES (5, 1,'identity12', 'deviceId1', '2022-11-02 18:33:54', '2022-11-02 18:33:54');
-- 申请账号id=1，企业证书identity13，下载到设备deviceId1
INSERT INTO `test_cert_deviceId` VALUES (6, 1,'identity13', 'deviceId1', '2022-11-02 18:33:54', '2022-11-02 18:33:54');



SET FOREIGN_KEY_CHECKS = 1;