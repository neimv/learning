/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 08_arrays_2d
 * @created     : domingo feb 05, 2023 14:09:51 CST
 */

#include <stdio.h>
#include <stdlib.h>

#define YEARS 5
#define MONTHS 12


int main()
{
    float rainfallAmounts[YEARS][MONTHS] = {
        {54.33, 52.19, 43.55, 48.5, 36.84, 49.69, 49.04, 51.8, 53.39, 37.15, 38.0, 36.94},
        {45.71, 38.05, 52.63, 54.19, 43.58, 39.71, 36.97, 51.79, 35.99, 47.86, 49.02, 37.94},
        {54.1, 50.9, 35.34, 35.91, 53.72, 42.32, 36.94, 53.78, 44.04, 53.17, 41.73, 41.1},
        {52.92, 37.54, 39.63, 51.15, 49.34, 53.47, 46.4, 50.91, 38.53, 48.21, 44.05, 38.32},
        {35.76, 43.06, 50.24, 43.18, 44.49, 41.37, 43.11, 40.08, 41.55, 53.64, 53.35, 43.15}
    };
    float yearAvg[YEARS], monthAvg[MONTHS];

    // get the average by year
    for (int i = 0; i < YEARS; i++)
    {
        /* printf("Calculing year: %d\n", i); */
        float total = 0, avg = 0.0;
        for (int j = 0; j < MONTHS; j++)
        {
            total += rainfallAmounts[i][j];
        }
        avg = total / MONTHS;
        /* printf("The total is: %f and avg is: %f\n", total, avg); */
        yearAvg[i] = avg;
    }


    printf("The results by year is: \n");
    printf("{");
    for (int i = 0; i<YEARS; i++)
    {
        if (i != 0) printf(", ");
        printf("%f", yearAvg[i]);
    }
    printf("}\n");

    // get the average by month
    for (int i = 0; i < MONTHS; i++)
    {
        /* printf("Calculing year: %d\n", i); */
        float total = 0, avg = 0.0;
        for (int j = 0; j < YEARS; j++)
        {
            total += rainfallAmounts[j][i];
        }
        avg = total / YEARS;
        /* printf("The total is: %f and avg is: %f\n", total, avg); */
        monthAvg[i] = avg;
    }

    printf("The results by month is: \n");
    printf("{");
    for (int i = 0; i<MONTHS; i++)
    {
        if (i != 0) printf(", ");
        printf("%f", monthAvg[i]);
    }
    printf("}\n");

    return EXIT_SUCCESS;
}

