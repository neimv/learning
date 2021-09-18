#include<stdio.h>

int main() {
    int counter, precount, postcount;
    counter = 10;
    precount = ++counter;
    postcount = counter++;

    printf("%d %d\n", precount, postcount);

    counter = 50;
    postcount = counter--;
    precount = --counter;

    printf("%d %d\n", postcount, precount);

    return 0;
}

