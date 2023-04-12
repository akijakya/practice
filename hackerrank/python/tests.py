import unittest
from iqtest import is_even
from repeatedstrings import repeatedString

class TestFunctiona(unittest.TestCase):

    def test_evens(self):
        self.assertTrue(is_even([2, 2, 3, 4, 8]))

    def test_unevens(self):
        self.assertFalse(is_even([2, 1, 3, 7, 3]))

    def test_three_chars(self):
        self.assertEqual(repeatedString('aba', 10), 7)

    def test_one_char(self):
        self.assertEqual(repeatedString('a', 100), 100)

# class TestFunctions(unittest.TestCase):
    
#     def three_chars(self):
#         self.assertEqual(repeatedString('aba', 10), 7)

#     def one_char(self):
#         self.assertEqual(repeatedString('a', 100), 100)

if __name__ == '__main__':
    # begin the unittest.main()
    unittest.main()