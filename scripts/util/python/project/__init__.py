from project.structures import Task, Env, ProjectDict, EnvDict, Template
from project.servicegen import ServiceGen, get_service_gen
from project.helpers import \
    write_task, add_to_main_include, add_to_main_task, get_main_file_name, get_task_file_name, create_dir, get_dict_from_file, save_dict_to_file, write_task_env, get_template, get_service_task_template, get_docker_task_template, get_service_template, get_service_name_by_location, get_env_prefix_by_location, get_task_name_by_location