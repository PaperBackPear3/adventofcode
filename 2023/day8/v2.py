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
GLOBAL_LETTER_LENGHT = len(letter_maps)


# read all lines and choose the next line based on the left_right list
def prepare_input(read_lines):
    inputs = {}
    for line in read_lines:
        [name, next_nodes] = line.split(" = ")
        inputs[name] = next_nodes.strip("()").split(", ")

    return inputs


parsed_input = prepare_input(read_lines)

# global_count = 0


# def get_next_move(node, index, count):
#     # check if its last index if so reset it to start
#     if index == len(letter_maps):
#         index = 0
#     count += 1
#     if letter_maps[index] == "L":
#         if node[0] == "ZZZ":
#             return count
#         return get_next_move(parsed_input[node[0]], index + 1, count)
#     elif letter_maps[index] == "R":
#         if node[1] == "ZZZ":
#             return count
#         return get_next_move(parsed_input[node[1]], index + 1, count)


start_time = time.time()
# res = get_next_move(parsed_input["AAA"], 0, global_count)
index = 0
current_node = parsed_input["AAA"]
letter_index = 0
while parsed_input[current_node[0]] != "ZZZ" or parsed_input[current_node[1]] != "ZZZ":
    if letter_index == GLOBAL_LETTER_LENGHT:
        letter_index = 0
    if letter_maps[letter_index] == "L":
        current_node = parsed_input[current_node[0]]
    elif letter_maps[letter_index] == "R":
        current_node = parsed_input[current_node[1]]
    index += 1
    letter_index += 1

end = time.time()
print(end - start_time)
print(index)
