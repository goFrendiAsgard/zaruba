import os, sys, json

toc_file_path = os.path.join(os.getcwd(), './docs/README.md')
task_list = json.loads(sys.argv[1])

toc_file = open(toc_file_path, 'rw')
toc_lines = toc_file.read().split('\n')
new_toc_lines = []
skip = False
for toc_line in toc_lines:
    if toc_line == '* [Core Tasks](core-tasks/README.md)':
        skip = True
        new_toc_lines.append(toc_line)
        task_lines = ['  * [{task_name}](core-tasks/{task_name}.md)'.format(task_name=task_name) for task_name in task_list]
        new_toc_lines += task_lines
        continue
    if skip:
        if toc_line.startswith('  *'):
            continue
        else:
            skip = False
    new_toc_lines.append(toc_line)
toc_file.write('\n'.join(new_toc_lines))
