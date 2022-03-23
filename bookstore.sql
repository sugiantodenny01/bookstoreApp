-- MySQL dump 10.13  Distrib 5.5.62, for Win64 (AMD64)
--
-- Host: localhost    Database: bookstore
-- ------------------------------------------------------
-- Server version	5.7.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `author`
--

DROP TABLE IF EXISTS `author`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `author` (
  `Author_ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) NOT NULL,
  `Pen_Name` varchar(100) NOT NULL,
  `Email` varchar(100) DEFAULT NULL,
  `Password` varchar(255) NOT NULL,
  `Is_Disabled` tinyint(1) NOT NULL DEFAULT '0',
  `Created_Time` datetime NOT NULL,
  PRIMARY KEY (`Author_ID`),
  UNIQUE KEY `Author_Author_ID_uindex` (`Author_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `author`
--

LOCK TABLES `author` WRITE;
/*!40000 ALTER TABLE `author` DISABLE KEYS */;
INSERT INTO `author` VALUES (2,'dennysugianto1@gmail.com','Denny','dennysugianto1@gmail.com','12345',0,'2022-03-19 07:09:35'),(3,'edo honda','edoHonda','edo1@gmail.com','1234567890',0,'2022-03-19 16:49:26'),(4,'david1@gmail.com','David','david1@gmail.com','12345',0,'2022-03-22 20:00:41');
/*!40000 ALTER TABLE `author` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `book`
--

DROP TABLE IF EXISTS `book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `book` (
  `Book_ID` int(11) NOT NULL AUTO_INCREMENT,
  `Author_ID` int(11) NOT NULL,
  `Title` varchar(255) DEFAULT NULL,
  `Summary` text,
  `Stock` int(11) NOT NULL,
  `Price` int(11) NOT NULL,
  `Cover_URL` varchar(255) NOT NULL,
  `Created_Time` datetime NOT NULL,
  PRIMARY KEY (`Book_ID`),
  UNIQUE KEY `Book_Book_ID_uindex` (`Book_ID`),
  KEY `FK_Author` (`Author_ID`),
  CONSTRAINT `FK_Author` FOREIGN KEY (`Author_ID`) REFERENCES `author` (`Author_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book`
--

LOCK TABLES `book` WRITE;
/*!40000 ALTER TABLE `book` DISABLE KEYS */;
INSERT INTO `book` VALUES (1,3,'The Northdanten Hero','I will see to bla bla',18,25000,'/assets/Detective Conan eps 150.jpg','2022-03-20 11:27:33'),(4,3,'Detective Conan eps 150','I\'m young detective in the world',18,25000,'/assets/Detective Conan eps 150.jpg','2022-03-22 12:48:46'),(5,3,'Beijing','Beijing Ni Hao',18,25000,'/assets/Beijing.jpg','2022-03-22 12:49:42'),(6,2,'Jakarta','Jakarta Cerah',18,25000,'/assets/550a1827e2.jpg','2022-03-22 12:51:29'),(7,2,'Luffy','Luffy Raja Bajak Laut',18,25000,'/assets/Luffy.jpg','2022-03-22 12:52:36'),(8,2,'Zoro','Zoro Samurai',18,25000,'/assets/Zoro.jpg','2022-03-22 12:53:30'),(9,2,'New jakarta','My city',18,25000,'/assets/New jakarta.jpg','2022-03-22 19:57:19');
/*!40000 ALTER TABLE `book` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sales`
--

DROP TABLE IF EXISTS `sales`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sales` (
  `Sales_ID` int(11) NOT NULL AUTO_INCREMENT,
  `Author_ID` int(11) NOT NULL,
  `Recipient_Name` varchar(255) NOT NULL,
  `Recipient_Email` varchar(100) NOT NULL,
  `Book_Title` varchar(255) NOT NULL,
  `Quantity` int(11) NOT NULL,
  `Price_Per_Unit` int(11) NOT NULL,
  `Price_Total` int(11) NOT NULL,
  PRIMARY KEY (`Sales_ID`),
  UNIQUE KEY `Sales_Sales_ID_uindex` (`Sales_ID`),
  KEY `FK_SalesAuthor` (`Author_ID`),
  CONSTRAINT `FK_SalesAuthor` FOREIGN KEY (`Author_ID`) REFERENCES `author` (`Author_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sales`
--

LOCK TABLES `sales` WRITE;
/*!40000 ALTER TABLE `sales` DISABLE KEYS */;
INSERT INTO `sales` VALUES (3,2,'Erick Borgsen','erbog@gmail.com','The Northdanten Hero',2,25000,50000);
/*!40000 ALTER TABLE `sales` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'bookstore'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-03-23  3:59:00
