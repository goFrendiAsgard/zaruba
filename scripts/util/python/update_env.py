from project.helpers import adjust_task_env
import sys
import os
import project

# USAGE
# python update_env.py <project_dir>

def update_env(project_dir:str):
    if not os.path.isabs(project_dir):
        project_dir = os.path.abspath(project_dir)
    for root, dir_names, file_names in os.walk(project_dir):
        for file_name in file_names:
            if not file_name.endswith('.zaruba.yaml'):
                continue
            task_file_name = os.path.join(root, file_name)
            task_file_dict = project.get_dict_from_file(task_file_name)
            if 'tasks' not in task_file_dict:
                continue
            for task_name, task_dict in task_file_dict['tasks'].items():
                task = project.Task(task_dict)
                task = adjust_task_env(task, task_file_name)
                project.write_task_env(project_dir, task)
                task_file_dict['tasks'][task_name] = task.as_dict()
            project.save_dict_to_file(task_file_name, task_file_dict)


if __name__ == '__main__':
    project_dir = sys.argv[1] if len(sys.argv) > 1 and sys.argv[1] != '' else '.'
    update_env(project_dir)