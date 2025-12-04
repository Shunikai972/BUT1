DROP TABLE Travail CASCADE CONSTRAINTS;
DROP TABLE Service CASCADE CONSTRAINTS;
DROP TABLE Employe CASCADE CONSTRAINTS;
DROP TABLE Projet CASCADE CONSTRAINTS;

CREATE TABLE Employe as select * from basetd.EMPLOYE;
CREATE TABLE Service as select * from basetd.SERVICE;
CREATE TABLE Projet as select * from basetd.PROJET;
CREATE TABLE Travail as select * from basetd.TRAVAIL;

SELECT * FROM Employe;
SELECT * FROM Service;
SELECT * FROM Projet;
SELECT * FROM Travail;

ALTER TABLE Employe ADD CONSTRAINT cleemploye PRIMARY KEY (NUEMPL);
ALTER TABLE Service ADD CONSTRAINT cleservice PRIMARY KEY (NUSERV);
ALTER TABLE Projet ADD CONSTRAINT cleprojet PRIMARY KEY (NUPROJ);

ALTER TABLE Employe ADD CONSTRAINT thebdo CHECK (hebdo <= 35);

ALTER TABLE Employe ADD CONSTRAINT refService FOREIGN KEY (affect) REFERENCES Service(NUSERV);
ALTER TABLE Projet ADD CONSTRAINT refEmploye FOREIGN KEY (resp) REFERENCES Employe(NUEMPL);
ALTER TABLE Travail ADD CONSTRAINT refTravProjet FOREIGN KEY (nuproj) REFERENCES Projet(nuproj);
ALTER TABLE Travail ADD CONSTRAINT refTravEmploye FOREIGN KEY (NUEMPL) REFERENCES Employe(NUEMPL);

---Ex1---
SELECT DISTINCT Employe.nomempl
FROM Employe 
JOIN Travail ON Employe.NUEMPL = Travail.NUEMPL ;

---OU NATURAL JOIN TRAVAIL;---
---Ex2---
SELECT nomempl
FROM Employe
WHERE NUEMPL IN (20, 30, 42);
---Ex3---
SELECT NUPROJ
FROM Projet
WHERE nomproj LIKE 'c%';
---Ex4---
SELECT NUEMPL
FROM Employe
WHERE hebdo IS NULL;
---Ex5---
SELECT e.nomempl, t.duree
FROM Employe e
LEFT JOIN Travail t ON e.NUEMPL = t.NUEMPL;
---Ex6---
SELECT DISTINCT t.nuproj
FROM Travail t
JOIN Employe e ON t.NUEMPL = e.NUEMPL;

---Ex7---
SELECT e.NUEMPL
FROM Employe e
LEFT JOIN Travail t ON e.NUEMPL = t.NUEMPL
WHERE t.NUEMPL IS NULL;
---Exercice 1.2---
---1---
SELECT NUEMPL, NOMEMPL
FROM Employe;
---2---
SELECT COUNT(*) AS nb_employes
FROM Employe;
---3---
SELECT AVG(hebdo) AS moyenne_hebdo
FROM Employe;
---4---
SELECT SUM(duree) AS total_durees
FROM Travail;
---5---
SELECT NOMEMPL
FROM Employe
ORDER BY NOMEMPL ASC;
---6---
SELECT NUEMPL, NUPROJ, DUREE
FROM Travail
ORDER BY NUEMPL DESC;
---7---
SELECT NOMSERV
FROM Service
WHERE NUSERV = 1;
---8---
SELECT NOMSERV
FROM Service
WHERE NUSERV <> 1;
---9---
SELECT e.NOMEMPL
FROM Employe e
LEFT JOIN Travail t ON e.NUEMPL = t.NUEMPL
WHERE t.NUEMPL IS NULL;
---10---
SELECT NOMEMPL
FROM Employe
WHERE hebdo BETWEEN 20 AND 30;
-- exercice 1.3 --

-- 1 -- Liste des noms de services avec le(s) employé(s) affecté(s)
-- Comme il n'y a pas de colonne "resp" dans Service, on prend les employés affectés au service
SELECT s.nomserv, e.nomempl AS employe
FROM Service s
JOIN Employe e ON e.affect = s.nuserv
ORDER BY s.nomserv, e.nomempl;

-- 2 -- Liste des employés avec les projets sur lesquels ils travaillent
SELECT e.nomempl, p.nomproj
FROM Employe e
JOIN Travail t ON e.NUEMPL = t.NUEMPL
JOIN Projet p ON t.NUPROJ = p.NUPROJ
ORDER BY e.nomempl;

-- 3 -- Pour le service "achat", les employés et les projets sur lesquels ils travaillent

SELECT Employe.nomempl AS chef, Projet.nuproj
FROM Service
JOIN Employe ON Service.chef = Employe.nuempl
LEFT JOIN Travail ON Employe.nuempl = Travail.nuempl
LEFT JOIN Projet ON Travail.nuproj = Projet.nuproj
WHERE Service.NOMSERV = 'achat';


-- 4 -- Liste des projets avec le responsable
SELECT p.nomproj, e.nomempl AS responsable
FROM Projet p
JOIN Employe e ON p.resp = e.NUEMPL
ORDER BY p.nomproj;

-- 5 -- Pour le projet "zorro", nom du responsable et employés qui y travaillent
SELECT r.nomempl AS responsable, e.nomempl AS employe
FROM Projet p
JOIN Employe r ON p.resp = r.NUEMPL
JOIN Travail t ON p.NUPROJ = t.NUPROJ
JOIN Employe e ON t.NUEMPL = e.NUEMPL
WHERE p.nomproj = 'zorro'
ORDER BY e.nomempl;

-- 6 -- Tous les employés avec les projets (même s'ils ne sont affectés à aucun projet)
SELECT e.nomempl, p.nomproj
FROM Employe e
LEFT JOIN Travail t ON e.NUEMPL = t.NUEMPL
LEFT JOIN Projet p ON t.NUPROJ = p.NUPROJ
ORDER BY e.nomempl;
--exercice 2--
--EXERCICE 2.1.1--
