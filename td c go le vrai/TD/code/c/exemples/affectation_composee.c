//affecation_composee.c

#include <assert.h>

//La fonction main 
int main(){
    int a =2;
    a +=10;
    //Si a ne vaut pas 12, le programme s'arréte 
    //un message d'erreur est affiché
    assert(a == 12);
    return 0;  
}