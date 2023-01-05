/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : triangle
 * @created     : lunes ene 02, 2023 16:20:21 CST
 */

#include <stdio.h>
#include <stdlib.h>


int main(int argc, char *argv[])
{
    if (argc != 3) {
        printf("Is necessary 2 args to calculate area and perimeter\n");
        printf("I have %d arguments\n", (argc - 1));

        return EXIT_FAILURE;
    }

    printf("number of argurments: %d\n", argc);
    printf("Hello I calculated area and perimeter with 2 args\n");
    double height = 0.0, width = 0.0, perimeter = 0.0, area = 0.0;

    height = atof(argv[1]);
    width = atof(argv[2]);

    perimeter = 2.0 * (height + width);
    area = width * height;

    printf("Width is: %f, height is: %f\n", height, width);
    printf("The area is: %f and perimeter is: %f\n", area, perimeter);

    return EXIT_SUCCESS;
}

