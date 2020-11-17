SET CHARSET UTF8;
use kbcblog-mysql;
SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;

CREATE TABLE IF NOT EXISTS `users`(
  `id` int AUTO_INCREMENT,
  `mail` varchar(100) NOT NULL UNIQUE,
  `passhash` varchar(255) NOT NULL,
  `name` varchar(20) NOT NULL,
  `comment` varchar(150),
  `github` varchar(100),
  `website` varchar(500),
  `languages` varchar(50),
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
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
update `articles` set `created` = CURRENT_TIMESTAMP where `created` is null;
update `articles` set `updated` = CURRENT_TIMESTAMP where `updated` is null;

CREATE TABLE IF NOT EXISTS `questions` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `title` varchar(100) NOT NULL,
  `body` mediumtext NOT NULL,
  `created` datetime,
  `updated` datetime,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
update `questions` set `created` = CURRENT_TIMESTAMP where `created` is null;
update `questions` set `updated` = CURRENT_TIMESTAMP where `updated` is null;

CREATE TABLE IF NOT EXISTS `article_likes` (
  `userid` int NOT NULL,
  `articleid` int NOT NULL,
  `created` datetime,
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`articleid`) REFERENCES `articles`(`id`) ON DELETE CASCADE,
  PRIMARY KEY(`userid`,`articleid`)
);
update `article_likes` set `created` = CURRENT_TIMESTAMP where `created` is null;

CREATE TABLE IF NOT EXISTS `question_likes` (
  `userid` int NOT NULL,
  `questionid` int NOT NULL,
  `created` datetime,
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`questionid`) REFERENCES `questions`(`id`) ON DELETE CASCADE,
  PRIMARY KEY(`userid`,`questionid`)
);
update `question_likes` set `created` = CURRENT_TIMESTAMP where `created` is null;

CREATE TABLE IF NOT EXISTS `article_comments` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `articleid` int NOT NULL,
  `text` varchar(255) NOT NULL,
  `created` datetime,
  `updated` datetime,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`articleid`) REFERENCES `articles`(`id`) ON DELETE CASCADE
);
update `article_comments` set `created` = CURRENT_TIMESTAMP where `created` is null;
update `article_comments` set `updated` = CURRENT_TIMESTAMP where `updated` is null;

CREATE TABLE IF NOT EXISTS `question_comments` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `questionid` int NOT NULL,
  `text` varchar(255) NOT NULL,
  `created` datetime,
  `updated` datetime,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`questionid`) REFERENCES `questions`(`id`) ON DELETE CASCADE
);
update `question_comments` set `created` = CURRENT_TIMESTAMP where `created` is null;
update `question_comments` set `updated` = CURRENT_TIMESTAMP where `updated` is null;