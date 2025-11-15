//exercice2.2.test.c
#include <assert.h>
#include <stdio.h>

#include "exercice2.2.lib.h"
 
int main(int argc, char *argv[]){
	//estBissextille
	assert(estBissextile(2000));
	assert(estBissextile(2020));
	assert(estBissextile(2024));
	assert(estBissextile(2028));
	assert(estBissextile(2400));

	assert(!estBissextile(1900));
	assert(!estBissextile(2100));
	assert(!estBissextile(2200));
	assert(!estBissextile(2300));

	assert(!estBissextile(0));
	printf("estBissextile est testée \n");

	assert(estJourValide(1,1,1582));
	assert(estJourValide(31,1,1582));
	assert(estJourValide(31,3,1582));
	assert(estJourValide(31,5,1582));
	assert(estJourValide(31,7,1582));
	assert(estJourValide(31,8,1582));
	assert(estJourValide(31,10,1582));
	assert(estJourValide(31,12,1582));
	assert(!estJourValide(31,4,1582));
	assert(!estJourValide(31,6,1582));
	assert(!estJourValide(31,9,1582));
	assert(!estJourValide(31,11,1582));
	
	assert(estJourValide(29,2,2004));
	assert(!estJourValide(29,2,1900));
	assert(!estJourValide(1,1,1581));
	assert(!estJourValide(32,1,1582));
	assert(!estJourValide(0,1,1582));
	assert(!estJourValide(1,13,1582));

	printf("estJourValide est testée \n");
	
	assert(31==joursDansMois(2000,1));
	assert(31==joursDansMois(2001,1));
	assert(28==joursDansMois(2001,2));
	assert(29==joursDansMois(2000,2));
	assert(31==joursDansMois(2000,3));
	assert(30==joursDansMois(2000,4));
	assert(31==joursDansMois(2000,5));
	assert(30==joursDansMois(2000,6));
	assert(31==joursDansMois(2000,7));
	assert(31==joursDansMois(2000,8));
	assert(30==joursDansMois(2000,9));
	assert(31==joursDansMois(2000,10));
	assert(30==joursDansMois(2000,11));
	assert(31==joursDansMois(2000,12));

	assert(-1 ==joursDansMois(0,12));
	assert(-1 ==joursDansMois(1,0));
	assert(-1 ==joursDansMois(1,13));
	printf("joursDansMois est testée\n");

	assert(0==joursDepuisOrigine(1,1,1582));
	assert(30==joursDepuisOrigine(31,1,1582));
	assert(243==joursDepuisOrigine(1,9,1582));
	assert(608==joursDepuisOrigine(1,9,1583));
	assert(974==joursDepuisOrigine(1,9,1584));
	assert(364==joursDepuisOrigine(31,12,1582));
	assert(162046==joursDepuisOrigine(1,9, 2025));

	assert(-1 == joursDepuisOrigine(1,1,1581));
	assert(-1 == joursDepuisOrigine(0,2,1582));
	assert(-1 == joursDepuisOrigine(1,0,1582));
	assert(-1 == joursDepuisOrigine(31,2,1582));
	printf("joursDepuisOrigine est testée \n");

	assert(0==jourDepuisNaissanceTuring(23,6,1912));
	assert(1==jourDepuisNaissanceTuring(24,6,1912));
	assert(1==jourDepuisNaissanceTuring(22,6,1912));
	assert(41343==jourDepuisNaissanceTuring(1,9,2025));
	assert(-1 == jourDepuisNaissanceTuring(1,1,1581));
	assert(-1 == jourDepuisNaissanceTuring(0,2,1582));
	assert(-1 == jourDepuisNaissanceTuring(1,0,1582));
	assert(-1 == jourDepuisNaissanceTuring(31,2,1582));
	printf("jourDepuisNaissanceTuring est testée \n");
	return 0;
}
