Drop table  aime ;
Drop table repas ;
Drop table nourritures ;
Drop table Dragons ;
Create table Dragons 
 ( 
 Dragon varchar(30) primary key,
 sexe char, 
 Longueur number(6) check (longueur >0), 
 NombreEcailles number(6) check (NombreEcailles >0), 
 CracheduFeu char(1) check (Crachedufeu in ('o','n')), 
 ComportementAmoureux varchar(30) check (ComportementAmoureux in ('macho','timide','sincere','volage')), 
 constraint monstyle check (longueur < NombreEcailles)

 ) ;

Create table nourritures 
 ( 
 produit varchar(30), 
 calories number(3), 
 primary key (produit) 
 );
 
 Create table repas 
 ( 
 dragon varchar(30) references Dragons(dragon) on delete cascade, 
 produit varchar(30), 
 quantite number(5), 
 primary key (dragon,produit) /*contrainte portant sur la table OBLIGATOIRE car concerne au moins deux 
attributs*/, 
 foreign key(produit) references nourritures (produit) on delete cascade 
 ) ;

Create table aime 
(dragonAimant varchar(30) references Dragons(dragon) on delete cascade,
dragonAime varchar(30) references Dragons(dragon) on delete cascade,
Force varchar(30),
 primary key (dragonAimant,dragonAime)
);

Insert into dragons values ('Smeagol' , 'M' , 132 , 1765 , 'o', 'macho');
Insert into dragons values ('Birduth' , 'M' , 260 , 4897 , 'n' , 'timide' );
Insert into dragons values ('Negueth' , 'F', 128 , 1561 , 'o' , 'sincere');
Insert into dragons values ('MissToc' , 'F' , 232 , 3500 , 'n' , 'sincere' );
Insert into dragons values ('Belong' , 'M', 234 , 3214 , 'o', 'macho');
Insert into dragons values ( 'Miloch' , 'M' , 65 , 675 , 'o' , 'timide' );
Insert into dragons values ('Nessie' , 'M' , 139 , 1365 , 'n', 'volage' );
Insert into dragons values ('Tarak', 'F' , 178 , 765 , 'o' , 'timide');
Insert into dragons values ('Solong' , 'M' , 125 , 1430 , 'o' , 'sincere' );

Insert into NOURRITUREs  values ('pomme' , 7);
Insert into NOURRITUREs  values ('cacahuete' , 10 );
Insert into NOURRITUREs  values ('orange' , 25 );
Insert into NOURRITUREs  values ('oeuf', 9 );
Insert into NOURRITUREs  values ('ver', 3 );
Insert into NOURRITUREs  values ('poisson' , 35);

Insert into REPAS values ('Smeagol', 'cacahuete' , 1000 );
Insert into REPAS values ('Smeagol' , 'pomme' , 16 );
Insert into REPAS values ('Birduth' , 'oeuf' , 6 );
Insert into REPAS values ('Negueth' , 'orange' , 3 );
Insert into REPAS values ('Negueth' , 'oeuf' , 2 );
Insert into REPAS values ('Miloch' , 'cacahuete', 100 );
Insert into REPAS values ('Miloch' , 'ver' , 50 );
Insert into REPAS values ('Nessie' , 'poisson' , 30 );
Insert into REPAS values ('Tarak' , 'pomme' , 60 );
Insert into REPAS values ('Tarak', 'orange' , 67 );
Insert into REPAS values ('Solong' , 'oeuf', 7 );
Insert into REPAS values ('Solong', 'poisson' , 7 );
Insert into REPAS values ('Solong' , 'orange' ,9); 

Insert into AIME values ('Smeagol' , 'Negueth' , 'passionnement');
Insert into AIME values ('Birduth' , 'Negueth' , 'beaucoup' );
Insert into AIME values ('Negueth' , 'Miloch' , 'a la folie' );
Insert into AIME values ('Belong' , 'Negueth' ,'a la folie');
Insert into AIME values ('Tarak', 'Belong', 'un peu');
Insert into AIME values ('Solong' , 'Tarak', 'beaucoup');

--1
SELECT dragon, sexe, crachedufeu
FROM Dragons;
--2
SELECT DISTINCT ComportementAmoureux
FROM Dragons;
--3
SELECT dragon
FROM Dragons
WHERE sexe = 'M';
--4
SELECT dragon
FROM Dragons
WHERE sexe = 'M'
AND CracheduFeu = 'o';
--5
SELECT dragon
FROM Dragons
WHERE ComportementAmoureux IN ('timide', 'sincere');
--6
SELECT dragon
FROM Dragons
WHERE longueur > 200;
--7
SELECT dragon
FROM Dragons
WHERE longueur BETWEEN 100 AND 200;
--8
SELECT dragon
FROM Dragons
WHERE ComportementAmoureux NOT IN ('macho', 'timide');
--9
SELECT dragon
FROM Dragons
WHERE dragon LIKE '%oc%';
--10
SELECT dragon
FROM Dragons
WHERE longueur IS NULL;
--11
SELECT dragon, (NombreEcailles - longueur) AS difference
FROM Dragons;
--12 
SELECT d.dragon, d.sexe, r.produit
FROM Dragons d
JOIN Repas r ON d.dragon = r.dragon;
--13
SELECT DISTINCT d.dragon, d.sexe
FROM Dragons d
JOIN Repas r ON d.dragon = r.dragon;
--14
SELECT dragon
FROM Dragons
WHERE longueur > (
    SELECT longueur
    FROM Dragons
    WHERE dragon = 'Birduth'
);
--15
SELECT dragon
FROM Dragons
WHERE longueur > (
    SELECT longueur FROM Dragons WHERE dragon = 'Birduth'
)
AND NombreEcailles = (
    SELECT NombreEcailles FROM Dragons WHERE dragon = 'Birduth'
);
--16
SELECT d.dragon, d.sexe, r.produit
FROM Dragons d
LEFT JOIN Repas r ON d.dragon = r.dragon;
--17*
SELECT AVG(longueur) AS moyenne_longueur
FROM Dragons;
--18
SELECT ComportementAmoureux, AVG(longueur) AS moyenne
FROM Dragons
GROUP BY ComportementAmoureux;
--19
SELECT ComportementAmoureux, AVG(longueur) AS moyenne
FROM Dragons
WHERE sexe = 'M'
GROUP BY ComportementAmoureux;
--20
SELECT ComportementAmoureux, AVG(longueur) AS moyenne
FROM Dragons
WHERE sexe = 'M'
GROUP BY ComportementAmoureux
HAVING COUNT(*) > 2;
--21
SELECT d.dragon
FROM Dragons d
WHERE NOT EXISTS (
    SELECT *
    FROM Nourritures n
    WHERE NOT EXISTS (
        SELECT *
        FROM Repas r
        WHERE r.dragon = d.dragon
        AND r.produit = n.produit
    )
);
--22
SELECT dragon, longueur
FROM Dragons
ORDER BY longueur DESC
FETCH FIRST 2 ROWS ONLY;
--23
SELECT dragon
FROM Repas
GROUP BY dragon
HAVING COUNT(DISTINCT produit) >= 2;

