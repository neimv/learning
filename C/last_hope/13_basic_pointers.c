/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 08_arrays_2d
 * @created     : 
 */

#include <stdio.h>
#include <stdlib.h>


int main()
{
    int number = 20;
    int *pNumber;
    pNumber = &number;

    printf("Showing pointer things\n");
    printf("the value is: %d\n", number);
    printf("the value of pointer is: %d\n", *pNumber);
    printf("The address of pointer is: %p\n", &pNumber);
    printf("The address of variable is: %p\n", &number);

    return EXIT_SUCCESS;
}
