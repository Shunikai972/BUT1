-- Script complet (à exécuter en tant que UTILISATEUR user2)
-- Remarques :
--  - je suis user2 (nom Oracle : USER2) ; le binôme est s2a01b (nom Oracle : s2a01b).
--  - adapter les noms d'utilisateurs si nécessaire.
--  - ce fichier contient : création tables (Vendée), contraintes, vues, grants, tests, rôles, distribution.
--  - exécutez les commandes dans l'ordre. Certaines actions (GRANT REFERENCES sur OperateurLA) doivent être faites par s2a01b (indiqué en commentaire).

/**************************************************************************
-- Préambule : nettoyage (on ignore les erreurs si l'objet n'existe pas)

/* ============================
-- Exercice 1 : Droits Select / Insert / Update
-- (Vous êtes USER2 ; vous créez les tables pour la Vendée et testez les droits)
============================ */

/* Création des tables pour le département Vendée (CommuneVE, OperateurVE, DistributionVE)
   à partir du schéma basetd (on suppose basetd.commune, basetd.operateur, basetd.distribution existent).
   Adapter noms colonnes si nécessaire (ici on assume : COMMUNE(CODE_INSEE, NOMCOMMUNE, NOMDEP),
   OPERATEUR(NUMFO, NOMFO, GENERATION, TECHNOLOGIE), DISTRIBUTION(ID, NUMFO, CODE_INSEE, ADRESSE, STATUT, ADR_NM_CP, EMR_DT_SERVICE, COORDONNEESL, COORDONNEESLT, COORD)
*/

-- 1) Création CommuneVE
CREATE TABLE CommuneVE AS
SELECT *
FROM basetd.commune
WHERE nomdep = 'Vendée';
-- Si la colonne s'appelle NOM_DEP ou autre, adapter la condition.

-- 2) Création OperateurVE (opérateurs 4G et 5G)
CREATE TABLE OperateurVE AS
SELECT *
FROM basetd.operateur
WHERE generation IN ('4G','5G');

-- 3) Création DistributionVE (liée aux communes Vendée et aux opérateurs 4G/5G)
CREATE TABLE DistributionVE AS
SELECT d.*
FROM basetd.distribution d
WHERE d.code_insee IN (SELECT code_insee FROM CommuneVE)
  AND d.numfo IN (SELECT numfo FROM OperateurVE);


-- 4) Ajout des clés primaires et étrangères locales (ex : PK sur CommuneVE.code_insee, OperateurVE.numfo, DistributionVE.id)
ALTER TABLE CommuneVE ADD CONSTRAINT PK_CommuneVE PRIMARY KEY (code_insee);
ALTER TABLE OperateurVE ADD CONSTRAINT PK_OperateurVE PRIMARY KEY (numfo);

-- Supposons que DistributionVE a colonne ID comme PK ; sinon utiliser une combinaison (NUMFO, CODE_INSEE, ID)
ALTER TABLE DistributionVE ADD CONSTRAINT PK_DistributionVE PRIMARY KEY (id);

-- FKs locales : DistributionVE.code_insee -> CommuneVE.code_insee ; DistributionVE.numfo -> OperateurVE.numfo
ALTER TABLE DistributionVE
  ADD CONSTRAINT FK_DistribVE_CommuneVE FOREIGN KEY (code_insee)
    REFERENCES CommuneVE (code_insee);

ALTER TABLE DistributionVE
  ADD CONSTRAINT FK_DistribVE_OperateurVE FOREIGN KEY (numfo)
    REFERENCES OperateurVE (numfo);
    
    

/* ------------ Tests d'accès (énoncés / commandes à faire par s2a01b pour vérifier) ------------
   USER2 va accorder les droits demandés à s2a01b ; ensuite s2a01b exécutera les tests indiqués.
----------------------------------------------------------------------------------------------- */

-- 1) Donnez le droit de faire des SELECT sur vos tables à votre binôme (s2a01b)
GRANT SELECT ON CommuneVE TO s2a01b;
GRANT SELECT ON OperateurVE TO s2a01b;
GRANT SELECT ON DistributionVE TO s2a01b;

