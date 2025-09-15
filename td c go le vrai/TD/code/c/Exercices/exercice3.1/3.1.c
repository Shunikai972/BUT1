#include <stdio.h>

int main(void)
{
    int a = 0;

    // For loop
    for (int i = 0; i <= 10; i++)
    {
        printf("%d", i);
    }
    printf("\n");

    // While loop
    a = 0;
    while (a <= 10)
    {
        printf("%d", a);
        a++;
    }
    printf("\n");

    // Do-while loop
    a = 0;
    do
    {
        printf("%d", a);
        a++;
    } while (a <= 10);
    printf("\n");

    return 0;
}
