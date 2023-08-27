/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : array
 * @created     : Thursday Aug 24, 2023 16:29:54 CST
 */
#include<stdio.h>


int main() 
{
    int A[100], i, size;

    printf("Size of integer datatype: %ld\n", sizeof(int));

    printf("Enter array size: \n");
    scanf("%d", &size);

    printf("Enter array elements: \n");
    for (i=0;i<size;i++)
    {
        scanf("%d", &A[i]);
    }

    printf("Index  -  Address   - Values\n");

    for (i=0;i<size;i++)
        printf("%d       - %u     - %d\n", i, &A[i], A[i]);
    
    return 0;
}
