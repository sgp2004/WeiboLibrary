drop database  `weibo_library`;
CREATE DATABASE IF NOT EXISTS `weibo_library`;
USE `weibo_library`;

CREATE TABLE IF NOT EXISTS `user` (
 `id` bigint(20) unsigned NOT NULL auto_increment,
  `username` char(50) NOT NULL,
  `name`  char(50),
  `own_book_count` int(10) signed NOT NULL DEFAULT 0,
  `borrow_book_count` int(10) signed NOT NULL DEFAULT 0,
  `allow_book_count` int(10) signed NOT NULL DEFAULT 0,
  `allow_borrow_days` tinyint(10) signed NOT NULL DEFAULT 30,
  `user_level`  tinyint(10) signed NOT NULL DEFAULT 0,
  PRIMARY KEY `username`(`username`),KEY `name` (`name`),KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='user';

CREATE TABLE IF NOT EXISTS `books` (
  `id` bigint(20) unsigned NOT NULL auto_increment,
  `username` char(50) NOT NULL,
  `name` char(100) NOT NULL,
  `author` char(50) ,
  `info` text,
  `img` char(200) ,
  `modified_time`  timestamp NOT NULL default  CURRENT_TIMESTAMP,
  `count` int(20) unsigned DEFAULT 1,
  `status` tinyint(1) unsigned DEFAULT 0,  /* can be borrow 0; borrowed 1; forbidden 2*/
  PRIMARY KEY `id` (`id`),KEY `author_index` (`author`),KEY `user_index` (`username`),KEY `book_index` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='books';


CREATE TABLE IF NOT EXISTS `borrow_info` (
  `id` bigint(20) unsigned NOT NULL auto_increment,
  `book_id` bigint(20) unsigned  NOT NULL,
  `owner_id` char(200) NOT NULL,
  `borrower_id` char(50) NOT NULL,
  `note` text NOT NULL default "",
  `start_time`  timestamp NOT NULL default  CURRENT_TIMESTAMP,
  `end_time` timestamp ,
  `is_back` tinyint(1) unsigned DEFAULT 0, /*0 not back , 1 back */
  PRIMARY KEY `id` (`id`),KEY `book_index` (`book_id`),KEY `owner_index` (`owner_id`),KEY `borrow_index` (`borrower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='borrow_info';

