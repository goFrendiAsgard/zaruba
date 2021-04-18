from typing import Mapping, Any, List
from dotenv import dotenv_values
from ruamel.yaml import YAML

import os, re, shutil, sys, traceback


def read_text(file_name: str) -> str:
    f_read = open(file_name, 'r')
    text = f_read.read()
    f_read.close()
    return text


def write_text(file_name: str, text: str):
    create_file_parent_dir(file_name)
    f_write = open(file_name, 'w')
    f_write.write(text)
    f_write.close()


def read_lines(file_name: str) -> List[str]:
    f_read = open(file_name, 'r')
    lines = f_read.readlines()
    f_read.close()
    return lines


def write_lines(file_name: str, lines: List[str]):
    create_file_parent_dir(file_name)
    f_write = open(file_name, 'w')
    f_write.writelines(lines)
    f_write.close()


def get_alphanum_basename(location: str) -> str:
    abs_location = os.path.abspath(location)
    base_name = os.path.basename(abs_location)
    return re.sub(r'[^A-Za-z0-9]+', ' ', base_name)


def capitalize(text: str) -> str:
    return ' '.join(sub[:1].upper() + sub[1:] for sub in text.split(' '))


def get_env_prefix(location: str) -> str:
    return ''.join(['_' + ch.lower() if ch.isupper() else ch for ch in get_alphanum_basename(location)]).lstrip('_').upper()


def get_service_name(location: str) -> str:
    capital_service_name = capitalize(get_alphanum_basename(location)).replace(' ', '')
    return capital_service_name[0].lower() + capital_service_name[1:]


def replace_str(string: str, replace_dict: Mapping[str, str]):
    new_string = string
    for key, val in replace_dict.items():
        new_string = new_string.replace(key, val)
    return new_string


def replace_in_file(file_name: str, replace_dict: Mapping[str, str]):
    if not replace_dict:
        return
    content = read_text(file_name)
    new_content = replace_str(content, replace_dict)
    if new_content == content:
        return
    write_text(file_name, new_content)


def replace_all(location: str, replace_dict: Mapping[str, str]):
    if os.path.isfile(location):
        replace_in_file(location, replace_dict)
        return
    for root, dir_names, file_names in os.walk(location):
        for file_name in file_names:
            replace_in_file(os.path.join(root, file_name), replace_dict)


def create_file_parent_dir(destination: str):
    destination = os.path.abspath(destination)
    destination_dir = os.path.dirname(destination)
    if not os.path.exists(destination_dir):
        os.makedirs(destination_dir)
