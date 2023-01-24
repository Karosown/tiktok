create database if not exists TikTok;
-- 用户
create table if not exists TikTok.`user`
(
    `id` varchar(256) not null comment 'id' primary key,
    `userAccount` varchar(256) not null unique comment '账号',
    `userName` varchar(256) null comment '用户昵称',
    `userPassword` varchar(512) not null comment '密码',
    `userAvatar` varchar(1024) null comment '用户头像',
    `gender` tinyint null comment '性别',
    `followCount` bigint default 0 null comment '关注数',
    `flowerCount` bigint default 0 null comment '粉丝数',
    `updateTime` datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `createTime` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `isDelete` tinyint default 0 not null comment '是否删除'
) comment '用户';
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('DOS', '黎思淼', '侯昊焱', 'jNDQ', 'www.napoleon-boyer.net', 1, 400323, 118514, '2022-12-29 09:50:44', '2022-07-29 13:18:13');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('acT', '范潇然', '赖鑫磊', '2hyv', 'www.shandi-hauck.io', 1, 99253, 3469, '2022-04-14 10:10:11', '2022-03-02 23:19:39');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('2Mxh', '吕鑫鹏', '何乐驹', 'PIPRl', 'www.kathy-turcotte.info', 1, 6690, 788951, '2022-11-27 18:02:56', '2022-06-19 01:51:41');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('Tmi2u', '赵正豪', '戴智辉', 'biz', 'www.huey-cruickshank.info', 0, 4, 8645540, '2022-05-05 05:54:52', '2022-01-01 23:37:46');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('LMxAo', '程智辉', '陆修杰', 'MbEZg', 'www.noel-waelchi.net', 1, 3195, 60231779, '2022-11-16 04:07:53', '2022-07-19 17:03:00');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('3Q', '熊子涵', '孙健雄', 'tnBe', 'www.donya-harris.name', 0, 5, 64, '2022-02-05 13:56:37', '2022-02-09 21:02:48');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('J5f', '李鸿煊', '吴振家', '8CWJF', 'www.roman-funk.co', 1, 4, 150706, '2022-11-30 10:07:10', '2022-10-17 04:33:02');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('z4CEd', '王烨伟', '任文博', '96ZyR', 'www.whitney-stokes.com', 1, 1, 143922395, '2022-09-28 05:15:51', '2022-09-19 13:19:16');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('vsEJj', '秦懿轩', '贾风华', '4N', 'www.norman-wisozk.net', 1, 1, 18629835, '2022-09-01 17:46:15', '2022-01-01 17:44:21');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('TE6vg', '洪鹭洋', '毛昊天', '3k', 'www.thad-shields.name', 0, 360, 385263, '2022-05-07 06:09:22', '2022-08-12 10:54:35');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('UjSc', '阎鹏', '黄鹤轩', '4yG', 'www.dewitt-rogahn.io', 1, 4618, 4, '2022-05-25 16:34:04', '2022-12-19 06:40:56');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('b89Y', '李明辉', '卢伟泽', '0b', 'www.shantae-koss.net', 0, 5, 708117953, '2022-03-03 08:31:30', '2022-04-04 09:07:26');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('6kIR', '董鹤轩', '韩展鹏', 'iuOW', 'www.tommy-hudson.name', 1, 212503, 8760248, '2022-07-29 06:18:58', '2022-11-02 13:32:53');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('3eM', '薛嘉懿', '姜明', 'GElGq', 'www.morton-klocko.net', 1, 505495801, 11, '2022-10-21 16:33:42', '2022-07-22 23:41:17');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('NzFD', '覃思远', '谢哲瀚', 'z2rd', 'www.alla-wehner.co', 0, 570891054, 8, '2022-11-25 13:21:39', '2022-12-24 22:15:50');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('Cv1y', '周荣轩', '吴峻熙', 'Gzxr', 'www.dessie-rodriguez.co', 1, 4940836, 663047527, '2022-12-09 02:59:56', '2022-10-07 14:23:34');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('DzgP', '贾思远', '贺皓轩', '6YyR', 'www.virgil-emmerich.com', 0, 9064, 3603637, '2022-12-25 15:55:49', '2022-04-22 11:57:02');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('jXX6e', '徐鹭洋', '沈博超', 'tr1Lv', 'www.delbert-price.com', 1, 61273727, 10269332, '2022-01-17 05:06:53', '2022-07-31 15:20:18');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('Pp', '顾彬', '钱天磊', 'dxN', 'www.temple-osinski.org', 0, 795927429, 5, '2022-07-24 14:32:29', '2022-12-08 08:00:46');
insert into TikTok.`user` (`id`, `userAccount`, `userName`, `userPassword`, `userAvatar`, `gender`, `followCount`, `flowerCount`, `updateTime`, `createTime`) values ('PtaI', '汪文博', '曾修洁', 'E4E', 'www.waylon-lang.info', 0, 32262, 61, '2022-06-06 18:55:00', '2022-03-03 06:05:35');