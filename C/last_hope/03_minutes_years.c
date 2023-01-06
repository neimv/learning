/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 03_minutes_years
 * @created     : jueves ene 05, 2023 13:07:19 CST
 */
#include <stdio.h>
#include <stdlib.h>


int main() {
    int minutes;
    double minutes_in_year = (60 * 24) * 365;
    double minutes_in_days = 60 * 24;
    double years, days;

    printf("ingrese total de minutos: \n");
    scanf("%d", &minutes);

    printf("total of minutes: %d\n", minutes);
    years = minutes / minutes_in_year;
    days = minutes / minutes_in_days;

    printf("total of year is: %lf\n", years);
    printf("total of days is: %lf\n", days);

    return EXIT_SUCCESS;
}

