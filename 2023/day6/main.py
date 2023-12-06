from sys import argv
import os
import math

file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)

#to win the boat needs to go farther then the record holder in each race
#time spend holding the button counts and the more is pressed the faster the boat goes

# Time:      7  15   30  race 1-2-3 in milliseconds
# Distance:  9  40  200  record distances in millimiters

#1 millisecond holds makes you go 1 millimiter per milliseconds so 
#tempo totale gara 7ms - tempo speso a caricare 1ms = 6ms quindi distanza = 6ms * 1mm/s = 6mm

def calc_diff_min_max(time,distance):
  delta = (time ** 2) - (4 * distance)
  delta = math.sqrt(delta)
  delta = round(delta)
  min_val = (-time - delta) / 2
  max_val = (-time + delta) / 2

  return (max_val-min_val)

times =[ int(character) for character in read_lines.pop(0).split(" ") if character.isdigit()]
record_distances = [ int(character) for character in read_lines.pop(0).split(" ") if character.isdigit()]

available_vals = []
for index,time in enumerate(times):
  available_vals.append( calc_diff_min_max(time, record_distances[index]))

mult =1

for val in available_vals:
  mult *= val

print(mult)

