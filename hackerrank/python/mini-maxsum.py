def miniMaxSum(arr):
    arr.sort()
    print(f"{sum(arr[:4])} {sum(arr[1:])}")

miniMaxSum([1, 3, 4, 2, 5])