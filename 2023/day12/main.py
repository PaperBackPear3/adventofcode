import os
import time
import sys
import numpy as np


sys.setrecursionlimit(10000)
file_name = sys.argv[1]


dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)
file1.close()

def find_possible_patters_in_spring_string(spring_string,line_quantities:list):
  for quantity in line_quantities:
    #first check for how many # closer (len + or - 1 to the quantity) to numers in line_quantities
    #then checks for the number of ? in the string closer to the quantity (len + or - 1 to the quantity)
    

springs = []
springs_quantities = []
for line in read_lines:
  split_line = line.split(" ")
  springs.append(split_line[0])
  springs_quantities.append([int(x) for x in split_line[1].split(",")])
  
print(springs)
print(line_characters)


# check for already existing patterns like
# ?###???????? 3,2,1 => ### is 3 