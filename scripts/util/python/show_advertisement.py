from typing import List, Mapping
from helper import cli
from ruamel.yaml import YAML
import helper.decoration as d
import os, random, re, time

# USAGE
# python show_advertisement.py

default_advertisement='''
 __________________________________
< Nothing to show, have a nice day >
 ----------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
'''

@cli
def show_advertisement():
    try:
        message = get_random_advertisement_message()
        print(decorate_message(message))
    except:
        print(decorate_message(default_advertisement))


def read_advertisement_dict() -> Mapping[str, Mapping[str, str]]:
    yaml = YAML()
    file_name = os.path.join(
        os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__)))),
        'advertisement.yaml'
    )
    f = open(file_name, 'r')
    adv_dict = yaml.load(f)
    f.close()
    return adv_dict


def complete_advertisement(adv: Mapping[str, str]) -> Mapping[str, str]:
    if 'start' not in adv:
        adv['start'] = '1900-01-01'
    adv['start'] = str(adv['start'])
    if 'end' not in adv:
        adv['end'] = '2900-01-01'
    adv['end'] = str(adv['end'])
    if 'pattern' not in adv:
        adv['pattern'] = '.*'
    return adv


def get_advertisement_list(adv_dict: Mapping[str, Mapping[str, str]]) -> List[Mapping[str, str]]:
    date_str = time.strftime('%Y-%m-%d')
    complete_adv_list = [complete_advertisement(adv) for adv in adv_dict.values()]
    return [
        adv for adv in complete_adv_list 
        if adv['start'] <= date_str and adv['end'] >= date_str and adv['pattern'] and re.search(adv['pattern'], date_str)
    ]


def get_random_advertisement_message():
    system_random = random.SystemRandom()
    adv_dict = read_advertisement_dict()
    adv_list = get_advertisement_list(adv_dict)
    adv = system_random.choice(adv_list)
    return adv['message'].strip()


def get_max_line_length(lines: List[str]) -> int:
    max_line_length = 0
    for line in lines:
        if len(line) > max_line_length:
            max_line_length = len(line)
    return max_line_length


def decorate_message(message: str) -> str:
    lines = message.split('\n')
    max_line_length = get_max_line_length(lines)
    horizontal_border = ''.ljust(max_line_length + 4, '*')
    decorated_horizontal_border = '{faint}{horizontal_border}{normal}'.format(faint=d.faint, horizontal_border=horizontal_border, normal=d.normal)
    decorated_lines = ['  {bold}{line}{normal}'.format(bold=d.bold, normal=d.normal, line=line)  for line in lines]
    return '\n'.join([
        decorated_horizontal_border,
        '\n'.join(decorated_lines),
        decorated_horizontal_border,
    ])


if __name__ == '__main__':
    show_advertisement()