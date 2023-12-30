# 5b.py -- PART 2

seeds = input()
seeds = seeds.split(':')[1].split()
seeds = [int(v) for v in seeds]
input()
map_names = []
maps = []

done = False
while True:
    try:
        map_name = input()
    except EOFError:
        break

    ranges = []

    line = input()
    while line.strip() != '':
        ds, ss, l = map(int, line.split())
        ranges.append([[ss, ss+l-1], [ds, ds+l-1]])

        try:
            line = input()
        except EOFError:
            done = True
            break
    ranges.sort()
    maps.append(ranges.copy())
    if done:
        break

# for m in maps:
#     print(m)


def seed_to_location(maps, seed):
    curr = seed
    for m in maps:
        for iv in m:
            if iv[0][0] <= curr <= iv[0][1]:
                curr = curr-iv[0][0] + iv[1][0]
                break
    return curr


def location_to_seed(maps, loc):
    curr = loc
    for m in maps[::-1]:
        for iv in m:
            if iv[1][0] <= curr <= iv[1][1]:
                curr = curr-iv[1][0] + iv[0][0]
                break
    return curr


def seed_in_ivs(ivs, seed):
    for iv in ivs:
        if iv[0] <= seed <= iv[1]:
            return True
    return False


seed_ivs = []

for i in range(0, len(seeds), 2):
    seed_ivs.append([seeds[i], seeds[i] + seeds[i+1]])

for loc in range(int(5e9)):
    if loc % int(1e6) == 0:
        print(loc)
    seed = location_to_seed(maps, loc)
    if seed_in_ivs(seed_ivs, seed):
        print("min_loc:", loc)
        break
