
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

CREATE TABLE `messages` (
  `id` int NOT NULL,
  `user_id` int NOT NULL,
  `user_name` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `message_type` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `message_content` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `media_file_name` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO `messages` (`id`, `user_id`, `user_name`, `message_type`, `message_content`, `media_file_name`, `created_at`, `updated_at`) VALUES
(22, 5, 'anhle150199', 'message', 'chào buổi sáng', '', '2021-12-09 09:49:54', '2021-12-09 09:49:54'),
(23, 5, 'anhle150199', 'message', 'hello, xin chao moi nguoi', '', '2021-12-09 10:00:28', '2021-12-09 10:00:28'),
(24, 4, 'Test 21', 'message', 'xin chao cac ban', '', '2021-12-09 10:35:06', '2021-12-09 10:35:06'),
(25, 3, 'test1', 'message', 'chào mọi người nhé', '', '2021-12-09 11:06:11', '2021-12-09 11:06:11'),
(37, 3, 'test1', 'message', 'Chao nhe', '', '2021-12-09 13:21:48', '2021-12-09 13:21:48'),
(39, 3, 'test1', 'image', '', 'xQswQz0A-onedrive_frdq.jpg', '2021-12-09 13:22:17', '2021-12-09 13:22:17'),
(40, 3, 'test1', 'video', '', 'GHNTTkmQ-BANLANHAT.mp4', '2021-12-09 13:22:28', '2021-12-09 13:22:28'),
(41, 6, 'qqqq', 'message', 'chao mn nhe', '', '2021-12-09 14:23:37', '2021-12-09 14:23:37'),
(42, 6, 'qqqq', 'message', 'hi', '', '2021-12-09 14:24:00', '2021-12-09 14:24:00'),
(46, 4, 'Test 21', 'message', 'okokok', '', '2021-12-09 15:16:03', '2021-12-09 15:16:03'),
(50, 7, 'admin', 'message', 'ok', '', '2021-12-09 16:47:15', '2021-12-09 16:47:15'),
(51, 7, 'admin', 'message', 'alo co ai o day khong', '', '2021-12-10 09:17:25', '2021-12-10 09:17:25'),
(52, 7, 'admin', 'image', '', 'SWwicYD5-3.2.png', '2021-12-10 09:18:55', '2021-12-10 09:18:55'),
(66, 4, 'Test 21', 'stamps', '', '21.png', '2021-12-10 14:29:00', '2021-12-10 14:29:00'),
(67, 7, 'admin', 'stamps', '', '22.png', '2021-12-10 14:35:29', '2021-12-10 14:35:29'),
(68, 7, 'admin', 'image', '', 'nFvFDiL9-2.1.png', '2021-12-10 14:36:11', '2021-12-10 14:36:11');

-- --------------------------------------------------------

CREATE TABLE `users` (
  `id` int UNSIGNED NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `email` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `users` (`id`, `name`, `email`, `password`, `created_at`, `updated_at`) VALUES
(2, 'anhle1', 'anhle1@gmail.com', '$2a$14$/uxTN1DQ4pmME0htIBD.xObBJ00grmGn3pGenHy1X.GKuEFO33H76', '2021-11-18 15:38:51', '2021-11-18 15:38:51'),
(3, 'test1', 'test1@gmail.com', '$2a$14$TYG4bXU0aRbWL9NjA/IJ8ukpgQ0aaKm65eMqHg71UMOiPOEBTIQ6W', '2021-11-18 16:57:02', '2021-11-18 16:57:02'),
(4, 'Test 21', 'test2@gmail.com', '$2a$14$E5MRBzd2YpFz6LEA.VmU6.o098pGXG6sN12SX36y.BSFjx87htING', '2021-11-18 20:49:19', '2021-11-18 20:49:19'),
(5, 'anhle150199', 'anhle150199@gmail.com', '$2a$14$77jFjMf7Hh9pNY.uo8pnfuwk9kmpq1Vur2KgmB3qVL6RgYqIIIdF.', '2021-11-30 10:32:59', '2021-11-30 10:32:59'),
(6, 'qqqq', 'test11@gmail.com', '$2a$14$nT9TxXGxIu3dARQoMEt69eAA123Q9pP/.RHAKVqXzG9UpnkoxLvkm', '2021-12-09 14:22:58', '2021-12-09 14:22:58'),
(7, 'admin', 'admin@gmail.com', '$2a$14$trMOARGKWuMNg3kaqGmFXOtudGhAoGGd9rmhILBph0qW/c4AYDJKe', '2021-12-09 14:31:28', '2021-12-09 14:31:28'),
(8, 'anhle1999', 'anhle111@gmail.com', '$2a$14$LIuixojoitGgnyzQxxGaWO0fgVIwQdEPHpdvFJtPJeyig5Ub0.8nW', '2021-12-09 15:52:34', '2021-12-09 15:52:34'),
(9, 'test3', 'test3@gmail.com', '$2a$14$jZQHOYqR9Jph3NuhtPsXluDbwOxDT1RU0SJ4vYa3Da4Hl6sgJLD7i', '2021-12-09 15:59:58', '2021-12-09 15:59:58'),
(10, 'text111', 'text@gmail.com', '$2a$14$ePQ6ZAx6WCmVjsDeuC6ws./JoviARf2ejR1fDGJQ4Yb.BWpIbminm', '2021-12-09 16:03:17', '2021-12-09 16:03:17');

ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

ALTER TABLE `messages`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=69;

ALTER TABLE `users`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;
