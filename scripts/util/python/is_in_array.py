from common_helper import get_argv
import sys, traceback

# USAGE
# python is_in_array.py <needle> <haystack> [separator=\n]

if __name__ == '__main__':
    try:
        needle = get_argv(1)
        haystack = get_argv(2)
        separator = get_argv(3, '\n')
        in_array = 1 if needle in haystack else 0
        print(in_array)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)
