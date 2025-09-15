//enum.c

//Inclusion de 
//En-tête Standard d'Entrée/Sortie
#include <stdio.h>

//La fonction main 
int main(int argc, char *argv[]){
    enum naturelPair { ZERO, DEUX, QUATRE, SIX };
    
    enum naturelPair z = ZERO;
    enum naturelPair d = DEUX;
    
    printf("%d \n", (int) z);
    printf("%d \n", ZERO);
    printf("%d \n", (int) d);
    printf("%d \n", DEUX);
    
	return 0;
}
