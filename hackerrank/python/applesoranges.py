def countApplesAndOranges(s, t, a, b, apples, oranges):
    appleCount = 0
    orangeCount = 0

    for x in apples:
        if (a + x) >= s and (a + x) <= t:
            appleCount += 1

    for x in oranges:
        if (b + x) >= s and (b + x) <= t:
            orangeCount += 1

    print(f"{appleCount}\n{orangeCount}")

countApplesAndOranges(7, 11, 5, 15, [-2, 2, 1], [5, -6])