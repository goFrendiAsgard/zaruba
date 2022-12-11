from typing import Any, Mapping
from pulumi_kubernetes.helm.v3 import Chart, LocalChartOpts
from helper import Config, Probe

import pulumi
import os

config = Config('./config/config.json')
probe = Probe(config)

app = Chart(
    'ztpl-deployment-name', 
    config=LocalChartOpts(
        path = './chart',
        namespace = os.getenv('NAMESPACE', 'default'),
        values = {
            'image': {
                'repository': config.get('image.repository'),
                'tag': config.get('image.tag', 'latest')
            },
            'fullnameOverride': os.getenv('FULLNAME_OVERRIDE'),
            'replicaCount': int(os.getenv('REPLICA_COUNT', '1')),
            'env': config.get('env', []),
            'ports': config.get('ports', []),
            'service': {
                'ports': config.get('service.ports', []),
                'type': os.getenv('SERVICE_TYPE', 'ClusterIP'),
                'enabled': os.getenv('SERVICE_ENABLED', 'True') == 'True',
            },
            'resources': {
                'limits': {
                    'cpu': os.getenv('RESOURCES_LIMITS_CPU', '100m'),
                    'memory': os.getenv('RESOURCES_LIMITS_MEMORY', '128Mi'),
                },
                'requests': {
                    'cpu': os.getenv('RESOURCES_REQUESTS_CPU', '100m'),
                    'memory': os.getenv('RESOURCES_REQUESTS_MEMORY', '128Mi'),
                },
            },
            'autoscaling': {
                'enabled': os.getenv('AUTOSCALING_ENABLED', 'True') == 'True',
                'minReplicas': int(os.getenv('AUTOSCALING_MIN_REPLICAS', '1')),
                'maxReplicas': int(os.getenv('AUTOSCALING_MAX_REPLICAS', '1')),
                'targetCPUUtilizationPercentage': int(os.getenv('AUTOSCALING_TARGET_CPU_UTILIZATION_PERCENTAGE', '80')),
                'targetMemorytilizationPercentage': int(os.getenv('AUTOSCALING_TARGET_MEMORY_UTILIZATION_PERCENTAGE', '80')),
            },
            'livenessProbe': probe.get_liveness_config(),
            'readinessProbe': probe.get_readiness_config(),
        },
        skip_await = True
    )
)

pulumi.export('app', app)
