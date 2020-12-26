from common_helper import get_argv

import os, sys, traceback

# USAGE
# python get_segment.py <string> <separator> <index>

def get_segment(s: str, sep: str, index: int) -> str:
    return s.split(sep)[index]

if __name__ == '__main__':
    try:
        s = get_argv(1)
        sep = get_argv(2)
        index = int(get_argv(3, '0'))
        print(get_segment(s, sep, index))
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)