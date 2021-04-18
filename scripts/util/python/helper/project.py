from dotenv import dotenv_values
from typing import List, Mapping, Any
from ruamel.yaml import YAML
from io import StringIO
from .config import YamlConfig
from .decoration import generate_icon

import os
import shutil


def get_env_file_names(location: str) -> List[str]:
    abs_location = os.path.abspath(location)
    env_file_names = [os.path.join(abs_location, f) for f in os.listdir(abs_location) if os.path.isfile(os.path.join(abs_location, f)) and (f.endswith('.env') or f.endswith('env.template'))]
    env_file_names.sort(key=lambda s: len(s))
    return env_file_names


class Project(YamlConfig):


    def __init__(self):
        super().__init__({})
        self.replacement_dict: Mapping[str, str] = {}
        self.file_name: str = ''

    
    def _set_default_properties(self):
        pass


    def generate(self, file_name: str):
        self._set_default_properties()
        # create yaml_text
        string_stream = StringIO()
        yaml=YAML()
        yaml.dump(self.data, string_stream)
        yaml_text = self._pretify_yaml(string_stream.getvalue())
        string_stream.close()
        # perform replacement
        for key, val in self.replacement_dict.items():
            yaml_text = yaml_text.replace(key, val)
        # write yaml
        abs_file_name = os.path.abspath(file_name)
        self._create_dir_if_not_exist(file_name)
        f_write = open(abs_file_name, 'w')
        f_write.write(yaml_text)
        f_write.close()


