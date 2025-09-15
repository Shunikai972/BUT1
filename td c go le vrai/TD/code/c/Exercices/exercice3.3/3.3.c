#include <stdio.h>

int main(void)
{
    int max_table;
    printf("Table de quoi ? ");
    scanf("%d", &max_table);

    for (int table = 1; table <= max_table; table++)
    {
        printf("Table de %d :\n", table);
        for (int nb = 1; nb <= 10; nb++)
        {
            int result = table * nb;
            printf("%d x %d = %d\n", table, nb, result);
        }
        printf("\n");
    }

    return 0;
}
