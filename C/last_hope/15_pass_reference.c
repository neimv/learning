/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 15_pass_reference
 * @created     : domingo feb 19, 2023 14:44:13 CST
 */

#include <stdio.h>
#include <stdlib.h>

void squareNumber(int * const);

int main()
{
    int number = 5;

    printf("The value of number is: %d\n", number);

    squareNumber(&number);

    printf("The value before square is: %d\n", number);

    return EXIT_SUCCESS;
}

void squareNumber(int * const number)
{
    *number = (*number) * (*number);
}
