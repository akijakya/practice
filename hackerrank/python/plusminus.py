def plusMinus(arr):
    pos = 0
    neg = 0
    zero = 0
    
    for x in arr:
        if x < 0:
            neg += 1
        elif x > 0:
            pos += 1
        else:
            zero += 1

    print ("{:.6f}".format(pos / len(arr)))
    print ("{:.6f}".format(neg / len(arr)))
    print ("{:.6f}".format(zero / len(arr)))

if __name__ == '__main__':
    plusMinus([-1, 1, 0, 4, 3, 0, -3])