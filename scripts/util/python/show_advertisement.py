from typing import List, Mapping
from ruamel.yaml import YAML
import os, random, re, time

# USAGE
# python show_advertisement.py

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


def decorate_message(message: str) -> str:
    lines = message.split('\n')
    max_line_length = 0
    for line in lines:
        if len(line) > max_line_length:
            max_line_length = len(line)
    decorated_lines = ['  ' + line for line in lines]
    horizontal_border = '****' + ''.ljust(max_line_length, '*')
    return '\n'.join([
        horizontal_border,
        '\n'.join(decorated_lines),
        horizontal_border,
    ])


if __name__ == '__main__':
    try:
        message = get_random_advertisement_message()
        print(decorate_message(message))
    except:
        print('Failed to show advertisement')