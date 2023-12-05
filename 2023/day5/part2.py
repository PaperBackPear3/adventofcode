from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)

total_points = 0
# seeds = int(read_lines[0].split(" "))


def filter_candidates_and_seed_maps(best_matching_points, seeds_and_offsets):
    results = set()
    for point in best_matching_points:
        for seed, offset in seeds_and_offsets.items():
            if point >= seed and point <= (seed + offset):
                results.add(point + offset)
    return results


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


def update_best_matching(best_matching_points, maps):
    new_best_matching_points = set()
    for map_data in maps:
        for point in best_matching_points:
            if map_data[0] <= point and point <= (map_data[0] + map_data[2]):
                new_best_matching_points.add(point + map_data[2])
    # add boundaries
    for map_data in maps:
        new_best_matching_points.add(map_data[1])
        new_best_matching_points.add(map_data[1] + map_data[2] - 1)
    return new_best_matching_points


seeds = [
    int(character) for character in read_lines.pop(0).split(" ") if character.isdigit()
]

seeds_and_offsets = {}
for index in range(0, len(seeds) - 1, 2):
    seeds_and_offsets[int(seeds[index])] = seeds[index + 1]

steps_map_dest_sour_off = get_steps_maps_from_lines(read_lines)

res = []

best_matching_points = set()
for maps in steps_map_dest_sour_off[::-1]:
    best_matching_points = update_best_matching(best_matching_points, maps)

best_matching_points = filter_candidates_and_seed_maps(
    best_matching_points, seeds_and_offsets
)

print(min(best_matching_points))
