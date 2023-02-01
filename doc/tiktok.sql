DROP DATABASE IF EXISTS TIKTOK;
CREATE DATABASE IF NOT EXISTS TIKTOK;
USE TIKTOK;
DROP TABLE IF EXISTS TIKTOK_VIDEO;
CREATE TABLE TIKTOK_VIDEO
(
    `vid`            int UNSIGNED UNIQUE AUTO_INCREMENT NOT NULL COMMENT '视频唯一标识',
    `author`         int UNSIGNED COMMENT '视频作者id',
    `play_url`       VARCHAR(255) COMMENT '视频播放地址',
    `cover_url`      VARCHAR(255) COMMENT '视频封面地址',
    `favorite_count` long COMMENT '视频的点赞总数',
    `comment_count`  long COMMENT '视频的评论总数',
    `is_favorite`    Bool COMMENT '是否点赞',
    `title`          VARCHAR(255) COMMENT '视频标题',
    `next_time`      int UNSIGNED DEFAULT NULL COMMENT '时间',
    PRIMARY KEY (VID)
) COMMENT = '视频';

DROP TABLE IF EXISTS TIKTOK_USER;
CREATE TABLE TIKTOK_USER
(
    `uid`            int UNSIGNED UNIQUE AUTO_INCREMENT NOT NULL COMMENT '用户id',
    `password`       VARCHAR(255) COMMENT '用户密码',
    `name`           VARCHAR(255) COMMENT '用户名称',
    `follow_count`   Long COMMENT '关注总数',
    `follower_count` Long COMMENT '粉丝总数',
    `is_follow`      Bool COMMENT '是否关注',
    `token`          varchar(255) COMMENT '用户授权token',
    `love_video`    text(65535) COMMENT '喜欢视频列表',
    PRIMARY KEY (UID)
) COMMENT = '用户';

DROP TABLE IF EXISTS TIKTOK_COMMENT;
CREATE TABLE TIKTOK_COMMENT
(
    `cid`         int UNSIGNED UNIQUE AUTO_INCREMENT NOT NULL COMMENT '评论id',
    `uid`         int UNSIGNED COMMENT '评论用户id',
    `content`     VARCHAR(255) COMMENT '评论内容',
    `create_date` VARCHAR(255) COMMENT '评论发布日期',
    PRIMARY KEY (CID)
) COMMENT = '评论表';

DROP TABLE IF EXISTS TIKTOK_MESSAGE;
CREATE TABLE TIKTOK_MESSAGE
(
    `mid`         int UNSIGNED UNIQUE AUTO_INCREMENT NOT NULL COMMENT '消息id',
    `content`     VARCHAR(255) COMMENT '消息内容',
    `create_time` VARCHAR(255) COMMENT '消息发送时间',
    PRIMARY KEY (MID)
) COMMENT = '消息';
