SELECT * FROM s2a03a.parcours_velo;

DESC s2a03a.parcours_velo;
-- supp ancienne tables:
DROP TABLE departement CASCADE CONSTRAINTS PURGE;
DROP TABLE region CASCADE CONSTRAINTS PURGE;


--création table région
CREATE TABLE region (
    code_region NUMBER(38),
    nom_region  VARCHAR2(128)
);
--alimentation table région:
INSERT INTO region (code_region, nom_region)
SELECT DISTINCT
       code_regio,
       region
FROM s2a03a.parcours_velo
WHERE code_regio IS NOT NULL
  AND region IS NOT NULL;
  
  -- primary key région
  ALTER TABLE region
ADD CONSTRAINT pk_region PRIMARY KEY (code_region);


--vérification : 
SELECT * FROM region ORDER BY code_region;


-- supp departement

DROP TABLE departement CASCADE CONSTRAINTS PURGE;
--création département 


CREATE TABLE departement (
    code_depar NUMBER(38),
    nom_departement VARCHAR2(128),
    dep_label VARCHAR2(128),
    code_region NUMBER(38)
);
--insertion

INSERT INTO departement (code_depar, nom_departement, dep_label, code_region)
SELECT
    code_depar,
    MAX(departement),
    MAX(dep_label),
    MAX(code_regio)
FROM s2a03a.parcours_velo
WHERE code_depar IS NOT NULL
GROUP BY code_depar;

--primary key
ALTER TABLE departement
ADD CONSTRAINT pk_departement PRIMARY KEY (code_depar);

--fk vers région
ALTER TABLE departement
ADD CONSTRAINT fk_departement_region
FOREIGN KEY (code_region)
REFERENCES region(code_region);
COMMIT;


--don du droit

GRANT SELECT, REFERENCES ON departement TO s2a01b;
GRANT SELECT, REFERENCES ON departement TO s2a02b;
GRANT SELECT, REFERENCES ON departement TO s2a03a;
