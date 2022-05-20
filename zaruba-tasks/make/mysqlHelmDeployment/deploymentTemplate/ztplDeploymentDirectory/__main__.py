import pulumi
from pulumi_kubernetes.helm.v3 import Chart, ChartOpts, LocalChartOpts
import os

app = Chart(
    'ztpl-app-name', 
    config=LocalChartOpts(
        path = './chart',
        namespace = os.getenv('NAMESPACE', 'default'),
        values = {
            'fullnameOverride': os.getenv('FULLNAME_OVERRIDE', ''),
            'image': {
                'repository': os.getenv('IMAGE_REGISTRY', 'docker.io'),
                'repository': os.getenv('IMAGE_REPOSITORY', 'bitnami/mysql'),
                'tag': os.getenv('IMAGE_TAG', '8.0.29'),
            },
            'auth': {
                'rootPassword': os.getenv('AUTH_ROOT_PASSWORD', 'toor'),
                'database': os.getenv('AUTH_DATABASE', 'sample'),
                'username': os.getenv('AUTH_USERNAME', 'mysql'),
                'password': os.getenv('AUTH_PASSWORD', 'mysql'),
                'replicationUser': os.getenv('AUTH_REPLICATION_USER', 'replicator'),
                'replicationPassword': os.getenv('AUTH_REPLICATION_PASSWORD', 'replicator'),
            },
            'primary': {
                'service': {
                    'port': os.getenv('PRIMARY_SERVICE_PORT', '3306'),
                    'type': os.getenv('PRIMARY_SERVICE_TYPE', 'ClusterIP'),
                },
            },
        },
        skip_await = True
    )
)

pulumi.export('app', app)
