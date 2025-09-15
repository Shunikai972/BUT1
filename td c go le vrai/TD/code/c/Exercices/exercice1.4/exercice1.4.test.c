//exercice1.4.test.c
#include <assert.h>
#include "exercice1.4.lib.h"
int main(int argc, char *argv[]){
assert(carreEntier(0)==0);
assert(carreEntier(1)==1);
assert(carreEntier(-1)==1);
assert(carreEntier(-2.5)==4);
assert(carreEntier(2.5)==4);
}