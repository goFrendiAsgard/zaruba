from typing import List

import re
import os
import subprocess
import sys


def render(zaruba_home: str, toc_file_path: str) -> List[str]:
    # get items
    core_task_file_name =  os.path.join(zaruba_home, 'core.zaruba.yaml')
    print('Extract tasks')
    tasks = extract_tasks(core_task_file_name)
    zaruba_bin_path = os.path.join(zaruba_home, 'zaruba')
    print('Extract utils')
    utils = extract_utils(zaruba_bin_path)
    print('Create items')
    items = get_items(tasks, utils, 1)
    print('Override TOC file')
    update_toc_file(toc_file_path, items)
    print('Render docs based on TOC')
    render_toc(zaruba_bin_path, toc_file_path)
    # find built-in location
    built_in_docs_dir = find_built_in_docs_dir(toc_file_path)
    print('Render tasks')
    render_task_docs(zaruba_home, zaruba_bin_path, built_in_docs_dir, tasks)
    print('Render utils')
    render_util_docs(built_in_docs_dir, utils)


def render_util_docs(built_in_docs_dir: str, utils: List[List[str]]):
    for util in utils:
        print('render util', util)
        doc_content = get_command_output(get_help_command(util))
        doc_content = '\n'.join(['```', doc_content, '```'])
        doc_path = get_util_doc_path(built_in_docs_dir, util)
        # read old content
        doc_file = open(doc_path, 'r')
        doc_lines = doc_file.readlines()
        doc_file.close()
        new_doc_lines: List[str] = []
        for doc_line in doc_lines:
            if doc_line.startswith('> TODO'):
                new_doc_lines.append(doc_content + '\n')
                continue
            new_doc_lines.append(doc_line)
        # override old content
        doc_file = open(doc_path, 'w')
        doc_file.writelines(new_doc_lines)
        doc_file.close()


def get_util_doc_path(built_in_docs_dir: str, util: List[str]) -> str:
    kebab_utils = [camel_to_kebab(part) for part in util]
    readme_part = kebab_utils[1:] 
    readme_part.append('README.md')
    readme_part = [built_in_docs_dir, 'utils'] + readme_part
    readme_path = os.path.join(*readme_part)
    if os.path.exists(readme_path):
        return readme_path
    doc_part = kebab_utils[1:len(kebab_utils)-1] 
    doc_part.append(kebab_utils[-1] + '.md')
    doc_part = [built_in_docs_dir, 'utils'] + doc_part
    doc_path = os.path.join(*doc_part)
    return doc_path


def render_task_docs(zaruba_home: str, zaruba_bin_path: str, built_in_docs_dir: str, tasks: List[str]):
    for task in tasks:
        print('render task', task)
        # get doc content
        doc_content = get_command_output([zaruba_bin_path, 'please', task, '-x', '-d=colorless'])
        doc_content = prepare_task_doc_content(zaruba_home, doc_content)
        kebab_task = camel_to_kebab(task)
        doc_path = os.path.join(built_in_docs_dir, 'tasks', kebab_task+'.md')
        # read old content
        doc_file = open(doc_path, 'r')
        doc_lines = doc_file.readlines()
        doc_file.close()
        # create new content
        new_doc_lines: List[str] = []
        for doc_line in doc_lines:
            if doc_line.startswith('> TODO'):
                new_doc_lines.append(doc_content + '\n')
                continue
            new_doc_lines.append(doc_line)
        # override old content
        doc_file = open(doc_path, 'w')
        doc_file.writelines(new_doc_lines)
        doc_file.close()


def prepare_task_doc_content(zaruba_home: str, doc_content: str) -> str:
    zaruba_home_trail_slash = zaruba_home.rstrip('/') + '/'
    zaruba_home_pattern = re.compile(zaruba_home_trail_slash)
    doc_content = zaruba_home_pattern.sub('${ZARUBA_HOME}', doc_content)
    doc_lines = doc_content.split('\n')
    is_processing_dependencies, is_processing_extends = False, False
    bullet_task_pattern = re.compile('^- `(.*)`$')
    for index, line in enumerate(doc_lines):
        if line == '## Extends':
            is_processing_extends = True
            continue
        if line == '## Dependencies':
            is_processing_dependencies = True
            continue
        if line.startswith('#'):
            is_processing_dependencies, is_processing_extends = False, False
            continue
        # it is a bullet pattern while we process extends/dependencies
        if (is_processing_dependencies or is_processing_extends) and line.startswith('- `'):
            match = bullet_task_pattern.findall(line)
            if len(match) > 0:
                task_name = match[0]
                kebab_task_name = camel_to_kebab(task_name)
                new_line = '- [{}]({})'.format(task_name, kebab_task_name + '.md')
                doc_lines[index] = new_line
    return '\n'.join(doc_lines)


def camel_to_kebab(name: str) -> str:
    name = re.sub('(.)([A-Z][a-z]+)', r'\1-\2', name)
    return re.sub('([a-z0-9])([A-Z])', r'\1-\2', name).lower()


