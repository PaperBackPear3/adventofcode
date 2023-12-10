import os
import time
import sys

sys.setrecursionlimit(10000000)
file_name = sys.argv[1]


dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)
file1.close()


# Pipes are represented by the following characters:
# | is a vertical pipe connecting north and south.
# - is a horizontal pipe connecting east and west.
# L is a 90-degree bend connecting north and east.
# J is a 90-degree bend connecting north and west.
# 7 is a 90-degree bend connecting south and west.
# F is a 90-degree bend connecting south and east.
# . is ground; there is no pipe in this tile.
# S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
def get_starting_coordinates(line):
    """returns the x coordinate of the starting pipe"""
    for index, value in enumerate(line):
        if line[index] == "S":
            return index


def create_map(input_files):
    """creates map from input file and returns it as a list of lists containig all the pipes and starting coordinates"""
    pipe_map = []
    starting_coords = []
    for y_cord, line in enumerate(input_files):
        pipe_map.append(list(line))
        if not starting_coords:
            x_cord = get_starting_coordinates(line)
            if x_cord is not None:
                starting_coords = [y_cord, x_cord]
    return pipe_map, starting_coords


def find_pipe_at_starting_point(pipe_map, starting_coords):
    """finds the pipe at the starting point by finding what are the connections to the pipe and getting the one mathcing"""

    upper_point = pipe_map[starting_coords[0] - 1][starting_coords[1]]
    lower_point = pipe_map[starting_coords[0] + 1][starting_coords[1]]
    previous_point = pipe_map[starting_coords[0]][starting_coords[1] - 1]
    next_point = pipe_map[starting_coords[0]][starting_coords[1] + 1]

    if upper_point == "|" or upper_point == "7" or upper_point == "F":
        if lower_point == "|" or lower_point == "L" or lower_point == "J":
            return "|"  # is a vertival pipe
        if previous_point == "-" or previous_point == "F" or previous_point == "L":
            return "J"  # is a 90-degree bend connecting north and west.
        if next_point == "-" or next_point == "7" or next_point == "J":
            return "L"  # is a 90-degree bend connecting north and east.
    if lower_point == "|" or lower_point == "L" or lower_point == "J":
        if previous_point == "-" or previous_point == "F" or previous_point == "L":
            return "7"
        if next_point == "-" or next_point == "7" or next_point == "J":
            return "F"
    if (
        previous_point == "-" or previous_point == "F" or previous_point == "L"
    ):  # this is probably not neeed, if it cant connect to top or bottom is always a horizontal pipe ( - )
        if next_point == "-" or next_point == "7" or next_point == "J":
            return "-"


def get_direction_from_pipe(pipe):
    """returns the directions that the pipe can go to"""
    if pipe == "|":
        return ["up", "down"]
    if pipe == "-":
        return ["left", "right"]
    if pipe == "L":
        return ["up", "right"]
    if pipe == "J":
        return ["up", "left"]
    if pipe == "7":
        return ["down", "left"]
    if pipe == "F":
        return ["down", "right"]


def walk_pipes_to_find_furthest_point(current_coords, prev_coords, count):
    """walks the pipes to find the furthest points from the starting point"""
    current_point = WORLD_MAP[current_coords[0]][current_coords[1]]
    copy_current_coords = current_coords.copy()
    available_directions = get_direction_from_pipe(current_point)
    if available_directions == ["up", "down"]:
        # if prev_coords == copy_current_coords:  # its the starting poing
        if prev_coords[0] < copy_current_coords[0]:
            copy_current_coords[0] += 1
        else:
            copy_current_coords[0] -= 1
    if available_directions == ["left", "right"]:
        # check where we're coming from to know where to go
        if prev_coords[1] < copy_current_coords[1]:
            copy_current_coords[1] += 1
        else:
            copy_current_coords[1] -= 1
    if available_directions == ["up", "right"]:
        if prev_coords[0] < copy_current_coords[0]:
            copy_current_coords[1] += 1
        else:
            copy_current_coords[0] -= 1
    if available_directions == ["up", "left"]:
        if prev_coords[0] == copy_current_coords[0]:
            copy_current_coords[0] -= 1
        else:
            copy_current_coords[1] -= 1
    if available_directions == ["down", "left"]:
        if prev_coords[0] == copy_current_coords[0]:
            copy_current_coords[0] += 1
        else:
            copy_current_coords[1] -= 1
    if available_directions == ["down", "right"]:
        if prev_coords[0] == copy_current_coords[0]:
            copy_current_coords[0] += 1
        else:
            copy_current_coords[1] += 1
    count += 1
    if copy_current_coords == starting_coords:
        return count + 1
    prev_coords = current_coords
    loop.append(copy_current_coords)
    return walk_pipes_to_find_furthest_point(
        current_coords=copy_current_coords, prev_coords=prev_coords, count=count
    )


WORLD_MAP, starting_coords = create_map(read_lines)
WORLD_MAP[starting_coords[0]][starting_coords[1]] = find_pipe_at_starting_point(
    WORLD_MAP, starting_coords
)
count = 0
loop = [starting_coords]
FURTHEST_POINT = walk_pipes_to_find_furthest_point(
    current_coords=starting_coords, prev_coords=starting_coords, count=count
)
print(loop)
print(FURTHEST_POINT / 2)
