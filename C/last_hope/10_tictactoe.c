/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 08_arrays_2d
 * @created     : domingo feb 05, 2023 14:09:51 CST
 */

#include <stdio.h>
#include <stdlib.h>

#define ELEMENTS 10
#define SIZE_BOARD 11

void drawboard(char*);

int main()
{
    char spaces[ELEMENTS] = "123456789";
    int score[2] = {0, 0};

    drawboard(spaces);

    return EXIT_SUCCESS;
}

void drawboard(char* elements)
{
    char charsBoard[3] = {' ', '|', '-'}, tmpChar;
    int element = 0;

    for (int i = 0; i < SIZE_BOARD - 2; i++)
    {
        for (int j = 0; j < SIZE_BOARD; j++)
        {
            if (j == 3 || j == 7) tmpChar = '|';
            else if (i == 2 || i == 5) tmpChar = '_';
            else if (i == 1 || i == 4 || i == 7)
            {
                if (j == 1 || j == 5 || j == 9)
                {
                    tmpChar = elements[element];
                    element++;
                }
                else tmpChar = ' ';
            }
            else tmpChar = ' ';
            printf("%c", tmpChar);
        }
        printf("\n");
    }

    // for (int i = 0; i < ELEMENTS - 1; i++)
    //     printf("value: %c\n", elements[i]);
}

