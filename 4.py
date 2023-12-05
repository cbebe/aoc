from collections import defaultdict

text = """Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
"""

text = open('input.txt').read()

# 5744980 too high

s = 0
d = defaultdict(lambda: 1)
for i, x in enumerate(text.split("\n")):
    if not x:
        continue
    d[i]
    a, b = x.split(":")
    w, n = [
        *map(lambda g: [
            *map(int, (h for h in g.split(" ") if h))
        ], b.split("|"))
    ]
    p = len(w) - len(set(w).difference(n))
    print(p)
    print(a, i, d)
    if p:
        for y in range(p):
            print(p, i+y+1)
            d[i+y+1] += d[i]
s = 0
for _, v in d.items():
    s += v
print(sum(d.values()))
