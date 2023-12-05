from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)

total_points = 0
# seeds = int(read_lines[0].split(" "))


def get_steps_maps_from_lines(lines):
    steps = {}
    index = 0
    prev_line_key = ""
    for line in lines:
        if line != "":
            if index == 0:
                steps[line] = []
                prev_line_key = line
            if index > 0:
                steps[prev_line_key].append(
                    [
                        int(character)
                        for character in line.split()
                        if character.isdigit()
                    ]
                )

            index += 1
        else:
            index = 0
    return steps


def get_map_for_next_step(step, seed):
    next_map = seed
    for index, value in enumerate(step):
        if seed >= value[1] and seed < (value[1] + value[2]):
            next_map = seed + (value[0] - value[1])
            #
    return next_map


seeds = [
    int(character) for character in read_lines.pop(0).split(" ") if character.isdigit()
]
steps = get_steps_maps_from_lines(read_lines)

res = []
for seed in seeds:
    next = seed
    for key, values in steps.items():
        next = get_map_for_next_step(values, next)
    res.append(next)

print(min(res))
