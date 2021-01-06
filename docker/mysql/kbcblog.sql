SET CHARSET UTF8;
use kbcblog-mysql;
SET CHARACTER_SET_CLIENT = utf8;
SET CHARACTER_SET_CONNECTION = utf8;

CREATE TABLE IF NOT EXISTS `users`(
  `id` int AUTO_INCREMENT,
  `mail` varchar(100) NOT NULL UNIQUE,
  `passhash` varchar(255) NOT NULL,
  `name` varchar(10) NOT NULL,
  `comment` varchar(150),
  `github` varchar(100),
  `website` varchar(500),
  `languages` varchar(50),
  -- `imgpath` varchar(200),
  `imgfile` mediumblob,
  `sex` int, -- 1:男 2:女
  PRIMARY KEY(`id`)
);

CREATE TABLE IF NOT EXISTS `articles` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `title` varchar(100) NOT NULL,
  `body` mediumtext NOT NULL,
  `tag`  varchar(120) NOT NULL,
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
  `tag`  varchar(120) NOT NULL,
  `category`  varchar(10) NOT NULL,
  `created` datetime,
  `updated` datetime,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
update `questions` set `created` = CURRENT_TIMESTAMP where `created` is null;
update `questions` set `updated` = CURRENT_TIMESTAMP where `updated` is null;

CREATE TABLE IF NOT EXISTS `article_likes` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `articleid` int NOT NULL,
  `created` datetime,
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`articleid`) REFERENCES `articles`(`id`) ON DELETE CASCADE,
  PRIMARY KEY(`id`),
);
update `article_likes` set `created` = CURRENT_TIMESTAMP where `created` is null;

CREATE TABLE IF NOT EXISTS `question_likes` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `questionid` int NOT NULL,
  `created` datetime,
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`questionid`) REFERENCES `questions`(`id`) ON DELETE CASCADE,
  PRIMARY KEY(`id`),
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

CREATE TABLE IF NOT EXISTS `article_comment_likes` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `article_commentid` int NOT NULL,
  `created` datetime,
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`article_commentid`) REFERENCES `article_comments`(`id`) ON DELETE CASCADE,
  PRIMARY KEY(`id`)
);
update `article_comment_likes` set `created` = CURRENT_TIMESTAMP where `created` is null;

CREATE TABLE IF NOT EXISTS `question_comment_likes` (
  `id` int AUTO_INCREMENT,
  `userid` int NOT NULL,
  `question_commentid` int NOT NULL,
  `created` datetime,
  FOREIGN KEY (`userid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`question_commentid`) REFERENCES `question_comments`(`id`) ON DELETE CASCADE,
  PRIMARY KEY(`id`)
);
update `question_comment_likes` set `created` = CURRENT_TIMESTAMP where `created` is null;

CREATE TABLE IF NOT EXISTS `follows` (
  `id` int AUTO_INCREMENT,
  `followerid` int NOT NULL,
  `followedid` int NOT NULL,
  `created` datetime,
  FOREIGN KEY (`followerid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`followedid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  PRIMARY KEY(`id`)
);
update `follows` set `created` = CURRENT_TIMESTAMP where `created` is null;

CREATE TABLE IF NOT EXISTS `notifications` (
  `id` int AUTO_INCREMENT,
  `articleid` int,
  `questionid` int,
  `visiterid` int NOT NULL,
  `visitedid` int NOT NULL,
  `a_commentid` int,
  `q_commentid` int,
  `action` varchar(15) NOT NULL,
  `checked` boolean NOT NULL default false,
  `created` datetime,
  `updated` datetime,
  PRIMARY KEY(`id`),
  FOREIGN KEY (`articleid`) REFERENCES `articles`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`questionid`) REFERENCES `questions`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`visiterid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`visitedid`) REFERENCES `users`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`a_commentid`) REFERENCES `article_comments`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`q_commentid`) REFERENCES `question_comments`(`id`) ON DELETE CASCADE,
);
update `notifications` set `created` = CURRENT_TIMESTAMP where `created` is null;
update `notifications` set `updated` = CURRENT_TIMESTAMP where `updated` is null;