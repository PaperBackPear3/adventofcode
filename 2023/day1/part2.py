from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
Lines = file1.readlines()

DIGIT_MAPPING = {
  "one" :1,
  "two" :2,
  "three" :3,
  "four" :4,
  "five" :5,
  "six" :6,
  "seven" :7,
  "eight" :8,
  "nine" :9
  }

REVERSED_DIGIT_MAPPING ={
  "nine":9,
  "eight":8,
  "seven":7,
  "six":6,
  "five":5,
  "four":4,
  "three":3,
  "two":2,
  "one":1
  }


# extensionsToCheck = ['.pdf', '.doc', '.xls']
# url_string = 'http://.../foo.doc'

# print([extension for extension in extensionsToCheck if(extension in url_string)])

#convert number words to digits
Lines2 = []
for line in Lines:
  for key, value in REVERSED_DIGIT_MAPPING.items():
    line= line.replace(key, str(value))
  Lines2.append(line)

sum = 0

for line in Lines2:
  # sum first and last numeric values in the string
  vals = [int(character) for character in line if character.isdigit()]
  print((vals[0]*10) + vals[-1])
        
print(sum)