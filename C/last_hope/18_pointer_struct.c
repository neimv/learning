/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 18_pointer_struct
 * @created     : miércoles feb 22, 2023 18:31:11 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BUFFER_SIZE 1000

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

    readItem(ptrItem);
    printItem(ptrItem);

    free(ptrItem->itemName);

    return EXIT_SUCCESS;
}

void readItem(stcItem *item)
{
    char buffer[BUFFER_SIZE];
    size_t sizeStr;

    printf("Get values of item\n");
    printf("name of item\n");
    fgets(buffer, BUFFER_SIZE, stdin);
    printf("quantity\n");
    scanf("%d", &item->quantity);
    printf("price\n");
    scanf("%f", &item->price);

    sizeStr = strlen(buffer);
    item->itemName = (char *)malloc(sizeStr);

    if(item->itemName == NULL) exit(-1);

    strcpy(item->itemName, buffer);
    item->amount = item->quantity * item->price;
}

void printItem(stcItem *item)
{
    printf("The item values is:\n");
    printf("The name is: %s", item->itemName);
    printf("The price is: %f\n", item->price);
    printf("The quantity is: %d\n", item->quantity);
    printf("The amount is: %f\n", item->amount);
}

