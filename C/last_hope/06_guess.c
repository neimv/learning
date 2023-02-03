/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 06_guess
 * @created     : jueves feb 02, 2023 19:28:25 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define NUM_ATTEMPS 5


int main()
{
    time_t t;
    srand((unsigned) time(&t));

    int randomNumber = rand() % 21, attemps = 0, option;

    do {
        attemps++;

        printf("Guess the number (0-20)\n you have %d, this is your attemp: %d>> \n", NUM_ATTEMPS, attemps);
        scanf("%d", &option);

        if (option == randomNumber)
        {
            printf("You win\n");
            return EXIT_SUCCESS;
        }

        /* printf("%d", randomNumber); */
    }
    while(attemps < NUM_ATTEMPS);

    printf("you failed the number is: %d\n", randomNumber);

    return EXIT_SUCCESS;
}

