from itertools import accumulate, combinations
import os
import sys


def solve(file_name, spacing):
    rows = file_name
    rng = range(len(rows))
    cols = [[row[i] for row in rows] for i in rng]
    ys = list(accumulate(1 if "#" in y else spacing for y in rows))
    xs = list(accumulate(1 if "#" in x else spacing for x in cols))
    points = [(xs[x], ys[y]) for x in rng for y in rng if rows[y][x] == "#"]

    return sum(
        abs(x1 - x0) + abs(y1 - y0) for (x0, y0), (x1, y1) in combinations(points, 2)
    )


file_name = sys.argv[1]


dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)
file1.close()

res1 = solve(read_lines, 2)
res2 = solve(read_lines, 1000000)
print(f"res1: {res1}")
print(f"res2: {res2}")
