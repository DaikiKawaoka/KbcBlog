SET CHARSET UTF8;
use kbcblog-mysql;
SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;

CREATE TABLE IF NOT EXISTS `users`(
  `id` int AUTO_INCREMENT,
  `mail` varchar(100) NOT NULL,
  `passhash` varchar(255) NOT NULL,
  `mailname`char(10) NOT NULL,
  `name` varchar(20) NOT NULL,

  PRIMARY KEY(`id`)
);

CREATE TABLE IF NOT EXISTS `articles` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `title` varchar(100) NOT NULL,
  `body` mediumtext NOT NULL,
  `created` datetime,
  `updated` datetime,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`)
);
update `articles` set `created` = CURRENT_TIMESTAMP where `created` is null;
update `articles` set `updated` = CURRENT_TIMESTAMP where `updated` is null;

CREATE TABLE IF NOT EXISTS `favorites` (
  `userid` int NOT NULL,
  `articleid` int NOT NULL,
  `created` datetime,
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`),
  FOREIGN KEY (`articleid`) REFERENCES `articles`(`id`),
  PRIMARY KEY(`userid`,`articleid`)
);
update `favorites` set `created` = CURRENT_TIMESTAMP where `created` is null;

CREATE TABLE IF NOT EXISTS `comments` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `articleid` int NOT NULL,
  `text` varchar(255) NOT NULL,
  `created` datetime,
  `updated` datetime,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`),
  FOREIGN KEY (`articleid`) REFERENCES `articles`(`id`)
);
update `comments` set `created` = CURRENT_TIMESTAMP where `created` is null;
update `comments` set `updated` = CURRENT_TIMESTAMP where `updated` is null;