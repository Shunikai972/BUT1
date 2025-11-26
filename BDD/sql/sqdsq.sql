DROP TABLE service CASCADE CONSTRAINTS;
CREATE TABLE service (
    nuserv NUMBER(2) CONSTRAINT cle_primaire_service PRIMARY KEY,
    nomserv VARCHAR2(20),
    chef NUMBER(2)
);
