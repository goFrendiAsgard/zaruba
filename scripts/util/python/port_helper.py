from typing import List, Mapping

def get_possible_ports_env(env_dict: Mapping[str, str], includes: List[int] = [21, 22, 23, 25, 53, 80, 443], excludes: List[int] = [], minimum: int = 3000, maximum: int = 100000) -> Mapping[str, str]:
    possible_ports_env: Mapping[str, str] = {}
    for env_key, env_val in env_dict.items():
        if not env_val.isnumeric():
            continue
        env_val_int = int(env_val)
        if env_val_int in excludes:
            continue
        if env_val_int in includes or (env_val_int >= 3000 and env_val_int < maximum):
            possible_ports_env[env_key] = env_val
    return possible_ports_env


def get_service_possible_ports_env(env_dict: Mapping[str, str]) -> Mapping[str, str]:
    return get_possible_ports_env(
        env_dict, 
        excludes=[3306, 5432, 5672, 15672, 27017, 6379, 9200, 9300, 7001, 7199, 9042, 9160], 
        minimum=3000, 
        maximum=100000
    )
