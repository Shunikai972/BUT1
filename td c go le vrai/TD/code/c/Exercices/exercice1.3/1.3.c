#include <assert.h>
#include <stdio.h>
int carreEntier(float n) {
int unEntier;
unEntier = n;
unEntier = unEntier * unEntier;
return(unEntier);
}
int main(int argc, char *argv[]){
assert(carreEntier(0)==0);
assert(carreEntier(1)==1);
assert(carreEntier(-1)==1);
assert(carreEntier(-2.5)==4);
assert(carreEntier(2.5)==4);
}