def find_built_in_docs_dir(toc_file_path: str) -> str:
    toc_dir_path = os.path.dirname(toc_file_path)
    toc_file = open(toc_file_path, 'r')
    toc_lines = toc_file.readlines()
    toc_file.close()
    built_in_pattern = re.compile('^[ \t]*- \[[^a-zA-Z0-9_\-]*built-in.*\]\((.*)\).*$', re.IGNORECASE)
    for toc_line in toc_lines:
        built_in_match = built_in_pattern.findall(toc_line)
        if len(built_in_match) > 0:
            built_in_dir_name = os.path.dirname(built_in_match[0])
            return os.path.join(toc_dir_path, built_in_dir_name)
    return ''


def render_toc(zaruba_bin_path: str, toc_file_path: str):
    # call zaruba toc
    get_command_output([zaruba_bin_path, 'toc', toc_file_path])


def update_toc_file(toc_file_path: str, items: List[str]):
    new_toc_lines = get_new_toc_lines(toc_file_path, items)
    toc_file = open(toc_file_path, 'w')
    toc_file.writelines(new_toc_lines)
    toc_file.close() 


def get_new_toc_lines(toc_file_path: str, items: List[str]) -> List[str]:
    toc_file = open(toc_file_path, 'r')
    toc_lines = toc_file.readlines()
    toc_file.close()
    new_toc_lines: List[str] = []
    start_toc_pattern = re.compile('<!--startToc-->')
    end_toc_pattern = re.compile('<!--endToc-->')
    built_in_pattern = re.compile('^([ \t]*)- [^a-zA-Z0-9_\-]*built-in.*$', re.IGNORECASE)
    item_pattern = re.compile('^([ \t]*)-.*$')
    built_in_indentation = ''
    start_toc_found, end_toc_found, handle_built_in = False, False, False
    for toc_line in toc_lines:
        if start_toc_pattern.match(toc_line):
            start_toc_found = True
        if end_toc_pattern.match(toc_line):
            end_toc_found = True
        built_in_match = built_in_pattern.findall(toc_line)
        # process builtin
        if start_toc_found and (not handle_built_in) and (not end_toc_found) and len(built_in_match) > 0:
            handle_built_in = True
            built_in_indentation = built_in_match[0]
            new_toc_lines.append(toc_line)
            for item in items:
                new_toc_lines.append(built_in_indentation + item + '\n')
            continue
        # ignore old children of builtin
        if start_toc_found and handle_built_in and (not end_toc_found):
            item_match = item_pattern.findall(toc_line)
            if len(item_match) > 0:
                item_indentation = item_match[0]
                if len(item_indentation) > len(built_in_indentation):
                    continue
                else:
                    handle_built_in = False
        new_toc_lines.append(toc_line)
    return new_toc_lines


def get_items(tasks: List[str], utils: List[List[str]], level: int) -> List[str]:
    result = []
    # tasks
    result.append(get_item('Tasks', level))
    for task in tasks:
        result.append(get_item(task, level + 1))
    # utils
    result.append(get_item('Utils', level))
    for util in utils:
        result.append(get_item(util[-1], level + len(util)-1))
    return result


def get_item(line: str, level: int) -> str:
    return ''.join([' ' for _ in range(level*2)]) + '- ' + line


def extract_tasks(core_task_file_name: str) -> List[str]:
    tasks: List[str] = []
    core_task_file = open(core_task_file_name, 'r')
    lines = core_task_file.readlines()
    r = re.compile('^[ \t]*- ./zaruba-tasks/.*/task\.(.*).y[a]?ml$')
    for line in lines:
        match = r.findall(line)
        if len(match) > 0:
            tasks.append(match[0])
    return tasks


def extract_utils(zaruba_bin_path: str) -> List[List[str]]:
    return extract_subutils([zaruba_bin_path])


def extract_subutils(command_parts: List[str]) -> List[List[str]]:
    command = get_help_command(command_parts)
    output = get_command_output(command)
    lines = output.split("\n")
    is_subutil = False
    r = re.compile('^[ ]+([a-zA-Z0-9_\-]+)[ ]*.*$')
    result: List[List[str]] = []
    for line in lines:
        # subutil started with 'Available Commands:' and ended with empty line
        if line == 'Available Commands:':
            is_subutil = True
            continue
        if is_subutil and line == '':
            is_subutil = False
            continue
        if not is_subutil:
            continue
        match = r.findall(line)
        if len(match) > 0:
            subcommand = match[0]
            new_command_parts = list(command_parts)
            new_command_parts.append(subcommand)
            result.append(new_command_parts)
            result += extract_subutils(new_command_parts)
    return result


def get_help_command(command_parts: List[str]) -> List[str]:
    command_parts_copy = list(command_parts)
    command_parts_copy.append('--help')
    return command_parts_copy


def get_command_output(command: List[str]) -> str:
    output = subprocess.check_output(command)
    return output.decode('utf-8')


if __name__ == '__main__':
    zaruba_home = sys.argv[1]
    toc_file_path = sys.argv[2]
    render(zaruba_home, toc_file_path)