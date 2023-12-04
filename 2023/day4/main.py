from sys import argv
import os


file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)

#checks if choosen numbers are in winning numbers and each match after the first one doubles the points
# def get_points(choosen_numbers, winning_numbers):
#   points = 0
#   isInStreak= False
#   for choosen_number in choosen_numbers:
#     if choosen_number in winning_numbers:
#       if isInStreak:
#         points *= 2
#       else:
#         isInStreak = True
#       points += 1
#     else:
#       isInStreak = False
#   return points
total_points = 0


def get_points(choosen_numbers, winning_numbers):
  count = 0
  for choosen_number in choosen_numbers:
    if choosen_number in winning_numbers:
      count += 1
  if(count == 0):
    return 0
  return pow(2,(count-1))


for line in read_lines:
    [card,values] = line.split(":")
    [choosen_numbers, winning_numbers] = values.split("|")
    choosen_numbers = [int(character) for character in choosen_numbers.split() if character.isdigit()]
    winning_numbers = [int(character) for character in winning_numbers.split() if character.isdigit()]
    total_points = total_points+ get_points(choosen_numbers, winning_numbers)
  
print(total_points)