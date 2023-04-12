import math
import unittest

def repeatedString(s, n):
    ending = s[:n%len(s)]
    aCount = (list(s).count('a') * (math.floor(n/len(s)))) + list(ending).count('a')
    return aCount

class TestFunctions(unittest.TestCase):

    def test_three_chars(self):
        self.assertEqual(repeatedString('aba', 10), 7)

    def test_one_char(self):
        self.assertEqual(repeatedString('a', 100), 100)

if __name__ == '__main__':
    # begin the unittest.main()
    unittest.main()