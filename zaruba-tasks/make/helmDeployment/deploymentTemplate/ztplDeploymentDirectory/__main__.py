import pulumi
from pulumi_kubernetes.helm.v3 import Chart, ChartOpts, LocalChartOpts
import json

namespace = 'default'

app = Chart(
    'ztpl-deployment-name', 
    config=LocalChartOpts(
        path='./chart',
        namespace = namespace,
        values = {
            # TODO: put your helm values
        },
        skip_await =True
    )
)

pulumi.export("app", app)
