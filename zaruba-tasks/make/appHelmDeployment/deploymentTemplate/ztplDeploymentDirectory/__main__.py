from typing import Any, Mapping
from pulumi_kubernetes.helm.v3 import Chart, ChartOpts, LocalChartOpts

import pulumi
import json
import os

# define config
config: Mapping[str, Any]
with open('./config/config.json') as f:
    config = json.load(f)

app = Chart(
    'ztpl-deployment-name', 
    config=LocalChartOpts(
        path = './chart',
        namespace = os.getenv('NAMESPACE', 'default'),
        values = {
            'image': {
                'repository': os.getenv('IMAGE_REPOSITORY'),
                'tag': os.getenv('IMAGE_TAG', 'latest')
            },
            'fullnameOverride': os.getenv('FULLNAME_OVERRIDE'),
            'replicaCount': int(os.getenv('REPLICA_COUNT', '1')),
            'env': config.get('env', []),
            'ports': config.get('ports', []),
            'service': {
                'ports': config.get('service.ports', []),
                'type': os.getenv('SERVICE_TYPE', 'ClusterIP'),
                'enabled': os.getenv('SERVICE_ENABLED', 'True') == 'True',
            }
        },
        skip_await = True
    )
)

pulumi.export('app', app)
