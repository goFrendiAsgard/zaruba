from helper import cli
import helper.generator as generator
import helper.task as task


@cli
def create_helm_task():
    main_file_name = 'main.zaruba.yaml'
    main_config = generator.read_config(main_file_name)
    if 'tasks' not in main_config:
        main_config['tasks'] = {}
    # helm apply
    helm_apply_task = task.Task({}).set_icon('🚢').set_extend('core.helmApply').set_location('helm-deployments').set_description('Deploy helm charts')
    main_config['tasks']['helmApply'] = helm_apply_task.as_dict()
    # helm destroy
    helm_destroy_task = task.Task({}).set_icon('🚢').set_extend('core.helmDestroy').set_location('helm-deployments').set_description('Destroy helm release')
    main_config['tasks']['helmDestroy'] = helm_destroy_task.as_dict()
    # save config
    generator.write_config(main_file_name, main_config)


if __name__ == '__main__':
    create_helm_task()