#include<stdio.h>
#include<stdlib.h>


int main() {
    int numb;
    printf("\nEnter your number: ");
    scanf("%d", &numb);

    if(numb >= 0) {
        printf("\nIt is a positive number");
    }

    return EXIT_SUCCESS;
}

