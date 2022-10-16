from prefect import flow, task, get_run_logger

@task
def say_hello(name):
    logger = get_run_logger()
    hello = f"hello {name}"
    print(hello)
    logger.info(hello)

@task
def say_goodbye(name):
    logger = get_run_logger()
    goodbye = f"goodbye {name}"
    print(goodbye)
    logger.info(goodbye)


@flow(name="test flow")
def greetings(names=["arthur", "trillian", "ford", "marvin"]):
    for name in names:
        say_hello(name)
        say_goodbye(name)

if __name__ == "__main__":
    greetings(["arthur", "trillian", "ford", "marvin"])
