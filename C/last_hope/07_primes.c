/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 07_primes
 * @created     : viernes feb 03, 2023 21:55:14 CST
 */

#include <stdio.h>
#include <stdlib.h>

#define MAX_NUMBER 100


int main()
{
    int primes[MAX_NUMBER], i, counts = 0;

    for (i = 0; i<MAX_NUMBER; i++) 
    {
        if (i == 1 || i == 2 || i == 3)
        {
            primes[counts] = i;
            counts++;
        }

        //check another numbers
        for (int j=0; j<=counts; j++) 
        {

        }
    }

    for (i = 0; i<counts; i++)
    {
        printf("%d, ", primes[i]);
    }

    printf("\n");

    return EXIT_SUCCESS;
}
