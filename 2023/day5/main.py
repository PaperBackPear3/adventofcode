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
                steps.append(line)
            else:
                steps[index-1] += (int(line.split(" ")))
        
    return steps

get_steps_maps_from_lines(read_lines)

#[destination_range_start ,source_range_start, range_length] = 