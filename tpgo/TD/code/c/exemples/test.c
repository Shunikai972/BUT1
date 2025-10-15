#include <stdio.h>
#include <string.h>
 
int main(int argc, char *argv[]){
	printf("Nombre de paramètre : %d\n",argc-1);
	for (int i=1; i<argc; i++)
		printf("Paramètre %d : %s \n",i,argv[i]);
}