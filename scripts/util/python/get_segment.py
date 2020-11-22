import sys
from sys import argv

# USAGE
# python get_segment.py <string> <separator> <index>

def get_segment(s: str, sep: str, index: int) -> str:
    return s.split(sep)[index]

if __name__ == '__main__':
    s = sys.argv[1] if len(sys.argv) > 1 else ''
    sep = sys.argv[2] if len(sys.argv) > 2 else ''
    index = int(sys.argv[3]) if len(sys.argv) > 3 else 0
    print(get_segment(s, sep, index))
