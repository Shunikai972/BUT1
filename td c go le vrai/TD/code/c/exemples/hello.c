//hello.c

//Inclusion de 
//En-tête Standard d'Entrée/Sortie
#include <stdio.h>
//Définition d'une macro
#define USER "jub" 

//La fonction main 
int main(int argc, char *argv[]){
	//l'instruction d'appel à la fonction  printf de la bibliothéque stdio
	printf("Bonjour %s\n", USER);
	//fin du main et renvoie du code retour 0
	return 0;
}