-- Tests (à exécuter en tant que s2a01b) :
-- SELECT COUNT(*) FROM USER2.CommuneVE;
-- SELECT * FROM USER2.DistributionVE WHERE ROWNUM <= 5;
-- (exemples) :
-- SELECT COUNT(*) FROM s2a01b.CommuneVE;
-- SELECT * FROM s2a01b.DistributionVE WHERE ROWNUM <= 5;

-- 2) Donnez le droit d'UPDATE du champ ADRESSE de la table DistributionVE à s2a01b
GRANT UPDATE (adresse) ON DistributionVE TO s2a01b;

-- Test (s2a01b) :
-- UPDATE USER2.DistributionVE SET adresse = adresse || ' -- TEST' WHERE id = (SELECT id FROM USER2.DistributionVE WHERE ROWNUM = 1);
-- -- vérifier avant commit
-- SELECT adresse FROM USER2.DistributionVE WHERE id = <id>;
-- COMMIT;
-- -- vérifier après commit
-- SELECT adresse FROM USER2.DistributionVE WHERE id = <id>;

-- 3) Donnez le droit d'INSERT à la table DistributionVE
GRANT INSERT ON DistributionVE TO s2a01b;
-- Test (s2a01b) :
-- INSERT INTO USER2.DistributionVE (id, numfo, code_insee, adresse, statut) VALUES (999999, (SELECT numfo FROM USER2.OperateurVE WHERE ROWNUM=1), (SELECT code_insee FROM USER2.CommuneVE WHERE ROWNUM=1), 'Adresse test', 'TEST');
-- -- vérifier avant commit
-- SELECT * FROM USER2.DistributionVE WHERE id = 999999;
-- COMMIT;
-- SELECT * FROM USER2.DistributionVE WHERE id = 999999;
-- -- puis cleanup (s2a01b ou USER2) :
-- DELETE FROM USER2.DistributionVE WHERE id = 999999; COMMIT;

-- 4) Supprimez (REVOKE) les droits donnés en 1,2,3
REVOKE SELECT ON CommuneVE FROM s2a01b;
REVOKE SELECT ON OperateurVE FROM s2a01b;
REVOKE SELECT ON DistributionVE FROM s2a01b;

REVOKE UPDATE (adresse) ON DistributionVE FROM s2a01b;
REVOKE INSERT ON DistributionVE FROM s2a01b;

COMMIT;

-- Test (s2a01b doit vérifier qu'il n'a plus accès) :
-- SELECT * FROM USER2.CommuneVE;  -- doit échouer (permission denied)
-- UPDATE USER2.DistributionVE SET adresse = 'x' WHERE id = <id>; -- doit échouer
-- INSERT INTO USER2.DistributionVE (...) VALUES (...); -- doit échouer

-- 5) Donner le droit de voir tous les attributs sauf (ADR_NM_CP, EMR_DT_SERVICE, COORDONNEESL, COORDONNEESLT, COORD)
-- On crée une vue qui exclut ces colonnes et on donne SELECT sur la vue.
CREATE OR REPLACE VIEW DistributionVE_NO_SENSITIVE AS
SELECT id, numfo, code_insee, adresse, statut
       -- ajouter ici d'autres colonnes si présentes et si non sensibles
FROM DistributionVE;

GRANT SELECT ON DistributionVE_NO_SENSITIVE TO s2a01b;

-- Test (s2a01b) :
SELECT * FROM s2a01b.DistributionLA_vue; -- doit montrer colonnes sans les champs sensibles
SELECT * FROM s2a01b.DistributionLA_vue; -- doit échouer (si SELECT sur table non accordé)

-- 6) Donnez le droit à votre binôme de voir le nombre d'antennes 4G et 5G des villes de votre département via 2 vues différentes
CREATE OR REPLACE VIEW V_COUNT_ANTENNES_4G AS
SELECT c.nom_commune,
       COUNT(d.id) AS nb_4g
FROM DistributionVE d
JOIN CommuneVE c ON d.code_insee = c.code_insee
JOIN OperateurVE o ON d.numfo = o.numfo
WHERE o.generation = '4G'
GROUP BY c.nom_commune;

CREATE OR REPLACE VIEW V_COUNT_ANTENNES_5G AS
SELECT c.nom_commune,
       COUNT(d.id) AS nb_5g
FROM DistributionVE d
JOIN CommuneVE c ON d.code_insee = c.code_insee
JOIN OperateurVE o ON d.numfo = o.numfo
WHERE o.generation = '5G'
GROUP BY c.nom_commune;

