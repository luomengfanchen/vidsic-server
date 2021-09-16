/*
    数据库
*/
CREATE DATABASE vidsic

/*
    用户信息表
*/
CREATE TABLE user_t (
    id          INT             AUTO_INCREMENT PRIMARY KEY,
    nickname    VARCHAR(255)    NOT NULL,
    email       VARCHAR(255)    UNIQUE,
    avator      VARCHAR(255)    DEFAULT '/api/static/avator/default_avator.png',
    password    VARCHAR(255)    NOT NULL,
    competence  VARCHAR(64)     DEFAULT 'user',
    birthday    DATE            NOT NULL,
    intro       VARCHAR(255)    NOT NULL,
    token       VARCHAR(255)    UNIQUE
)AUTO_INCREMENT=1000;

/*
    专栏文件信息表
*/
CREATE TABLE article_t (
    id      INT             AUTO_INCREMENT PRIMARY KEY,
    name    VARCHAR(255)    NOT NULL,
    date    DATETIME        NOT NULL,
    author  INT             NOT NULL,
    cover   VARCHAR(255)    DEFAULT '/api/static/image/article_cover.jpg',
    text    VARCHAR(255)    NOT NULL,
    views   INT             DEFAULT 0
)AUTO_INCREMENT=1000;

/*
    图像文件信息表
*/
CREATE TABLE image_t (
    id      INT             AUTO_INCREMENT PRIMARY KEY,
    name    VARCHAR(255)    NOT NULL,
    date    DATETIME        NOT NULL,
    author  INT             NOT NULL,
    path    VARCHAR(255)    NOT NULL
)AUTO_INCREMENT=1000;

/*
    音乐文件信息表
*/
CREATE TABLE music_t (
    id          INT             AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255)    NOT NULL,
    descript    VARCHAR(255)    NOT NULL,
    date        DATETIME        NOT NULL,
    singer      VARCHAR(255)    NOT NULL,
    cover       VARCHAR(255)    DEFAULT '/api/static/image/music_cover.jpg',
    type        VARCHAR(255)    NOT NULL,
    path        VARCHAR(255)    NOT NULL,
    views       INT             DEFAULT 0
)AUTO_INCREMENT=1000;

/*
    视频文件信息表
*/
CREATE TABLE video_t (
    id          INT             AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255)    NOT NULL,
    descript    VARCHAR(255)    NOT NULL,
    date        DATETIME        NOT NULL,
    author      INT             NOT NULL,
    cover       VARCHAR(255)    DEFAULT '/api/static/image/video_cover.jpg',
    type        VARCHAR(255)    NOT NULL,
    path        VARCHAR(255)    NOT NULL,
    views       INT             DEFAULT 0
)AUTO_INCREMENT=1000;

/*
    提交队列表
*/
CREATE TABLE commit_t (
    id          INT             AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255)    NOT NULL,
    date        DATETIME        NOT NULL,
    path        VARCHAR(255)    NOT NULL,
    type        VARCHAR(255)    NOT NULL,
    status      BOOLEAN         DEFAULT false,
    commiter    INT             NOT NULL
)AUTO_INCREMENT=1000;
