from dotenv import dotenv_values
from typing import List, Mapping, Any
from ruamel.yaml import YAML
from io import StringIO
from .config import YamlConfig

import os

class Project(YamlConfig):


    def __init__(self):
        super().__init__({})
        self.replacement_dict: Mapping[str, str] = {}
        self.file_name: str = ''

    
    def _set_default_properties(self):
        pass


    def include(self, file_name: str):
        if self.exist(['includes']) and file_name in self.get(['includes']):
            return
        self.append(['includes'], file_name) 


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
        self.include('${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['tasks', 'run', 'dependencies'], [])
        self.set_default(['tasks', 'runContainer', 'dependencies'], [])
        self.set_default(['tasks', 'stopContainer', 'dependencies'], [])
        self.set_default(['tasks', 'removeContainer', 'dependencies'], [])
        self.set_default(['tasks', 'buildImage', 'dependencies'], [])
        self.set_default(['tasks', 'pushImage', 'dependencies'], [])
        self.set_default(['tasks', 'helmApply', 'extend'], 'core.helmApply')
        self.set_default(['tasks', 'helmApply', 'location'], 'helm-deployments')
        self.set_default(['tasks', 'helmDestroy', 'extend'], 'core.helmDestroy')
        self.set_default(['tasks', 'helmDestroy', 'location'], 'helm-deployments')


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
        self.location = ''
        self.main_project: MainProject = MainProject()


    def _set_default_properties(self):
        service_name = 'zarubaServiceName'
        run_task_name = 'runZarubaServiceName'
        self.include('${ZARUBA_HOME}/scripts/core.zaruba.yaml')
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
        if self.exist(['tasks', run_task_name, 'location']):
            self.location = self.get(['tasks', run_task_name, 'location'])


    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str, location: str):
        file_name = self._get_file_name(dir_name, service_name)
        self.main_project.load(dir_name)
        self._set_service_name(service_name)
        self.location = location
        self.replacement_dict = {
            'zarubaServiceName': self.service_name,
            'ZarubaServiceName': self.capital_service_name,
            'zarubaContainerName': container_name,
            'zarubaImageName': image_name,
            'zarubaTaskLocation': location,
            'ZARUBA_ENV_PREFIX': self.service_name.upper().replace(' ', '_')
        }
        super().generate(file_name)
        self.main_project.include(os.path.join('zaruba-tasks', '{}.zaruba.yaml'.format(service_name)))
        self.main_project.save(dir_name)


class ServiceProject(TaskProject):

    
    def _set_default_properties(self):
        super()._set_default_properties()
        # container related settings
        self.set_default(['configs', 'zarubaServiceNameContainer', 'containerName'], 'zarubaContainerName')
        self.set_default(['configs', 'zarubaServiceNameContainer', 'imageName'], 'zarubaImageName')
        self.set_default(['lconfigs', 'zarubaServiceNameContainer'], {})
        # run
        self.set_default(['tasks', 'runZarubaServiceName', 'extend'], 'core.startService')
        # buildImage
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'extend'], 'core.buildDockerImage')
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'timeout'], '1h')
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'configRef'], 'zarubaServiceNameContainer')
        # pushImage
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'extend'], 'core.pushDockerImage')
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'dependencies'], ['buildZarubaServiceNameImage'])
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'timeout'], '1h')
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'configRef'], 'zarubaServiceNameContainer')
        # runContainer
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'extend'], 'core.startDockerContainer')
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'dependencies'], ['buildZarubaServiceNamelmage'])
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'configRef'], 'zarubaServiceNameContainer')
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'lconfigRef'], 'zarubaServiceNameContainer')
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'envRef'], 'zarubaServiceName')
        # stopContainer
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'extend'], 'core.stopDockerContainer')
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'configRef'], 'zarubaServiceNameContainer')
        # removeContainer
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'extend'], 'core.removeDockerContainer')
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'configRef'], 'zarubaServiceNameContainer')

    
    def _get_env_dict(self, location: str) -> Mapping[str, str]:
        env_dict: Mapping[str, str] = {}
        for env_file in ('sample.env', 'template.env', 'env.template', '.env'):
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


    def load_env(self, location: str, service_name: str, env_prefix: str=''):
        env_dict = self._get_env_dict(location)
        if env_prefix == '':
            env_prefix = service_name.upper().replace(' ', '_')
        for key, val in env_dict.items():
            if self.exist(['envs', service_name, key]):
                continue
            self.set_default(['envs', service_name, key, 'from'], '{}_{}'.format(env_prefix, key))
            self.set_default(['envs', service_name, key, 'default'], val)

    
    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str, location: str, ports: List[str]):
        self.load_env(location, 'zarubaServiceName', 'ZARUBA_ENV_PREFIX')
        if not os.path.isabs(location):
            location = os.path.join('..', location)
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
        super().generate(dir_name, service_name, image_name, container_name, location)
        


class DockerProject(TaskProject):

    def _set_default_properties(self):
        super()._set_default_properties()
        # container related settings
        self.set_default(['configs', 'zarubaServiceName', 'containerName'], 'zarubaContainerName')
        self.set_default(['configs', 'zarubaServiceName', 'imageName'], 'zarubaImageName')
        # stopContainer
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'extend'], 'core.stopDockerContainer')
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'configRef'], 'zarubaServiceName')
        # removeContainer
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'extend'], 'core.removeDockerContainer')
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'configRef'], 'zarubaServiceName')

 
    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str, location: str):
        if service_name == '':
            service_name = container_name
        self._set_service_name(service_name)
        self.main_project.load(dir_name)
        self.main_project.register_run_task('run{}'.format(self.capital_service_name))
        self.main_project.register_run_container_task('run{}'.format(self.capital_service_name))
        self.main_project.register_stop_container_task('stop{}Container'.format(self.capital_service_name))
        self.main_project.register_remove_container_task('remove{}Container'.format(self.capital_service_name))
        self.main_project.save(dir_name)
        super().generate(dir_name, service_name, image_name, container_name, location)
        

       