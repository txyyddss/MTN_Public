-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2026-04-15 14:22:02
-- 服务器版本： 9.4.0
-- PHP 版本： 8.5.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `McMMO`
--

-- --------------------------------------------------------

--
-- 表的结构 `mcmmo_cooldowns`
--

CREATE TABLE `mcmmo_cooldowns` (
  `user_id` int UNSIGNED NOT NULL,
  `taming` int UNSIGNED NOT NULL DEFAULT '0',
  `mining` int UNSIGNED NOT NULL DEFAULT '0',
  `woodcutting` int UNSIGNED NOT NULL DEFAULT '0',
  `repair` int UNSIGNED NOT NULL DEFAULT '0',
  `unarmed` int UNSIGNED NOT NULL DEFAULT '0',
  `herbalism` int UNSIGNED NOT NULL DEFAULT '0',
  `excavation` int UNSIGNED NOT NULL DEFAULT '0',
  `archery` int UNSIGNED NOT NULL DEFAULT '0',
  `swords` int UNSIGNED NOT NULL DEFAULT '0',
  `axes` int UNSIGNED NOT NULL DEFAULT '0',
  `acrobatics` int UNSIGNED NOT NULL DEFAULT '0',
  `blast_mining` int UNSIGNED NOT NULL DEFAULT '0',
  `chimaera_wing` int UNSIGNED NOT NULL DEFAULT '0',
  `crossbows` int UNSIGNED NOT NULL DEFAULT '0',
  `tridents` int UNSIGNED NOT NULL DEFAULT '0',
  `maces` int UNSIGNED NOT NULL DEFAULT '0',
  `spears` int UNSIGNED NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `mcmmo_cooldowns`
--

INSERT INTO `mcmmo_cooldowns` (`user_id`, `taming`, `mining`, `woodcutting`, `repair`, `unarmed`, `herbalism`, `excavation`, `archery`, `swords`, `axes`, `acrobatics`, `blast_mining`, `chimaera_wing`, `crossbows`, `tridents`, `maces`, `spears`) VALUES
(1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);

-- --------------------------------------------------------

--
-- 表的结构 `mcmmo_experience`
--

CREATE TABLE `mcmmo_experience` (
  `user_id` int UNSIGNED NOT NULL,
  `taming` int UNSIGNED NOT NULL DEFAULT '0',
  `mining` int UNSIGNED NOT NULL DEFAULT '0',
  `woodcutting` int UNSIGNED NOT NULL DEFAULT '0',
  `repair` int UNSIGNED NOT NULL DEFAULT '0',
  `unarmed` int UNSIGNED NOT NULL DEFAULT '0',
  `herbalism` int UNSIGNED NOT NULL DEFAULT '0',
  `excavation` int UNSIGNED NOT NULL DEFAULT '0',
  `archery` int UNSIGNED NOT NULL DEFAULT '0',
  `swords` int UNSIGNED NOT NULL DEFAULT '0',
  `axes` int UNSIGNED NOT NULL DEFAULT '0',
  `acrobatics` int UNSIGNED NOT NULL DEFAULT '0',
  `fishing` int UNSIGNED NOT NULL DEFAULT '0',
  `alchemy` int UNSIGNED NOT NULL DEFAULT '0',
  `crossbows` int UNSIGNED NOT NULL DEFAULT '0',
  `tridents` int UNSIGNED NOT NULL DEFAULT '0',
  `maces` int UNSIGNED NOT NULL DEFAULT '0',
  `spears` int UNSIGNED NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `mcmmo_experience`
--

INSERT INTO `mcmmo_experience` (`user_id`, `taming`, `mining`, `woodcutting`, `repair`, `unarmed`, `herbalism`, `excavation`, `archery`, `swords`, `axes`, `acrobatics`, `fishing`, `alchemy`, `crossbows`, `tridents`, `maces`, `spears`) VALUES
(1, 0, 135, 278, 0, 450, 215, 0, 0, 0, 187, 550, 0, 0, 0, 0, 0, 0),
(2, 0, 0, 0, 0, 75, 1002, 0, 0, 0, 0, 281, 0, 0, 0, 0, 0, 0),
(3, 105, 0, 105, 105, 105, 105, 105, 105, 0, 0, 0, 105, 105, 105, 105, 105, 0),
(4, 1623, 351, 0, 0, 1778, 1170, 650, 461, 1217, 0, 795, 0, 0, 595, 0, 0, 1291),
(5, 0, 0, 1488, 0, 0, 926, 1814, 0, 296, 556, 434, 0, 0, 678, 0, 0, 236),
(6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(9, 0, 562, 850, 0, 84, 0, 30, 0, 688, 0, 1596, 0, 0, 0, 0, 0, 0);

-- --------------------------------------------------------

--
-- 表的结构 `mcmmo_huds`
--

CREATE TABLE `mcmmo_huds` (
  `user_id` int UNSIGNED NOT NULL,
  `mobhealthbar` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `scoreboardtips` int NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `mcmmo_huds`
--

INSERT INTO `mcmmo_huds` (`user_id`, `mobhealthbar`, `scoreboardtips`) VALUES
(1, 'HEARTS', 0),
(2, 'HEARTS', 0),
(3, 'HEARTS', 0),
(4, 'HEARTS', 0),
(5, 'HEARTS', 0),
(6, 'HEARTS', 0),
(7, 'HEARTS', 0),
(8, 'HEARTS', 0),
(9, 'HEARTS', 0);

-- --------------------------------------------------------

--
-- 表的结构 `mcmmo_skills`
--

CREATE TABLE `mcmmo_skills` (
  `user_id` int UNSIGNED NOT NULL,
  `taming` int UNSIGNED NOT NULL DEFAULT '0',
  `mining` int UNSIGNED NOT NULL DEFAULT '0',
  `woodcutting` int UNSIGNED NOT NULL DEFAULT '0',
  `repair` int UNSIGNED NOT NULL DEFAULT '0',
  `unarmed` int UNSIGNED NOT NULL DEFAULT '0',
  `herbalism` int UNSIGNED NOT NULL DEFAULT '0',
  `excavation` int UNSIGNED NOT NULL DEFAULT '0',
  `archery` int UNSIGNED NOT NULL DEFAULT '0',
  `swords` int UNSIGNED NOT NULL DEFAULT '0',
  `axes` int UNSIGNED NOT NULL DEFAULT '0',
  `acrobatics` int UNSIGNED NOT NULL DEFAULT '0',
  `fishing` int UNSIGNED NOT NULL DEFAULT '0',
  `alchemy` int UNSIGNED NOT NULL DEFAULT '0',
  `crossbows` int UNSIGNED NOT NULL DEFAULT '0',
  `tridents` int UNSIGNED NOT NULL DEFAULT '0',
  `maces` int UNSIGNED NOT NULL DEFAULT '0',
  `spears` int UNSIGNED NOT NULL DEFAULT '0',
  `total` int UNSIGNED NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 转存表中的数据 `mcmmo_skills`
--

INSERT INTO `mcmmo_skills` (`user_id`, `taming`, `mining`, `woodcutting`, `repair`, `unarmed`, `herbalism`, `excavation`, `archery`, `swords`, `axes`, `acrobatics`, `fishing`, `alchemy`, `crossbows`, `tridents`, `maces`, `spears`, `total`) VALUES
(1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 3, 18, 0, 0, 0, 0, 0, 0, 22),
(2, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 7),
(3, 0, 7, 0, 0, 0, 0, 0, 0, 25, 1, 20, 0, 0, 0, 0, 0, 1, 54),
(4, 5, 0, 0, 0, 0, 3, 2, 0, 6, 0, 19, 0, 0, 1, 0, 0, 15, 51),
(5, 0, 0, 0, 0, 0, 1, 1, 0, 13, 1, 34, 0, 0, 0, 0, 0, 1, 51),
(6, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 15, 0, 0, 0, 0, 0, 0, 16),
(7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0),
(9, 0, 9, 1, 0, 0, 0, 5, 0, 1, 0, 21, 0, 0, 0, 0, 0, 0, 37);

-- --------------------------------------------------------

--
-- 表的结构 `mcmmo_users`
--

CREATE TABLE `mcmmo_users` (
  `id` int NOT NULL,
  `user` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `uuid` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lastlogin` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转存表中的数据 `mcmmo_users`
--

INSERT INTO `mcmmo_users` (`id`, `user`, `uuid`, `lastlogin`) VALUES
(1, 'txrog', '28fcb19f-8a07-3f1c-9403-66c0f1275aa2', 1776170383),
(2, 'BE_MADIAM2395', '00000000-0000-0000-0009-01f6b8cd28b2', 1776167513),
(3, 'Andras', 'dd94d51e-2195-38e1-9d59-31b82b4b715f', 1776254986),
(4, 'wor114514', 'a3ec4008-6aaa-3fb9-b46a-36fc3360df12', 1776238052),
(5, 'mc9957', 'bcc9c4ca-ae03-34d8-b7e6-a842b9d9c057', 1776224840),
(6, 'lu', '2114cd3b-96c1-348e-9c5b-1144fa099884', 1776254987),
(7, 'BE_DF738383', '00000000-0000-0000-0009-01f2441ec47d', 1776073649),
(8, 'txro', 'f66eed8e-d372-36fb-9e93-77b49695de0f', 1776085297),
(9, 'HEHAWellgood', 'a20b9b3f-9edd-34c2-a360-a58100a8af25', 1776260988);

--
-- 转储表的索引
--

--
-- 表的索引 `mcmmo_cooldowns`
--
ALTER TABLE `mcmmo_cooldowns`
  ADD PRIMARY KEY (`user_id`);

--
-- 表的索引 `mcmmo_experience`
--
ALTER TABLE `mcmmo_experience`
  ADD PRIMARY KEY (`user_id`);

--
-- 表的索引 `mcmmo_huds`
--
ALTER TABLE `mcmmo_huds`
  ADD PRIMARY KEY (`user_id`);

--
-- 表的索引 `mcmmo_skills`
--
ALTER TABLE `mcmmo_skills`
  ADD PRIMARY KEY (`user_id`);

--
-- 表的索引 `mcmmo_users`
--
ALTER TABLE `mcmmo_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uuid` (`uuid`),
  ADD KEY `user_index` (`user`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `mcmmo_users`
--
ALTER TABLE `mcmmo_users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
