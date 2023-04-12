def birthdayCakeCandles(ar):
    ar.sort(reverse = True)
    tallCount = 1
    if ar[0] == ar[-1]:
        tallCount = len(ar)
    else:
        for i, x in enumerate(ar, start=1):
            if x == ar[i]:
                tallCount += 1
            else:
                break
    return tallCount

arr = [5, 3, 4, 2, 5]
print (arr[0] == arr[-1])