from flask import request


def get_name(name) -> str:
    if not name:
        name = request.args.get("name")
    if not name:
        name = request.form.get("name")
    if not name:
        name = ""
    return name