class MainProject(Project):

    def __init__(self):
        super().__init__()


    def _set_default_properties(self):
        self.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['tasks', 'run', 'icon'], generate_icon())
        self.set_default(['tasks', 'run', 'dependencies'], [])
        self.set_default(['tasks', 'runContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'runContainer', 'dependencies'], [])
        self.set_default(['tasks', 'stopContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'stopContainer', 'dependencies'], [])
        self.set_default(['tasks', 'removeContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'removeContainer', 'dependencies'], [])
        self.set_default(['tasks', 'buildImage', 'icon'], generate_icon())
        self.set_default(['tasks', 'buildImage', 'dependencies'], [])
        self.set_default(['tasks', 'pushImage', 'icon'], generate_icon())
        self.set_default(['tasks', 'pushImage', 'dependencies'], [])
        


    def _get_file_name(self, dir_name: str) -> str:
        return os.path.join(dir_name, 'main.zaruba.yaml')


    def load(self, dir_name: str):
        super().load(self._get_file_name(dir_name))


    def save(self, dir_name: str):
        super().save(self._get_file_name(dir_name))


    def generate(self, dir_name: str):
        super().generate(self._get_file_name(dir_name))

    
    def register_run_task(self, task_name):
        self.append(['tasks', 'run', 'dependencies'], task_name)

    
    def register_run_container_task(self, task_name):
        self.append(['tasks', 'runContainer', 'dependencies'], task_name)

    
    def register_stop_container_task(self, task_name):
        self.append(['tasks', 'stopContainer', 'dependencies'], task_name)

    
    def register_remove_container_task(self, task_name):
        self.append(['tasks', 'removeContainer', 'dependencies'], task_name)

    
    def register_build_image_task(self, task_name):
        self.append(['tasks', 'buildImage', 'dependencies'], task_name) 

    
    def register_push_image_task(self, task_name):
        self.append(['tasks', 'pushImage', 'dependencies'], task_name) 


class TaskProject(Project):

    def __init__(self):
        super().__init__()
        self.service_name = ''
        self.capital_service_name = ''
        self.main_project: MainProject = MainProject()


    def _set_default_properties(self):
        service_name = 'zarubaServiceName'
        run_task_name = 'runZarubaServiceName'
        self.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['envs', service_name], {})
        self.set_default(['configs', service_name], {})
        self.set_default(['lconfigs', service_name], {})
        self.set_default(['tasks', run_task_name, 'configRef'], service_name)
        self.set_default(['tasks', run_task_name, 'envRef'], service_name)
        self.set_default(['tasks', run_task_name, 'lconfRef'], service_name)
    

    def _set_service_name(self, service_name):
        self.service_name = service_name
        self.capital_service_name = self.service_name[0].upper() + self.service_name[1:]
    

    def _get_file_name(self, dir_name: str, service_name: str):
        return os.path.join(dir_name, 'zaruba-tasks', '{}.zaruba.yaml'.format(service_name))
    

    def load_from_template(self, template_file_name: str):
        super().load(template_file_name)


    def load(self, dir_name: str, service_name: str):
        file_name = self._get_file_name(dir_name, service_name)
        self.main_project.load(dir_name)
        super().load(file_name)
        run_task_name = 'run{}'.format(self.service_name)


    def save(self, dir_name: str, service_name: str):
        file_name = self._get_file_name(dir_name, service_name)
        super().save(self._get_file_name(file_name))


    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str, location: str, start_command: str):
        file_name = self._get_file_name(dir_name, service_name)
        self.main_project.load(dir_name)
        self._set_service_name(service_name)
        self.replacement_dict = {
            'zarubaServiceName': self.service_name,
            'ZarubaServiceName': self.capital_service_name,
            'ZARUBA_SERVICE_NAME': self.service_name.upper().replace(' ', '_'),
            'zarubaContainerName': container_name,
            'zarubaImageName': image_name,
            'zarubaServiceLocation': location,
            'zarubaStartCommand': start_command,
        }
        super().generate(file_name)
        self.main_project.append_if_not_exist(['includes'], os.path.join('zaruba-tasks', '{}.zaruba.yaml'.format(service_name)))
        self.main_project.save(dir_name)
        self.load(dir_name, service_name)


class ServiceProject(TaskProject):

    
    def _set_default_properties(self):
        super()._set_default_properties()
        # container related settings
        self.set_default(['configs', 'zarubaServiceNameContainer', 'containerName'], 'zarubaContainerName')
        self.set_default(['configs', 'zarubaServiceNameContainer', 'imageName'], 'zarubaImageName')
        self.set_default(['lconfigs', 'zarubaServiceNameContainer'], {})
        # run
        self.set_default(['tasks', 'runZarubaServiceName', 'icon'], generate_icon())
        self.set_default(['tasks', 'runZarubaServiceName', 'extend'], 'core.startService')
        # runContainer
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'extend'], 'core.startDockerContainer')
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'dependencies'], ['buildZarubaServiceNameImage'])
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'configRef'], 'zarubaServiceNameContainer')
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'lconfigRef'], 'zarubaServiceNameContainer')
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'envRef'], 'zarubaServiceName')
        # stopContainer
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'extend'], 'core.stopDockerContainer')
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'configRef'], 'zarubaServiceNameContainer')
        # removeContainer
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'extend'], 'core.removeDockerContainer')
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'configRef'], 'zarubaServiceNameContainer')
        # buildImage
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'icon'], generate_icon())
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'extend'], 'core.buildDockerImage')
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'timeout'], '1h')
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'configRef'], 'zarubaServiceNameContainer')
        # pushImage
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'icon'], generate_icon())
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'extend'], 'core.pushDockerImage')
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'dependencies'], ['buildZarubaServiceNameImage'])
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'timeout'], '1h')
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'configRef'], 'zarubaServiceNameContainer')
        
    
    def _get_env_dict(self, location: str) -> Mapping[str, str]:
        env_dict: Mapping[str, str] = {}
        for env_file in get_env_file_names(location):
            env_path = os.path.abspath(os.path.join(location, env_file))
            if not os.path.isfile(env_path):
                continue
            local_env: Mapping[str, str] = dotenv_values(env_path)
            for env_key, env_value in local_env.items():
                env_dict[env_key] = env_value
        return env_dict
    

    def _get_possible_ports_env(self, service_name: str) -> List[str]:
        if not self.exist(['envs', service_name]):
            return []
        ports = []
        for env_key in self.get(['envs', service_name]):
            if not self.exist(['envs', service_name, env_key, 'default']):
                continue
            env_val = self.get(['envs', service_name, env_key, 'default'])
            if not env_val.isnumeric():
                continue
            env_val_int = int(env_val)
            if env_val_int in [21, 22, 23, 25, 53, 80, 443] or (env_val_int >=3000 and env_val_int not in [3306, 5432, 5672, 15672, 27017, 6379, 9200, 9300, 7001, 7199, 9042, 9160]):
                ports.append('{{ .GetEnv "' + env_key +'" }}')
        return ports


    def _load_env(self, service_name: str, location: str='', env_prefix: str=''):
        env_dict = self._get_env_dict(location)
        if env_prefix == '':
            env_prefix = service_name.upper().replace(' ', '_')
        for key, val in env_dict.items():
            if self.exist(['envs', service_name, key]):
                continue
            self.set_default(['envs', service_name, key, 'from'], '{}_{}'.format(env_prefix, key))
            self.set_default(['envs', service_name, key, 'default'], val)


    def load(self, dir_name: str, service_name: str):
        super().load(dir_name, service_name)
        self._set_service_name(service_name)
        task_name = 'run{}'.format(self.capital_service_name)
        if not self.exist(['tasks', task_name, 'location']):
            return
        task_location = self.get(['tasks', task_name, 'location'])
        if not os.path.isabs(task_location):
            task_location = os.path.abspath(os.path.join(dir_name, 'zaruba_tasks', task_location))
        self._load_env(service_name, task_location)
    

    def save_env(self, dir_name: str, service_name: str):
        env_file_names = get_env_file_names(dir_name)
        default_env_file_name = os.path.join(dir_name, 'template.env')
        if default_env_file_name not in env_file_names:
            env_file_names.append(default_env_file_name)
        for file_name in env_file_names:
            existing_envvars: Mapping[str, str] = dotenv_values(file_name) if os.path.isfile(file_name) else {}
            is_first_writing = True
            f_write = open(file_name, 'a')
            env_map: Mapping[str, Mapping[str, str]] = self.get(['envs', service_name]) if self.exist(['envs', service_name]) else {}
            for env_key, env in env_map.items():
                env_from = env.get('from')
                if env_from == '' or env_from in existing_envvars:
                    continue
                if is_first_writing:
                    is_first_writing = False
                    f_write.write('\n# {}\n'.format(service_name))
                env_value = env.get('default', '')
                f_write.write('{}={}\n'.format(env_from, env_value))
            f_write.close()

    
    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str, location: str, start_command: str, ports: List[str]):
        self._load_env('zarubaServiceName', location=location, env_prefix='ZARUBA_SERVICE_NAME')
        if not os.path.isabs(location):
            location = os.path.relpath(os.path.abspath(location), os.path.abspath(os.path.join(dir_name, 'zaruba-tasks')))
        if container_name == '':
            container_name = service_name
        if image_name == '':
            image_name = container_name
        image_name = image_name.lower()
        # handle ports
        if len(ports) == 0:
            ports = self._get_possible_ports_env('zarubaServiceName')
        self.set(['lconfigs', 'zarubaServiceName', 'ports'], ports)
        self._set_service_name(service_name)
        self.main_project.load(dir_name)
        self.main_project.register_run_task('run{}'.format(self.capital_service_name))
        self.main_project.register_run_container_task('run{}Container'.format(self.capital_service_name))
        self.main_project.register_stop_container_task('stop{}Container'.format(self.capital_service_name))
        self.main_project.register_remove_container_task('remove{}Container'.format(self.capital_service_name))
        self.main_project.register_build_image_task('build{}Image'.format(self.capital_service_name))
        self.main_project.register_push_image_task('push{}Image'.format(self.capital_service_name))
        self.main_project.save(dir_name)
        super().generate(dir_name, service_name, image_name, container_name, location, start_command)
        

