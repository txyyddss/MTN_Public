-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2026-04-15 14:27:44
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
-- 数据库： `Floodgate`
--

-- --------------------------------------------------------

--
-- 表的结构 `LinkedPlayers`
--

CREATE TABLE `LinkedPlayers` (
  `bedrockId` binary(16) NOT NULL,
  `javaUniqueId` binary(16) NOT NULL,
  `javaUsername` varchar(16) COLLATE utf8mb4_general_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转存表中的数据 `LinkedPlayers`
--

INSERT INTO `LinkedPlayers` (`bedrockId`, `javaUniqueId`, `javaUsername`) VALUES
(0x0000000000000000000901f6b8cd28b2, 0x28fcb19f8a073f1c940366c0f1275aa2, 'txrog');

-- --------------------------------------------------------

--
-- 表的结构 `LinkedPlayersRequest`
--

CREATE TABLE `LinkedPlayersRequest` (
  `javaUsername` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `javaUniqueId` binary(16) NOT NULL,
  `linkCode` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `bedrockUsername` varchar(16) COLLATE utf8mb4_general_ci NOT NULL,
  `requestTime` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 转储表的索引
--

--
-- 表的索引 `LinkedPlayers`
--
ALTER TABLE `LinkedPlayers`
  ADD PRIMARY KEY (`bedrockId`),
  ADD KEY `bedrockId` (`bedrockId`,`javaUniqueId`);

--
-- 表的索引 `LinkedPlayersRequest`
--
ALTER TABLE `LinkedPlayersRequest`
  ADD PRIMARY KEY (`javaUsername`),
  ADD KEY `requestTime` (`requestTime`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
