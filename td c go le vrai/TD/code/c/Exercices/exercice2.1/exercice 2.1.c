#include <stdio.h>
int main(void)
{
    
    int age;
    float prix;
    printf("Quel est le prix de vos achats ? ");
    scanf("%f",&prix);
    printf("Quel est votre age ?");
    scanf("%d",&age);
    if(age < 18)
    {
        prix = prix * 0.95;
    }
    if(prix > 300)
    {
        prix = prix -5;
    }
    printf("Le montant a payer et de : %f \n",prix);
}
