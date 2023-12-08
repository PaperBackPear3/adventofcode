import os
import time
import sys

sys.setrecursionlimit(100000)
file_name = sys.argv[1]


dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)
file1.close()


# parse input
letter_maps = [char for char in read_lines.pop(0)]
read_lines.pop(0)


# read all lines and choose the next line based on the left_right list
def prepare_input(read_lines):
    inputs = {}
    for line in read_lines:
        [name, next_nodes] = line.split(" = ")
        inputs[name] = next_nodes.strip("()").split(", ")

    return inputs


parsed_input = prepare_input(read_lines)

global_count = 0


def get_next_move(node, index, count):
    # check if its last index if so reset it to start
    if index == len(letter_maps):
        index = 0
    count += 1
    if letter_maps[index] == "L":
        if node[0] == "ZZZ":
            return count
        return get_next_move(parsed_input[node[0]], index + 1, count)
    elif letter_maps[index] == "R":
        if node[1] == "ZZZ":
            return count
        return get_next_move(parsed_input[node[1]], index + 1, count)


res = get_next_move(parsed_input["AAA"], 0, global_count)
print(res)
