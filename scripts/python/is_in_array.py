from typing import List
from helper import cli


@cli
def show_is_in_array(needle: str, separator: str, haystacks: str = ''):
    haystack_list = [haystack.strip() for haystack in haystacks.split(separator)]
    is_in_array = 1 if needle in haystack_list else 0
    print(is_in_array)


if __name__ == '__main__':
    show_is_in_array()