class DockerProject(TaskProject):

    def _set_default_properties(self):
        super()._set_default_properties()
        # container related settings
        self.set_default(['configs', 'zarubaServiceName', 'useImagePrefix'], False)
        self.set_default(['configs', 'zarubaServiceName', 'imageName'], 'zarubaImageName')
        self.set_default(['configs', 'zarubaServiceName', 'containerName'], 'zarubaContainerName')
        # run
        self.set_default(['tasks', 'runZarubaServiceName', 'icon'], generate_icon())
        self.set_default(['tasks', 'runZarubaServiceName', 'extend'], 'core.startDockerService')
        # stopContainer
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'extend'], 'core.stopDockerContainer')
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'configRef'], 'zarubaServiceName')
        # removeContainer
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'extend'], 'core.removeDockerContainer')
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'configRef'], 'zarubaServiceName')

 
    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str, location: str, start_command: str):
        if service_name == '':
            service_name = container_name
        image_name = image_name.lower()
        self._set_service_name(service_name)
        self.main_project.load(dir_name)
        self.main_project.register_run_task('run{}'.format(self.capital_service_name))
        self.main_project.register_run_container_task('run{}'.format(self.capital_service_name))
        self.main_project.register_stop_container_task('stop{}Container'.format(self.capital_service_name))
        self.main_project.register_remove_container_task('remove{}Container'.format(self.capital_service_name))
        self.main_project.save(dir_name)
        super().generate(dir_name, service_name, image_name, container_name, location, start_command)
        

