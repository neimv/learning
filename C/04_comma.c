#include<stdio.h>

int main() {
    int x, y, z;
    x = y = 20;

    z = (x++, y--, x+y);

    printf("%d", z);

    return 0;
}

