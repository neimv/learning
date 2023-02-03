/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 05_pays
 * @created     : miércoles feb 01, 2023 18:39:19 CST
 */

#include <stdio.h>
#include <stdlib.h>

#define TAXRATE_300 .15
#define TAXRATE_150 .20
#define TAXRATE_REST .25

int main()
{
    float basicPay = 12.0, total, taxes = 0.0, netPay = 0.0;
    int timeLimit = 40, numberHours, extraHours = 0;

    printf("cuantas horas trabajo? >> \n");
    scanf("%d", &numberHours);

    if (numberHours < 0)
    {
        printf("El valor %d, no es valido", numberHours);
        return EXIT_FAILURE;
    }
    else if (numberHours > timeLimit)
    {
        extraHours = numberHours - timeLimit;
        numberHours = 40;
    }

    total = (numberHours * basicPay) + (extraHours * (basicPay * 1.5));

    if (total <= 300)
    {
        taxes = total * TAXRATE_300;
    }
    else if (total > 300 && total <= 450)
    {
        taxes = 300 * TAXRATE_300;
        taxes += (total - 300) * TAXRATE_150;
    }
    else if (total > 450) 
    {
        taxes = 300 * TAXRATE_300;
        taxes += 150 * TAXRATE_150;
        taxes += (total - 450) * TAXRATE_REST;
    }

    // netpay
    netPay = total - taxes;

    printf("el total es: %f\n", total);
    printf("el total de impuestos es: %f\n", taxes);
    printf("el pago total despues de impuestos es: %f\n", netPay);

    return EXIT_SUCCESS;
}
