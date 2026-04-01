--s2a01a.borne_de_recharge
--4 tables
--opérateur
--aménageur
--station
--borne
--3NF
-- ============================================================
-- TABLE OPERATEUR
-- ============================================================
select * from s2a01b.borne_de_recharge;

create table borne as
select * from s2a01a.borne;
DESC s2a01b.borne_de_recharge;
select * from s2a01a.operateur;





SET t.id_operateur_fk = (
    SELECT o.id_operateur 
    FROM OPERATEUR o
    WHERE o.nom_operateur = t.nom_operateur
      AND (o.contact_operateur = t.contact_operateur OR (o.contact_operateur IS NULL AND t.contact_operateur IS NULL))
      AND (o.telephone_operateur = t.telephone_operateur OR (o.telephone_operateur IS NULL AND t.telephone_operateur IS NULL))
);

DROP TABLE OPERATEUR;



CREATE TABLE operateur (
    id_operateur NUMBER GENERATED ALWAYS AS IDENTITY(START WITH 1 INCREMENT BY 1),
    nom_operateur       VARCHAR2(128),
    contact_operateur   VARCHAR2(128),
    telephone_operateur VARCHAR2(128)
);

INSERT INTO operateur (nom_operateur, contact_operateur, telephone_operateur)
SELECT DISTINCT nom_operateur, contact_operateur, telephone_operateur
FROM borne;

COMMIT;


grant all on Operateur to S2a01b;
grant all on Operateur to S2a02b;
grant all on Operateur to S2a03a;
select * from operateur;    




-- PK sur operateur
ALTER TABLE operateur ADD CONSTRAINT pk_operateur PRIMARY KEY (id_operateur);



-- S2a01b a besoin de référencer operateur depuis station
GRANT SELECT, REFERENCES ON operateur TO S2a01b;

-- Les autres ont besoin de SELECT pour les requêtes
GRANT SELECT ON operateur TO S2a02b, S2a03a;


 commit;
 
 -- Regarde les colonnes de borne_de_recharge
DESC S2a01b.borne_de_recharge;



DESC S2a01b.Station;
DESC S2a02b.Amenageur;
DESC S2a03a.Commune;
DESC S2a03a.Borne;
Desc Operateur;
select * from S2a03a.Borne;
Select * from S2a03a.Commune;
 -- 5
SELECT o.nom_operateur,
       c.libelle_departement AS departement,
       COUNT(*) AS nb_prises
FROM operateur o
JOIN S2a03a.borne b ON b.id_operateur = o.id_operateur
JOIN S2a01b.station s ON s.id_station = b.id_station
JOIN S2a03a.commune c ON c.code_insee_commune = s.code_insee_commune
GROUP BY o.nom_operateur, c.libelle_departement
ORDER BY departement, nb_prises DESC;

SELECT o.nom_operateur,
       b.puissance_nominale,
       COUNT(*) AS nb_prises
FROM operateur o
JOIN S2a03a.borne b ON b.id_operateur = o.id_operateur
GROUP BY o.nom_operateur, b.puissance_nominale
ORDER BY o.nom_operateur, nb_prises DESC;





--7 
--vue1
--vue locale departement 44
CREATE VIEW vue_prises_operateur_44 AS
SELECT o.nom_operateur,
       c.libelle_departement AS departement,
       COUNT(*) AS nb_prises
FROM operateur o
JOIN S2a03a.borne b ON b.id_operateur = o.id_operateur
JOIN S2a01b.station s ON s.id_station = b.id_station
JOIN S2a03a.commune c ON c.code_insee_commune = s.code_insee_commune
WHERE c.libelle_departement = 'Loire-Atlantique'
GROUP BY o.nom_operateur, c.libelle_departement;

-- vue locale départmeent 49 : 
CREATE VIEW vue_prises_operateur_49 AS
SELECT o.nom_operateur,
       c.libelle_departement AS departement,
       COUNT(*) AS nb_prises
FROM operateur o
JOIN S2a03a.borne b ON b.id_operateur = o.id_operateur
JOIN S2a01b.station s ON s.id_station = b.id_station
JOIN S2a03a.commune c ON c.code_insee_commune = s.code_insee_commune
WHERE c.libelle_departement = 'Maine-et-Loire'
GROUP BY o.nom_operateur, c.libelle_departement;
--vue globale sytnhèse
CREATE VIEW vue_prises_operateur_global AS
SELECT * FROM vue_prises_operateur_44
UNION ALL
SELECT * FROM vue_prises_operateur_49;


--vue2
--vue locale 44
CREATE VIEW vue_puissance_operateur_44 AS
SELECT o.nom_operateur,
       b.puissance_nominale,
       COUNT(*) AS nb_prises
FROM operateur o
JOIN S2a03a.borne b ON b.id_operateur = o.id_operateur
JOIN S2a01b.station s ON s.id_station = b.id_station
JOIN S2a03a.commune c ON c.code_insee_commune = s.code_insee_commune
WHERE c.libelle_departement = 'Loire-Atlantique'
GROUP BY o.nom_operateur, b.puissance_nominale;
--vue locale 49
CREATE VIEW vue_puissance_operateur_49 AS
SELECT o.nom_operateur,
       b.puissance_nominale,
       COUNT(*) AS nb_prises
FROM operateur o
JOIN S2a03a.borne b ON b.id_operateur = o.id_operateur
JOIN S2a01b.station s ON s.id_station = b.id_station
JOIN S2a03a.commune c ON c.code_insee_commune = s.code_insee_commune
WHERE c.libelle_departement = 'Maine-et-Loire'
GROUP BY o.nom_operateur, b.puissance_nominale;
--vue locale synthese

CREATE VIEW vue_puissance_operateur_global AS
SELECT * FROM vue_puissance_operateur_44
UNION ALL
SELECT * FROM vue_puissance_operateur_49;

--donner droit autres comptes : 
-- Donner accès aux vues à un autre compte
GRANT select ON vue_prises_operateur_global TO S2a03b;
GRANT select ON vue_puissance_operateur_global TO S2a03b;
GRANT select ON vue_prises_operateur_global TO s2a01b;
GRANT select ON vue_puissance_operateur_global TO s2a01b;
GRANT select ON vue_prises_operateur_global TO s2a02b;
GRANT select ON vue_puissance_operateur_global TO s2a02b;
