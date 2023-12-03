from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)

[line_before, current_line, line_after] = [""] + read_lines[0:2]


def get_line_symbols_indexes(line):
    # check if line has symbols
    symbols_index = []
    for index, char in enumerate(line):
        if not char.isdigit() and char != ".":
            symbols_index.append(index)
    return symbols_index


def get_line_numbers_indexes(line):
    # check if line has numbers
    numbers_index = {"numbers": [], "indexs": []}
    number = ""
    for index, char in enumerate(line):
        if char.isdigit():
            number += char
            if index == len(line) - 1:
                numbers_index["numbers"].append(number)
                numbers_index["indexs"].append({"start": index - len(number) + 1, "end": index-1})
        if char == ".":
            if number:
                numbers_index["numbers"].append(number)
                numbers_index["indexs"].append({"start": index - len(number), "end": index-1})
                number = ""
    return numbers_index


def check_if_prec_ext(
    line_before_numbers_indexes,
    current_line_numbers_indexes,
    line_after_numbers_indexes,
    indexes,
):
    sum_of_numbers = 0
    # check each character is digit, if yes take the whole number and check if there are nearby symbols that are not "." if condition is matched return the number
    for index in indexes:
        #check if line before has numbers at given index
        for number_index in line_before_numbers_indexes["indexs"]:
            if number_index - index ==0:
                sum_of_numbers += int(line_before_numbers_indexes["numbers"][line_before_numbers_indexes["indexs"].index(number_index)])
        if index in line_before_numbers_indexes["indexs"]:
            sum_of_numbers += int(line_before_numbers_indexes["numbers"][line_before_numbers_indexes["indexs"].index(index)])
    return sum_of_numbers

         


def iterate_over_lines(lines):
    for index, line in enumerate(lines):
        if index == 0:
            continue
        if index == len(lines) - 1:
            break
        [line_before, current_line, line_after] = lines[index - 1 : index + 2]
        [
            line_before_numbers_indexes,
            current_line_numbers_indexes,
            line_after_numbers_indexes,
        ] = [
            get_line_numbers_indexes(line_before),
            get_line_numbers_indexes(current_line),
            get_line_numbers_indexes(line_after),
        ]
        indexes = get_line_symbols_indexes(current_line)
        res = check_if_prec_ext(
            line_before_numbers_indexes,
            current_line_numbers_indexes,
            line_after_numbers_indexes,
            indexes,
        )

        print(res)


iterate_over_lines(read_lines)
