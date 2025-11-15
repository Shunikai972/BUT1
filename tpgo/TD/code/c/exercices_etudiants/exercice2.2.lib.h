//exercice2.2.lib.h
/**
 * Prend en paramètre l'année
 * Renvoie true si l'année est bissextile et false sinon
 */
bool estBissextile(unsigned int); 

/** 
 * Prend en paramètres l'année et le mois
 * Renvoie le npmbre de jours du mois ou
 * -1 si l'année est antérieur à 1582 ou
 * si le mois n'est pas valide
 * 1 janvier, ..., 12 décembre 
 */ 
short joursDansMois(unsigned int, unsigned short);

/**
 * Prend en paramètre un jour, un mois, une année
 * Renvoie true si la date est valide false sinon
 * les année doivent être supérieure ou égale à 1582
 */ 
bool estJourValide(unsigned short, unsigned short,unsigned int);

/** 
 * Prend en paramètre un jour, un mois, une année
 * Renvoie le nombre de jours depuis le 1 janvier 1582 ou
 * -1 si la date n'est pas valide
 */  
int joursDepuisOrigine(unsigned short, unsigned short,unsigned int);
/** 
 * Prend en paramètre un jour, un mois, une année
 * Renvoie le nombre de jours depuis le 23 juin 1912
 * -1 si la date n'est pas valide
 */
int jourDepuisNaissanceTuring(unsigned short, unsigned short,unsigned int);