from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
Lines = file1.read().splitlines(keepends=False)

# 12 red cubes, 13 green cubes, and 14 blue cubes.
AVAILABLE_CUBES_PER_COLOR = {"red": 12, "green": 13, "blue": 14}


# 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
# 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
# 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
# 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
# 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
def check_if_game_is_possible(line):
    colors = line.split(";")
    for color in colors:
        cubes = color.split(",")
        for cube in cubes:
            [extra_space, number, color] = cube.split(" ")
            if int(number) > AVAILABLE_CUBES_PER_COLOR[color]:
                return False
    return True


def get_game_id_if_possible(line):
    [game_name, game_values] = line.split(":")
    game_id = int(game_name.split(" ")[1])
    if check_if_game_is_possible(game_values):
        return game_id
    else:
        return None


res = 0
for line in Lines:
    game_id = get_game_id_if_possible(line)
    if game_id:
        res += game_id
print(res)
