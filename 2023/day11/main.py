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

universe = []
for line in read_lines:
    universe.append(list(line))
    
expanded_universe = []
for line in universe:
    expanded_universe.append(line)
    if "#" not in line:
        expanded_universe.append(line)
        
tmp = np.array(universe)
universe = tmp.transpose()
universe = universe.tolist()
for line in universe:
    expanded_universe.append(line)
    if "#" not in line:
        expanded_universe.append(line)
print(expanded_universe)