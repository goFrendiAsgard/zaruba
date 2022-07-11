from typing import Any, Mapping
from pulumi_kubernetes.helm.v3 import Chart, ChartOpts, LocalChartOpts

import pulumi
import json

# define config
config: Mapping[str, Any]
with open('./config/config.json') as f:
    config = json.load(f)

app = Chart(
    'ztpl-deployment-name', 
    config=LocalChartOpts(
        path = './chart',
        namespace = config.get('namespace', 'default'),
        values = {
            'image': {
                'repository': config.get('image.repository'),
                'tag': config.get('image.tag', 'latest')
            },
            'replicaCount': config.get('replicaCount', '1'),
            'env': config.get('env', []),
            'ports': config.get('ports', []),
            'service': {
                'ports': config.get('service.ports', []),
                'type': config.get('service.type', 'ClusterIP'),
                'enabled': config.get('service.enabled', 'true'),
            }
        },
        skip_await = True
    )
)

pulumi.export('app', app)
