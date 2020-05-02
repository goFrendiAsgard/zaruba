from flask import Flask


class Something():

    def __init__(self, app):
        self.app = app

    def setup(self):
        self.app.route("/")(self.doSomething)

    def doSomething(self):
        return "hello"


app = Flask(__name__)
Something(app).setup()

app.run("localhost", 3000)
