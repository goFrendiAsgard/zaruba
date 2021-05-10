from dotenv import dotenv_values
from typing import List, Mapping, Any
from ruamel.yaml import YAML
from io import StringIO
from .config import YamlConfig
from .decoration import generate_icon
from .text import get_env_file_names, capitalize, snake, dash

import os
import re
import shutil


class Project(YamlConfig):


    def __init__(self):
        super().__init__()
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

    
    def register_run_task(self, task_name: str):
        self.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['tasks', 'run', 'icon'], generate_icon())
        self.append(['tasks', 'run', 'dependencies'], task_name)

    
    def register_run_container_task(self, task_name: str):
        self.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['tasks', 'runContainer', 'icon'], generate_icon())
        self.append(['tasks', 'runContainer', 'dependencies'], task_name)

    
    def register_stop_container_task(self, task_name: str):
        self.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['tasks', 'stopContainer', 'icon'], generate_icon())
        self.append(['tasks', 'stopContainer', 'dependencies'], task_name)

    
    def register_remove_container_task(self, task_name: str):
        self.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['tasks', 'removeContainer', 'icon'], generate_icon())
        self.append(['tasks', 'removeContainer', 'dependencies'], task_name)

    
    def register_build_image_task(self, task_name: str):
        self.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['tasks', 'buildImage', 'icon'], generate_icon())
        self.append(['tasks', 'buildImage', 'dependencies'], task_name) 

    
    def register_push_image_task(self, task_name: str):
        self.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.set_default(['tasks', 'pushImage', 'icon'], generate_icon())
        self.append(['tasks', 'pushImage', 'dependencies'], task_name) 


    def update_env(self, dir_name: str):
        includes = self.get(['includes']) if self.exist(['includes']) else []
        for include in includes:
            match = re.findall(r'^(\./)?zaruba-tasks/([a-zA-Z0-9_]+).zaruba.yaml$', include)
            if len(match) == 0:
                continue
            service_name = match[0][1]
            capital_service_name = service_name[0].upper() + service_name[1:]
            task_name = 'run{}'.format(capital_service_name)
            service_project = ServiceProject()
            service_project.load(dir_name, service_name)
            if not service_project.exist(['tasks', task_name, 'location']):
                continue
            # service task found
            service_project.load(dir_name, service_name, reload_env=True)
            service_project.save(dir_name, service_name)
            service_project.save_env(dir_name, service_name)
            # update helm service if file exist
            helm_service_project = HelmServiceProject()
            if not helm_service_project.file_exist(dir_name, service_name):
                continue
            helm_service_project.load(dir_name, service_name, reload_env=True)
            helm_service_project.save(dir_name, service_name)


