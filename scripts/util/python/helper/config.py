from typing import Any, List, Mapping
from ruamel.yaml import YAML
from io import StringIO
import os

class YamlConfig:


    def __init__(self, data: Mapping[str, Any]):
        self.data: Mapping[str, Any] = data
    

    def _create_dir_if_not_exist(self, file_name: str):
        abs_file_name = os.path.abspath(file_name)
        abs_dir = os.path.dirname(abs_file_name)
        if not os.path.exists(abs_dir):
            os.makedirs(abs_dir)
    
 
    def _get_indentation_level(self, line: str) -> int:
        return (len(line) - len(line.lstrip(' -')))


    def _pretify_yaml(self, text: str) -> str:
        lines = text.split('\n')
        non_empty_lines = []
        prev_line = ''
        for line_index, line in enumerate(lines):
            if line.strip() == '':
                continue
            if len(non_empty_lines) != 0:
                level = self._get_indentation_level(line)
                prev_level = self._get_indentation_level(prev_line)
                if (level < prev_level and (level == 0 or level == 2)) or (level == 2 and (prev_level == 0 or (line.startswith('-') and not prev_line.startswith('-')))):
                    non_empty_lines.append('')
            non_empty_lines.append(line)
            prev_line = line
        return '\n'.join(non_empty_lines)


    def load(self, file_name: str):
        yaml=YAML()
        f = open(file_name, 'r')
        data = yaml.load(f)
        f.close()
        if data is not None:
            self.data = data


    def save(self, file_name: str):
        string_stream = StringIO()
        yaml=YAML()
        yaml.dump(self.data, string_stream)
        yaml_text = self._pretify_yaml(string_stream.getvalue())
        string_stream.close()
        # write file
        self._create_dir_if_not_exist(file_name)
        f = open(file_name, 'w')
        f.write(yaml_text)
        f.close()


    def append(self, keys: List[str], new_value: str):
        if not self.exist(keys):
            self.set(keys, [])
        values = self.get(keys)
        values.append(new_value)

    
    def append_if_not_exist(self, keys: List[str], new_value: str):
        if not self.exist(keys):
            self.set(keys, [])
        values = self.get(keys)
        if new_value not in values:
            values.append(new_value)
    

    def exist(self, keys: List[str]) -> bool:
        value: Any = self.data
        for key in keys:
            if isinstance(value, list):
                if len(value) > key:
                    value = value[key]
                    continue
                return False
            elif isinstance(value, dict):
                if key in value:
                    value = value[key]
                    continue
                return False
            return False
        return True
    
    
    def get(self, keys: List[str]) -> Any:
        value: Any = self.data
        for key in keys:
            if isinstance(value, list):
                if len(value) > key:
                    value = value[key]
                    continue
                raise ValueError('`{value}` has no index `{key}`'.format(key=key, value=value))
            elif isinstance(value, dict):
                if key in value:
                    value = value[key]
                    continue
                raise ValueError('`{value}` has no key `{key}`'.format(key=key, value=value))
            raise ValueError('`{value}` is neither list or dictionary'.format(value=value))
        return value

    
    def set(self, keys: List[str], value: Any):
        data: Any = self.data
        last_key = keys[-1]
        prefix_keys = keys[:len(keys)-1]
        for key in prefix_keys:
            if key not in data:
                data[key] = {}
            data = data[key]
        data[last_key] = value
    
 
    def set_default(self, keys: List[str], value: Any):
        if not self.exist(keys):
            self.set(keys, value)


    def unset(self, keys: List[str]):
        data: Any = self.data
        last_key = keys[-1]
        prefix_keys = keys[:len(keys)-1]
        for key in prefix_keys:
            if key not in data:
                return
            data = data[key]
        if last_key in data:
            del data[last_key]

