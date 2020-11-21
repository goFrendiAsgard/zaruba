import sys

# USAGE
# python get_segment.py <string> <separator> <index>

def get_segment(s: str, sep: str, index: int) -> str:
    return s.split(sep)[index]

if __name__ == "__main__":
    s = sys.argv[1]
    sep = sys.argv[2]
    index = int(sys.argv[3])
    print(get_segment(s, sep, index))
