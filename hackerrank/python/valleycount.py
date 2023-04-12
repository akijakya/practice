def countingValleys(n, s):
    arr = list(s)
    level = 0
    valleyCount = 0
    for i, x in enumerate(arr):
        if x == "U":
            level += 1
        else:
            level -= 1
        if i > 0 and level == 0 and arr[i] == "U":
            valleyCount += 1
    return valleyCount

# print(countingValleys(8, "UDDDUDUU"))
print(countingValleys(12, "DDUUDDUDUUUD"))