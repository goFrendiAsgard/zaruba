from task import Task
from generator_helper import read_config, write_config

import os, sys, traceback

# USAGE
# python create_helm_task.py

def create_helm_task():
    main_file_name = 'main.zaruba.yaml'
    main_config = read_config(main_file_name)
    if 'tasks' not in main_config:
        main_config['tasks'] = {}
    # helm apply
    helm_apply_task = Task({}).set_icon('🚢').set_extend('core.helmApply').set_location('helm-deployments').set_description('Deploy helm charts')
    main_config['tasks']['helmApply'] = helm_apply_task.as_dict()
    # helm destroy
    helm_destroy_task = Task({}).set_icon('🚢').set_extend('core.helmDestroy').set_location('helm-deployments').set_description('Destroy helm release')
    main_config['tasks']['helmDestroy'] = helm_destroy_task.as_dict()
    # save config
    write_config(main_file_name, main_config)


if __name__ == '__main__':
    try:
        create_helm_task()
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)