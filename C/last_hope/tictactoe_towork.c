/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 08_arrays_2d
 * @created     : domingo feb 05, 2023 14:09:51 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ELEMENTS 10
#define SIZE_BOARD 11
#define WIN_TIMES 5

void drawboard(char*);
void drawscore(int[2]);
int checkWinner(char, char*);

int main()
{
    char spaces[ELEMENTS] = "123456789", ficha, quit;
    int score[2] = {0, 0}, tiro, p1orp2 = 0, tiros = 0, winSome;

    while (1)
    {
        drawscore(score);
        drawboard(spaces);

        printf("\nYour turn (select between 1-9)>> ");
        scanf("%d", &tiro);

        // check if index is valid
        if (tiro < 1 || tiro > 9) 
        {
            printf("The value of 'tiro' is not valid: %d\n", tiro);
            continue;
        }
        tiro--;

        // check if does not exists
        if (spaces[tiro] == 'X' || spaces[tiro] == 'O')
        {
            printf("The value of 'tiro' is not valid: %d\n", tiro);
            continue;
        }

        ficha = p1orp2 == 0 ? 'X' : 'O';
        spaces[tiro] = ficha;
        p1orp2 = !p1orp2;
        tiros++;

        if (tiros >= 5) {
            winSome = checkWinner(ficha, spaces);
            printf("result: %d\n", winSome);

            if (winSome && ficha == 'X')
            {
                printf("Player 1 wins\n");
                score[0]++;
            }
            else if (winSome && ficha == 'O')
            {
                printf("Player 2 wins\n");
                score[1]++;
            }
        }
        
        if (tiros == 9 && winSome == 0) printf("TIE\n");

        if (tiros == 9 || winSome)
        {
            printf("Do you wish game again? [n/y]\n");
            quit = getc(stdin);
            scanf("%c", &quit);

            if (quit == 'n' || quit == 'N') return EXIT_FAILURE;
            else {
                stpcpy(spaces, "123456789");
                winSome = 0;
                tiros = 0;
                p1orp2 = 0;
            }
        }
    }

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
}

void drawscore(int score[2])
{
    printf("Player 1: X\t\tPlayer 2: O\n");
    printf("Score P1: %d\t\tP2: %d", score[0], score[1]);
    printf("\n\n");
}

int checkWinner(char tiro, char* elements)
{
    int counter = 0;

    // check columns
    for (int i = 0; i < ELEMENTS ; i++)
    {
        if (!(i % 3)) counter = 0;

        if (tiro == elements[i]) counter++;

        if (counter == 3) return 1;
    }

    // check rows
    if (elements[0] == tiro && elements[3] == tiro && elements[6] == tiro) return 1;
    else if (elements[1] == tiro && elements[4] == tiro && elements[7] == tiro) return 1;
    else if (elements[2] == tiro && elements[5] == tiro && elements[8] == tiro) return 1;
    // check X
    else if (elements[0] == tiro && elements[4] == tiro && elements[8] == tiro) return 1;
    else if (elements[2] == tiro && elements[4] == tiro && elements[6] == tiro) return 1;

    return 0;
}
