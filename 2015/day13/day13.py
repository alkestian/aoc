from itertools import permutations

def calculate_happiness(seating, happiness):
    total_happiness = 0
    num_guests = len(seating)

    for i in range(num_guests):
        left = seating[i]
        right = seating[(i + 1) % num_guests]
        total_happiness += happiness.get((left, right), 0)
        total_happiness += happiness.get((right, left), 0)

    return total_happiness

def part1(names, happiness):
    max_happiness = float('-inf')
    for perm in permutations(names):
        conf_happiness = calculate_happiness(perm, happiness)
        max_happiness = max(max_happiness, conf_happiness)
    return max_happiness

def part2(names, happiness):
    for name in names:
        happiness[("Cameron", name)] = 0
        happiness[(name, "Cameron")] = 0
    names = list(names)
    names.append("Cameron")
    max_happiness = float('-inf')
    for perm in permutations(names):
        conf_happiness = calculate_happiness(perm, happiness)
        max_happiness = max(max_happiness, conf_happiness)
    return max_happiness

f = open("./2015/Day13/input.txt")
lines = f.readlines()
happiness = {}
names = []
for line in lines:
    pieces = line.split(" ")
    if pieces[2] == "gain":
        score = int(pieces[3])
    else:
        score = -int(pieces[3])
    happiness[(pieces[0], pieces[-1][:-2])] = score
    names.append(pieces[0])
names = set(names)
print(part2(names, happiness))







