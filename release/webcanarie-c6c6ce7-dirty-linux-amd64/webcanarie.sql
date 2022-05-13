SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `webcanarie`;
CREATE DATABASE `webcanarie` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci */;
USE `webcanarie`;

DROP TABLE IF EXISTS `Prenotazione`;
CREATE TABLE `Prenotazione` (
  `Codice` int(11) NOT NULL AUTO_INCREMENT,
  `Inizio` int(11) NOT NULL,
  `Fine` int(11) NOT NULL,
  PRIMARY KEY (`Codice`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `Prenotazione` (`Codice`, `Inizio`, `Fine`) VALUES
(1, 1650916512, 1651002912);

DROP TABLE IF EXISTS `Stanza`;
CREATE TABLE `Stanza` (
  `Codice` int(11) NOT NULL AUTO_INCREMENT,
  `Nome` varchar(32) NOT NULL,
  `Dotazioni` varchar(2048) NOT NULL,
  PRIMARY KEY (`Codice`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

TRUNCATE `Stanza`;
INSERT INTO `Stanza` (`Codice`, `Nome`, `Dotazioni`) VALUES
(1, 'Cucina ', 'Lorem ipsum dolor sit amet'),
(2, 'Soggiorno ', 'Lorem ipsum dolor sit amet'),
(3, 'Camera matrimoniale', 'Lorem ipsum dolor sit amet'),
(4, 'Seconda camera', 'Lorem ipsum dolor sit amet'),
(5, 'Bagno', 'Lorem ipsum dolor sit amet'),
(6, 'Terrazza', 'Lorem ipsum dolor sit amet');

DROP TABLE IF EXISTS `Immagine`;
CREATE TABLE `Immagine` (
  `Codice` int(11) NOT NULL AUTO_INCREMENT,
  `Percorso` varchar(40),
  `Descrizione` varchar(40) CHARACTER SET utf8 NOT NULL,
  `Stanza` int(11),
  PRIMARY KEY (`Codice`),
  KEY `Stanza` (`Stanza`),
  CONSTRAINT `Immagine_ibfk_1` FOREIGN KEY (`Stanza`) REFERENCES `Stanza` (`Codice`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

TRUNCATE `Immagine`;
INSERT INTO `Immagine` (`Percorso`, `Descrizione`, `Stanza`) VALUES
('/static/img/apartment/cucina1.jpeg',	'Cucina e soggiorno',	1),
('/static/img/apartment/cucina2.jpeg',	'Cucina e soggiorno',	1),
('/static/img/apartment/tv1.jpeg',	'Cucina e soggiorno',	2),
('/static/img/apartment/letto1.jpeg',	'Camera matrimoniale',	3),
('/static/img/apartment/letto2.jpeg',	'Camera matrimoniale',	3),
('/static/img/apartment/letto_luce2.jpeg',	'Camera matrimoniale',	3),
('/static/img/apartment/salotto1.jpeg',	'Seconda camera',	4),
('/static/img/apartment/salotto2.jpeg',	'Seconda camera',	4),
('/static/img/apartment/salotto_luce2.jpeg',	'Seconda camera',	4),
('/static/img/apartment/bagno1.jpeg',	'Bagno',	5),
('/static/img/apartment/piscina.png',	'Piscina',	6),
('/static/img/isle/playagrande.png', 'Playa Grande', NULL),
('/static/img/isle/lanzarote.png', 'Parco nazionale Timanfaya', NULL),
('/static/img/isle/papagayo.png', 'Playa de Papagayo', NULL),
('/static/img/isle/mare1.png', 'Playa de la Cera', NULL),
('/static/img/isle/spiaggia2.png', 'Playa de la Francesa', NULL),
('/static/img/isle/teguise.png', 'Teguise', NULL);


CREATE USER 'webcanarie'@'%' IDENTIFIED BY 'secret';
GRANT ALL PRIVILEGES ON webcanarie.* TO 'webcanarie'@'%';
