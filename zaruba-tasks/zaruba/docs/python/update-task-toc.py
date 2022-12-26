import os, sys, json

toc_file_path = os.path.join(os.getcwd(), './docs/README.md')
task_icon_map = json.loads(sys.argv[1])
task_kebab_name_map = json.loads(sys.argv[2])

old_toc_file = open(toc_file_path, 'r')
toc_lines = old_toc_file.read().split('\n')
new_toc_lines = []
skip = False
for toc_line in toc_lines:
    if toc_line == '* [🥝 Core Tasks](coreTasks/README.md)':
        skip = True
        new_toc_lines.append(toc_line)
        task_lines = [
            '  * [{task_icon} {task_name}](coreTasks/{task_kebab_name}.md)'.format(
                task_name=task_name, 
                task_icon=task_icon_map[task_name], 
                task_kebab_name=task_kebab_name_map[task_name]
            ) for task_name in task_icon_map
        ]
        new_toc_lines += task_lines
        continue
    if skip:
        if toc_line.startswith('  *'):
            continue
        else:
            skip = False
    new_toc_lines.append(toc_line)

new_toc_file = open(toc_file_path, 'w')
new_toc_file.write('\n'.join(new_toc_lines))
