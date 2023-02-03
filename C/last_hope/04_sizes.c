/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 04_sizes
 * @created     : lunes ene 30, 2023 18:08:06 CST
 * Get sizes of int, char, long, long long, double, long double
 */

#include <stdio.h>
#include <stdlib.h>

int main()
{
    printf("int: %zd\n", sizeof(int));
    printf("char: %zd\n", sizeof(char));
    printf("long: %zd\n", sizeof(long));
    printf("long long: %zd\n", sizeof(long long));
    printf("double: %zd\n", sizeof(double));
    printf("long double: %zd\n", sizeof(long double));

    return EXIT_SUCCESS;
}

