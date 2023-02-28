/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 18_pointer_struct
 * @created     : miércoles feb 22, 2023 18:31:11 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

#define BUFFER_SIZE 100

size_t totalNums(FILE*);
void convertToUpper(char*);

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

    printf("\nthe file %s, contains %lu lines\n", filename, lenNum);
    fclose(pFile);
    pFile = NULL;

    convertToUpper(filename);

    return EXIT_SUCCESS;
}

size_t totalNums(FILE *file)
{
    size_t len = 0;
    char ch;

    while (ch != EOF) 
    {
        ch = fgetc(file);
        if (ch == '\n') len++;
    };

    return len;
}

void convertToUpper(char *filename)
{
    FILE *pFile, *pCopy;
    int stop = 0;
    char ch, *nameBuffer = "tmp.program.c";
    pFile = fopen(filename, "r");
    pCopy = fopen(nameBuffer, "w");

    if (pFile == NULL)
    {
        printf("Error opening file %s", filename);
        exit(EXIT_FAILURE);
    }

    if (pCopy == NULL)
    {
        printf("Error opening file %s", nameBuffer);
        exit(EXIT_FAILURE);
    }

    while (ch != EOF) 
    {
        ch = fgetc(pFile);
        printf("%c", ch);

        if (isalpha(ch) != 0 && islower(ch) != 0)
        {
            ch -= 32;
        }

        putc(ch, pCopy);
    };

    fclose(pFile);
    fclose(pCopy);
    pFile = NULL;
    pCopy = NULL;
}
