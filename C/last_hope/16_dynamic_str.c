/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 16_dynamic_str
 * @created     : lunes feb 20, 2023 18:01:44 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define TRUE 1
#define FALSE 0


int main()
{
    int size;
    char *text = NULL;

    printf("Enter limit of text ");
    scanf("%d", &size);

    text = (char *) malloc(size * sizeof(char));

    if (text != NULL)
    {
        printf("Enter som text:\n ");
        scanf(" ");
        gets(text);

        printf("inputted text is: %s\n", text);
    }

    free(text);
    text = NULL;

    return EXIT_SUCCESS;
}

