from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
Lines = file1.readlines()

max_val = 0

bigger_max = 0
medium_max = 0
smaller_max = 0

for line in Lines:
    if line != "\n":
        max_val += int(line)
    if line == "\n":
        if max_val > bigger_max:
            smaller_max = medium_max
            medium_max = bigger_max
            bigger_max = max_val
        elif max_val > medium_max:
            smaller_max = medium_max
            medium_max = max_val
        elif max_val > smaller_max:
            smaller_max = max_val
        max_val = 0
file1.close()

print(
    "res: "
    + str(
        [
            bigger_max,
            medium_max,
            smaller_max,
            sum([bigger_max, medium_max, smaller_max]),
        ]
    )
)
