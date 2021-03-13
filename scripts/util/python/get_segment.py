from helper import cli


@cli
def show_segment(text: str, separator: str, index: int) -> str:
    print(text.split(separator)[index])


if __name__ == '__main__':
    show_segment()