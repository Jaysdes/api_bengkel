-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 15 Jul 2025 pada 03.08
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bengkel`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `cache`
--

CREATE TABLE `cache` (
  `key` varchar(255) NOT NULL,
  `value` mediumtext NOT NULL,
  `expiration` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `cache_locks`
--

CREATE TABLE `cache_locks` (
  `key` varchar(255) NOT NULL,
  `owner` varchar(255) NOT NULL,
  `expiration` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `customer`
--

CREATE TABLE `customer` (
  `id_customer` int(11) NOT NULL,
  `nama_customer` varchar(100) DEFAULT NULL,
  `id_jenis` int(11) DEFAULT NULL,
  `no_kendaraan` varchar(50) DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `telepon` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `customer`
--

INSERT INTO `customer` (`id_customer`, `nama_customer`, `id_jenis`, `no_kendaraan`, `alamat`, `telepon`) VALUES
(4, 'Budi Santoso', 2, 'B1234XYZ', 'Jl. Merdeka No. 19', '08123456789'),
(6, 'Bamabang', 1, 'B1231BB', 'Jl.Makmur N0.12', '08123311100');

-- --------------------------------------------------------

--
-- Struktur dari tabel `detail_transaksi`
--

CREATE TABLE `detail_transaksi` (
  `id_detail` int(11) NOT NULL,
  `no_spk` int(11) DEFAULT NULL,
  `id_customer` int(11) DEFAULT NULL,
  `no_kendaraan` varchar(50) DEFAULT NULL,
  `id_sparepart` int(11) DEFAULT NULL,
  `id_service` int(11) DEFAULT NULL,
  `id_jasa` int(11) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `failed_jobs`
--

CREATE TABLE `failed_jobs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `uuid` varchar(255) NOT NULL,
  `connection` text NOT NULL,
  `queue` text NOT NULL,
  `payload` longtext NOT NULL,
  `exception` longtext NOT NULL,
  `failed_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `jenis_jasa`
--

CREATE TABLE `jenis_jasa` (
  `id_jasa` int(11) NOT NULL,
  `nama_jasa` varchar(100) DEFAULT NULL,
  `harga_jasa` bigint(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `jenis_jasa`
--

INSERT INTO `jenis_jasa` (`id_jasa`, `nama_jasa`, `harga_jasa`) VALUES
(1, 'Service Ringan Motor', 75000),
(2, 'Service Lengkap Motor', 125000),
(3, 'Ganti Oli Mesin Motor', 40000),
(4, 'Ganti Kampas Rem Motor', 30000),
(5, 'Tune Up Motor', 85000),
(6, 'Service Ringan Mobil', 150000),
(7, 'Service Lengkap Mobil', 300000),
(8, 'Ganti Oli Mesin Mobil', 100000),
(9, 'Ganti Kampas Rem Mobil', 120000),
(10, 'Tune Up Mobil', 175000),
(11, 'Cek Kelistrikan', 50000),
(12, 'Spooring & Balancing', 200000),
(13, 'Cuci Mobil', 50000),
(14, 'Cuci Motor', 25000);

-- --------------------------------------------------------

--
-- Struktur dari tabel `jenis_kendaraan`
--

CREATE TABLE `jenis_kendaraan` (
  `id_jenis` int(11) NOT NULL,
  `jenis_kendaraan` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `jenis_kendaraan`
--

INSERT INTO `jenis_kendaraan` (`id_jenis`, `jenis_kendaraan`) VALUES
(1, 'Motor'),
(2, 'Mobil'),
(3, 'Motor');

-- --------------------------------------------------------

--
-- Struktur dari tabel `jenis_service`
--

CREATE TABLE `jenis_service` (
  `id_service` int(11) NOT NULL,
  `jenis_service` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `jenis_service`
--

INSERT INTO `jenis_service` (`id_service`, `jenis_service`) VALUES
(1, 'Berkala'),
(2, 'Tidak Berkala');

-- --------------------------------------------------------

--
-- Struktur dari tabel `jobs`
--

CREATE TABLE `jobs` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `queue` varchar(255) NOT NULL,
  `payload` longtext NOT NULL,
  `attempts` tinyint(3) UNSIGNED NOT NULL,
  `reserved_at` int(10) UNSIGNED DEFAULT NULL,
  `available_at` int(10) UNSIGNED NOT NULL,
  `created_at` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `job_batches`
--

CREATE TABLE `job_batches` (
  `id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `total_jobs` int(11) NOT NULL,
  `pending_jobs` int(11) NOT NULL,
  `failed_jobs` int(11) NOT NULL,
  `failed_job_ids` longtext NOT NULL,
  `options` mediumtext DEFAULT NULL,
  `cancelled_at` int(11) DEFAULT NULL,
  `created_at` int(11) NOT NULL,
  `finished_at` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `laporan`
