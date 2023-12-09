import os
import time
import sys

sys.setrecursionlimit(100000)
file_name = sys.argv[1]


dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)
file1.close()


def calulate_steps_diffrence(values):
    diffrences = []
    for index in range(len(values) - 1):
        diffrences.append(values[index + 1] - values[index])
    return diffrences


def calculate_next_prediction(diffrences):
    diffrences = diffrences[::-1]
    for index in range(len(diffrences) - 1):
        diffrences[index + 1].insert(0, diffrences[index + 1][0] - diffrences[index][0])
    return diffrences[-1][0]


next_prediction = 0
for line in read_lines:
    diffrences = []
    steps = [int(x) for x in line.split(" ")]
    diffrences.append(steps)
    index = 0
    # are_all_equal = all(x == diffrences[index][0] for x in diffrences[index])
    are_all_equal = False
    while not are_all_equal:
        diffrences.append(calulate_steps_diffrence(diffrences[index]))
        index += 1
        are_all_equal = all(x == diffrences[index][0] for x in diffrences[index])

    next_prediction += calculate_next_prediction(diffrences)

print(next_prediction)