GRANT SELECT ON V_COUNT_ANTENNES_4G TO s2a01b;
GRANT SELECT ON V_COUNT_ANTENNES_5G TO s2a01b;

-- Test (s2a01b) le fils du pauvre :
SELECT * FROM s2a01b.Nb_Antennes_4G_Ville_LA WHERE ROWNUM <= 10;
SELECT * FROM s2a01b.Nb_Antennes_5G_Ville_LA WHERE ROWNUM <= 10;

-- 7) Donnez le droit de voir le nombre d'antennes 4G et 5G de chaque opérateur pour les différentes villes.
CREATE OR REPLACE VIEW V_OP_CITY_BY_TECH AS
SELECT o.nomfo AS operateur,
       c.nom_commune AS commune,
       o.generation AS technologie,
       COUNT(d.id) AS nb_antennes
FROM DistributionVE d
JOIN OperateurVE o ON d.numfo = o.numfo
JOIN CommuneVE c ON d.code_insee = c.code_insee
GROUP BY o.nomfo, c.nom_commune, o.generation;

GRANT SELECT ON V_OP_CITY_BY_TECH TO s2a01b;

-- Test (s2a01b) :
SELECT * FROM s2a01b.Nb_Antennes_Operateur_Ville WHERE operateur = 'ORANGE' AND ROWNUM <= 10;

-- 8) Répartition des différentes technologies (ex : part 4G vs 5G par ville / opérateur)
CREATE OR REPLACE VIEW V_TECH_REPARTITION AS
SELECT c.nom_commune,
       o.nomfo AS operateur,
       o.generation AS technologie,
       COUNT(d.id) AS nb_antennes,
       (COUNT(d.id) / SUM(COUNT(d.id)) OVER (PARTITION BY c.nom_commune)) * 100 AS pct_par_ville
FROM DistributionVE d
JOIN OperateurVE o ON d.numfo = o.numfo
JOIN CommuneVE c ON d.code_insee = c.code_insee
GROUP BY c.nom_commune, o.nomfo, o.generation;

GRANT SELECT ON V_TECH_REPARTITION TO s2a01b;

-- Test (s2a01b) :
SELECT * FROM s2a01b.Nb_Antennes_Operateur_Ville WHERE ROWNUM <= 20;

-- 9) Top 20 meilleures villes avec la 5G de l'opérateur ORANGE ; ajouter nb 4G
CREATE OR REPLACE VIEW V_TOP20_ORANGE_5G AS
SELECT v5g.nom_commune,
       v5g.nb_5g,
       NVL(v4g.nb_4g,0) AS nb_4g
FROM (
  SELECT c.nom_commune, COUNT(d.id) AS nb_5g
  FROM DistributionVE d
  JOIN OperateurVE o ON d.numfo = o.numfo
  JOIN CommuneVE c ON d.code_insee = c.code_insee
  WHERE o.nomfo = 'ORANGE' AND o.generation = '5G'
  GROUP BY c.nom_commune
  ORDER BY COUNT(d.id) DESC
  FETCH FIRST 20 ROWS ONLY
) v5g
LEFT JOIN (
  SELECT c.nom_commune, COUNT(d.id) AS nb_4g
  FROM DistributionVE d
  JOIN OperateurVE o ON d.numfo = o.numfo
  JOIN CommuneVE c ON d.code_insee = c.code_insee
  WHERE o.nomfo = 'ORANGE' AND o.generation = '4G'
  GROUP BY c.nom_commune
) v4g ON v5g.nom_commune = v4g.nom_commune;

GRANT SELECT ON V_TOP20_ORANGE_5G TO s2a01b;

-- Test (s2a01b) :
-- SELECT * FROM USER2.V_TOP20_ORANGE_5G;

/* ============================
-- Exercice 2 : Rôles
============================ */

-- Avant de créer les rôles, supprimer tous les droits que vous avez donnés à votre binôme (déjà fait ci-dessus).
-- 1) Créez un rôle LoginMon_Ami (préfixé par votre login : USER2)
CREATE ROLE USER2_LoginMon_Ami;

