from .services import greet, greet_everyone


def test_greet_empty_parameter():
    greetings = greet("")
    assert greetings == "Hello world !!!"


def test_greet_non_empty_parameter():
    greetings = greet("Kouga")
    assert greetings == "Hello Kouga"


def test_greet_everyone_empty_parameter():
    greetings = greet_everyone([])
    assert greetings == "Hello everyone !!!"


def test_greet_everyone_non_empty_parameter():
    greetings = greet_everyone(["Kouga", "Kaoru"])
    assert greetings == "Hello Kouga, Kaoru, and everyone"
