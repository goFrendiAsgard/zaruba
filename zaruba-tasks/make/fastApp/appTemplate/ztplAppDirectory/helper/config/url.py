def url_path(url_path: str) -> str:
    '''
    Add preceeding slash, remove trailing slash
    '''
    return '/{}'.format(url_path.strip('/'))
