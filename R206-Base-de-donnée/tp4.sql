-- ============================================================
--  TD4 - Exercice 3 CORRIGÉ
--  User2 : S2A01A (Vendée)   |   Binôme : S2A01B (Loire-Atlantique)
--  Colonnes réelles :
--    CommuneVE    : CODE_INSEE, NOM_COMMUNE, NOMDEP
--    DistributionVE : ID, NUMFO, CODE_INSEE, ADRESSE, ADR_NM_CP,
--                     STATUT, EMR_DT_SERVICE, COORDONNEESL, COORDONNEESLT, COORD
--    OperateurLA  : NUMFO, NOMFO, GENERATION, TECHNOLOGIE
-- ============================================================


-- ============================================================
-- NETTOYAGE
-- ============================================================
DROP TABLE DistributionVE CASCADE CONSTRAINTS PURGE;
DROP TABLE CommuneVE      CASCADE CONSTRAINTS PURGE;

DROP VIEW v_4G_villes_VE;
DROP VIEW v_5G_villes_VE;
DROP VIEW v_4G_operateur_VE;
DROP VIEW v_5G_operateur_VE;
DROP VIEW v_4G_techno_VE;
DROP VIEW v_5G_techno_VE;


-- ============================================================
-- Q2 : Création des tables (Vendée) en utilisant OperateurLA de S2A01B
-- (S2A01B doit avoir fait : GRANT SELECT ON OperateurLA TO S2A01A)
-- ============================================================

CREATE TABLE CommuneVE AS
    SELECT * FROM basetd.commune
    WHERE nomdep = 'Vendée';

CREATE TABLE DistributionVE AS
    SELECT * FROM basetd.distribution
    WHERE code_insee IN (SELECT code_insee FROM CommuneVE)
      AND numfo      IN (SELECT numfo FROM S2A01B.OperateurLA);


-- ============================================================
-- Q3 : Contraintes locales (PK et FK locales)
-- ============================================================

ALTER TABLE CommuneVE      ADD CONSTRAINT pk_communeVE      PRIMARY KEY (code_insee);
ALTER TABLE DistributionVE ADD CONSTRAINT pk_distributionVE PRIMARY KEY (id);

ALTER TABLE DistributionVE ADD CONSTRAINT fk_distVE_commune
    FOREIGN KEY (code_insee) REFERENCES CommuneVE(code_insee);


-- ============================================================
-- Q4 : Contrainte FK distante vers OperateurLA de S2A01B
-- (S2A01B doit avoir fait : GRANT REFERENCES ON OperateurLA TO S2A01A)
-- ============================================================

ALTER TABLE DistributionVE ADD CONSTRAINT fk_distVE_operateurLA
    FOREIGN KEY (numfo) REFERENCES S2A01B.OperateurLA(numfo);


-- ============================================================
-- Q5 : Vues (questions 6, 7, 8 de l'Exercice 1) avec les vrais noms
-- ============================================================

-- Q6 : Nombre d'antennes 4G par ville
CREATE OR REPLACE VIEW v_4G_villes_VE AS
    SELECT c.NOM_COMMUNE, COUNT(*) AS nb_antennes_4G
    FROM DistributionVE d
    JOIN CommuneVE c          ON d.code_insee = c.code_insee
    JOIN S2A01B.OperateurLA o ON d.numfo      = o.numfo
    WHERE o.generation = '4G'
    GROUP BY c.NOM_COMMUNE;

-- Q6 : Nombre d'antennes 5G par ville
CREATE OR REPLACE VIEW v_5G_villes_VE AS
    SELECT c.NOM_COMMUNE, COUNT(*) AS nb_antennes_5G
    FROM DistributionVE d
    JOIN CommuneVE c          ON d.code_insee = c.code_insee
    JOIN S2A01B.OperateurLA o ON d.numfo      = o.numfo
    WHERE o.generation = '5G'
    GROUP BY c.NOM_COMMUNE;

-- Q7 : Nombre d'antennes 4G par opérateur et par ville
CREATE OR REPLACE VIEW v_4G_operateur_VE AS
    SELECT c.NOM_COMMUNE, o.NOMFO, COUNT(*) AS nb_antennes_4G
    FROM DistributionVE d
    JOIN CommuneVE c          ON d.code_insee = c.code_insee
    JOIN S2A01B.OperateurLA o ON d.numfo      = o.numfo
    WHERE o.generation = '4G'
    GROUP BY c.NOM_COMMUNE, o.NOMFO;

-- Q7 : Nombre d'antennes 5G par opérateur et par ville
CREATE OR REPLACE VIEW v_5G_operateur_VE AS
    SELECT c.NOM_COMMUNE, o.NOMFO, COUNT(*) AS nb_antennes_5G
    FROM DistributionVE d
    JOIN CommuneVE c          ON d.code_insee = c.code_insee
    JOIN S2A01B.OperateurLA o ON d.numfo      = o.numfo
    WHERE o.generation = '5G'
    GROUP BY c.NOM_COMMUNE, o.NOMFO;

-- Q8 : Répartition par technologie 4G
CREATE OR REPLACE VIEW v_4G_techno_VE AS
    SELECT c.NOM_COMMUNE, o.NOMFO, o.TECHNOLOGIE, COUNT(*) AS nb_antennes
    FROM DistributionVE d
    JOIN CommuneVE c          ON d.code_insee = c.code_insee
    JOIN S2A01B.OperateurLA o ON d.numfo      = o.numfo
    WHERE o.generation = '4G'
    GROUP BY c.NOM_COMMUNE, o.NOMFO, o.TECHNOLOGIE;

-- Q8 : Répartition par technologie 5G
CREATE OR REPLACE VIEW v_5G_techno_VE AS
    SELECT c.NOM_COMMUNE, o.NOMFO, o.TECHNOLOGIE, COUNT(*) AS nb_antennes
    FROM DistributionVE d
    JOIN CommuneVE c          ON d.code_insee = c.code_insee
    JOIN S2A01B.OperateurLA o ON d.numfo      = o.numfo
    WHERE o.generation = '5G'
    GROUP BY c.NOM_COMMUNE, o.NOMFO, o.TECHNOLOGIE;


-- ============================================================
-- Droits SELECT sur les tables ET les vues pour S2A01B
-- ============================================================
GRANT SELECT ON CommuneVE          TO S2A01B;
GRANT SELECT ON DistributionVE     TO S2A01B;
GRANT SELECT ON v_4G_villes_VE     TO S2A01B;
GRANT SELECT ON v_5G_villes_VE     TO S2A01B;
GRANT SELECT ON v_4G_operateur_VE  TO S2A01B;
GRANT SELECT ON v_5G_operateur_VE  TO S2A01B;
GRANT SELECT ON v_4G_techno_VE     TO S2A01B;
GRANT SELECT ON v_5G_techno_VE     TO S2A01B;