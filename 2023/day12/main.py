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

springs = []
line_characters = []
for line in read_lines:


# check for already existing patterns like
# ?###???????? 3,2,1 => ### is 3 