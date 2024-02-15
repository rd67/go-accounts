-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: mysql
-- Generation Time: Feb 15, 2024 at 05:08 PM
-- Server version: 8.2.0
-- PHP Version: 7.4.20
SET
    SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

START TRANSACTION;

SET
    time_zone = "+00:00";

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */
;

/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */
;

/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */
;

/*!40101 SET NAMES utf8mb4 */
;

--
-- Database: `go_accounts`
--
CREATE DATABASE IF NOT EXISTS `go_accounts` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

USE `go_accounts`;

-- --------------------------------------------------------
--
-- Table structure for table `accounts`
--
DROP TABLE IF EXISTS `accounts`;

CREATE TABLE `accounts` (
    `id` bigint UNSIGNED NOT NULL,
    `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_unicode_ci NOT NULL,
    `balance` float UNSIGNED ZEROFILL NOT NULL DEFAULT '000000000000',
    `currency` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'INR',
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0',
    `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- --------------------------------------------------------
--
-- Table structure for table `entries`
--
DROP TABLE IF EXISTS `entries`;

CREATE TABLE `entries` (
    `id` bigint UNSIGNED NOT NULL,
    `account_id` bigint UNSIGNED NOT NULL,
    `amount` float UNSIGNED ZEROFILL NOT NULL,
    `currency` varchar(3) NOT NULL DEFAULT 'INR',
    `exchange_rate` float NOT NULL DEFAULT '1',
    `isDeleted` tinyint(1) NOT NULL DEFAULT '0',
    `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

-- --------------------------------------------------------
--
-- Table structure for table `transfers`
--
DROP TABLE IF EXISTS `transfers`;

CREATE TABLE `transfers` (
    `id` bigint UNSIGNED NOT NULL,
    `sender_account_id` bigint UNSIGNED NOT NULL,
    `receiver_account_id` bigint UNSIGNED NOT NULL,
    `amount` float UNSIGNED ZEROFILL NOT NULL,
    `currency` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'INR',
    `exchange_rate` float NOT NULL DEFAULT '1',
    `status` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--
--
-- Indexes for table `accounts`
--
ALTER TABLE
    `accounts`
ADD
    PRIMARY KEY (`id`);

--
-- Indexes for table `entries`
--
ALTER TABLE
    `entries`
ADD
    PRIMARY KEY (`id`),
ADD
    KEY `account_id` (`account_id`);

--
-- Indexes for table `transfers`
--
ALTER TABLE
    `transfers`
ADD
    PRIMARY KEY (`id`),
ADD
    KEY `sender_account_id` (`sender_account_id`, `receiver_account_id`),
ADD
    KEY `receiver_account_id` (`receiver_account_id`),
ADD
    KEY `sender_account_id_2` (`sender_account_id`);

--
-- AUTO_INCREMENT for dumped tables
--
--
-- AUTO_INCREMENT for table `accounts`
--
ALTER TABLE
    `accounts`
MODIFY
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `entries`
--
ALTER TABLE
    `entries`
MODIFY
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `transfers`
--
ALTER TABLE
    `transfers`
MODIFY
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--
--
-- Constraints for table `transfers`
--
ALTER TABLE
    `transfers`
ADD
    CONSTRAINT `transfers_ibfk_1` FOREIGN KEY (`receiver_account_id`) REFERENCES `accounts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
ADD
    CONSTRAINT `transfers_ibfk_2` FOREIGN KEY (`sender_account_id`) REFERENCES `accounts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */
;

/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */
;

/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */
;
