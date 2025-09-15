#include <stdio.h>

int main(void)
{
    int a = 1;
    int nb;
    int result;
    result = 0;
    nb = 1;
    printf("table de quoi");
    scanf("%d", &a);
    while(nb < 10)
    {
        result = a  * nb,
        printf("%d", result);
        printf("\n");
        nb++;
    }
    
}
