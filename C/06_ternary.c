#include<stdio.h>
#include<stdlib.h>


int main() {
    int u, v, w;
    u = 20;
    v = 30;
    w = (u > v ? u : v);

    printf("\nGreatest out of the two is %d", w);

    return EXIT_SUCCESS;
}

