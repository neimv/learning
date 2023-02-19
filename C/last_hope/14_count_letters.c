/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 08_arrays_2d
 * @created     : domingo feb 05, 2023 14:09:51 CST
 */

#include <stdio.h>
#include <stdlib.h>

int strLenP(const char*);

int main()
{
    const char testingWord[] = "Hello World Viri";
    int lenStrInt;
    // char *pTesting = testingWord;

    lenStrInt = strLenP(testingWord);
    printf("len of %s is %d\n", testingWord, lenStrInt);

    return EXIT_SUCCESS;
}

int strLenP(const char *pStr)
{
    const char *lastAddress = pStr;

    while (*lastAddress)
        ++lastAddress;

    return lastAddress - pStr;
}