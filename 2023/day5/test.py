from typing import List, Set

# Represents one specific map
class Interval:
    def __init__(self, start=0, length=0, offset=0, line=''):
        if len(line) > 0:
            dest, start, length = [int(x) for x in line.split(' ')]
            offset = dest - start
        self.start = start
        self.end = start + length -1
        self.shift = offset

    def contains(self, point: int):
        return self.start <= point <= self.end
    
    def outputs(self, output: int):
        return self.start + self.shift <= output <= self.end + self.shift
    
    def offset(self, point: int):
        return point + self.shift
    
    def undo_offset(self, offset_point: int):
        return offset_point - self.shift
    
# Represents one layer of Intervals
class IntervalList:
    def __init__(self, intervals: List[Interval]):
        self.intervals  = intervals

    def process(self, point: int):
        for interval in self.intervals:
            if interval.contains(point):
                return interval.offset(point)
        return point

    def undo_process(self, outputs: Set[int]):
        potential_inputs = set()
        for interval in self.intervals:
            for output in outputs:
                if interval.outputs(output):
                    potential_inputs.add(interval.undo_offset(output))
        outputs.update(potential_inputs)

    def filter(self, points: Set[int]):
        results = set()
        for point in points:
            if any(interval.contains(point) for interval in self.intervals):
                results.add(point)
        return results
    
    def boundaries(self, candidate_points: Set[int]):
        for interval in self.intervals:
            candidate_points.add(interval.start)
            candidate_points.add(interval.end)
    
def parse_input(input):
    input += '\n'
    lines = iter(input.split('\n'))

    seed_input, _ = next(lines).split(': ')[1], next(lines)
    seeds = [int(seed) for seed in seed_input.split(' ')]
    seed_intervals = [Interval(seeds[i], seeds[i+1]) for i in range(0, len(seeds), 2)]
    seed_intervals = IntervalList(seed_intervals)

    interval_lists: List[IntervalList] = []

    for x in range(7):
        _, line = next(lines), next(lines) #Skip the blank line and header lines

        intervals = []
        while len(line) != 0:
            intervals.append(Interval(line=line))
            line = next(lines)

        interval_list = IntervalList(intervals)
        interval_lists.append(interval_list)

    return seeds, seed_intervals, interval_lists

def part1(seeds: List[int], interval_lists: List[IntervalList]):
    min_seed = float('inf')
    for seed in seeds:
        for intervalList in interval_lists:
            seed = intervalList.process(seed)
        min_seed = seed if seed < min_seed else min_seed
    return min_seed

def part2(seed_intervals: IntervalList, interval_lists: List[IntervalList]):
    
    candidate_points = set() # Set of input seeds which can potentially results in the optimal location
    for intervalList in interval_lists[::-1]: # Go in reverse order from the last map to the first map
        intervalList.undo_process(candidate_points) # Undo the processing for candidate_points from lower maps
        intervalList.boundaries(candidate_points) # This map layer contributes its boundaries as candidates

    candidate_points = seed_intervals.filter(candidate_points) # Remove candidates that are not valid

    return part1(candidate_points, interval_lists) # Test the candidates and find the optimal one

def solution(input):
    seeds, seed_intervals, interval_lists = parse_input(input)

    part1_ans = part1(seeds, interval_lists)
    print(f'Part 1: {part1_ans}')

    part2_ans = part2(seed_intervals, interval_lists)
    print(f'Part 2: {part2_ans}')

input = """seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4"""


# 56 start +37 offset = 92 end 
# 60 destination

# 93 start +4 offset = 96 end
# 56 destination shift = 93 start - 56 destination = 37 offset sihfted

solution(input)