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
    int primes[MAX_NUMBER], i, counts = 0, isPrime = 1;

    for (i = 1; i<MAX_NUMBER; i++) 
    {
        if (i == 1 || i == 2 || i == 3)
        {
            primes[counts] = i;
            counts++;
            continue;
        }

        //check another numbers
        for (int j=0; j<counts; j++) 
        {
            if (primes[j] == 1) continue;

            isPrime = i % primes[j];
            /* printf("number: %d Is prime: %d of: %d\n", i, isPrime, primes[j]); */

            if (!isPrime) break;
        }

        if (isPrime)
        {
            primes[counts] = i;
            counts++;
            isPrime = 1;
        }
    }

    /* for (i = 0; i<counts; i++) */
    /* { */
    /*     printf("%d, ", primes[i]); */
    /* } */

    printf("{");
    for (i = 0; i<counts; i++)
    {
        if (i != 0) printf(", ");
        printf("%d", primes[i]);
    }
    printf("}\n");

    return EXIT_SUCCESS;
}
