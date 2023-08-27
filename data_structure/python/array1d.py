
import array as arr

a = arr.array('i', [1, 2, 3])

print("Before insertion - Given array: ", end=" ")

for i in range(0, 3):
    print(a[i], end=" ")

print()

a.insert(1, 4)
a.append(5)
a.insert(3, 2)

print("After insertion - Given array: ", end=" ")

for i in range(0, 6):
    print(a[i], end=" ")

print()

a.pop(1)
a.remove(2)

print("After pop - Given array: ", end=" ")

for i in range(0, 4):
    print(a[i], end=" ")

print()
