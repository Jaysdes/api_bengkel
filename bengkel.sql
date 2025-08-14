-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 14 Agu 2025 pada 02.47
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
  `telepon` varchar(20) DEFAULT NULL,
  `tanggal_masuk` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `customer`
--

INSERT INTO `customer` (`id_customer`, `nama_customer`, `id_jenis`, `no_kendaraan`, `alamat`, `telepon`, `tanggal_masuk`) VALUES
(4, 'Budi Santoso', 1, 'B1234XYZ', 'Jl. Merdeka No. 19', '082123231', '2025-08-09'),
(6, 'Bamabang', 1, 'B1231BB', 'Jl.Makmur N0.12', '08123311100', NULL),
(7, 'Sam Padn', 2, 'B2123X', 'Jln.Samarinda 21', '0832312313', '2025-08-07'),
(8, 'Massa', 2, 'A092321W', 'Jln.Samarinda 21', '0832312313', '2025-08-07'),
(9, 'Massa', 2, 'B2123X', 'Jln. Makmur 22', '0832312313', '2025-08-12');

-- --------------------------------------------------------

--
-- Struktur dari tabel `detail_transaksi`
--

CREATE TABLE `detail_transaksi` (
  `id_detail` int(11) NOT NULL,
  `id_transaksi` int(11) DEFAULT NULL,
  `no_spk` int(11) DEFAULT NULL,
  `id_customer` int(11) DEFAULT NULL,
  `no_kendaraan` varchar(50) DEFAULT NULL,
  `id_sparepart` int(11) DEFAULT NULL,
  `id_service` int(11) DEFAULT NULL,
  `id_jasa` int(11) DEFAULT NULL,
  `total` bigint(20) DEFAULT NULL,
  `bayar` bigint(20) DEFAULT NULL,
  `kembalian` bigint(20) DEFAULT NULL
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
(2, 'Mobil');

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
  `status` varchar(50) NOT NULL DEFAULT 'transaksi di proses',
  `keterangan` varchar(255) DEFAULT 'menunggu konfirmasi',
  `waktu_mulai` datetime NOT NULL,
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
('IPihCknric6fSFuf9xreaGe3LXDqzNMbVEdtO6vG', NULL, '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36', 'YTo1OntzOjY6Il9mbGFzaCI7YToyOntzOjM6Im9sZCI7YTowOnt9czozOiJuZXciO2E6MDp7fX1zOjY6Il90b2tlbiI7czo0MDoicFo5OU5HSG5JbkdUWEV3ZTZkNU0zcEtaZmw1WmtMWlc3bTlTUTk1dSI7czo5OiJfcHJldmlvdXMiO2E6MTp7czozOiJ1cmwiO3M6MzE6Imh0dHA6Ly8xMjcuMC4wLjE6ODAwMC9kYXNoYm9hcmQiO31zOjU6InRva2VuIjtzOjE0MToiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SmxlSEFpT2pFM05UVXhOelEyTXpjc0luSnZiR1VpT2lKaFpHMXBiaUlzSW5WelpYSmZhV1FpT2pKOS43T2wwUFpfaE01ZkFpTXNFRFhPb2wyakJGdXl1Vll2VGw4bWhaSFlveTNZIjtzOjQ6InVzZXIiO2E6NDp7czo1OiJlbWFpbCI7czoxMjoianlAZ21haWwuY29tIjtzOjI6ImlkIjtpOjI7czo0OiJuYW1lIjtzOjEzOiJKYXlhZGkgZGluYXRhIjtzOjQ6InJvbGUiO3M6NToiYWRtaW4iO319', 1755101344);

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

--
-- Dumping data untuk tabel `spk`
--

INSERT INTO `spk` (`id_spk`, `tanggal_spk`, `id_service`, `id_jasa`, `id_customer`, `id_jenis`, `no_kendaraan`, `keluhan`) VALUES
(1, '2025-07-31', 2, 1, 4, 1, 'B1234XYZ', 'Suara BIsing'),
(2, '2025-08-02', 1, 3, 6, 1, 'B1231BB', 'Telat ganti'),
(5, '2025-08-07', 1, 2, 6, 1, 'B1231BB', 'DAsa'),
(6, '2025-08-10', 1, 13, 8, 2, 'A092321W', 'mandi'),
(7, '2025-08-13', 1, 1, 4, 1, 'B1234XYZ', '');

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
  `telepon` varchar(20) DEFAULT NULL,
  `id_mekanik` int(11) DEFAULT NULL,
  `harga_jasa` int(110) DEFAULT 0,
  `harga_sparepart` int(110) DEFAULT 0,
  `total` int(110) DEFAULT 0,
  `status_pembayaran` varchar(255) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `transaksi`
--

INSERT INTO `transaksi` (`id_transaksi`, `id_spk`, `id_customer`, `id_jenis`, `no_kendaraan`, `telepon`, `id_mekanik`, `harga_jasa`, `harga_sparepart`, `total`, `status_pembayaran`) VALUES
(1, 1, 4, 1, 'B1234XYZ', '08123456789', 1, 75000, 0, 75000, 'lunas'),
(2, 1, 6, 1, 'B1234XYZ', '08123456789', 1, 75000, 120000, 195000, 'lunas'),
(3, 1, 4, 1, 'B1234XYZ', '082123231', 1, 75000, 120000, 195000, 'lunas'),
(4, 1, 4, 1, 'B1234XYZ', '082123231', 1, 75000, 120000, 195000, ''),
(5, 1, 4, 1, 'B1234XYZ', '082123231', 1, 75000, 120000, 195000, ''),
(6, 5, 6, 2, 'B1231BB', '08123311100', 1, 125000, 120000, 245000, '');

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
(8, 'Mas Kera', 'ker@gmail.com', NULL, '$2a$10$IQTQB8Gz6QwQKs0TNJW7ce/5evwDgHCe/KuL.gCd60FByLB7vGJBO', 'gudang', '', '2025-06-23 09:34:03.790', '2025-08-07 11:46:05.138'),
(9, 'Kalas', 'kal@gmail.com', NULL, '$2a$10$9h6WMb8et2/Ax7tXUFKvIOvkmxE5cBbxROKPTRqyPKdF1mh9pVEQu', 'customer', '', '2025-08-11 03:00:28.480', '2025-08-12 10:14:19.036'),
(10, 'jaas', 'jay2@gmail.com', NULL, '$2a$10$5VL1hhMi1Lh6Q3BzeI9xx.ic4RhwHf60P.hxhwBzYZVBvaHaGbLpe', 'customer', '', '2025-08-12 10:16:03.683', '2025-08-13 12:23:04.620');

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
  ADD KEY `id_jasa` (`id_jasa`),
  ADD KEY `fk_detail_transaksi_transaksi` (`id_transaksi`);

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
  MODIFY `id_customer` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

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
  MODIFY `id_spk` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT untuk tabel `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `id_transaksi` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

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
  ADD CONSTRAINT `detail_transaksi_ibfk_5` FOREIGN KEY (`id_jasa`) REFERENCES `jenis_jasa` (`id_jasa`),
  ADD CONSTRAINT `fk_detail_transaksi_transaksi` FOREIGN KEY (`id_transaksi`) REFERENCES `transaksi` (`id_transaksi`) ON DELETE CASCADE ON UPDATE CASCADE;

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
