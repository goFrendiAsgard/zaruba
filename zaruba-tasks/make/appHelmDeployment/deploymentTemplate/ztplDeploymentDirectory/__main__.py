import pulumi
from pulumi_kubernetes.helm.v3 import Chart, ChartOpts, LocalChartOpts
import json

# define config
config = {
    'namespace': 'default',
    'image.respository': 'nginx',
    'image.tag': 'latest',
    'replicaCount': 1,
    'env': [],
    'ports': [],
    'service.ports': [],
}

# merge config with user's config
with open('./config/config.json') as f:
    user_config = json.load(f)
    config.update(user_config)

app = Chart(
    'ztpl-app-name', 
    config=LocalChartOpts(
        path = './chart',
        namespace = config.get('namespace'),
        values = {
            'image': {
                'repository': config.get('image.repository'),
                'tag': config.get('image.tag')
            },
            'replicaCount': config.get('replicaCount'),
            'env': config.get('env'),
            'ports': config.get('ports'),
            'service': {
                'ports': config.get('service.ports')
            }
        },
        skip_await = True
    )
)

pulumi.export("app", app)