-- 2) Affectez à ce rôle les privilèges correspondant aux questions 1,2 et 3 de l'exercice 1.
--   -> privileges : SELECT sur mes tables (CommuneVE, OperateurVE, DistributionVE),
--                  UPDATE (adresse) sur DistributionVE,
--                  INSERT sur DistributionVE.

GRANT SELECT ON CommuneVE TO USER2_LoginMon_Ami;
GRANT SELECT ON OperateurVE TO USER2_LoginMon_Ami;
GRANT SELECT ON DistributionVE TO USER2_LoginMon_Ami;

GRANT UPDATE (adresse) ON DistributionVE TO USER2_LoginMon_Ami;
GRANT INSERT ON DistributionVE TO USER2_LoginMon_Ami;

select * from s2a01b.communela;

-- 3) Activez et affectez ce rôle à votre binôme (s2a01b). NOTE : s2a01b devra se déconnecter/reconnecter pour voir le rôle actif.
GRANT USER2_LoginMon_Ami TO s2a01b;

-- 4) Supprimer le privilège de SELECT des différentes tables au rôle loginMon_Ami.
REVOKE SELECT ON CommuneVE FROM USER2_LoginMon_Ami;
REVOKE SELECT ON OperateurVE FROM USER2_LoginMon_Ami;
REVOKE SELECT ON DistributionVE FROM USER2_LoginMon_Ami;

-- 5) Supprimer le rôle LoginMon_Ami à votre binôme (retirer le rôle de l'utilisateur)
REVOKE USER2_LoginMon_Ami FROM s2a01b;
-- (ou DROP ROLE USER2_LoginMon_Ami si vous souhaitez le supprimer complètement)
DROP ROLE USER2_LoginMon_Ami;

COMMIT;

-- 6) Créez un rôle LoginVoir_mes_tables (privleges de la question 1) et LoginUpdate_mes_tables (privleges de questions 2 & 3)
CREATE ROLE USER2_LoginVoir_mes_tables;
CREATE ROLE USER2_LoginUpdate_mes_tables;

-- Donner les privilèges :
GRANT SELECT ON CommuneVE TO USER2_LoginVoir_mes_tables;
GRANT SELECT ON OperateurVE TO USER2_LoginVoir_mes_tables;
GRANT SELECT ON DistributionVE TO USER2_LoginVoir_mes_tables;

GRANT UPDATE (adresse) ON DistributionVE TO USER2_LoginUpdate_mes_tables;
GRANT INSERT ON DistributionVE TO USER2_LoginUpdate_mes_tables;

-- 7) Créez un rôle LoginVoir_Update qui regroupe les deux rôles précédents.
CREATE ROLE USER2_LoginVoir_Update;

-- on "donne" les rôles aux rôle composite : (grant role to role)
GRANT USER2_LoginVoir_mes_tables TO USER2_LoginVoir_Update;
GRANT USER2_LoginUpdate_mes_tables TO USER2_LoginVoir_Update;

-- 8) Affectez ce rôle à votre binôme et testez
GRANT USER2_LoginVoir_Update TO s2a01b;

-- Tests (s2a01b) : après reconnexion
-- SET ROLE USER2_LoginVoir_Update;    -- (dans Oracle, user peut SET ROLE ... si rôle est activé)
-- SELECT * FROM USER2.CommuneVE;
-- UPDATE USER2.DistributionVE SET adresse = 'MAJ ROLE' WHERE ROWNUM = 1; COMMIT;
-- INSERT INTO USER2.DistributionVE (id, numfo, code_insee, adresse, statut) VALUES (888888, (SELECT numfo FROM USER2.OperateurVE WHERE ROWNUM=1), (SELECT code_insee FROM USER2.CommuneVE WHERE ROWNUM=1), 'Ins role', 'TEST'); COMMIT;

-- 9) Supprimer le rôle LoginUpdate_mes_tables du rôle LoginVoir_Update et vérifiez les actions.
REVOKE USER2_LoginUpdate_mes_tables FROM USER2_LoginVoir_Update;
-- Maintenant s2a01b (ayant le rôle composite) n'aura plus les privilèges d'INSERT/UPDATE via le rôle composite.

