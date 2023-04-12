# 24. Pyramid sort
# Create a function that takes an array of numbers(inputArray) and sorts them in the following manner:
# Odd numbers should be ascending
# Even numbers should be descending
# Odd numbers come before even numbers
# Example case:
# pyramidSort([1,2,3,4,5,6,7]);
# Should return [1,3,5,7,6,4,2].

def pyramidSort(arr):
    evens = []
    odds = []

    for x in arr:
        if x%2 == 0:
            evens.append(x)
        else:
            odds.append(x)
    
    # sorted(odds)
    # sorted(evens)

    # for x in range(len(evens)):
    #     odds.append(evens[-x-1])

    odds.sort()
    evens.sort(reverse=True)

    odds.extend(evens)
    print(odds)
    

pyramidSort([1,3,4,5,6,2,7])