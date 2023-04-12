import unittest

def timeConversion(s):
    hour = s[:2]
    if s[-2:] == "AM":
        if hour == "12":
            return f"00{s[2:-2]}"
        else:
            return s[:-2]
    else:
        if hour == "12":
            return f"12{s[2:-2]}"
        else:
            return f"{int(hour) + 12}{s[2:-2]}"


# class TestTime(unittest.TestCase):

#     def test_AM(self):
#         self.assertEqual(timeConversion("12:34:23AM"), "00:34:23")



print(timeConversion("12:34:23AM"))

# if __name__ == "__main__":
#     unittest.main()