from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
Lines = file1.readlines()

sum = 0

for line in Lines:
  # sum first and last numeric values in the string
  vals = [int(character) for character in line if character.isdigit()]
  sum += (vals[0]*10) + vals[-1]
        
print(sum)