from typing import List, Mapping
from .text import get_env_file_names, capitalize, snake, add_python_indentation

import os


class CodeGen():

    def __init__(self):
        self.file_content_dict: Mapping[str, str] = {}
        self.replacement_dict: Mapping[str, str] = {}


    def _create_dir_if_not_exist(self, file_name: str):
        abs_file_name = os.path.abspath(file_name)
        abs_dir = os.path.dirname(abs_file_name)
        if not os.path.exists(abs_dir):
            os.makedirs(abs_dir)
    

    def read_file_from_disk(self, file_name: str) -> str:
        f_read = open(file_name, 'r')
        content = f_read.read()
        f_read.close()
        return content
    

    def write_file_to_disk(self, file_name: str, content: str):
        f_write = open(file_name, 'w')
        f_write.write(content)
        f_write.close()

    
    def get_content(self, file_name: str) -> str:
        if file_name not in self.file_content_dict:
            raise ValueError('{} is not exist in file_content_dict'.format(file_name))
        content = self.file_content_dict[file_name]
        return content

    
    def set_content(self, file_name: str, content: str):
        self.file_content_dict[file_name] = content
    

    def replace_text(self, text: str, replacement_dict: Mapping[str, str]) -> str:
        new_text = text
        for key, val in replacement_dict.items():
            new_text = new_text.replace(key, val)
        return new_text


    def load_from_template(self, template_dir_name: str):
        self._load(template_dir_name)


    def load(self, dir_name: str):
        return self._load(dir_name)
    

    def _load(self, dir_name: str):
        for file_name in self.file_content_dict:
            abs_file_name = os.path.abspath(os.path.join(dir_name, file_name))
            content = self.read_file_from_disk(abs_file_name)
            self.file_content_dict[file_name] = content


    def save(self, dir_name: str):
        for file_name, content in self.file_content_dict.items():
            abs_file_name = os.path.abspath(os.path.join(dir_name, file_name))
            self._create_dir_if_not_exist(abs_file_name)
            self.write_file_to_disk(abs_file_name, content)


    def generate(self, dir_name: str):
        new_file_content_dict: Mapping[str, str] = {}
        for file_name, content in self.file_content_dict.items():
            for key, val in self.replacement_dict.items():
                content = content.replace(key, val)
                file_name = file_name.replace(key, val)
            new_file_content_dict[file_name] = content
        self.file_content_dict = new_file_content_dict
        self.save(dir_name)


class FastApiService(CodeGen):

    def __init__(self, service_name: str):
        super().__init__()
        self.service_name = service_name

    def load(self, dir_name: str):
        self.file_content_dict = {
            '{}/main.py'.format(self.service_name): '',
        }
        super().load(dir_name)
    

    def load_from_template(self, template_dir_name: str):
        self.file_content_dict = {
            'zarubaServiceName/Dockerfile': '',
            'zarubaServiceName/main.py': '',
            'zarubaServiceName/Pipfile': '',
            'zarubaServiceName/Pipfile.lock': '',
            'zarubaServiceName/start.sh': '',
            'zarubaServiceName/template.env': '',
            'zarubaServiceName/helpers/__init__.py': '',
            'zarubaServiceName/helpers/transport/__init__.py': '',
            'zarubaServiceName/helpers/transport/interface.py': '',
            'zarubaServiceName/helpers/transport/local.py': '',
            'zarubaServiceName/helpers/transport/rmq.py': '',
            'zarubaServiceName/repos/__init__.py': '',
            'zarubaServiceName/schemas/__init__.py': '',
        }
        super().load_from_template(template_dir_name)

    
    def generate(self, dir_name: str):
        self.replacement_dict = {
            'zarubaServiceName': self.service_name,
            'ZARUBA_SERVICE_NAME': snake(self.service_name).upper(),
        }
        super().generate(dir_name)