class TaskProject(Project):

    def __init__(self):
        super().__init__()
        self.main_project: MainProject = MainProject()


    def _set_default_properties(self):
        service_name = 'zarubaServiceName'
        run_task_name = 'runZarubaServiceName'
        self.set_default(['envs', service_name], {})
        self.set_default(['configs', service_name], {})
        self.set_default(['lconfigs', service_name], {})
        self.set_default(['tasks', run_task_name, 'configRef'], service_name)
        self.set_default(['tasks', run_task_name, 'envRef'], service_name)
        self.set_default(['tasks', run_task_name, 'lconfRef'], service_name)
    

    def _get_file_name(self, dir_name: str, service_name: str):
        return os.path.join(dir_name, 'zaruba-tasks', '{}.zaruba.yaml'.format(service_name))
    

    def load_from_template(self, template_file_name: str):
        super().load(template_file_name)


    def load(self, dir_name: str, service_name: str):
        file_name = self._get_file_name(dir_name, service_name)
        self.main_project.load(dir_name)
        super().load(file_name)


    def save(self, dir_name: str, service_name: str):
        file_name = self._get_file_name(dir_name, service_name)
        super().save(file_name)


    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str, location: str, start_command: str, runner_version: str):
        file_name = self._get_file_name(dir_name, service_name)
        self.main_project.load(dir_name)
        self.replacement_dict = {
            'zarubaServiceName': service_name,
            'zarubaservicename': service_name.lower(),
            'ZarubaServiceName': capitalize(service_name),
            'ZARUBA_SERVICE_NAME': snake(service_name).upper(),
            'zarubaContainerName': container_name,
            'zarubaImageName': image_name,
            'zarubaServiceLocation': location,
            'zarubaStartCommand': start_command,
            'zarubaRunnerVersion': runner_version,
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
        self.set_default(['configs', 'zarubaServiceNameContainer', 'expose'], 'lconfig.ports')
        self.set_default(['configs', 'zarubaServiceNameContainer', 'localhost'], 'host.docker.internal')
        # run
        self.set_default(['tasks', 'runZarubaServiceName', 'icon'], generate_icon())
        self.set_default(['tasks', 'runZarubaServiceName', 'extend'], 'core.startService')
        self.set_default(['tasks', 'runZarubaServiceName', 'location'], 'zarubaServiceLocation')
        # runContainer
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'extend'], 'core.startDockerContainer')
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'dependencies'], ['buildZarubaServiceNameImage'])
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'configRef'], 'zarubaServiceNameContainer')
        self.set_default(['tasks', 'runZarubaServiceNameContainer', 'lconfigRef'], 'zarubaServiceName')
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
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'location'], 'zarubaServiceLocation')
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'inputs'], ['docker.env'])
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'timeout'], '1h')
        self.set_default(['tasks', 'buildZarubaServiceNameImage', 'configRef'], 'zarubaServiceNameContainer')
        # pushImage
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'icon'], generate_icon())
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'extend'], 'core.pushDockerImage')
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'dependencies'], ['buildZarubaServiceNameImage'])
        self.set_default(['tasks', 'pushZarubaServiceNameImage', 'inputs'], ['docker.env'])
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
            env_val = str(self.get(['envs', service_name, env_key, 'default']))
            if not env_val.isnumeric():
                continue
            env_val_int = int(env_val)
            if env_val_int in [21, 22, 23, 25, 53, 80, 443] or (env_val_int >=3000 and env_val_int not in [3306, 5432, 5672, 15672, 27017, 6379, 9200, 9300, 7001, 7199, 9042, 9160]):
                ports.append('{{ .GetEnv "' + env_key +'" }}')
        return ports


    def _load_env(self, service_name: str, service_location: str, env_prefix: str):
        env_dict = self._get_env_dict(service_location)
        for key, val in env_dict.items():
            if self.exist(['envs', service_name, key]):
                continue
            env_key = key if key.startswith(env_prefix + '_') else '{}_{}'.format(env_prefix, key)
            self.set_default(['envs', service_name, key, 'from'], env_key)
            self.set_default(['envs', service_name, key, 'default'], val)


    def load(self, dir_name: str, service_name: str, reload_env: bool=False):
        super().load(dir_name, service_name)
        task_name = 'run{}'.format(capitalize(service_name))
        if not self.exist(['tasks', task_name, 'location']):
            return
        service_location = self.get(['tasks', task_name, 'location'])
        if not os.path.isabs(service_location):
            service_location = os.path.abspath(os.path.join(dir_name, 'zaruba_tasks', service_location))
        if reload_env:
            self._load_env(service_name, service_location, snake(service_name).upper())
    

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
                if 'from' not in env:
                    continue
                env_from = env.get('from')
                if env_from == '' or env_from in existing_envvars:
                    continue
                if is_first_writing:
                    is_first_writing = False
                    f_write.write('\n# {}\n'.format(service_name))
                env_value = env.get('default', '')
                f_write.write('{}={}\n'.format(env_from, env_value))
            f_write.close()

    
    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str, location: str, start_command: str, ports: List[str], runner_version: str):
        #self._load_env('zarubaServiceName', service_location, env_prefix='ZARUBA_SERVICE_NAME')
        service_location = location
        if not os.path.isabs(location):
            location = os.path.relpath(os.path.abspath(location), os.path.abspath(os.path.join(dir_name, 'zaruba-tasks')))
        if service_name == '':
            service_name = os.path.basename(os.path.abspath(service_location))
            service_name = re.sub(r'[^a-zA-Z0-9]', '', service_name)
        if container_name == '':
            container_name = service_name
        if image_name == '':
            image_name = container_name
        image_name = dash(image_name)
        self._load_env('zarubaServiceName', service_location, snake(service_name).upper())
        # handle ports
        if len(ports) == 0:
            ports = self._get_possible_ports_env('zarubaServiceName')
        self.set(['lconfigs', 'zarubaServiceName', 'ports'], ports)
        capital_service_name = capitalize(service_name)
        self.main_project.load(dir_name)
        self.main_project.register_run_task('run{}'.format(capital_service_name))
        self.main_project.register_run_container_task('run{}Container'.format(capital_service_name))
        self.main_project.register_stop_container_task('stop{}Container'.format(capital_service_name))
        self.main_project.register_remove_container_task('remove{}Container'.format(capital_service_name))
        self.main_project.register_build_image_task('build{}Image'.format(capital_service_name))
        self.main_project.register_push_image_task('push{}Image'.format(capital_service_name))
        self.main_project.save(dir_name)
        super().generate(dir_name, service_name, image_name, container_name, location, start_command, runner_version=runner_version)
        

