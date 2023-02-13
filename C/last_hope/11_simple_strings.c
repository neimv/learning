/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 08_arrays_2d
 * @created     : domingo feb 05, 2023 14:09:51 CST
 */

#include <stdio.h>
#include <stdlib.h>

#define BUFFER_SIZE 250


int strownlen(char*);
void strowncat(char*, char*, char*);
int strowncmp(char*, char*);
void showReverse(char*);

int main() {
    char *arr[] = { "Geek", "Geeks", "Geekfor", "GeekForDevops" };
    char *arrCmp[] = { "Geek", "GeeK" };
    char result[BUFFER_SIZE];
    int sizeStr;

    // for (int i = 0; i < 4; i++)
    // {
    //     sizeStr = strownlen(arr[i]);
    //     printf("The size of string is: %d\n", sizeStr);
    // }

    // strowncat(result, arr[0], arr[1]);
    // printf("The string result is: %s\n", result);

    // printf("Check the string: %s vs %s: %d\n", arrCmp[0], arrCmp[0], strowncmp(arrCmp[0], arrCmp[0]));
    // printf("Check the string: %s vs %s: %d\n", arrCmp[0], arrCmp[1], strowncmp(arrCmp[0], arrCmp[1]));
    // printf("Check the string: %s vs %s: %d\n", arrCmp[0], arr[1], strowncmp(arrCmp[0], arr[1]));

    for (int i = 0; i < 4; i++) 
    {
        printf("The reverse of string is:\n");
        showReverse(arr[i]);
    }

    return EXIT_SUCCESS;
}

int strownlen(char *string)
{
    int size = 0;
    // printf("the string is: %s\n", string);
    while(string[size] != '\0') size++;

    return size;
}

void strowncat(char *destination, char *str1, char *str2)
{
    int indexStr = 0, indexDst = 0;

    while (str1[indexStr] != '\0')
    {
        destination[indexDst] = str1[indexStr];
        indexDst++;
        indexStr++;
    }

    indexStr = 0;
    while (str2[indexStr] != '\0')
    {
        destination[indexDst] = str2[indexStr];
        indexStr++;
        indexDst++;
    }
}

int strowncmp(char *str1, char *str2)
{
    int sizeStr1, sizeStr2, i = 0;

    sizeStr1 = strownlen(str1);
    sizeStr2 = strownlen(str2);

    if (sizeStr1 < sizeStr2) return -1;
    else if (sizeStr1 > sizeStr2) return 1;

    while (str1[i] != '\0')
    {
        if (str1[i] < str2[i]) return -1;
        if (str1[i] > str2[i]) return 1;

        i++;
    }

    return 0;
}

void showReverse(char *str)
{
    for (int i = strownlen(str); i >= 0; i--) putchar(str[i]);

    printf("\n");
}