# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.17)
# Database: goblog
# Generation Time: 2019-12-05 06:29:45 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table hello
# ------------------------------------------------------------

DROP TABLE IF EXISTS `hello`;

CREATE TABLE `hello` (
  `world` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_option
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_option`;

CREATE TABLE `t_option` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL DEFAULT '',
  `value` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_option` WRITE;
/*!40000 ALTER TABLE `t_option` DISABLE KEYS */;

INSERT INTO `t_option` (`id`, `name`, `value`)
VALUES
	(1,'sitename','博客演示'),
	(2,'siteurl','http://127.0.0.1:8080/'),
	(3,'subtitle','基于Go语言和Beego框架的博客系统'),
	(4,'pagesize','10'),
	(5,'keywords','Go语言,博客程序,GoBlog'),
	(6,'description','基于Go语言和Beego框架的博客系统'),
	(7,'email',''),
	(8,'theme','default'),
	(9,'timezone','8'),
	(10,'stat','');

/*!40000 ALTER TABLE `t_option` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_post`;

CREATE TABLE `t_post` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `author` varchar(15) NOT NULL DEFAULT '' COMMENT '作者',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
  `color` varchar(7) NOT NULL DEFAULT '' COMMENT '标题颜色',
  `url_name` varchar(100) NOT NULL DEFAULT '' COMMENT 'url名',
  `url_type` tinyint(3) NOT NULL DEFAULT '0' COMMENT 'url访问形式',
  `content` mediumtext NOT NULL COMMENT '内容',
  `tags` varchar(100) NOT NULL DEFAULT '' COMMENT '标签',
  `views` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '查看次数',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态{0:正常,1:草稿,2:回收站}',
  `post_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  `update_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `is_top` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否置顶',
  PRIMARY KEY (`id`),
  KEY `userid` (`user_id`),
  KEY `posttime` (`post_time`),
  KEY `urlname` (`url_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_post` WRITE;
/*!40000 ALTER TABLE `t_post` DISABLE KEYS */;

INSERT INTO `t_post` (`id`, `user_id`, `author`, `title`, `color`, `url_name`, `url_type`, `content`, `tags`, `views`, `status`, `post_time`, `update_time`, `is_top`)
VALUES
	(1,1,'admin','关于我','','about',1,'<p>个人简介</p>','',7,0,'2013-12-31 10:27:49','2013-12-31 10:27:53',0),
	(2,1,'admin','友情链接','','links',1,'','',6,0,'2014-01-01 02:29:54','2017-11-23 22:09:04',0),
	(3,1,'admin','友情链接','','links',1,'','',6,0,'2014-01-01 02:29:54','2017-11-23 22:09:04',0),
	(4,1,'admin','友情链接','','links',1,'','',6,0,'2014-01-01 02:29:54','2017-11-23 22:09:04',0),
	(5,1,'admin','友情链接','','links',1,'','',6,0,'2014-01-01 02:29:54','2017-11-23 22:09:04',0),
	(6,1,'admin','友情链接','','links',1,'','',6,0,'2014-01-01 02:29:54','2017-11-23 22:09:04',0),
	(7,1,'admin','友情链接','','links',1,'','',6,0,'2014-01-01 02:29:54','2017-11-23 22:09:04',0),
	(8,1,'admin','友情链接','','links',1,'','',6,0,'2014-01-01 02:29:54','2017-11-23 22:09:04',0),
	(10,1,'admin','测试标题','','',0,'<p>这是测试</p>','',0,0,'0000-00-00 00:00:00','2019-12-04 16:39:56',1),
	(11,1,'admin','','','',0,'<p>这是测试</p>','',0,0,'0000-00-00 00:00:00','2019-05-10 17:30:25',1),
	(12,1,'admin','测试标题','','',0,'<p>这是测试</p>','',0,0,'0000-00-00 00:00:00','2019-05-10 17:31:02',1),
	(13,1,'admin','测试标题','','',0,'<p>这是测试</p>','',0,0,'0000-00-00 00:00:00','2019-12-04 16:39:47',1);

/*!40000 ALTER TABLE `t_post` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_tag
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_tag`;

CREATE TABLE `t_tag` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '' COMMENT '标签名',
  `count` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '使用次数',
  PRIMARY KEY (`id`),
  KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_tag` WRITE;
/*!40000 ALTER TABLE `t_tag` DISABLE KEYS */;

INSERT INTO `t_tag` (`id`, `name`, `count`)
VALUES
	(1,'go',1);

/*!40000 ALTER TABLE `t_tag` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table t_tag_post
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_tag_post`;

CREATE TABLE `t_tag_post` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '标签id',
  `post_id` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '内容id',
  `post_status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '内容状态',
  `post_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '发布时间',
  PRIMARY KEY (`id`),
  KEY `tagid` (`tag_id`),
  KEY `postid` (`post_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table t_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `t_user`;

CREATE TABLE `t_user` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(15) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(15) NOT NULL COMMENT '手机号',
  `login_count` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `last_ip` varchar(15) NOT NULL DEFAULT '0' COMMENT '最后登录ip',
  `last_login` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最后登录时间',
  `authkey` char(10) NOT NULL DEFAULT '' COMMENT '登录key',
  `active` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否激活',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `t_user` WRITE;
/*!40000 ALTER TABLE `t_user` DISABLE KEYS */;

INSERT INTO `t_user` (`id`, `username`, `password`, `email`, `phone`, `login_count`, `last_ip`, `last_login`, `authkey`, `active`)
VALUES
	(1,'测试1','','','1595203910',0,'127.0.0.1','2018-01-25 23:07:21','',1);

/*!40000 ALTER TABLE `t_user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
