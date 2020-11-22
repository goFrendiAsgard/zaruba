import sys
from ruamel.yaml import YAML

# USAGE
# python add_project_kwarg.py <key> <value> [file]

def set_project_kwarg(key: str, value: str, filename: str):
    yaml=YAML()
    obj = yaml.load(open(filename, 'r'))
    obj[key] = value
    yaml.dump(obj, open(filename, 'w'))
    

if __name__ == "__main__":
    key = sys.argv[1]
    value = sys.argv[2]
    filename = sys.argv[3] if len(sys.argv) > 3 else "default.kwargs.yaml"
    set_project_kwarg(key, value, filename)
