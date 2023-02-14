
/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 08_arrays_2d
 * @created     : 
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFFER_SIZE 255

void swap(char*, char*);
void bubbleSort(char[][BUFFER_SIZE], int);
void printList(char[][BUFFER_SIZE], int);

int main()
{
    char stringList[][BUFFER_SIZE] = {
        "Hello",
        "darkness",
        "stephi",
        "testing",
        "I"
    };

    bubbleSort(stringList, 6);
    printList(stringList, 6);

    return EXIT_SUCCESS;
}

void swap(char *str1, char *str2)
{
    printf("Swapping strings\n");
}

void bubbleSort(char listStr[][BUFFER_SIZE], int size)
{
    char buff[BUFFER_SIZE];
    printf("sorting string\n");

    for (int i = 0; i < size - 1; i++)
        // Last i elements are already in place
        for (int j = 0; j < size - i - 1; j++)
            if (strcmp(listStr[j], listStr[j + 1]) > 0)
            {
                strcpy(buff, listStr[j]);
                strcpy(listStr[j], listStr[j + 1]);
                strcpy(listStr[j + 1], buff);
            }
}

void printList(char listStr[][BUFFER_SIZE], int size)
{
    for (int i = 0; i < size - 1; i++)
        printf("%s\n", listStr[i]);
}
