from flask import request


def get_name(name) -> str:
    if name == "" or name is None:
        name = request.args.get("name")
    if name == "" or name is None:
        name = request.form.get("name")
    if name is None:
        name = ""
    return name
