/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : 17_simple_struct
 * @created     : miércoles feb 22, 2023 18:06:41 CST
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Employee
{
    char name[200];
    int hireDate;
    float salary;
} Employee;


int main()
{
    struct Employee emps[2];

    strcpy(emps[0].name, "neimv");
    emps[0].hireDate = 2008;
    emps[0].salary = 1.3;

    printf("Get data:\n");
    printf("what is the name: \n");
    scanf("%s", emps[1].name);
    printf("what is the hire date: \n");
    scanf("%d", &emps[1].hireDate);
    printf("what is the salary: \n");
    scanf("%f", &emps[1].salary);

    printf("The employee 1 are:\n");
    printf("name: %s\n", emps[0].name);
    printf("hire date: %d\n", emps[0].hireDate);
    printf("salary: %f\n", emps[0].salary);
    printf("The employee 2 are:\n");
    printf("name: %s\n", emps[1].name);
    printf("hire date: %d\n", emps[1].hireDate);
    printf("salary: %f\n", emps[1].salary);

    return EXIT_SUCCESS;
}

