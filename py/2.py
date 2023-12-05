from collections import defaultdict
import re
text = """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."""

text = open("input.txt").read()
mat = [*map(lambda x: [*x], text.split("\n"))]

num = [
    (int(m.group(0)), x[0], m.span())
    for x in enumerate(text.split('\n'))
    for m in re.finditer('[0-9]+', x[1])
]
by_y = defaultdict(dict)
for n in num:
    for e in range(n[2][0], n[2][1]):
        by_y[n[1]][e] = n

found = {}
s = 0
for y, row in enumerate(mat):
    for x, col in enumerate(row):
        if col == '*':
            gears = []
            for (i, j) in [
                (0, -1), (0, 1), (1, 0), (-1, 0),
                (1, -1), (-1, 1), (1, 1), (-1, -1),
            ]:
                if ((d := by_y.get(y+j))
                        and (e := d.get(x+i))
                        and (e[1], e[2]) not in found):
                    found[(e[1], e[2])] = True
                    gears.append(e[0])
            if len(gears) == 2:
                s += gears[0]*gears[1]
print(s)
