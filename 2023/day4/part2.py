from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)

total_points = 0


def get_points(choosen_numbers, winning_numbers):
    count = 0
    for choosen_number in choosen_numbers:
        if choosen_number in winning_numbers:
            count += 1
    if count == 0:
        return 0
    return count


card_copy_tracker = [1] * len(read_lines)

for line_index, line in enumerate(read_lines):
    [card, values] = line.split(":")
    [choosen_numbers, winning_numbers] = values.split("|")
    choosen_numbers = [
        int(character) for character in choosen_numbers.split() if character.isdigit()
    ]
    winning_numbers = [
        int(character) for character in winning_numbers.split() if character.isdigit()
    ]
    total_points = get_points(choosen_numbers, winning_numbers)

    startIdx = line_index + 1
    endIdx = line_index + total_points
    currentCardCopies = card_copy_tracker[line_index]
    for i in range(startIdx, endIdx + 1):
        card_copy_tracker[i] += currentCardCopies

print(sum(card_copy_tracker))
