import math
import os
import random
import re
import sys

# Complete the sockMerchant function below.
def sockMerchant(n, ar):
    pairs = {}
    result = 0
    for x in ar:
        if x in pairs:
            if pairs[x] == 1:
                result += 1
                del pairs[x]
            else:
                pairs[x] = 1
        else:
            pairs[x] = 1
    return result

print(sockMerchant(9, [3, 3, 2, 1, 2, 3, 3, 2, 3]))