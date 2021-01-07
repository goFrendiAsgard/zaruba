import random

evangelion_angels = [
    'Adam', 'Lilith', 'Sachiel', 'Shamshel', 'Ramiel', 'Gaghiel', 'Israfel',
    'Sandalphon', 'Matarael', 'Sahaquiel', 'Ireul', 'Leliel', 'Bardiel',
    'Zeruel', 'Arael', 'Armisael', 'Tabris', 'Lilin']

ubuntu_adjectives = [
    'warty', 'hoary', 'breezy', 'dapper', 'edgy', 'Feisty', 'gutsy', 'hardy', 
    'interpid', 'jaunty', 'karmic', 'lucid', 'maverick', 'natty', 'oneiric', 
    'precise', 'quantal', 'raring', 'saucy', 'trusty', 'utopic', 'vivid', 'wily', 
    'xenial', 'yakkety', 'zesty', 'artful', 'bionic', 'cosmic', 'disco', 'eoam',
    'focal', 'groovy']

def get_random_name():
    system_random = random.SystemRandom()
    adjective = system_random.choice(ubuntu_adjectives)
    angel = system_random.choice(evangelion_angels)
    return '{}{}'.format(adjective, angel)

if __name__ == '__main__':
    print(get_random_name())