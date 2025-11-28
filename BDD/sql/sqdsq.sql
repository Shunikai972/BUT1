------------------------------------------------------------
-- 1. CREATION DES TABLES
------------------------------------------------------------

-- SERVICE
DROP TABLE Concerne CASCADE CONSTRAINTS;
DROP TABLE Travail CASCADE CONSTRAINTS;
DROP TABLE Employe CASCADE CONSTRAINTS;
DROP TABLE Projet CASCADE CONSTRAINTS;
DROP TABLE Service CASCADE CONSTRAINTS;
DROP TABLE Adresse CASCADE CONSTRAINTS;

CREATE TABLE Service (
    numS NUMBER PRIMARY KEY,
    nomS VARCHAR2(30)
);

-- EMPLOYE
CREATE TABLE Employe (
    numE NUMBER PRIMARY KEY,
    nomE VARCHAR2(30),
    prenomE VARCHAR2(30),
    ville VARCHAR2(30),
    hebdo NUMBER CHECK (hebdo <= 25),   -- <= 25h/sem
    numS NUMBER,
    CONSTRAINT fk_emp_service FOREIGN KEY (numS)
        REFERENCES Service(numS)
);

-- PROJET
CREATE TABLE Projet (
    numP NUMBER PRIMARY KEY,
    nomP VARCHAR2(30)
);

-- TRAVAIL
CREATE TABLE Travail (
    numE NUMBER,
    numP NUMBER,
    duree NUMBER,
    PRIMARY KEY (numE, numP),   
    CONSTRAINT fk_travail_emp FOREIGN KEY (numE) REFERENCES Employe(numE),
    CONSTRAINT fk_travail_projet FOREIGN KEY (numP) REFERENCES Projet(numP)
);

-- CONCERNE
CREATE TABLE Concerne (
    numS NUMBER,
    numP NUMBER,
    PRIMARY KEY (numS, numP),
    CONSTRAINT fk_concerne_service FOREIGN KEY (numS) REFERENCES Service(numS),
    CONSTRAINT fk_concerne_projet  FOREIGN KEY (numP) REFERENCES Projet(numP)
);

------------------------------------------------------------
-- 2. INSERTIONS
------------------------------------------------------------

INSERT INTO Service VALUES (1, 'RH');
INSERT INTO Service VALUES (2, 'Compta');
INSERT INTO Service VALUES (3, 'Informatique');

INSERT INTO Employe VALUES (10, 'Durand', 'Paul', 'Paris', 20, 1);
INSERT INTO Employe VALUES (20, 'Martin', 'Julie', 'Lyon', 25, 2);
INSERT INTO Employe VALUES (37, 'Dupont', 'Jules', 'Rennes', 18, 3);

INSERT INTO Projet VALUES (101, 'Apollo');
INSERT INTO Projet VALUES (103, 'Zeus');

INSERT INTO Travail VALUES (10, 103, 6);
INSERT INTO Travail VALUES (20, 103, 2);
INSERT INTO Travail VALUES (37, 101, 4);

INSERT INTO Concerne VALUES (2, 103);
INSERT INTO Concerne VALUES (3, 101);

------------------------------------------------------------
-- 3. MISES A JOUR (PARTIE TUPLES)
------------------------------------------------------------

-- Insérer un employé avec hebdo > 25 (doit échouer)
-- INSERT INTO Employe VALUES (50, 'Test', 'Michel', 'Tours', 30, 1);

-- Jules habite maintenant à Nantes
UPDATE Employe SET ville = 'Nantes'
WHERE numE = 37;

-- Diminuer de 2h les employés du service 2
UPDATE Employe SET hebdo = hebdo - 2
WHERE numS = 2;

-- Supprimer l’employé n°37
DELETE FROM Travail WHERE numE = 37;
DELETE FROM Employe WHERE numE = 37;

-- Durée de travail = 4 pour projet 103
UPDATE Travail SET duree = 4
WHERE numP = 103;

------------------------------------------------------------
-- 4. MODIFICATIONS DE STRUCTURE (PARTIE STRUCTURES)
------------------------------------------------------------

-- Ajouter salaire
ALTER TABLE Employe ADD salaire NUMBER;

UPDATE Employe SET salaire = 3000 WHERE numS = 3;
UPDATE Employe SET salaire = 2500 WHERE salaire IS NULL;

-- Ajouter attribut nb dans table Service
ALTER TABLE Service ADD nb NUMBER;

-- Remplir nb = nb d'employés du service
UPDATE Service S
SET nb = (SELECT COUNT(*) FROM Employe E WHERE E.numS = S.numS);

-- Ajouter table Adresse
CREATE TABLE Adresse (
    idA NUMBER PRIMARY KEY,
    rue  VARCHAR2(50),
    ville VARCHAR2(30)
);

-- Ajouter idA dans Employe
ALTER TABLE Employe ADD idA NUMBER;

ALTER TABLE Employe ADD CONSTRAINT fk_emp_adresse
    FOREIGN KEY (idA) REFERENCES Adresse(idA);

-- Insérer adresses et associer employés
INSERT INTO Adresse VALUES (1, '12 rue Victor Hugo', 'Nantes');
INSERT INTO Adresse VALUES (2, '5 avenue Foch', 'Paris');
INSERT INTO Adresse VALUES (3, '23 rue Lafayette', 'Lyon');

UPDATE Employe SET idA = 1 WHERE ville = 'Nantes';
UPDATE Employe SET idA = 2 WHERE ville = 'Paris';
UPDATE Employe SET idA = 3 WHERE ville = 'Lyon';

-- Ajouter trav (total heures travaillées)
ALTER TABLE Employe ADD trav NUMBER;

UPDATE Employe E
SET trav = (
    SELECT NVL(SUM(T.duree), 0)
    FROM Travail T
    WHERE T.numE = E.numE
);

------------------------------------------------------------
COMMIT;
------------------------------------------------------------