class HelmProject(Project):

    def __init__(self):
        super().__init__()
        self.main_project: MainProject = MainProject()

    def _get_file_name(self, dir_name: str) -> str:
        return os.path.join(dir_name, 'helm-deployments', 'helmfile.yaml')


    def load(self, dir_name: str):
        super().load(self._get_file_name(dir_name))
        self.main_project.load(dir_name)


    def save(self, dir_name: str):
        super().save(self._get_file_name(dir_name))


    def generate(self, dir_name: str):
        script_path = os.path.dirname(os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__)))))
        helm_template_dir = os.path.join(script_path, 'templates', 'helmDeployments')
        shutil.copytree(helm_template_dir, os.path.join(dir_name, 'helm-deployments'))
        # copy helm deployments
        self.main_project.load(dir_name)
        self.main_project.set_default(['tasks', 'helmApply', 'icon'], generate_icon())
        self.main_project.set_default(['tasks', 'helmApply', 'extend'], 'core.helmApply')
        self.main_project.set_default(['tasks', 'helmApply', 'location'], 'helm-deployments')
        self.main_project.set_default(['tasks', 'helmDestroy', 'icon'], generate_icon())
        self.main_project.set_default(['tasks', 'helmDestroy', 'extend'], 'core.helmDestroy')
        self.main_project.set_default(['tasks', 'helmDestroy', 'location'], 'helm-deployments')
        self.main_project.save(dir_name)
        self.load(dir_name)


class HelmServiceProject(Project):

    def __init__(self):
        super().__init__()
        self.service_name = ''
        self.capital_service_name = ''
        self.service_project: ServiceProject = ServiceProject()
        self.helm_project: HelmProject = HelmProject()


    def _get_file_name(self, dir_name: str, service_name: str) -> str:
        return os.path.join(dir_name, 'helm-deployments', 'values', '{}.yaml.gotmpl'.format(service_name))


    def _set_service_name(self, service_name):
        self.service_name = service_name
        self.capital_service_name = self.service_name[0].upper() + self.service_name[1:]
    

    def load_env(self):
        pass


    def load(self, dir_name: str, service_name: str):
        super().load(self._get_file_name(dir_name))
        self.helm_project.load(dir_name)
        self.service_project.load(dir_name, service_name)


    def save(self, dir_name: str, service_name: str):
        super().save(self._get_file_name(dir_name))


    def generate(self, dir_name: str, service_name: str):
        self.helm_project.load(dir_name)
        self.service_project.load(dir_name, service_name)
        self.load_env()
        self._set_service_name(service_name)
        service_container_config: Mapping[str, str] = self.service_project.get(['configs', service_name]) if self.service_project.exist(['configs', service_name]) else {}
        self.replacement_dict = {
            'zarubaServiceName': self.service_name,
            'ZarubaServiceName': self.capital_service_name,
            'ZARUBA_SERVICE_NAME': self.service_name.upper().replace(' ', '_'),
            'zarubaContainerName': service_container_config.get('containerName', service_name),
            'zarubaImageName': service_container_config.get('imageName', service_name),
        }
        file_name = self._get_file_name(dir_name, service_name)
        super().generate(file_name)
