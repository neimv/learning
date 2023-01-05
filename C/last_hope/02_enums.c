/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 02_enums
 * @created     : lunes ene 02, 2023 17:22:12 CST
 */

#include <stdio.h>
#include <stdlib.h>

enum Company{GOOGLE, FACEBOOK, XEROX, YAHOO, EBAY, MICROSOFT};

int main() {
    enum Company xerox = XEROX, google = GOOGLE, ebay = EBAY;

    printf("xerox is: %d\n", xerox);
    printf("google is: %d\n", google);
    printf("ebay is: %d\n", ebay);
}