-- Test (s2a01b) :
-- SET ROLE USER2_LoginVoir_Update;
-- UPDATE USER2.DistributionVE SET adresse = 'X' WHERE ROWNUM = 1; -- doit échouer
-- INSERT INTO USER2.DistributionVE (...) VALUES (...); -- doit échouer
-- SELECT * FROM USER2.CommuneVE; -- doit toujours fonctionner (si LoginVoir_mes_tables toujours actif)

/* ============================
-- Exercice 3 : Distribution (Partage) des tables
============================ */

-- Suppression de toutes les tables créées localement (pour repartir à zéro)
DROP TABLE DistributionVE CASCADE CONSTRAINTS PURGE;
DROP TABLE CommuneVE CASCADE CONSTRAINTS PURGE;
DROP TABLE OperateurVE CASCADE CONSTRAINTS PURGE;

COMMIT;

-- L'objectif : répartir les tables entre s2a01b et USER2
-- - s2a01b : OperateurLA, DistributionLA, CommuneLA  (créées sur s2a01b)
-- - USER2 : DistributionVE, CommuneVE (créées sur USER2)
-- Dans ce scénario la table OperateurLA n'est pas disponible sur USER2 ; elle est détenue par s2a01b.
-- Pour que USER2 puisse créer une table DistributionVE qui référence OperateurLA (clé étrangère) il faut que s2a01b
-- donne au minimum le privilège REFERENCES (ou SELECT+REFERENCES) sur s2a01b.OperateurLA à USER2.

-- 1) USER2 : (re)crée CommuneVE et DistributionVE à partir de basetd (pour les villes de Vendée)
CREATE TABLE CommuneVE AS
SELECT *
FROM basetd.commune
WHERE nomdep = 'Vendée';

-- Opérateurs : dans ce scénario OperateurLA est sur s2a01b ; USER2 peut cependant créer un sous-ensemble d'OperateurVE local
-- si besoin. Ici on suppose qu'OperateurLA est sur s2a01b et contient les opérateurs de Loire-Atlantique.
-- Si USER2 veut référencer s2a01b.OperateurLA dans une FK distante, s2a01b doit exécuter (sur son compte) :
--    GRANT REFERENCES ON OperateurLA TO USER2;
-- Optionnellement :
--    GRANT SELECT ON OperateurLA TO USER2;

-- 2) USER2 : création de DistributionVE en utilisant s2a01b.OperateurLA (nécessite GRANT REFERENCES de s2a01b)
-- (si s2a01b a donné REFERENCES, on peut créer la table avec la colonne numfo et ensuite créer la contrainte FK distante)
CREATE TABLE DistributionVE AS
SELECT d.*
FROM basetd.distribution d
WHERE d.code_insee IN (SELECT code_insee FROM CommuneVE)
  AND d.numfo IN (SELECT numfo FROM s2a01b.OperateurLA);  -- nécessite SELECT (ou s2a01b doit autoriser)

COMMIT;

-- 3) Création des contraintes d'intégrité PK et FK locales (sur USER2)
ALTER TABLE CommuneVE ADD CONSTRAINT PK_CommuneVE PRIMARY KEY (code_insee);
ALTER TABLE DistributionVE ADD CONSTRAINT PK_DistributionVE PRIMARY KEY (id);

-- FK locale vers CommuneVE
ALTER TABLE DistributionVE
  ADD CONSTRAINT FK_DistribVE_CommuneVE_local FOREIGN KEY (code_insee)
    REFERENCES CommuneVE (code_insee);

COMMIT;

-- 4) Création des contraintes distantes : DistributionVE.numfo -> s2a01b.OperateurLA.numfo
-- Pour créer cette contrainte distante, s2a01b doit faire (sur son compte) :
--    GRANT REFERENCES ON OperateurLA TO USER2;
-- (ou GRANT SELECT, REFERENCES ON s2a01b.OperateurLA TO USER2; si vous voulez aussi permettre select)
--
-- Une fois que s2a01b a donné REFERENCES, USER2 peut créer la contrainte comme suit :