--

CREATE TABLE `laporan` (
  `id_laporan` int(11) NOT NULL,
  `id_transaksi` int(11) DEFAULT NULL,
  `id_customer` int(11) DEFAULT NULL,
  `total_biaya` bigint(20) DEFAULT NULL,
  `tanggal_laporan` date DEFAULT curdate(),
  `catatan` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `mekanik`
--

CREATE TABLE `mekanik` (
  `id_mekanik` int(11) NOT NULL,
  `nama_mekanik` varchar(100) DEFAULT NULL,
  `jenis_kelamin` enum('L','P') DEFAULT NULL,
  `alamat` text DEFAULT NULL,
  `telepon` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `mekanik`
--

INSERT INTO `mekanik` (`id_mekanik`, `nama_mekanik`, `jenis_kelamin`, `alamat`, `telepon`) VALUES
(1, 'maman', 'L', 'jln.Bojong No.15', '0849124232');

-- --------------------------------------------------------

--
-- Struktur dari tabel `migrations`
--

CREATE TABLE `migrations` (
  `id` int(10) UNSIGNED NOT NULL,
  `migration` varchar(255) NOT NULL,
  `batch` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `migrations`
--

INSERT INTO `migrations` (`id`, `migration`, `batch`) VALUES
(1, '0001_01_01_000001_create_cache_table', 1),
(2, '0001_01_01_000002_create_jobs_table', 1),
(3, '2025_06_13_091728_create_sessions_table', 2);

-- --------------------------------------------------------

--
-- Struktur dari tabel `proses`
--

CREATE TABLE `proses` (
  `id_proses` int(11) NOT NULL,
  `id_transaksi` int(11) DEFAULT NULL,
  `id_mekanik` int(11) DEFAULT NULL,
  `status` enum('dalam_antrian','proses','selesai') DEFAULT 'dalam_antrian',
  `keterangan` text DEFAULT NULL,
  `waktu_mulai` datetime DEFAULT current_timestamp(),
  `waktu_selesai` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `sessions`
--

CREATE TABLE `sessions` (
  `id` varchar(255) NOT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `ip_address` varchar(45) DEFAULT NULL,
  `user_agent` text DEFAULT NULL,
  `payload` longtext NOT NULL,
  `last_activity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data untuk tabel `sessions`
--

INSERT INTO `sessions` (`id`, `user_id`, `ip_address`, `user_agent`, `payload`, `last_activity`) VALUES
('045HLMYf2aq1jh9RBTBZYpM0xNY3xS6pVokI3ZXL', NULL, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36', 'YTozOntzOjY6Il90b2tlbiI7czo0MDoiSHNWUW9MZTdSbm5wbG9IcFp2bzdtMG84ZW1hWURPQWgxY2szMHFIRCI7czo5OiJfcHJldmlvdXMiO2E6MTp7czozOiJ1cmwiO3M6MjE6Imh0dHA6Ly8xMjcuMC4wLjE6ODAwMCI7fXM6NjoiX2ZsYXNoIjthOjI6e3M6Mzoib2xkIjthOjA6e31zOjM6Im5ldyI7YTowOnt9fX0=', 1751430004),
('7rfqypCHIjk6qfo0WQJRSXzAaA7jjmmIojRhDc44', NULL, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36', 'YTo1OntzOjY6Il9mbGFzaCI7YToyOntzOjM6Im9sZCI7YTowOnt9czozOiJuZXciO2E6MDp7fX1zOjY6Il90b2tlbiI7czo0MDoiRXluM2JhSjQ3cGFLQzBKV2pXbjF5TFhWYzl5b3N5SXNyZU04Q0R0ZCI7czo5OiJfcHJldmlvdXMiO2E6MTp7czozOiJ1cmwiO3M6MjY6Imh0dHA6Ly8xMjcuMC4wLjE6ODAwMC9ob21lIjt9czo1OiJ0b2tlbiI7czoxNDU6ImV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpsZUhBaU9qRTNOVEUxTVRZMU1qQXNJbkp2YkdVaU9pSmpkWE4wYjIxbGNpSXNJblZ6WlhKZmFXUWlPalY5Lkx3THZQZ2NBdlZmeDgzTXd5U1ZsNEJVb251SUpudno0OG1NSTU4TTJjT2ciO3M6NDoidXNlciI7YTo0OntzOjU6ImVtYWlsIjtzOjE1OiJqYW1hbEBnbWFpbC5jb20iO3M6MjoiaWQiO2k6NTtzOjQ6Im5hbWUiO3M6NToiamFtYWwiO3M6NDoicm9sZSI7czo4OiJjdXN0b21lciI7fX0=', 1751430128),
('wlxGWfs87o8G50PcnzQGYKH0Ud2fW2hgAiO8mC12', NULL, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36', 'YTo1OntzOjY6Il9mbGFzaCI7YToyOntzOjM6Im9sZCI7YTowOnt9czozOiJuZXciO2E6MDp7fX1zOjY6Il90b2tlbiI7czo0MDoiaEFSdEo1Z1JGVmxqOGRDN3V4bHcxdTlaYW1ZUWRFajBFckhzdU9xNyI7czo5OiJfcHJldmlvdXMiO2E6MTp7czozOiJ1cmwiO3M6Mjk6Imh0dHA6Ly8xMjcuMC4wLjE6ODAwMC9tZWthbmlrIjt9czo1OiJ0b2tlbiI7czoxNDE6ImV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpsZUhBaU9qRTNOVEl3TXpVNU1EVXNJbkp2YkdVaU9pSmhaRzFwYmlJc0luVnpaWEpmYVdRaU9qSjkuWnJMNzFQcE0zUXVNd2dyRTlxTnI5M0djT2VhbGx0dWl0VnYwR0pyNlYyayI7czo0OiJ1c2VyIjthOjQ6e3M6NToiZW1haWwiO3M6MTI6Imp5QGdtYWlsLmNvbSI7czoyOiJpZCI7aToyO3M6NDoibmFtZSI7czoxMzoiSmF5YWRpIGRpbmF0YSI7czo0OiJyb2xlIjtzOjU6ImFkbWluIjt9fQ==', 1751949524);

-- --------------------------------------------------------

--
-- Struktur dari tabel `sparepart`
--

CREATE TABLE `sparepart` (
  `id_sparepart` int(11) NOT NULL,
  `nama_sparepart` varchar(100) DEFAULT NULL,
  `harga_beli` bigint(20) NOT NULL,
  `harga_jual` bigint(20) NOT NULL,
  `stok` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `sparepart`
--

INSERT INTO `sparepart` (`id_sparepart`, `nama_sparepart`, `harga_beli`, `harga_jual`, `stok`) VALUES
(1, 'Ban Dalam', 100000, 120000, 20);

-- --------------------------------------------------------

--
-- Struktur dari tabel `spk`
--

CREATE TABLE `spk` (
  `id_spk` int(11) NOT NULL,
  `tanggal_spk` date DEFAULT NULL,
  `id_service` int(11) DEFAULT NULL,
  `id_jasa` int(11) DEFAULT NULL,
  `id_customer` int(11) DEFAULT NULL,
  `id_jenis` int(11) DEFAULT NULL,
  `no_kendaraan` varchar(50) DEFAULT NULL,
  `keluhan` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaksi`
--

CREATE TABLE `transaksi` (
  `id_transaksi` int(11) NOT NULL,
  `id_spk` int(11) DEFAULT NULL,
  `id_customer` int(11) DEFAULT NULL,
  `id_jenis` int(11) DEFAULT NULL,
  `no_kendaraan` varchar(50) DEFAULT NULL,
  `pemilik` varchar(100) DEFAULT NULL,
  `telepon` varchar(20) DEFAULT NULL,
  `id_mekanik` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `email_verified_at` datetime(3) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `role` enum('mekanik','admin','keuangan','gudang','customer') DEFAULT 'customer',
  `remember_token` varchar(100) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `email_verified_at`, `password`, `role`, `remember_token`, `created_at`, `updated_at`) VALUES
(2, 'Jayadi dinata', 'jy@gmail.com', NULL, '$2a$10$jxkjQfgqf6Zta.nZ9B3VlubP.lLJW5aS5gwxCmAenL4u4i2WlclK.', 'admin', '', '2025-05-26 03:51:36.078', '2025-05-26 03:51:36.078'),
(4, 'Budi Santoso', 'budi@gmail.com', NULL, '$2a$10$UBVEpeRpWqJQbuligSaKpeT4wfx1IkJeNiKxkLsgI3wlQj1r/ShTa', 'mekanik', '', '2025-06-14 05:29:47.101', '2025-06-14 05:29:47.101'),
(5, 'jamal', 'jamal@gmail.com', NULL, '$2a$10$5wx84alAUlgPnNqPyjoFQeGyaQjFz5cHjj4UfyUnI5yL2yvcruk5a', 'customer', '', '2025-06-16 08:29:26.342', '2025-06-16 08:29:26.342'),
(7, 'jumint', 'jumin@gmail.com', NULL, '$2a$10$SLUkyL2ShAad2au8ltVkj.PGmA2p5dbPvWF0VYOKXUL51H0LibcIi', 'keuangan', '', '2025-06-17 03:58:47.103', '2025-06-17 03:58:47.103'),
(8, 'Mas Ker', 'ker@gmail.com', NULL, '$2a$10$IQTQB8Gz6QwQKs0TNJW7ce/5evwDgHCe/KuL.gCd60FByLB7vGJBO', 'gudang', '', '2025-06-23 09:34:03.790', '2025-06-23 09:34:03.790');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `cache`
--
ALTER TABLE `cache`
  ADD PRIMARY KEY (`key`);

--
-- Indeks untuk tabel `cache_locks`
--
ALTER TABLE `cache_locks`
  ADD PRIMARY KEY (`key`);

--
-- Indeks untuk tabel `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`id_customer`),
  ADD KEY `id_jenis` (`id_jenis`);

--
-- Indeks untuk tabel `detail_transaksi`
--
ALTER TABLE `detail_transaksi`
  ADD PRIMARY KEY (`id_detail`),
  ADD KEY `no_spk` (`no_spk`),
  ADD KEY `id_customer` (`id_customer`),
  ADD KEY `id_sparepart` (`id_sparepart`),
  ADD KEY `id_service` (`id_service`),
  ADD KEY `id_jasa` (`id_jasa`);

--
-- Indeks untuk tabel `failed_jobs`
--
ALTER TABLE `failed_jobs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `failed_jobs_uuid_unique` (`uuid`);

--
-- Indeks untuk tabel `jenis_jasa`
--
ALTER TABLE `jenis_jasa`
  ADD PRIMARY KEY (`id_jasa`);

--
-- Indeks untuk tabel `jenis_kendaraan`
--
ALTER TABLE `jenis_kendaraan`
  ADD PRIMARY KEY (`id_jenis`);

--
-- Indeks untuk tabel `jenis_service`
--
ALTER TABLE `jenis_service`
  ADD PRIMARY KEY (`id_service`);

--
-- Indeks untuk tabel `jobs`
--
ALTER TABLE `jobs`
  ADD PRIMARY KEY (`id`),
  ADD KEY `jobs_queue_index` (`queue`);

--
-- Indeks untuk tabel `job_batches`
--
ALTER TABLE `job_batches`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `laporan`
--
ALTER TABLE `laporan`
  ADD PRIMARY KEY (`id_laporan`),
  ADD KEY `id_transaksi` (`id_transaksi`),
  ADD KEY `id_customer` (`id_customer`);

--
-- Indeks untuk tabel `mekanik`
--
ALTER TABLE `mekanik`
  ADD PRIMARY KEY (`id_mekanik`);

--
-- Indeks untuk tabel `migrations`
--
ALTER TABLE `migrations`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `proses`
--
ALTER TABLE `proses`
  ADD PRIMARY KEY (`id_proses`),
  ADD KEY `id_transaksi` (`id_transaksi`),
  ADD KEY `id_mekanik` (`id_mekanik`);

--
-- Indeks untuk tabel `sessions`
--
ALTER TABLE `sessions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `sessions_user_id_index` (`user_id`),
  ADD KEY `sessions_last_activity_index` (`last_activity`);

--
-- Indeks untuk tabel `sparepart`
--
ALTER TABLE `sparepart`
  ADD PRIMARY KEY (`id_sparepart`);

--
-- Indeks untuk tabel `spk`
--
ALTER TABLE `spk`
  ADD PRIMARY KEY (`id_spk`),
  ADD KEY `id_service` (`id_service`),
  ADD KEY `id_jasa` (`id_jasa`),
  ADD KEY `id_customer` (`id_customer`),
  ADD KEY `id_jenis` (`id_jenis`);

--
-- Indeks untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`id_transaksi`),
  ADD KEY `id_spk` (`id_spk`),
  ADD KEY `id_customer` (`id_customer`),
  ADD KEY `id_jenis` (`id_jenis`),
  ADD KEY `id_mekanik` (`id_mekanik`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_users_email` (`email`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `customer`
--
ALTER TABLE `customer`
  MODIFY `id_customer` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `detail_transaksi`
--
ALTER TABLE `detail_transaksi`
  MODIFY `id_detail` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `failed_jobs`
--
ALTER TABLE `failed_jobs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `jenis_jasa`
--
ALTER TABLE `jenis_jasa`
  MODIFY `id_jasa` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT untuk tabel `jobs`
--
ALTER TABLE `jobs`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `laporan`
--
ALTER TABLE `laporan`
  MODIFY `id_laporan` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `mekanik`
--
ALTER TABLE `mekanik`
  MODIFY `id_mekanik` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT untuk tabel `migrations`
--
ALTER TABLE `migrations`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `proses`
--
ALTER TABLE `proses`
  MODIFY `id_proses` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `sparepart`
--
ALTER TABLE `sparepart`
  MODIFY `id_sparepart` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT untuk tabel `spk`
--
ALTER TABLE `spk`
  MODIFY `id_spk` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `id_transaksi` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `customer`
--
ALTER TABLE `customer`
  ADD CONSTRAINT `customer_ibfk_1` FOREIGN KEY (`id_jenis`) REFERENCES `jenis_kendaraan` (`id_jenis`);

--
-- Ketidakleluasaan untuk tabel `detail_transaksi`
--
ALTER TABLE `detail_transaksi`
  ADD CONSTRAINT `detail_transaksi_ibfk_1` FOREIGN KEY (`no_spk`) REFERENCES `spk` (`id_spk`),
  ADD CONSTRAINT `detail_transaksi_ibfk_2` FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id_customer`),
  ADD CONSTRAINT `detail_transaksi_ibfk_3` FOREIGN KEY (`id_sparepart`) REFERENCES `sparepart` (`id_sparepart`),
  ADD CONSTRAINT `detail_transaksi_ibfk_4` FOREIGN KEY (`id_service`) REFERENCES `jenis_service` (`id_service`),
  ADD CONSTRAINT `detail_transaksi_ibfk_5` FOREIGN KEY (`id_jasa`) REFERENCES `jenis_jasa` (`id_jasa`);

--
-- Ketidakleluasaan untuk tabel `laporan`
--
ALTER TABLE `laporan`
  ADD CONSTRAINT `laporan_ibfk_1` FOREIGN KEY (`id_transaksi`) REFERENCES `transaksi` (`id_transaksi`),
  ADD CONSTRAINT `laporan_ibfk_2` FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id_customer`);

--
-- Ketidakleluasaan untuk tabel `proses`
--
ALTER TABLE `proses`
  ADD CONSTRAINT `proses_ibfk_1` FOREIGN KEY (`id_transaksi`) REFERENCES `transaksi` (`id_transaksi`),
  ADD CONSTRAINT `proses_ibfk_2` FOREIGN KEY (`id_mekanik`) REFERENCES `mekanik` (`id_mekanik`);

--
-- Ketidakleluasaan untuk tabel `spk`
--
ALTER TABLE `spk`
  ADD CONSTRAINT `spk_ibfk_1` FOREIGN KEY (`id_service`) REFERENCES `jenis_service` (`id_service`),
  ADD CONSTRAINT `spk_ibfk_2` FOREIGN KEY (`id_jasa`) REFERENCES `jenis_jasa` (`id_jasa`),
  ADD CONSTRAINT `spk_ibfk_3` FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id_customer`),
  ADD CONSTRAINT `spk_ibfk_4` FOREIGN KEY (`id_jenis`) REFERENCES `jenis_kendaraan` (`id_jenis`);

--
-- Ketidakleluasaan untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  ADD CONSTRAINT `transaksi_ibfk_1` FOREIGN KEY (`id_spk`) REFERENCES `spk` (`id_spk`),
  ADD CONSTRAINT `transaksi_ibfk_2` FOREIGN KEY (`id_customer`) REFERENCES `customer` (`id_customer`),
  ADD CONSTRAINT `transaksi_ibfk_3` FOREIGN KEY (`id_jenis`) REFERENCES `jenis_kendaraan` (`id_jenis`),
  ADD CONSTRAINT `transaksi_ibfk_4` FOREIGN KEY (`id_mekanik`) REFERENCES `mekanik` (`id_mekanik`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
