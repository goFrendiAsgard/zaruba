from ruamel.yaml import YAML

def include(filename: str):
    yaml = YAML()
    main_filename = 'main.zaruba.yaml'
    main_obj = yaml.load(open(main_filename, 'r'))
    main_obj['includes'].append(filename)
    yaml.dump(main_obj, open(main_filename, 'w'))
