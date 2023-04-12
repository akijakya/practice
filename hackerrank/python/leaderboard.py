import copy

def climbingLeaderboard(scores, alice):
    keys = list(dict.fromkeys(scores))
    result = []
    i = len(keys) - 1
            
    def findRank(x, i):
        if x > keys[0]:
            result.append(1)
        elif x < keys[i]:
            result.append(i + 2)
        elif x == keys[i]:
            result.append(i + 1)
        else: 
            i -= 1
            i = findRank(x, i)
        return i

    for x in alice:
        i = findRank(x, i)

    return result



print(climbingLeaderboard([100, 100, 50, 40, 40, 20, 10], [5, 25, 50, 120]))