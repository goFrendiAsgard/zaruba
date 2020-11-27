import sys
from ruamel.yaml import YAML

# USAGE
# python add_project_kwarg.py <key> <value> [file]

def set_project_kwarg(key: str, value: str, file_name: str):
    yaml=YAML()
    obj = yaml.load(open(file_name, 'r'))
    obj[key] = value
    yaml.dump(obj, open(file_name, 'w'))
    

if __name__ == '__main__':
    key = sys.argv[1]
    value = sys.argv[2]
    file_name = sys.argv[3] if len(sys.argv) > 3 else 'default.kwargs.yaml'
    set_project_kwarg(key, value, file_name)
