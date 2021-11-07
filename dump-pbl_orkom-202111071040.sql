-- MySQL dump 10.13  Distrib 8.0.21, for Linux (x86_64)
--
-- Host: localhost    Database: pbl_orkom
-- ------------------------------------------------------
-- Server version	8.0.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `change_component`
--

DROP TABLE IF EXISTS `change_component`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `change_component` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `id_troubleshooting` int unsigned NOT NULL,
  `id_component` int unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_id_troubleshooting` (`id_troubleshooting`),
  KEY `FK_id_component` (`id_component`),
  CONSTRAINT `FK_id_component` FOREIGN KEY (`id_component`) REFERENCES `components` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FK_id_troubleshooting` FOREIGN KEY (`id_troubleshooting`) REFERENCES `troubleshooting` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=97 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `change_component`
--

LOCK TABLES `change_component` WRITE;
/*!40000 ALTER TABLE `change_component` DISABLE KEYS */;
INSERT INTO `change_component` VALUES (1,1,3),(3,6,2),(4,6,3),(5,7,3),(6,8,2),(7,8,4),(8,9,2),(9,9,3),(10,9,4),(11,10,1),(12,10,2),(13,10,3),(30,16,4),(31,16,3),(32,16,1),(33,16,2),(83,18,4),(84,18,2),(85,18,3),(95,19,3),(96,19,2);
/*!40000 ALTER TABLE `change_component` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `components`
--

DROP TABLE IF EXISTS `components`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `components` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `component` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `components`
--

LOCK TABLES `components` WRITE;
/*!40000 ALTER TABLE `components` DISABLE KEYS */;
INSERT INTO `components` VALUES (1,'RAM','2021-09-28 03:19:30'),(2,'GPU','2021-09-28 03:19:30'),(3,'MONITOR','2021-09-28 03:19:30'),(4,'VGA','2021-09-28 03:19:30');
/*!40000 ALTER TABLE `components` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `detail_maintenance`
--

DROP TABLE IF EXISTS `detail_maintenance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `detail_maintenance` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `id_maintenance` int unsigned NOT NULL,
  `nama_petugas` varchar(100) NOT NULL DEFAULT 'Jhon Doe',
  `lama_pengerjaan` int unsigned NOT NULL DEFAULT '1',
  `informasi_lainnya` text,
  `detail_pengerjaan` text,
  PRIMARY KEY (`id`),
  KEY `FK_id_maintenance` (`id_maintenance`),
  CONSTRAINT `FK_id_maintenance` FOREIGN KEY (`id_maintenance`) REFERENCES `maintenance` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `detail_maintenance`
--

LOCK TABLES `detail_maintenance` WRITE;
/*!40000 ALTER TABLE `detail_maintenance` DISABLE KEYS */;
INSERT INTO `detail_maintenance` VALUES (2,2,'muhammad al farizzi',17,'perferendis autem no','perspiciatis fuga '),(3,3,'hic corporis dolor e',7,'non id eum magni eni','elit non minim inci'),(4,4,'maiores harum invent',2,'nulla nihil fugiat a','aliquid assumenda mo'),(5,5,'aute quos sint offic',60,'deserunt ipsam ut ni','ipsum voluptatem v'),(6,6,'quo minus voluptatem',17,'autem illum quos is','alias est adipisicin');
/*!40000 ALTER TABLE `detail_maintenance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `maintenance`
--

DROP TABLE IF EXISTS `maintenance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `maintenance` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `motherboard` varchar(100) NOT NULL DEFAULT 'Tidak Diketahui',
  `ram_vendor` varchar(100) NOT NULL DEFAULT 'Tidak Diketahui',
  `ram_amount` enum('2','4','6','8','lainnya') DEFAULT NULL,
  `gpu` varchar(100) NOT NULL DEFAULT 'Tidak Diketahui',
  `nic` varchar(100) NOT NULL DEFAULT 'Tidak Diketahui',
  `os` enum('windows','linux','mac') DEFAULT NULL,
  `arsitektur_os` enum('32','64') DEFAULT NULL,
  `tgl_masuk` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `tgl_keluar` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `maintenance`
--

LOCK TABLES `maintenance` WRITE;
/*!40000 ALTER TABLE `maintenance` DISABLE KEYS */;
INSERT INTO `maintenance` VALUES (2,'aerox','gigabyte','8','nvidia 3069 ti','none','linux','64','2021-09-23 14:04:23','2021-09-23 14:04:23'),(3,'pariatur voluptatem','voluptatibus eligend','4','ea vitae id do hic c','velit et ducimus al','windows','64','2021-09-27 04:38:14','2021-09-27 04:38:14'),(4,'enim et consectetur','dolor inventore quod','4','eligendi nam labore ','porro aut amet cupi','linux','64','2021-09-27 14:07:03','2021-09-27 14:07:03'),(5,'accusamus reprehende','neque similique natu','8','enim quis tempore a','ipsum rem cupidatat ','linux','64','2021-10-01 09:28:08','2021-10-01 09:28:08'),(6,'magnam voluptatum vo','in rem anim veniam ','6','fugiat voluptate rep','laudantium at eos n','mac','64','2021-10-01 09:28:17','2021-10-01 09:28:17');
/*!40000 ALTER TABLE `maintenance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `troubleshooting`
--

DROP TABLE IF EXISTS `troubleshooting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `troubleshooting` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `nama_customer` varchar(50) NOT NULL,
  `lama_pengerjaan` int unsigned NOT NULL,
  `biaya` int unsigned NOT NULL DEFAULT '0',
  `permasalahan` text,
  `detail_pengerjaan` text,
  `informasi_lainnya` text,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `troubleshooting`
--

LOCK TABLES `troubleshooting` WRITE;
/*!40000 ALTER TABLE `troubleshooting` DISABLE KEYS */;
INSERT INTO `troubleshooting` VALUES (1,'Quisquam ea quidem a',64,71,'Minim non aut numqua','Eius sed animi poss','Sit aliquid est fa','2021-09-28 09:02:48'),(6,'Id fuga Minim beata',96,89,'Praesentium consecte','Delectus ea non opt','Dignissimos dolorem ','2021-09-30 15:15:19'),(7,'Saepe voluptates rep',15,29,'Rerum ut reiciendis ','Consequatur Volupta','Vel amet sequi maxi','2021-09-30 15:15:26'),(8,'Lorem in officiis et',95,89,'Culpa quia minima ea','Magna id voluptatem','Laborum Suscipit nu','2021-09-30 15:15:31'),(9,'Eu architecto conseq',28,57,'Cupidatat eos dolori','Enim qui deleniti ut','Officia perspiciatis','2021-10-01 04:28:48'),(10,'Architecto optio et',44,15,'Accusamus et ullamco','Voluptatem ad et mag','Laudantium consequa','2021-10-01 04:32:31'),(11,'Aute ea et voluptati',6,29,'Odit fugiat ut iust','Quasi ea cum neque l','Quisquam officia rep','2021-10-01 04:42:13'),(16,'Autem aut eu volupta',2,24,'Sint officiis possim','Sint ut aut nostrud ','A omnis voluptatum c','2021-10-01 04:49:17'),(17,'Repudiandae reprehen',14,33,'Mollitia sit quisqua','Sunt et voluptas odi','Aut consequatur offi','2021-10-01 04:49:48'),(18,'Ujang',30,31,'Tempore magnam aut ','Eum sunt voluptate a','Voluptatem porro exp','2021-10-01 06:41:54'),(19,'Consequatur Fuga T',45,37,'Nihil quis et irure ','Natus nostrud maxime','Delectus itaque com','2021-10-01 07:06:07');
/*!40000 ALTER TABLE `troubleshooting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','$2a$10$Vb.jVKxey7HiGWbtc26sTe0mzgr0XES8CEUN8J53QfFpJ75iHeERu','2021-10-03 02:32:52');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'pbl_orkom'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-11-07 10:40:35
