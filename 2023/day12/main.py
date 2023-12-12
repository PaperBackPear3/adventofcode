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


def remove_non_changing_positions(spring_string: str, max_line_quantity: int) -> str:
    # removes from string all the positions that are not changing through all the arrangements
    check_string = ["#" * max_line_quantity]
    hyphotetical_full_string = ["?" * max_line_quantity]
    startin_index = 0

    # count max number of # that are next to each other in the string
    max_hashtags = 0
    tmp_max_hashtags = 0
    start_index = 0
    tmp_start_index = 0

    for index, char in enumerate(spring_string):
        if char == "#" or char == "?":
            if tmp_start_index == 0:
                tmp_start_index = index
            tmp_max_hashtags += 1
        if char == ".":
            if tmp_max_hashtags > max_hashtags:
                max_hashtags = tmp_max_hashtags
                start_index = tmp_start_index
            tmp_max_hashtags = 0
            tmp_start_index = 0
        

    if max_hashtags == max_line_quantity:
        # remove those characters from the string and return it
        new_string = ""
        for index, char in enumerate(spring_string):
            if index < start_index or index > start_index + max_hashtags:
                new_string += char
        return new_string

    if spring_string.find(check_string):
        startin_index = spring_string.find(check_string)
        cutted_string = ""
        for stirng_index, char in enumerate(spring_string):
            if (
                stirng_index >= startin_index
                and stirng_index < startin_index + max_line_quantity
            ):
                cutted_string += char
        return {
            "cutted_string": cutted_string,
            "starting_index": startin_index,
            "ending_index": startin_index + max_line_quantity,
            "count": 1,
        }
    # check if there is a spot with all ? that can hold the max_line_quantity
    if hyphotetical_full_string in spring_string:
        startin_index = spring_string.find(hyphotetical_full_string)
        for stirng_index, char in enumerate(spring_string):
            if (
                stirng_index >= startin_index
                and stirng_index < startin_index + max_line_quantity
            ):
                cutted_string += char
        return {
            "cutted_string": spring_string,
            "starting_index": startin_index,
            "ending_index": startin_index + max_line_quantity,
            "count": 1,
        }
    # check if there is a possible string that can contain the max_line_quantity (for example ??## can contain 3 but also 4 as .### or ####)
    # if there is no possible string that can contain the max_line_quantity then return the string as it is


def find_possible_patters_in_spring_string(spring_string: str, line_quantities: list):
    res = []
    while max(line_quantities) > 1:
        res.append(remove_non_changing_positions(spring_string, max(line_quantities)))
        line_quantities.remove(max(line_quantities))
        # first check for how many # closer (len quantity) to numers in line_quantities
        # then checks for the number of ? in the string closer to the quantity (len + or - 1 to the quantity)
    return res


springs_lines = []
springs_quantities = []
for line in read_lines:
    split_line = line.split(" ")
    springs_lines.append(split_line[0])
    springs_quantities.append([int(x) for x in split_line[1].split(",")])

arrangement = []
for index, springs in enumerate(springs_lines):
    arrangement.append(
        find_possible_patters_in_spring_string(springs, springs_quantities[index])
    )

print(sum(arrangement))

print(springs)
print(springs_quantities)


# check for already existing patterns like
# ?###???????? 3,2,1 => ### is 3
