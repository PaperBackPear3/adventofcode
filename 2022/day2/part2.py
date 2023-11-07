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

conversion = {"X": "A", "Y": "B", "Z": "C"}


def getMatchPoints(opponent: str, mine: str):
    mypoints = symbol_points[mine]

    if opponent == mine:
        mypoints += 3
    if opponent == "C" and mine == "A":
        mypoints += 6

    if opponent < mine:
        if mine == "B":
            mypoints += 6
        if opponent == "B":
            mypoints += 6

    return mypoints


total_points = 0

for line in Lines:
    opponent_symbol, my_symbol = line.split(None, maxsplit=2)
    total_points += getMatchPoints(opponent_symbol, conversion[my_symbol])

print(total_points)