class FastApiModule(CodeGen):

    def __init__(self, service_name: str, module_name: str):
        super().__init__()
        self.service_name = service_name
        self.module_name = module_name
        self.import_module_partial = ''
        self.load_module_partial = ''
        self.handle_rpc_partial = ''
        self.handle_event_parial = ''
        self.handle_route = ''


    def load(self, dir_name: str):
        self.file_content_dict = {
            '{}/{}/controller.py'.format(self.service_name, self.module_name): '',
        }
        super().load(dir_name)


    def load_from_template(self, template_dir_name: str):
        self.file_content_dict = {
            'zarubaServiceName/zarubaModuleName/__init__.py': '',
            'zarubaServiceName/zarubaModuleName/controller.py': '',
        }
        super().load_from_template(template_dir_name)
        partial_path = os.path.join(os.path.abspath(template_dir_name), 'partials')
        self.import_module_partial = self.read_file_from_disk(os.path.join(partial_path, 'import_module.py'))
        self.load_module_partial = self.read_file_from_disk(os.path.join(partial_path, 'load_module.py'))
        self.handle_rpc_partial = self.read_file_from_disk(os.path.join(partial_path, 'handle_rpc.py'))
        self.handle_event_partial = self.read_file_from_disk(os.path.join(partial_path, 'handle_event.py'))
        self.handle_route_partial = self.read_file_from_disk(os.path.join(partial_path, 'handle_route.py'))
    

    def add_route(self, dir_name: str, http_method: str, url: str):
        handle_route_script = self.replace_text(
            add_python_indentation(self.handle_route_partial, 2), 
            {
                'zarubaHttpMethod': http_method,
                'zarubaUrl': url,
                'zaruba_url': snake(url.replace('/', '_').replace('-', '_')).strip('_'),
            }
        )
        handle_route_script = '\n{}\n'.format(handle_route_script)
        controller_file_name = self._get_controller_file_name(dir_name)
        controller_file_content = self.read_file_from_disk(controller_file_name)
        controller_script_lines = controller_file_content.split('\n')
        insert_index = -1
        for line_index, line in enumerate(controller_script_lines):
            if line.startswith(add_python_indentation('def handle_route', 1)):
                insert_index = line_index + 1
                break
        if insert_index == -1:
            raise ValueError('Cannot find handle_route method in {}'.format(controller_file_name))
        controller_script_lines.insert(insert_index, handle_route_script)
        controller_file_content = '\n'.join(controller_script_lines)
        self.write_file_to_disk(controller_file_name, controller_file_content)


    def add_event_handler(self, dir_name: str, event_name: str):
        handle_event_script = self.replace_text(
            add_python_indentation(self.handle_event_partial, 2), 
            {
                'zarubaEventName': event_name,
                'zaruba_event_name': snake(event_name),
            }
        )
        handle_event_script = '\n{}\n'.format(handle_event_script)
        controller_file_name = self._get_controller_file_name(dir_name)
        controller_file_content = self.read_file_from_disk(controller_file_name)
        controller_script_lines = controller_file_content.split('\n')
        insert_index = -1
        for line_index, line in enumerate(controller_script_lines):
            if line.startswith(add_python_indentation('def handle_event', 1)):
                insert_index = line_index + 1
                break
        if insert_index == -1:
            raise ValueError('Cannot find handle_event method in {}'.format(controller_file_name))
        controller_script_lines.insert(insert_index, handle_event_script)
        controller_file_content = '\n'.join(controller_script_lines)
        self.write_file_to_disk(controller_file_name, controller_file_content)


    def add_rpc_handler(self, dir_name: str, event_name: str):
        handle_rpc_script = self.replace_text(
            add_python_indentation(self.handle_rpc_partial, 2), 
            {
                'zarubaEventName': event_name,
                'zaruba_event_name': snake(event_name),
            }
        )
        handle_rpc_script = '\n{}\n'.format(handle_rpc_script)
        controller_file_name = self._get_controller_file_name(dir_name)
        controller_file_content = self.read_file_from_disk(controller_file_name)
        controller_script_lines = controller_file_content.split('\n')
        insert_index = -1
        for line_index, line in enumerate(controller_script_lines):
            if line.startswith(add_python_indentation('def handle_event', 1)):
                insert_index = line_index + 1
                break
        if insert_index == -1:
            raise ValueError('Cannot find handle_event method {}'.format(controller_file_name))
        controller_script_lines.insert(insert_index, handle_rpc_script)
        controller_file_content = '\n'.join(controller_script_lines)
        self.write_file_to_disk(controller_file_name, controller_file_content)


    def _get_controller_file_name(self, dir_name: str) -> str:
        controller_file_name = '{}/{}/controller.py'.format(self.service_name, self.module_name)
        return os.path.abspath(os.path.join(dir_name, controller_file_name))
    

    def add_python_indentation(self, text: str, level: int) -> str:
        spaces = (level * 4) * ' '
        indented_lines = [spaces + line for line in text.split('\n')]
        return '\n'.join(indented_lines)

    
    def generate(self, dir_name: str):
        self.replacement_dict = {
            'zarubaServiceName': self.service_name,
            'zarubaModuleName': self.module_name,
        }
        super().generate(dir_name)
        self._register_module(dir_name)
    

    def _register_module(self, dir_name: str):
        # prepare partials
        replacement_dict = {
            'zarubaModuleName': self.module_name,
            'ZarubaModuleName': capitalize(self.module_name),
            'zaruba_module_name': snake(self.module_name),
        }
        import_module_script = self.replace_text(self.import_module_partial, replacement_dict) 
        load_module_script = self.replace_text(self.load_module_partial, replacement_dict)  
        # load service
        service = FastApiService(self.service_name)
        service.load(dir_name)
        main_file_name = '{}/main.py'.format(self.service_name)
        main_file_content = service.get_content(main_file_name)
        main_file_lines = main_file_content.split('\n')
        main_file_lines = self._insert_import_module_script(main_file_lines, import_module_script)
        main_file_lines = self._insert_load_module_script(main_file_lines, load_module_script)
        main_file_content = '\n'.join(main_file_lines)
        service.set_content(main_file_name, main_file_content)
        service.save(dir_name)
    

    def _insert_import_module_script(self, lines: List[str], import_module_script: str) -> List[str]:
        import_module_line_index = 0
        for line_index, line in enumerate(lines):
            if not line.startswith('from '):
                import_module_line_index = line_index
                break
        lines.insert(import_module_line_index, import_module_script)
        return lines
    
    def _insert_load_module_script(self, lines: List[str], load_module_script: str) -> List[str]:
        lines.append('')
        lines.append(load_module_script)
        return lines


