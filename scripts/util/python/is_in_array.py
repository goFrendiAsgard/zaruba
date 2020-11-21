import sys

# USAGE
# python is_in_array.py <needle> <haystack> [separator=\n]

def is_in_array(needle: str, haystack: str) -> bool:
    return needle in haystack

if __name__ == "__main__":
    needle = sys.argv[1]
    haystack = sys.argv[2]
    separator = sys.argv[3] if len(sys.argv) > 3 else "\n"
    if is_in_array(needle, haystack):
        print(1)
    else:
        print(0)