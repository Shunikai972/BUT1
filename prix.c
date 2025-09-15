#include <stdio.h>

int main(void)
{
    int age;
    int prix;

    prix = 180;
    printf('C quoi ton age frr?');
    age = scanf('b');
    if(age < 0)
        return(0);
    if (age < 18 && (prix > 0 && prix < 300))
        printf("%d", prix*0,95);
    else if (age < 18 && (prix > 0 && prix > 300))
        printf("%d", prix -5);
}