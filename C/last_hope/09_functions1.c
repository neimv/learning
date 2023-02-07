
/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 08_arrays_2d
 * @created     : domingo feb 05, 2023 14:09:51 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#define BUFFER_SIZE 40

int gcd(int, int);
float absolute(float);
float squareRoot(float);


int main()
{
    char exitOpt, buffer[BUFFER_SIZE], trash;
    int option, num1, num2, result, isFloat = 0, i = 0;
    float valueDirty;

    do
    {
        printf("What do you whish?\n");
        printf("1.- greatest common divisor\n");
        printf("2.- absolute value\n");
        printf("3.- square root\n");
        printf("4.- exit\n");
        scanf("%d", &option);

        switch (option)
        {
        case 1:
            printf("Calculing greatest common divisor\n");
            printf("Give me the first number: \n");
            scanf("%d", &num1);
            printf("Give me the second number: \n");
            scanf("%d", &num2);
            result = gcd(num1, num2);

            printf("The gcd is: %d\n", result);
            break;

        case 2:
            trash = getc(stdin);
            printf("Calculate absolute value\n");
            printf("Give me the number (float separated with .): \n");
            do
            {
                trash = getc(stdin);
                if (trash == '\n') break;
                buffer[i] = trash;

                if (trash == '.') isFloat=1;
                i++;
            }
            while(1);

            // printf("The value of buffer is: %s\n", buffer);
            // printf("The value is float: %d\n", isFloat);
            valueDirty = atof(buffer);
            valueDirty = absolute(valueDirty);

            if (isFloat) printf("The value is %f\n", valueDirty);
            else printf("The value is %d\n", (int)valueDirty);

            isFloat = 0;
            i = 0;

            break;

        case 3:
            printf("Calculing square root\n");
            printf("Give me the number: \n");
            scanf("%f", &valueDirty);

            float valueSqrt = sqrt(valueDirty);
            printf("The sqrt of %f is %f\n", valueDirty, valueSqrt);
            break;

        case 4:
            printf("Bye bye\n");
            break;
        
        default:
            printf("Option %d is not valid\n", option);
            break;
        }
    }
    while(option != 4);

    return EXIT_SUCCESS;
}

int gcd(int num1, int num2)
{
    if (num1 == 0 || num2 == 0) return 0;
    int min = num1 < num2 ? num1 : num2;
    int max = num1 > num2 ? num1 : num2;

    if (!(max % min)) return min;

    int middleNumber = min / 2;
    printf("Middle num: %d\n", middleNumber);

    for (int i = middleNumber; i >= 0; i--)
    {
        if (min % i == 0 && max % i == 0) return i;
    }

    return 0;
}

float absolute(float value)
{
    float returned = value < 0 ? value * -1 : value;

    return returned;
}
