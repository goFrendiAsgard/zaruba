from typing import List


def greet(name: str) -> str:
    if name == "":
        return "Hello world !!!"
    return "Hello {}".format(name)


def greet_everyone(names: List[str]) -> str:
    if len(names) == 0:
        return "Hello everyone !!!"
    return "Hello {}, and everyone".format(", ".join(names))
