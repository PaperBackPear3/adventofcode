from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)

total_points = 0
#seeds = int(read_lines[0].split(" "))

def get_steps_maps_from_lines(lines):
    steps = []
    index = 0
    for line in lines:
        if line is not "":
            if index == 0:
                steps.append([line])
            if index > 0:
                steps[-1].append([int(character) for character in line.split() if character.isdigit()])
            index += 1
        else:
            index = 0
    return steps

steps = get_steps_maps_from_lines(read_lines)
seeds = [int(character) for character in steps[0][0].split(" ") if character.isdigit()]

for seed in seeds:
    total_points += seed
