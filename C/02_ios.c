#include<stdio.h>
#include<string.h>


int main() {
    char c;
    printf("\nEnter the text: ");
    c = getchar();
    printf("The output text is: ");

    while(c != '\n') {
        if(c == ' ' || c == '\t') {
            c = ' ';
            putchar(c);
        }

        while(c == ' ' || c == '\t') {
            c = getchar();
        }

        putchar(c);
        c = getchar();
    }

    printf("\n");

    return 0;
}

