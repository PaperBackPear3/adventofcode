from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
Lines = file1.read().splitlines(keepends=False)


# 12 red cubes, 13 green cubes, and 14 blue cubes.


# 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
# 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
# 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
# 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
# 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
def get_max_by_color(line):
    max_per_color = {"blue": 0, "green": 0, "red": 0}
    colors = line.split(";")
    for color in colors:
        cubes = color.split(",")
        for cube in cubes:
            [extra_space, number, color] = cube.split(" ")
            if int(number) > max_per_color[color]:
                max_per_color[color] = int(number)

    return max_per_color


def find_power(line):
    game_values = line.split(":")[1]
    max = get_max_by_color(game_values)
    return max["blue"] * max["green"] * max["red"]


res = 0
for line in Lines:
    power = find_power(line)
    if power:
        res += power
print(res)
