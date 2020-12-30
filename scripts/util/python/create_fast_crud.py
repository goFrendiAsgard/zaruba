from common_helper import get_argv

import os, re, sys, traceback

# USAGE
# python create_fast_crud <location> <module> <entity> <fields>


def create_fast_crud(location: str, module: str, entity: str, fields: List[str]):
    print('Not implemented yet :(')


if __name__ == '__main__':
    location = get_argv(1)
    module = get_argv(2)
    entity = get_argv(3)
    str_fields = get_argv(4)
    fields = fields.split(',') if fields != '' else []
    try:
        create_fast_crud(location, module, entity)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)
