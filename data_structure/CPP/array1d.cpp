/**
 * @author      : neimv (neimv@dark-universe)
 * @file        : array1d
 * @created     : Thursday Aug 24, 2023 16:51:48 CST
 */
#include<iostream>

using namespace std;

int main()
{
    int a[100], i, size;

    cout << "Enter array size: " << endl;
    cin >> size;

    cout << "Enter array elements: " << endl;

    for (i=0;i<size;i++)
        cin >> a[i];

    cout << "Index     - Address    - Value" << endl;

    for (i=0;i<size;i++)
        cout << i << "    " << &a[i] << "     " << a[i] << endl;

    return 0;
}
