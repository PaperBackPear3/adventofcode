# A for Rock, B for Paper, and C for Scissors.
# my symbols X for Rock, Y for Paper, and Z for Scissors.

# symbol point 1 for Rock, 2 for Paper, and 3 for Scissors
# outcome of the round points 0 if you lost, 3 if the round was a draw, and 6 if you won

from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
Lines = file1.readlines()

symbol_points = {"A": 1, "B": 2, "C": 3}
win_points = {"X": 0, "Y": 3, "Z": 6}
conversion = {"X": "A", "Y": "B", "Z": "C"}


def getMatchPoints(opponent: str, game_res: str):
    mypoints = win_points[game_res]

    if game_res == "X":
        if opponent == "A":
            mypoints += 3
        if opponent == "B":
            mypoints += 1
        if opponent == "C":
            mypoints += 2
    if game_res == "Y":
        mypoints += symbol_points[opponent]
    if game_res == "Z":
        if opponent == "A":
            mypoints += 2
        if opponent == "B":
            mypoints += 3
        if opponent == "C":
            mypoints += 1

    return mypoints


total_points = 0

for line in Lines:
    opponent_symbol, game_res = line.split(None, maxsplit=2)
    total_points += getMatchPoints(opponent_symbol, game_res)

print(total_points)
