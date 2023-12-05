from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)

total_points = 0
# seeds = int(read_lines[0].split(" "))


def get_steps_maps_from_lines(lines):
    steps = []
    index = 0
    prev_line_key = ""
    for line in lines:
        if line != "":
            if index == 0:
                steps.append([])
            if index > 0:
                steps[-1].append(
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
    for value in step:
        if seed >= value[1] and seed < (value[1] + value[2]):
            next_map = seed + (value[0] - value[1])
    return next_map


def get_quantity_per_step(step, seed, offset):
    if seed >= step[1] and seed < (step[1] + step[2]):
        if (seed + offset) > (step[1] + step[2]):
            return (step[1] + step[2]) - (seed + offset)
        else:
            return offset


seeds = [
    int(character) for character in read_lines.pop(0).split(" ") if character.isdigit()
]

seeds_and_offsets = {}
for index in range(0, len(seeds) - 1, 2):
    seeds_and_offsets[int(seeds[index])] = seeds[index + 1]

steps_map_dest_sour_off = get_steps_maps_from_lines(read_lines)

res = []
for seed, offset in seeds_and_offsets.items():
    for values in steps_map_dest_sour_off:
        for step in values:
            qty_per_step = get_quantity_per_step(step, seed, offset)
    res.append(qty_per_step)

print(min(res))
