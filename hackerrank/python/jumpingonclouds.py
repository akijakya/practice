def jumpingOnClouds(c):
    # arr = list(map(int, c.split(' ')))
    result = 0
    doubleJump = False
    for i, e in enumerate (c):
        if doubleJump == True:
            doubleJump = False
        elif i < (len(c) - 2):
            if c[i + 2] == 0:
                result += 1
                doubleJump = True
            else:
                result += 1
        elif i == (len(c) - 2):
            result += 1
    return result

print(jumpingOnClouds([0, 0, 0, 0, 1, 0]))
#print(jumpingOnClouds([0, 0, 0, 1, 0, 0]))