-- (Cette commande échouera si s2a01b n'a pas donné REFERENCES)
ALTER TABLE DistributionVE
  ADD CONSTRAINT FK_DistribVE_OperateurLA_remote FOREIGN KEY (numfo)
    REFERENCES s2a01b.OperateurLA (numfo);

COMMIT;

-- 5) Refaire les mêmes questions (6,7,8 de l'exercice1) avec la gestion des droits.
-- En plus des SELECT sur les tables, donner aussi le droit de voir les différentes vues créées sur votre compte.

-- Création (ou recréation) des vues (mêmes vues mais maintenant DistributionVE/CommuneVE peuvent référencer s2a01b.OperateurLA)
CREATE OR REPLACE VIEW V_COUNT_ANTENNES_4G AS
SELECT c.nomcommune,
       COUNT(d.id) AS nb_4g
FROM DistributionVE d
JOIN CommuneVE c ON d.code_insee = c.code_insee
JOIN s2a01b.OperateurLA o ON d.numfo = o.numfo
WHERE o.generation = '4G'
GROUP BY c.nomcommune;

CREATE OR REPLACE VIEW V_COUNT_ANTENNES_5G AS
SELECT c.nomcommune,
       COUNT(d.id) AS nb_5g
FROM DistributionVE d
JOIN CommuneVE c ON d.code_insee = c.code_insee
JOIN s2a01b.OperateurLA o ON d.numfo = o.numfo
WHERE o.generation = '5G'
GROUP BY c.nomcommune;

CREATE OR REPLACE VIEW V_OP_CITY_BY_TECH AS
SELECT o.nomfo AS operateur,
       c.nomcommune AS commune,
       o.generation AS technologie,
       COUNT(d.id) AS nb_antennes
FROM DistributionVE d
JOIN s2a01b.OperateurLA o ON d.numfo = o.numfo
JOIN CommuneVE c ON d.code_insee = c.code_insee
GROUP BY o.nomfo, c.nomcommune, o.generation;

CREATE OR REPLACE VIEW V_TECH_REPARTITION AS
SELECT c.nomcommune,
       o.nomfo AS operateur,
       o.generation AS technologie,
       COUNT(d.id) AS nb_antennes,
       (COUNT(d.id) / SUM(COUNT(d.id)) OVER (PARTITION BY c.nomcommune)) * 100 AS pct_par_ville
FROM DistributionVE d
JOIN s2a01b.OperateurLA o ON d.numfo = o.numfo
JOIN CommuneVE c ON d.code_insee = c.code_insee
GROUP BY c.nomcommune, o.nomfo, o.generation;

COMMIT;

-- Donner les droits SELECT sur ces vues à s2a01b :
GRANT SELECT ON V_COUNT_ANTENNES_4G TO s2a01b;
GRANT SELECT ON V_COUNT_ANTENNES_5G TO s2a01b;
GRANT SELECT ON V_OP_CITY_BY_TECH TO s2a01b;
GRANT SELECT ON V_TECH_REPARTITION TO s2a01b;

-- Tests (s2a01b doit exécuter) :
-- SELECT * FROM USER2.V_COUNT_ANTENNES_4G WHERE ROWNUM <= 10;
-- SELECT * FROM USER2.V_OP_CITY_BY_TECH WHERE operateur = 'ORANGE' AND ROWNUM <= 10;

COMMIT;

/**************************************************************************
-- Fin du script
-- Notes / rappel d'actions qui doivent être faites par s2a01b :
--  - Sur s2a01b (compte du binôme) :
--      1) Créer / conserver OperateurLA, CommuneLA, DistributionLA si demandé.
--      2) Pour permettre à USER2 de créer FK distantes vers s2a01b.OperateurLA, s2a01b doit exécuter :
--           GRANT REFERENCES ON OperateurLA TO USER2;
--         (et éventuellement GRANT SELECT ON OperateurLA TO USER2;)
--      3) Tester les connexions et rôles côté s2a01b (SET ROLE, SELECT, UPDATE, INSERT selon cas).
--
--  - Remarques Oracle :
--      * Selon la version Oracle, la syntaxe FETCH FIRST N ROWS ONLY fonctionne (12c+).
--      * Si certains noms de colonnes diffèrent (ex. NOM_DEP vs NOMDEP), adaptez les WHERE et SELECT.
--      * Certaines opérations (GRANT ROLE TO s2a01b) nécessitent la connexion d'administration ou création/gestion rôles par l'utilisateur.
--
--  - Pour le rapport : intégrer toutes les commandes SQL exécutées, captures (résultats SELECT),
--    et indiquer qui (s2a01b / USER2) a exécuté chaque test. Garder la symétrie des exercices 1 & 2.
**************************************************************************/