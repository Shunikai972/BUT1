//incrementation.c
#include <assert.h>

int main(){
    int a =2;
    assert(++a == 3);
    assert(a == 3);
    assert(a++ == 3);
    assert(a == 4);
    return 0;  
}