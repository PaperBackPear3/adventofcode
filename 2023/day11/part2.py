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


# def get_expanded_universe(universe):
#     tmp_universe = []
#     for line in universe:
#         tmp_universe.append(list(line))
#         if "#" not in line:
#             tmp_universe.append(list(line))

#     expanded_universe = np.array(tmp_universe)
#     transposed_universe = expanded_universe.transpose().tolist()
#     transposed_copy = transposed_universe.copy()
#     tmp_index = 0
#     for index, transposed_line in enumerate(transposed_copy):
#         if "#" not in transposed_line:
#             transposed_universe.insert(index + tmp_index, list(transposed_line))
#             tmp_index += 1
#     expanded_universe = np.array(transposed_universe).transpose()
#     return expanded_universe.tolist()


def get_pairs(expanded_universe):
    pairs = {}
    for index, line in enumerate(expanded_universe):
        for char_index, char in enumerate(line):
            if char == "#":
                pairs[(index, char_index)] = []
                for index2 in range(index, len(expanded_universe)):
                    for line_char_index in range(len(expanded_universe[index2])):
                        if expanded_universe[index2][line_char_index] == "#":
                            if index2 == index and line_char_index == char_index:
                                continue
                            pairs[((index, char_index))].append(
                                (index2, line_char_index)
                            )
    return pairs


def calculate_distance(
    starting_coords, pair_coordinates, rows_to_duplicate, colums_to_duplicate
):
    distances = []
    distance = 0
    for pair in pair_coordinates:
        distance = 0
        x1 = starting_coords[0]
        y1 = starting_coords[1]
        x2 = pair[0]
        y2 = pair[1]
        distance = abs(x1 - x2) + abs(y1 - y2)
        intersections = 0
        for column in colums_to_duplicate:
            if (column > starting_coords[1] and column < pair[1]) or (
                column < starting_coords[1] and column > pair[1]
            ):
                intersections += 1
        for row in rows_to_duplicate:
            if (row > starting_coords[0] and row < pair[0]) or (
                row < starting_coords[0] and row > pair[0]
            ):
                intersections += 1
        distances.append(distance + (intersections * 1000000))
    return sum(distances)


def get_duplicate_coordinates(expanded_universe):
    colums_to_duplicate = []
    for index, line in enumerate(expanded_universe):
        if "#" not in line:
            colums_to_duplicate.append(index)
    return colums_to_duplicate


universe = []
for line in read_lines:
    universe.append(list(line))

# expanded_universe = get_expanded_universe(universe)

rows_to_duplicate = get_duplicate_coordinates(universe)
tmp_uni = np.array(universe).transpose().tolist()
colums_to_duplicate = get_duplicate_coordinates(tmp_uni)

pairs = get_pairs(universe)


distances = []

tmp = list(pairs)
tmp.pop(-1)
for pos in tmp:
    distances.append(
        calculate_distance(pos, pairs[pos], rows_to_duplicate, colums_to_duplicate)
    )

print(sum(distances))
