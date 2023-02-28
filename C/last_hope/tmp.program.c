/**
 * @AUTHOR      : NEIMV (NEIMV@DARK-UNIVERSE)
 * @FILE        : 08_ARRAYS_2D
 * @CREATED     : DOMINGO FEB 05, 2023 14:09:51 CST
 */

#INCLUDE <STDIO.H>
#INCLUDE <STDLIB.H>
#INCLUDE <STRING.H>

#DEFINE ELEMENTS 10
#DEFINE SIZE_BOARD 11
#DEFINE WIN_TIMES 5

VOID DRAWBOARD(CHAR*);
VOID DRAWSCORE(INT[2]);
INT CHECKWINNER(CHAR, CHAR*);

INT MAIN()
{
    CHAR SPACES[ELEMENTS] = "123456789", FICHA, QUIT;
    INT SCORE[2] = {0, 0}, TIRO, P1ORP2 = 0, TIROS = 0, WINSOME;

    WHILE (1)
    {
        DRAWSCORE(SCORE);
        DRAWBOARD(SPACES);

        PRINTF("\NYOUR TURN (SELECT BETWEEN 1-9)>> ");
        SCANF("%D", &TIRO);

        // CHECK IF INDEX IS VALID
        IF (TIRO < 1 || TIRO > 9) 
        {
            PRINTF("THE VALUE OF 'TIRO' IS NOT VALID: %D\N", TIRO);
            CONTINUE;
        }
        TIRO--;

        // CHECK IF DOES NOT EXISTS
        IF (SPACES[TIRO] == 'X' || SPACES[TIRO] == 'O')
        {
            PRINTF("THE VALUE OF 'TIRO' IS NOT VALID: %D\N", TIRO);
            CONTINUE;
        }

        FICHA = P1ORP2 == 0 ? 'X' : 'O';
        SPACES[TIRO] = FICHA;
        P1ORP2 = !P1ORP2;
        TIROS++;

        IF (TIROS >= 5) {
            WINSOME = CHECKWINNER(FICHA, SPACES);
            PRINTF("RESULT: %D\N", WINSOME);

            IF (WINSOME && FICHA == 'X')
            {
                PRINTF("PLAYER 1 WINS\N");
                SCORE[0]++;
            }
            ELSE IF (WINSOME && FICHA == 'O')
            {
                PRINTF("PLAYER 2 WINS\N");
                SCORE[1]++;
            }
        }
        
        IF (TIROS == 9 && WINSOME == 0) PRINTF("TIE\N");

        IF (TIROS == 9 || WINSOME)
        {
            PRINTF("DO YOU WISH GAME AGAIN? [N/Y]\N");
            QUIT = GETC(STDIN);
            SCANF("%C", &QUIT);

            IF (QUIT == 'N' || QUIT == 'N') RETURN EXIT_FAILURE;
            ELSE {
                STPCPY(SPACES, "123456789");
                WINSOME = 0;
                TIROS = 0;
                P1ORP2 = 0;
            }
        }
    }

    RETURN EXIT_SUCCESS;
}

VOID DRAWBOARD(CHAR* ELEMENTS)
{
    CHAR CHARSBOARD[3] = {' ', '|', '-'}, TMPCHAR;
    INT ELEMENT = 0;

    FOR (INT I = 0; I < SIZE_BOARD - 2; I++)
    {
        FOR (INT J = 0; J < SIZE_BOARD; J++)
        {
            IF (J == 3 || J == 7) TMPCHAR = '|';
            ELSE IF (I == 2 || I == 5) TMPCHAR = '_';
            ELSE IF (I == 1 || I == 4 || I == 7)
            {
                IF (J == 1 || J == 5 || J == 9)
                {
                    TMPCHAR = ELEMENTS[ELEMENT];
                    ELEMENT++;
                }
                ELSE TMPCHAR = ' ';
            }
            ELSE TMPCHAR = ' ';
            PRINTF("%C", TMPCHAR);
        }
        PRINTF("\N");
    }
}

VOID DRAWSCORE(INT SCORE[2])
{
    PRINTF("PLAYER 1: X\T\TPLAYER 2: O\N");
    PRINTF("SCORE P1: %D\T\TP2: %D", SCORE[0], SCORE[1]);
    PRINTF("\N\N");
}

INT CHECKWINNER(CHAR TIRO, CHAR* ELEMENTS)
{
    INT COUNTER = 0;

    // CHECK COLUMNS
    FOR (INT I = 0; I < ELEMENTS ; I++)
    {
        IF (!(I % 3)) COUNTER = 0;

        IF (TIRO == ELEMENTS[I]) COUNTER++;

        IF (COUNTER == 3) RETURN 1;
    }

    // CHECK ROWS
    IF (ELEMENTS[0] == TIRO && ELEMENTS[3] == TIRO && ELEMENTS[6] == TIRO) RETURN 1;
    ELSE IF (ELEMENTS[1] == TIRO && ELEMENTS[4] == TIRO && ELEMENTS[7] == TIRO) RETURN 1;
    ELSE IF (ELEMENTS[2] == TIRO && ELEMENTS[5] == TIRO && ELEMENTS[8] == TIRO) RETURN 1;
    // CHECK X
    ELSE IF (ELEMENTS[0] == TIRO && ELEMENTS[4] == TIRO && ELEMENTS[8] == TIRO) RETURN 1;
    ELSE IF (ELEMENTS[2] == TIRO && ELEMENTS[4] == TIRO && ELEMENTS[6] == TIRO) RETURN 1;

    RETURN 0;
}
ÿ