class DockerProject(TaskProject):

    def _set_default_properties(self):
        super()._set_default_properties()
        # container related settings
        self.set_default(['configs', 'zarubaServiceName', 'useImagePrefix'], False)
        self.set_default(['configs', 'zarubaServiceName', 'imageName'], 'zarubaImageName')
        self.set_default(['configs', 'zarubaServiceName', 'containerName'], 'zarubaContainerName')
        self.set_default(['configs', 'zarubaServiceName', 'expose'], 'config.port')
        # pull
        self.set_default(['tasks', 'pullZarubaServiceNameImage', 'icon'], generate_icon())
        self.set_default(['tasks', 'pullZarubaServiceNameImage', 'extend'], 'core.pullDockerImage')
        self.set_default(['tasks', 'pullZarubaServiceNameImage', 'configRef'], 'zarubaServiceName')
        # run
        self.set_default(['tasks', 'runZarubaServiceName', 'icon'], generate_icon())
        self.set_default(['tasks', 'runZarubaServiceName', 'extend'], 'core.startDockerContainer')
        self.append_if_not_exist(['tasks', 'runZarubaServiceName', 'dependencies'], 'pullZarubaServiceNameImage')
        # stopContainer
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'extend'], 'core.stopDockerContainer')
        self.set_default(['tasks', 'stopZarubaServiceNameContainer', 'configRef'], 'zarubaServiceName')
        # removeContainer
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'icon'], generate_icon())
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'extend'], 'core.removeDockerContainer')
        self.set_default(['tasks', 'removeZarubaServiceNameContainer', 'configRef'], 'zarubaServiceName')

 
    def generate(self, dir_name: str, service_name: str, image_name: str, container_name: str):
        if container_name == '':
            container_name = image_name
        if service_name == '':
            service_name = container_name
        image_name = dash(image_name)
        capital_service_name = capitalize(service_name)
        self.main_project.load(dir_name)
        self.main_project.register_run_task('run{}'.format(capital_service_name))
        self.main_project.register_run_container_task('run{}'.format(capital_service_name))
        self.main_project.register_stop_container_task('stop{}Container'.format(capital_service_name))
        self.main_project.register_remove_container_task('remove{}Container'.format(capital_service_name))
        self.main_project.save(dir_name)
        super().generate(dir_name, service_name, image_name, container_name, location='', start_command='', runner_version='')
        

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
        script_path = os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
        helm_template_dir = os.path.join(script_path, 'templates', 'helmDeployments')
        shutil.copytree(helm_template_dir, os.path.join(dir_name, 'helm-deployments'))
        # copy helm deployments
        self.main_project.load(dir_name)
        self.main_project.append_if_not_exist(['includes'], '${ZARUBA_HOME}/scripts/core.zaruba.yaml')
        self.main_project.set_default(['tasks', 'helmApply', 'icon'], generate_icon())
        self.main_project.set_default(['tasks', 'helmApply', 'extend'], 'core.helmApply')
        self.main_project.set_default(['tasks', 'helmApply', 'location'], 'helm-deployments')
        self.main_project.set_default(['tasks', 'helmApply', 'inputs'], ['helm.env', 'kube.context', 'docker.env'])
        self.main_project.set_default(['tasks', 'helmDestroy', 'icon'], generate_icon())
        self.main_project.set_default(['tasks', 'helmDestroy', 'extend'], 'core.helmDestroy')
        self.main_project.set_default(['tasks', 'helmDestroy', 'location'], 'helm-deployments')
        self.main_project.set_default(['tasks', 'helmDestroy', 'inputs'], ['helm.env', 'kube.context', 'docker.env'])
        self.main_project.save(dir_name)
        self.load(dir_name)
    
 
    def register_release(self, service_name: str):
        if not self.exist(['releases']):
            self.set(['releases'], [])
        releases: List[Mapping[str, str]] = self.get(['releases'])
        # check whethere service already registered or not
        registered = False
        for release in releases:
            if release['name'] == dash(service_name):
                registered = True
                break
        # do nothing if service already registered
        if registered:
            return
        # register new release
        self.append(['releases'], {
            'name': dash(service_name),
            'chart': './charts/app',
            'values': ['./values/{}.yaml.gotmpl'.format(dash(service_name))],
        })


