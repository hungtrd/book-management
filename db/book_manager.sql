/*
 Navicat Premium Data Transfer

 Source Server         : revel_db
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : localhost:3307
 Source Schema         : demo

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 04/12/2022 03:07:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `author` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `release_date` date DEFAULT NULL,
  `total_page` int(11) DEFAULT NULL,
  `category` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of books
-- ----------------------------
BEGIN;
INSERT INTO `books` VALUES (1, 'Naruto', 'Naruto là một loạt manga Nhật Bản được viết và minh họa bởi Kishimoto Masashi', 'Kishimoto Masashi', '1999-09-21', 101, 2, '2022-11-29 16:59:00', '2022-12-03 19:15:50');
INSERT INTO `books` VALUES (2, 'Doraemon', ' Trong truyện lấy bối cảnh ở thế kỷ 22, Doraemon là chú mèo robot của tương lai do xưởng Matsushiba', 'Fujiko F. Fujio', '2005-04-15', 50, 2, '2022-11-29 17:09:46', '2022-11-29 17:09:46');
INSERT INTO `books` VALUES (6, 'One Piece', 'One Piece, từng được xuất bản tại Việt Nam dưới tên gọi Đảo Hải Tặc là bộ manga dành cho lứa tuổi thiếu niên của tác giả Oda Eiichiro, được đăng định kì trên tạp chí Weekly Shōnen Jump', 'Oda Eiichiro', '1997-07-19', 1, 2, '2022-12-03 19:17:41', '2022-12-03 19:17:41');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fullname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'hungtran', 'Hung Tran', '$2a$10$uZp0msZFdhknFSaNGvF.x.MhcTBbjBwoNJNwI6lFRp5bX5TXT63G.', 'hungtran@gmail.com', '2022-12-03 17:34:20', '2022-12-03 17:34:20');
INSERT INTO `users` VALUES (2, 'hungtran2', 'Hung Tran 2', '$2a$10$I3WsG03BpfAlg/9h5BgB.e/4zlR/Q3PE9Bvs9ghM8YJEd3dOwoVP6', 'hungtran2@gmail.com', '2022-12-03 19:27:01', '2022-12-03 19:27:01');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
