from helper import cli


@cli
def show_segment(text: str, separator: str, index: str) -> str:
    list_index=int(index)
    print(text.split(separator)[list_index])


if __name__ == '__main__':
    show_segment()