class FastApiCrud(CodeGen):

    def __init__(self, service_name: str, module_name: str, entity_name: str, field_names: List[str]):
        super().__init__()
        self.service_name = service_name
        self.module_name = module_name
        self.entity_name = entity_name
        self.field_names = field_names
        self.controller_handle_event_partial = ''
        self.controller_import_partial = ''
        self.repo_field_declaration_partial = ''
        self.repo_field_update_partial = ''
        self.repo_field_insert_partial = ''
        self.init_repo_partial = ''
        self.controller_handle_route_partial = ''
        self.controller_init_property_partial = ''
        self.import_repo_partial = ''
        self.schema_field_declaration_partial = ''
    

    def load(self, dir_name: str):
        self.file_content_dict = {
            '{}/repos/db{}.py'.format(self.service_name, capitalize(self.entity_name)): '',
            '{}/repos/{}.py'.format(self.service_name, self.entity_name): '',
            '{}/schemas/{}.py'.format(self.service_name, self.entity_name): '',
            '{}/{}/handle{}Event.py'.format(self.service_name, self.module_name, capitalize(self.entity_name)): '',
            '{}/{}/handle{}Route.py'.format(self.service_name, self.module_name, capitalize(self.entity_name)): '',
        }
        super().load(dir_name)


    def load_from_template(self, template_dir_name: str):
        self.file_content_dict = {
            'zarubaServiceName/repos/dbZarubaEntityName.py': '',
            'zarubaServiceName/repos/zarubaEntityName.py': '',
            'zarubaServiceName/schemas/zarubaEntityName.py': '',
            'zarubaServiceName/zarubaModuleName/handleZarubaEntityNameEvent.py': '',
            'zarubaServiceName/zarubaModuleName/handleZarubaEntityNameRoute.py': '',
        }
        super().load_from_template(template_dir_name)
        partial_path = os.path.join(os.path.abspath(template_dir_name), 'partials')
        self.controller_handle_event_partial = self.read_file_from_disk(os.path.join(partial_path, 'controller_handle_event.py'))
        self.controller_import_partial = self.read_file_from_disk(os.path.join(partial_path, 'controller_import.py'))
        self.repo_field_declaration_partial = self.read_file_from_disk(os.path.join(partial_path, 'repo_field_declaration.py'))
        self.repo_field_update_partial = self.read_file_from_disk(os.path.join(partial_path, 'repo_field_update.py'))
        self.repo_field_insert_partial = self.read_file_from_disk(os.path.join(partial_path, 'repo_field_insert.py'))
        self.init_repo_partial = self.read_file_from_disk(os.path.join(partial_path, 'init_repo.py'))
        self.controller_handle_route_partial = self.read_file_from_disk(os.path.join(partial_path, 'controller_handle_route.py'))
        self.controller_init_property_partial = self.read_file_from_disk(os.path.join(partial_path, 'controller_init_property.py'))
        self.import_repo_partial = self.read_file_from_disk(os.path.join(partial_path, 'import_repo.py'))
        self.schema_field_declaration_partial = self.read_file_from_disk(os.path.join(partial_path, 'schema_field_declaration.py'))

    
    def generate(self, dir_name: str):
        self.replacement_dict = {
            'zarubaServiceName': self.service_name,
            'zarubaModuleName': self.module_name,
            'zarubaEntityName': self.entity_name,
            'ZarubaEntityName': capitalize(self.entity_name),
            'zaruba_entity_name': snake(self.entity_name),
            'zaruba_field_name': self.field_names[0] if len(self.field_names) > 0 else 'id',
        }
        self._complete_repo()
        self._complete_schema()
        super().generate(dir_name)
        self._register_handler(dir_name)
        self._adjust_service(dir_name)

    
    def _adjust_service(self, dir_name: str):
        replace_dict = {
            'zarubaEntityName': self.entity_name,
            'ZarubaEntityName': capitalize(self.entity_name),
            'zaruba_entity_name': snake(self.entity_name),
        }
        # get import script
        import_script = self.replace_text(self.import_repo_partial, replace_dict).strip()
        init_script = self.replace_text(self.init_repo_partial, replace_dict).strip()
        # load service
        service = FastApiService(self.service_name)
        service.load(dir_name)
        main_file_name = '{}/main.py'.format(self.service_name)
        main_script = service.get_content(main_file_name)
        main_lines = main_script.split('\n')
        # add import
        insert_import_index = 0
        for line_index, line in enumerate(main_lines):
            if not line.startswith('from '):
                insert_import_index = line_index
                break
        main_lines.insert(insert_import_index, import_script)
        # adjust init script
        controller_declaration_index = -1
        for line_index, line in enumerate(main_lines):
            if line.startswith('{}_controller ='.format(snake(self.module_name))):
                controller_declaration_index = line_index
                break
        if controller_declaration_index == -1:
            raise ValueError('Cannot find {}_controller declaration {}'.format(snake(self.module_name), main_file_name))
        # adjust controller declaration
        controller_declaration_line = main_lines[controller_declaration_index]
        controller_declaration_line = controller_declaration_line.replace(')', ', {entity_repo}={entity_repo})'.format(
            entity_repo = '{}_repo'.format(snake(self.entity_name)),
        ))
        main_lines[controller_declaration_index] = controller_declaration_line
        # add repo init
        main_lines.insert(controller_declaration_index, init_script)
        # save changes
        main_script = '\n'.join(main_lines)
        service.set_content(main_file_name, main_script)
        service.save(dir_name)
    

    def _register_handler(self, dir_name: str):
        module = FastApiModule(self.service_name, self.module_name)
        module.load(dir_name)
        controller_file_name = '{}/{}/controller.py'.format(self.service_name, self.module_name)
        self._insert_controller_import(module, controller_file_name)
        self._adjust_controller_constructor(module, controller_file_name)
        self._add_controller_event_handler(module, controller_file_name)
        self._add_controller_route_handler(module, controller_file_name)
        module.save(dir_name)


    def _insert_controller_import(self, module: FastApiModule, controller_file_name: str):
        import_script = self.replace_text(
            self.controller_import_partial,
            {
                'zarubaModuleName': self.module_name,
                'zarubaEntityName': self.entity_name,
                'ZarubaEntityName': capitalize(self.entity_name),
                'zaruba_entity_name': snake(self.entity_name),
            }
        ).strip()
        controller_script = module.get_content(controller_file_name)
        controller_script_lines = controller_script.split('\n')
        insert_index = 0
        for line_index, line in enumerate(controller_script_lines):
            if not line.startswith('from '):
                insert_index = line_index
                break
        controller_script_lines.insert(insert_index, import_script)
        controller_script = '\n'.join(controller_script_lines)
        module.set_content(controller_file_name, controller_script)
    

    def _adjust_controller_constructor(self, module: FastApiModule, controller_file_name: str):
        init_property_script = self.replace_text(
            self.controller_init_property_partial,
            {
                'zaruba_entity_name': snake(self.entity_name),
            }
        ).strip()
        init_property_script = add_python_indentation(init_property_script, 2)
        controller_script = module.get_content(controller_file_name)
        controller_script_lines = controller_script.split('\n')
        controller_class_index = -1
        constructor_index = -1
        insert_index = -1
        for line_index, line in enumerate(controller_script_lines):
            if line.startswith('class Controller('):
                controller_class_index = line_index
                continue
            if controller_class_index > -1 and line.startswith(add_python_indentation('def __init__(', 1)):
                constructor_index = line_index
                insert_index = line_index + 1
                continue
            if constructor_index > -1 and line.startswith(add_python_indentation('self.enable_event', 2)):
                insert_index = line_index + 1
                break
            if constructor_index > -1 and not line.startswith(add_python_indentation('', 2)):
                break 
        if insert_index == -1:
            raise ValueError('Cannot find Controller constructor in {}'.format(controller_file_name))
        # update constructor 
        constructor_line = controller_script_lines[constructor_index]
        constructor_line = constructor_line.replace('):', ', {entity_name}_repo: {EntityName}Repo):'.format(
            entity_name = snake(self.entity_name),
            EntityName = capitalize(self.entity_name)
        ))
        controller_script_lines[constructor_index] = constructor_line
        # insert
        controller_script_lines.insert(insert_index, init_property_script)
        controller_script = '\n'.join(controller_script_lines)
        module.set_content(controller_file_name, controller_script)
    
    
    def _add_controller_event_handler(self, module: FastApiModule, controller_file_name: str):
        handler_script = self.replace_text(
            self.controller_handle_event_partial,
            {
                'zaruba_entity_name': snake(self.entity_name),
            }
        ).strip()
        handler_script = add_python_indentation(handler_script, 3)
        controller_script = module.get_content(controller_file_name)
        controller_script_lines = controller_script.split('\n')
        controller_class_index = -1
        controller_start_index = -1
        for line_index, line in enumerate(controller_script_lines):
            if line.startswith('class Controller('):
                controller_class_index = line_index
                continue
            if controller_class_index > -1 and line.startswith(add_python_indentation('def start(', 1)):
                controller_start_index = line_index
                insert_index = line_index + 1
                continue
            if controller_start_index > -1 and line.startswith(add_python_indentation('if self.enable_event', 2)):
                insert_index = line_index + 1
                break
            if controller_start_index > -1 and not line.startswith(add_python_indentation('', 2)):
                break 
        if insert_index == -1:
            raise ValueError('Cannot find Controller constructor in {}'.format(controller_file_name))
        # insert
        controller_script_lines.insert(insert_index, handler_script)
        controller_script = '\n'.join(controller_script_lines)
        module.set_content(controller_file_name, controller_script)

 
    def _add_controller_route_handler(self, module: FastApiModule, controller_file_name: str):
        handler_script = self.replace_text(
            self.controller_handle_route_partial,
            {
                'zaruba_entity_name': snake(self.entity_name),
            }
        ).strip()
        handler_script = add_python_indentation(handler_script, 3)
        controller_script = module.get_content(controller_file_name)
        controller_script_lines = controller_script.split('\n')
        controller_class_index = -1
        controller_start_index = -1
        for line_index, line in enumerate(controller_script_lines):
            if line.startswith('class Controller('):
                controller_class_index = line_index
                continue
            if controller_class_index > -1 and line.startswith(add_python_indentation('def start(', 1)):
                controller_start_index = line_index
                insert_index = line_index + 1
                continue
            if controller_start_index > -1 and line.startswith(add_python_indentation('if self.enable_route', 2)):
                insert_index = line_index + 1
                break
            if controller_start_index > -1 and not line.startswith(add_python_indentation('', 2)):
                break 
        if insert_index == -1:
            raise ValueError('Cannot find Controller constructor in {}'.format(controller_file_name))
        # insert
        controller_script_lines.insert(insert_index, handler_script)
        controller_script = '\n'.join(controller_script_lines)
        module.set_content(controller_file_name, controller_script)
 
    
    def _complete_repo(self):
        if len(self.field_names) == 0:
            return
        self._complete_repo_field_declaration()
        self._complete_repo_field_insert()
        self._complete_repo_field_update() 
    
    def _complete_repo(self):
        if len(self.field_names) == 0:
            return
        self._complete_repo_field_declaration()
        self._complete_repo_field_insert()
        self._complete_repo_field_update()


    def _complete_repo_field_declaration(self):
        # get field declaration
        field_declaration_lines = []
        for field_name in self.field_names:
            new_line = self.replace_text(
                self.repo_field_declaration_partial,
                {
                    'zaruba_field_name': snake(field_name),
                }
            )
            new_line = add_python_indentation(new_line.strip('\n'), 1)
            field_declaration_lines.append(new_line)
        field_declaration_script = '\n'.join(field_declaration_lines)
        # get db repo script
        db_repo_file_name = 'zarubaServiceName/repos/dbZarubaEntityName.py'
        db_repo_script = self.get_content(db_repo_file_name)
        db_repo_lines = db_repo_script.split('\n')
        # look for insert index
        entity_class_index = -1
        insert_index = -1
        table_name_index = -1
        for line_index, line in enumerate(db_repo_lines):
            if line.startswith('class DBZarubaEntityNameEntity'):
                entity_class_index = line_index
            if entity_class_index != -1 and line.startswith(add_python_indentation('__tablename__', 1)):
                table_name_index = line_index
            if entity_class_index != -1 and line.startswith(add_python_indentation('id = Column(', 1)):
                insert_index = line_index + 1
                break
        if insert_index == -1 and table_name_index != -1:
            insert_index = table_name_index + 1
        if insert_index == -1 and entity_class_index != -1:
            insert_index = entity_class_index + 1
        if insert_index == -1:
            raise ValueError('Cannot find DBZarubaEntityNameEntity class in {}'.fomrat(db_repo_file_name))
        # insert new line
        db_repo_lines.insert(insert_index, field_declaration_script)
        db_repo_script = '\n'.join(db_repo_lines)
        self.set_content(db_repo_file_name, db_repo_script)


    def _complete_repo_field_insert(self):
        # get field insert
        field_insert_lines = []
        for field_name in self.field_names:
            new_line = self.replace_text(
                self.repo_field_insert_partial,
                {
                    'zaruba_field_name': snake(field_name),
                }
            )
            new_line = add_python_indentation(new_line.strip('\n'), 4)
            field_insert_lines.append(new_line)
        field_insert_script = '\n'.join(field_insert_lines)
        # get db repo script
        db_repo_file_name = 'zarubaServiceName/repos/dbZarubaEntityName.py'
        db_repo_script = self.get_content(db_repo_file_name)
        db_repo_lines = db_repo_script.split('\n')
        # look for insert index
        repo_class_index = -1
        method_index = -1
        instance_index = -1
        insert_index = -1
        for line_index, line in enumerate(db_repo_lines):
            if line.startswith('class DBZarubaEntityNameRepo'):
                repo_class_index = line_index
            if repo_class_index != -1 and line.startswith(add_python_indentation('def insert(self,', 1)):
                method_index = line_index
            if method_index != -1 and line.startswith(add_python_indentation('db_entity = DBZarubaEntityNameEntity', 3)):
                instance_index = line_index
            if instance_index != -1 and line.startswith(add_python_indentation('id=str(', 4)):
                insert_index = line_index + 1
                break
        if insert_index == -1 and instance_index != -1:
            insert_index = instance_index + 1
        if insert_index == -1:
            raise ValueError('Cannot find data-insert on DBZarubaEntityNameRepo.insert in {}'.fomrat(db_repo_file_name))
        # insert new line
        db_repo_lines.insert(insert_index, field_insert_script)
        db_repo_script = '\n'.join(db_repo_lines)
        self.set_content(db_repo_file_name, db_repo_script)


    def _complete_repo_field_update(self):
        # get field update
        field_update_lines = []
        for field_name in self.field_names:
            new_line = self.replace_text(
                self.repo_field_update_partial,
                {
                    'zaruba_field_name': snake(field_name),
                    'zaruba_entity_name': snake(self.entity_name),
                }
            )
            new_line = add_python_indentation(new_line.strip('\n'), 3)
            field_update_lines.append(new_line)
        field_update_script = '\n'.join(field_update_lines)
        # get db repo script
        db_repo_file_name = 'zarubaServiceName/repos/dbZarubaEntityName.py'
        db_repo_script = self.get_content(db_repo_file_name)
        db_repo_lines = db_repo_script.split('\n')
        # look for update index
        repo_class_index = -1
        method_index = -1
        insert_index = -1
        for line_index, line in enumerate(db_repo_lines):
            if line.startswith('class DBZarubaEntityNameRepo'):
                repo_class_index = line_index
            if repo_class_index != -1 and line.startswith(add_python_indentation('def update(self,', 1)):
                method_index = line_index
            if method_index != -1 and line.startswith(add_python_indentation('db_entity.updated_at', 3)):
                insert_index = line_index
                break
            if method_index != -1 and insert_index == -1 and line.startswith(add_python_indentation('db.add(db_entity)', 3)):
                insert_index = line_index
                break
        if insert_index == -1:
            raise ValueError('Cannot find data-update on DBZarubaEntityNameRepo.update in {}'.fomrat(db_repo_file_name))
        # update new line
        db_repo_lines.insert(insert_index, field_update_script)
        db_repo_script = '\n'.join(db_repo_lines)
        self.set_content(db_repo_file_name, db_repo_script)


    def _complete_schema(self):
        if len(self.field_names) == 0:
            return
        # get schema field declaration
        schema_field_declaration_lines = []
        for field_name in self.field_names:
            new_line = self.replace_text(
                self.schema_field_declaration_partial, 
                {
                    'zaruba_field_name': snake(field_name),
                }
            )
            new_line = add_python_indentation(new_line.strip('\n'), 1)
            schema_field_declaration_lines.append(new_line)
        schema_field_declaration = '\n'.join(schema_field_declaration_lines)
        # get schema script
        schema_script_file_name = 'zarubaServiceName/schemas/zarubaEntityName.py'
        schema_script = self.get_content(schema_script_file_name)
        schema_script_lines = schema_script.split('\n')
        # insert schema field declaration to schema script
        insert_index = -1
        for line_index, line in enumerate(schema_script_lines):
            if line.startswith('class ZarubaEntityNameData('):
                insert_index = line_index + 1
                if schema_script_lines[insert_index] == add_python_indentation('pass', 1):
                    schema_script_lines.pop(insert_index)
                break
        if insert_index == -1:
            raise ValueError('Cannot find ZarubaEntityNameData class in {}'.format(schema_script_file_name))
        schema_script_lines.insert(insert_index, schema_field_declaration)
        schema_script = '\n'.join(schema_script_lines)
        self.set_content(schema_script_file_name, schema_script)