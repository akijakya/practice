# def iq_test(numbers):
#     numArray = map(int, numbers.split(' '))
#     if (is_even(numArray)):
#         for i, num in enumerate(numArray, start=1):
#             if num % 2 == 1:
#                 print(i)
#     else:
#         for i, num in enumerate(numArray, start=1):
#             if num % 2 == 0:
#                 print(i)
            
def iq_test(numbers):
    numArray = list(map(int, numbers.split(' ')))
    for i, num in enumerate(numArray, start=1):
        if num % 2 == 1:
            return i
            
def is_even(numArray):
    even = 0
    uneven = 0
    for x in numArray:
        if x % 2 == 1:
            uneven += 1
        else:
            even += 1
    if even > uneven:
        return(True)
    else:
        return(False)

arr = "2 2 1 4 6"
arr2 = "1 3 7 6 5"
arr3 = arr.split(' ')
arr4 = list(map(int, arr3))
arr5 = list(map(int, arr.split(' ')))