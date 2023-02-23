/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 18_pointer_struct
 * @created     : miércoles feb 22, 2023 18:31:11 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFFER_SIZE 200

typedef struct Item
{
    char *itemName;
    int quantity;
    float price;
    float amount;
} stcItem;


void readItem(stcItem*);
void printItem(stcItem*);

int main()
{
    stcItem item;
    stcItem *ptrItem;

    ptrItem = &item;

    return EXIT_SUCCESS;
}

void readItem(stcItem *item)
{
    char buffer[BUFFER_SIZE];

    printf("Get values of item\n");
    printf("what is the name of item\n");
    scanf("%s", buffer);
    printf("")
}

void printItem(stcItem *item)
{

}

