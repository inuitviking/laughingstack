#!/usr/bin/env python3
#

from concurrent.futures import ThreadPoolExecutor


with ThreadPoolExecutor(max_workers=6) as executor:
    three = list(range(0, 3))
    h_s = executor.map(lambda _: "h", three)
    a_s = executor.map(lambda _: "a", three)
    print("".join('{}{}'.format(x, y) for (x,y) in zip(h_s, a_s)))
