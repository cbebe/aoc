text = """seeds: 79 14 55 13

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
56 93 4
"""

text = open('input.txt').read()

s, *maps = [s.split("\n") for s in text.strip().split("\n\n") if s]
ss = [*map(int, s[0].split(":")[1].strip().split(" "))]
maps = [*map(lambda x: [*map(lambda y: [int(z)
             for z in y.split(" ") if z], x[1:])], maps)]

# Modify ss to be pairs
new_ss = [(ss[i], ss[i+1]) for i in range(0, len(ss), 2)]
ss = new_ss

# NOT GONNA WORK
for x in maps:
    new_ss = []
    for a, ra in ss:
        to_find = {(a, a + ra - 1)}
        while len(to_find):
            start, end = to_find.pop()
            found = False
            for d in x:
                dest, src, r = d
                src_end = src + r
                start_in = start >= src and start < src_end
                end_in = end >= src and end < src_end
                # 3 cases:
                # In range (src + r)
                if start_in and end_in:
                    new_ss.append((dest + (start - src), ra))
                    found = True
                    break
                # Cut off left
                elif not start_in and end_in:
                    new_ss.append((dest, src - start))
                    to_find.add((start, src - 1))
                    found = True
                    break
                # Cut off right
                elif start_in and not end_in:
                    new_ss.append((dest + (start - src), src_end - start))
                    to_find.add((src_end, end))

            # if g >= src and g <= src + r:
            #     new_ss.append(dest + (g-src))
            #     found = True
            #     break

            # Not found in any of the ranges
            if not found:
                new_ss.append((start, (end - start) + 1))
        # if not found:
            # new_ss.append(g)
    ss = new_ss

# print(min(ss))

print(min(ss, key=lambda x: x[0])[0])