class HelmServiceProject(Project):

    def __init__(self):
        super().__init__()
        self.service_project: ServiceProject = ServiceProject()
        self.helm_project: HelmProject = HelmProject()


    def _get_file_name(self, dir_name: str, service_name: str) -> str:
        return os.path.join(dir_name, 'helm-deployments', 'values', '{}.yaml.gotmpl'.format(dash(service_name)))


    def _set_default_properties(self):
        super()._set_default_properties()
        self.set_default(['app', 'name'], 'zaruba-service-name')
        self.set_default(['app', 'group'], 'db')
        self.set_default(['app', 'container', 'imagePrefix'], '{{ .Values | get "commonImagePrefix" "local" }}')
        self.set_default(['app', 'container', 'imageTag'], '{{ .Values | get "commonImageTag" "latest" }}')
        self.set_default(['app', 'container', 'image'], 'zarubaImageName')
        self.set_default(['app', 'container', 'env'], [])
        self.set_default(['app', 'ports'], [])
        

    def _load_env(self, service_name: str):
        self.set_default(['app', 'container', 'env'], [])
        # get env from service
        service_envs: Mapping[str, Mapping[str, str]] = self.service_project.get(['envs', service_name]) if self.service_project.exist(['envs', service_name]) else {}
        # get current container env
        current_container_envs: List[Mapping[str, str]] = self.get(['app', 'container', 'env'])
        for env_key, env_map in service_envs.items():
            # check whether env is already registered or not
            env_registered = False
            for current_container_env in current_container_envs:
                current_container_env_name = current_container_env.get('name')
                if current_container_env_name == env_key:
                    env_registered = True
                    break
            # if env is already registered, do nothing
            if env_registered:
                continue
            # register
            self.append(['app', 'container', 'env'], {
                'name': env_key,
                'value': env_map.get('default', '')
            })
    

    def _append_port(self, port_str: str):
        self.set_default(['app', 'ports'], [])
        port: int = 0
        try:
            port = int(port_str)
        except:
            port = 80
        self.append(['app', 'ports'], {
            'containerPort': port,
            'servicePort': port,
        })
    

    def _assign_ports(self, service_name: str):
        # get port from service
        service_ports: List[str] = self.service_project.get(['lconfigs', service_name, 'ports']) if self.service_project.exist(['lconfigs', service_name, 'ports']) else []
        # get env from service
        service_envs: Mapping[str, Mapping[str, str]] = self.service_project.get(['envs', service_name]) if self.service_project.exist(['envs', service_name]) else {}
        # service port
        for port in service_ports:
            if port.isnumeric():
                self._append_port(port)
                continue
            # port is not numeric, probably go template
            match = re.findall(r'^\{\{ \.GetEnv "(.*)" \}\}$', port)
            if len(match) < 1:
                continue
            # env key found, assign default value as port
            env_key = match[0]
            if 'default' in service_envs[env_key]:
                self._append_port(service_envs[env_key]['default'])


    def file_exist(self, dir_name: str, service_name: str) -> bool:
        file_name = self._get_file_name(dir_name, service_name)
        return os.path.isfile(file_name)


    def load(self, dir_name: str, service_name: str, reload_env: bool=True):
        file_name = self._get_file_name(dir_name, service_name)
        super().load(file_name)
        self.helm_project.load(dir_name)
        self.service_project.load(dir_name, service_name)
        if reload_env:
            self._load_env(service_name)


    def save(self, dir_name: str, service_name: str):
        file_name = self._get_file_name(dir_name, service_name)
        super().save(file_name)


    def generate(self, dir_name: str, service_name: str):
        self.helm_project.load(dir_name)
        self.service_project.load(dir_name, service_name)
        self._set_default_properties()
        self._load_env(service_name)
        self._assign_ports(service_name)
        service_image_name_key = ['configs', '{}Container'.format(service_name), 'imageName']
        image_name = self.service_project.get(service_image_name_key) if self.service_project.exist(service_image_name_key) else service_name
        self.replacement_dict = {
            'zarubaServiceName': service_name,
            'zaruba-service-name': dash(service_name),
            'zarubaImageName': image_name,
        }
        # register service to helm deployment
        self.helm_project.register_release(service_name)
        self.helm_project.save(dir_name)
        # generate deployment
        file_name = self._get_file_name(dir_name, service_name)
        super().generate(file_name)
        self.load(dir_name, service_name)
