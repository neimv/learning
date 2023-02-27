/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 18_pointer_struct
 * @created     : miércoles feb 22, 2023 18:31:11 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFFER_SIZE 1000

size_t totalNums(FILE*);

int main()
{
    FILE *pFile;
    size_t lenNum;
    char filename[] = "tictactoe_towork.c";
    pFile = fopen(filename, "r");

    if (pFile == NULL)
    {
        printf("Error opening file %s", filename);
        return EXIT_FAILURE;
    }

    lenNum = totalNums(pFile);

    printf("the file %s, contains %d lines", filename, lenNum);

    return EXIT_SUCCESS;
}

size_t totalNums(FILE *file)
{
    size_t len = 0;
    char c;

    while()

    return len;
}