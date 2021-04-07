import re

script_template = '''
    @app.get('{url}')
    def {handler}():
        return 'response of {url}'
'''


def get_script(url: str) -> str:
    handler='handle_route_{}'.format(re.sub(r'[^A-Za-z0-9_]+', '_', url).lower())
    return script_template.format(url=url, handler=handler)