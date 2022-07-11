import pulumi
from pulumi_kubernetes.helm.v3 import Chart, ChartOpts, LocalChartOpts
import os

app = Chart(
    'ztpl-app-name', 
    config=LocalChartOpts(
        path = './chart',
        namespace = os.getenv('NAMESPACE', 'default'),
        values = {
            "architecture": os.getenv('ARCHITECTURE', 'standalone'),
            "auth": {
                "customPasswordFiles": {},
                "database": os.getenv('AUTH_DATABASE', 'my_database'),
                "existingSecret": os.getenv('AUTH_EXISTINGSECRET', ''),
                "forcePassword": os.getenv('AUTH_FORCEPASSWORD', 'False') == 'True',
                "password": os.getenv('AUTH_PASSWORD', ''),
                "replicationPassword": os.getenv('AUTH_REPLICATIONPASSWORD', ''),
                "replicationUser": os.getenv('AUTH_REPLICATIONUSER', 'replicator'),
                "rootPassword": os.getenv('AUTH_ROOTPASSWORD', ''),
                "usePasswordFiles": os.getenv('AUTH_USEPASSWORDFILES', 'False') == 'True',
                "username": os.getenv('AUTH_USERNAME', '')
            },
            "clusterDomain": os.getenv('CLUSTERDOMAIN', 'cluster.local'),
            "commonAnnotations": {},
            "commonLabels": {},
            "diagnosticMode": {
                "args": [
                    os.getenv('DIAGNOSTICMODE_ARGS_0', 'infinity')
                ],
                "command": [
                    os.getenv('DIAGNOSTICMODE_COMMAND_0', 'sleep')
                ],
                "enabled": os.getenv('DIAGNOSTICMODE_ENABLED', 'False') == 'True'
            },
            "extraDeploy": [],
            "fullnameOverride": os.getenv('FULLNAMEOVERRIDE', ''),
            "global": {
                "imagePullSecrets": [],
                "imageRegistry": os.getenv('GLOBAL_IMAGEREGISTRY', ''),
                "storageClass": os.getenv('GLOBAL_STORAGECLASS', '')
            },
            "image": {
                "debug": os.getenv('IMAGE_DEBUG', 'False') == 'True',
                "pullPolicy": os.getenv('IMAGE_PULLPOLICY', 'IfNotPresent'),
                "pullSecrets": [],
                "registry": os.getenv('IMAGE_REGISTRY', 'docker.io'),
                "repository": os.getenv('IMAGE_REPOSITORY', 'bitnami/mysql'),
                "tag": os.getenv('IMAGE_TAG', '8.0.28-debian-10-r41')
            },
            "initdbScripts": {},
            "initdbScriptsConfigMap": os.getenv('INITDBSCRIPTSCONFIGMAP', ''),
            "metrics": {
                "enabled": os.getenv('METRICS_ENABLED', 'False') == 'True',
                "extraArgs": {
                    "primary": [],
                    "secondary": []
                },
                "image": {
                    "pullPolicy": os.getenv('METRICS_IMAGE_PULLPOLICY', 'IfNotPresent'),
                    "pullSecrets": [],
                    "registry": os.getenv('METRICS_IMAGE_REGISTRY', 'docker.io'),
                    "repository": os.getenv('METRICS_IMAGE_REPOSITORY', 'bitnami/mysqld-exporter'),
                    "tag": os.getenv('METRICS_IMAGE_TAG', '0.14.0-debian-10-r11')
                },
                "livenessProbe": {
                    "enabled": os.getenv('METRICS_LIVENESSPROBE_ENABLED', 'True') == 'True',
                    "failureThreshold": int(os.getenv('METRICS_LIVENESSPROBE_FAILURETHRESHOLD', '3')),
                    "initialDelaySeconds": int(os.getenv('METRICS_LIVENESSPROBE_INITIALDELAYSECONDS', '120')),
                    "periodSeconds": int(os.getenv('METRICS_LIVENESSPROBE_PERIODSECONDS', '10')),
                    "successThreshold": int(os.getenv('METRICS_LIVENESSPROBE_SUCCESSTHRESHOLD', '1')),
                    "timeoutSeconds": int(os.getenv('METRICS_LIVENESSPROBE_TIMEOUTSECONDS', '1'))
                },
                "readinessProbe": {
                    "enabled": os.getenv('METRICS_READINESSPROBE_ENABLED', 'True') == 'True',
                    "failureThreshold": int(os.getenv('METRICS_READINESSPROBE_FAILURETHRESHOLD', '3')),
                    "initialDelaySeconds": int(os.getenv('METRICS_READINESSPROBE_INITIALDELAYSECONDS', '30')),
                    "periodSeconds": int(os.getenv('METRICS_READINESSPROBE_PERIODSECONDS', '10')),
                    "successThreshold": int(os.getenv('METRICS_READINESSPROBE_SUCCESSTHRESHOLD', '1')),
                    "timeoutSeconds": int(os.getenv('METRICS_READINESSPROBE_TIMEOUTSECONDS', '1'))
                },
                "resources": {
                    "limits": {},
                    "requests": {}
                },
                "service": {
                    "annotations": {
                        "prometheus.io/scrape": os.getenv('METRICS_SERVICE_ANNOTATIONS_PROMETHEUS_IO_SCRAPE', 'true')
                    },
                    "port": int(os.getenv('METRICS_SERVICE_PORT', '9104')),
                    "type": os.getenv('METRICS_SERVICE_TYPE', 'ClusterIP')
                },
                "serviceMonitor": {
                    "additionalLabels": {},
                    "enabled": os.getenv('METRICS_SERVICEMONITOR_ENABLED', 'False') == 'True',
                    "honorLabels": os.getenv('METRICS_SERVICEMONITOR_HONORLABELS', 'False') == 'True',
                    "interval": os.getenv('METRICS_SERVICEMONITOR_INTERVAL', '30s'),
                    "namespace": os.getenv('METRICS_SERVICEMONITOR_NAMESPACE', ''),
                    "relabellings": [],
                    "scrapeTimeout": os.getenv('METRICS_SERVICEMONITOR_SCRAPETIMEOUT', '')
                }
            },
            "nameOverride": os.getenv('NAMEOVERRIDE', ''),
            "networkPolicy": {
                "allowExternal": os.getenv('NETWORKPOLICY_ALLOWEXTERNAL', 'True') == 'True',
                "enabled": os.getenv('NETWORKPOLICY_ENABLED', 'False') == 'True',
                "explicitNamespacesSelector": {}
            },
            "primary": {
                "affinity": {},
                "args": [],
                "command": [],
                "containerSecurityContext": {
                    "enabled": os.getenv('PRIMARY_CONTAINERSECURITYCONTEXT_ENABLED', 'True') == 'True',
                    "runAsUser": int(os.getenv('PRIMARY_CONTAINERSECURITYCONTEXT_RUNASUSER', '1001'))
                },
                "customLivenessProbe": {},
                "customReadinessProbe": {},
                "customStartupProbe": {},
                "existingConfigmap": os.getenv('PRIMARY_EXISTINGCONFIGMAP', ''),
                "extraEnvVars": [],
                "extraEnvVarsCM": os.getenv('PRIMARY_EXTRAENVVARSCM', ''),
                "extraEnvVarsSecret": os.getenv('PRIMARY_EXTRAENVVARSSECRET', ''),
                "extraFlags": os.getenv('PRIMARY_EXTRAFLAGS', ''),
                "extraVolumeMounts": [],
                "extraVolumes": [],
                "hostAliases": [],
                "initContainers": [],
                "livenessProbe": {
                    "enabled": os.getenv('PRIMARY_LIVENESSPROBE_ENABLED', 'True') == 'True',
                    "failureThreshold": int(os.getenv('PRIMARY_LIVENESSPROBE_FAILURETHRESHOLD', '3')),
                    "initialDelaySeconds": int(os.getenv('PRIMARY_LIVENESSPROBE_INITIALDELAYSECONDS', '5')),
                    "periodSeconds": int(os.getenv('PRIMARY_LIVENESSPROBE_PERIODSECONDS', '10')),
                    "successThreshold": int(os.getenv('PRIMARY_LIVENESSPROBE_SUCCESSTHRESHOLD', '1')),
                    "timeoutSeconds": int(os.getenv('PRIMARY_LIVENESSPROBE_TIMEOUTSECONDS', '1'))
                },
                "nodeAffinityPreset": {
                    "key": os.getenv('PRIMARY_NODEAFFINITYPRESET_KEY', ''),
                    "type": os.getenv('PRIMARY_NODEAFFINITYPRESET_TYPE', ''),
                    "values": []
                },
                "nodeSelector": {},
                "pdb": {
                    "enabled": os.getenv('PRIMARY_PDB_ENABLED', 'False') == 'True',
                    "maxUnavailable": os.getenv('PRIMARY_PDB_MAXUNAVAILABLE', ''),
                    "minAvailable": int(os.getenv('PRIMARY_PDB_MINAVAILABLE', '1'))
                },
                "persistence": {
                    "accessModes": [
                        os.getenv('PRIMARY_PERSISTENCE_ACCESSMODES_0', 'ReadWriteOnce')
                    ],
                    "annotations": {},
                    "enabled": os.getenv('PRIMARY_PERSISTENCE_ENABLED', 'True') == 'True',
                    "existingClaim": os.getenv('PRIMARY_PERSISTENCE_EXISTINGCLAIM', ''),
                    "selector": {},
                    "size": os.getenv('PRIMARY_PERSISTENCE_SIZE', '8Gi'),
                    "storageClass": os.getenv('PRIMARY_PERSISTENCE_STORAGECLASS', '')
                },
                "podAffinityPreset": os.getenv('PRIMARY_PODAFFINITYPRESET', ''),
                "podAnnotations": {},
                "podAntiAffinityPreset": os.getenv('PRIMARY_PODANTIAFFINITYPRESET', 'soft'),
                "podLabels": {},
                "podSecurityContext": {
                    "enabled": os.getenv('PRIMARY_PODSECURITYCONTEXT_ENABLED', 'True') == 'True',
                    "fsGroup": int(os.getenv('PRIMARY_PODSECURITYCONTEXT_FSGROUP', '1001'))
                },
                "readinessProbe": {
                    "enabled": os.getenv('PRIMARY_READINESSPROBE_ENABLED', 'True') == 'True',
                    "failureThreshold": int(os.getenv('PRIMARY_READINESSPROBE_FAILURETHRESHOLD', '3')),
                    "initialDelaySeconds": int(os.getenv('PRIMARY_READINESSPROBE_INITIALDELAYSECONDS', '5')),
                    "periodSeconds": int(os.getenv('PRIMARY_READINESSPROBE_PERIODSECONDS', '10')),
                    "successThreshold": int(os.getenv('PRIMARY_READINESSPROBE_SUCCESSTHRESHOLD', '1')),
                    "timeoutSeconds": int(os.getenv('PRIMARY_READINESSPROBE_TIMEOUTSECONDS', '1'))
                },
                "resources": {
                    "limits": {},
                    "requests": {}
                },
                "rollingUpdatePartition": os.getenv('PRIMARY_ROLLINGUPDATEPARTITION', ''),
                "service": {
                    "annotations": {},
                    "clusterIP": os.getenv('PRIMARY_SERVICE_CLUSTERIP', ''),
                    "externalTrafficPolicy": os.getenv('PRIMARY_SERVICE_EXTERNALTRAFFICPOLICY', 'Cluster'),
                    "loadBalancerIP": os.getenv('PRIMARY_SERVICE_LOADBALANCERIP', ''),
                    "loadBalancerSourceRanges": [],
                    "nodePort": os.getenv('PRIMARY_SERVICE_NODEPORT', ''),
                    "port": int(os.getenv('PRIMARY_SERVICE_PORT', '3306')),
                    "type": os.getenv('PRIMARY_SERVICE_TYPE', 'ClusterIP')
                },
                "sidecars": [],
                "startupProbe": {
                    "enabled": os.getenv('PRIMARY_STARTUPPROBE_ENABLED', 'True') == 'True',
                    "failureThreshold": int(os.getenv('PRIMARY_STARTUPPROBE_FAILURETHRESHOLD', '10')),
                    "initialDelaySeconds": int(os.getenv('PRIMARY_STARTUPPROBE_INITIALDELAYSECONDS', '15')),
                    "periodSeconds": int(os.getenv('PRIMARY_STARTUPPROBE_PERIODSECONDS', '10')),
                    "successThreshold": int(os.getenv('PRIMARY_STARTUPPROBE_SUCCESSTHRESHOLD', '1')),
                    "timeoutSeconds": int(os.getenv('PRIMARY_STARTUPPROBE_TIMEOUTSECONDS', '1'))
                },
                "tolerations": [],
                "updateStrategy": os.getenv('PRIMARY_UPDATESTRATEGY', 'RollingUpdate')
            },
            "rbac": {
                "create": os.getenv('RBAC_CREATE', 'False') == 'True'
            },
            "schedulerName": os.getenv('SCHEDULERNAME', ''),
            "secondary": {
                "affinity": {},
                "args": [],
                "command": [],
                "containerSecurityContext": {
                    "enabled": os.getenv('SECONDARY_CONTAINERSECURITYCONTEXT_ENABLED', 'True') == 'True',
                    "runAsUser": int(os.getenv('SECONDARY_CONTAINERSECURITYCONTEXT_RUNASUSER', '1001'))
                },
                "customLivenessProbe": {},
                "customReadinessProbe": {},
                "customStartupProbe": {},
                "existingConfigmap": os.getenv('SECONDARY_EXISTINGCONFIGMAP', ''),
                "extraEnvVars": [],
                "extraEnvVarsCM": os.getenv('SECONDARY_EXTRAENVVARSCM', ''),
                "extraEnvVarsSecret": os.getenv('SECONDARY_EXTRAENVVARSSECRET', ''),
                "extraFlags": os.getenv('SECONDARY_EXTRAFLAGS', ''),
                "extraVolumeMounts": [],
                "extraVolumes": [],
                "hostAliases": [],
                "initContainers": [],
                "livenessProbe": {
                    "enabled": os.getenv('SECONDARY_LIVENESSPROBE_ENABLED', 'True') == 'True',
                    "failureThreshold": int(os.getenv('SECONDARY_LIVENESSPROBE_FAILURETHRESHOLD', '3')),
                    "initialDelaySeconds": int(os.getenv('SECONDARY_LIVENESSPROBE_INITIALDELAYSECONDS', '5')),
                    "periodSeconds": int(os.getenv('SECONDARY_LIVENESSPROBE_PERIODSECONDS', '10')),
                    "successThreshold": int(os.getenv('SECONDARY_LIVENESSPROBE_SUCCESSTHRESHOLD', '1')),
                    "timeoutSeconds": int(os.getenv('SECONDARY_LIVENESSPROBE_TIMEOUTSECONDS', '1'))
                },
                "nodeAffinityPreset": {
                    "key": os.getenv('SECONDARY_NODEAFFINITYPRESET_KEY', ''),
                    "type": os.getenv('SECONDARY_NODEAFFINITYPRESET_TYPE', ''),
                    "values": []
                },
                "nodeSelector": {},
                "pdb": {
                    "enabled": os.getenv('SECONDARY_PDB_ENABLED', 'False') == 'True',
                    "maxUnavailable": os.getenv('SECONDARY_PDB_MAXUNAVAILABLE', ''),
                    "minAvailable": int(os.getenv('SECONDARY_PDB_MINAVAILABLE', '1'))
                },
                "persistence": {
                    "accessModes": [
                        os.getenv('SECONDARY_PERSISTENCE_ACCESSMODES_0', 'ReadWriteOnce')
                    ],
                    "annotations": {},
                    "enabled": os.getenv('SECONDARY_PERSISTENCE_ENABLED', 'True') == 'True',
                    "selector": {},
                    "size": os.getenv('SECONDARY_PERSISTENCE_SIZE', '8Gi'),
                    "storageClass": os.getenv('SECONDARY_PERSISTENCE_STORAGECLASS', '')
                },
                "podAffinityPreset": os.getenv('SECONDARY_PODAFFINITYPRESET', ''),
                "podAnnotations": {},
                "podAntiAffinityPreset": os.getenv('SECONDARY_PODANTIAFFINITYPRESET', 'soft'),
                "podLabels": {},
                "podSecurityContext": {
                    "enabled": os.getenv('SECONDARY_PODSECURITYCONTEXT_ENABLED', 'True') == 'True',
                    "fsGroup": int(os.getenv('SECONDARY_PODSECURITYCONTEXT_FSGROUP', '1001'))
                },
                "readinessProbe": {
                    "enabled": os.getenv('SECONDARY_READINESSPROBE_ENABLED', 'True') == 'True',
                    "failureThreshold": int(os.getenv('SECONDARY_READINESSPROBE_FAILURETHRESHOLD', '3')),
                    "initialDelaySeconds": int(os.getenv('SECONDARY_READINESSPROBE_INITIALDELAYSECONDS', '5')),
                    "periodSeconds": int(os.getenv('SECONDARY_READINESSPROBE_PERIODSECONDS', '10')),
                    "successThreshold": int(os.getenv('SECONDARY_READINESSPROBE_SUCCESSTHRESHOLD', '1')),
                    "timeoutSeconds": int(os.getenv('SECONDARY_READINESSPROBE_TIMEOUTSECONDS', '1'))
                },
                "replicaCount": int(os.getenv('SECONDARY_REPLICACOUNT', '1')),
                "resources": {
                    "limits": {},
                    "requests": {}
                },
                "rollingUpdatePartition": os.getenv('SECONDARY_ROLLINGUPDATEPARTITION', ''),
                "service": {
                    "annotations": {},
                    "clusterIP": os.getenv('SECONDARY_SERVICE_CLUSTERIP', ''),
                    "externalTrafficPolicy": os.getenv('SECONDARY_SERVICE_EXTERNALTRAFFICPOLICY', 'Cluster'),
                    "loadBalancerIP": os.getenv('SECONDARY_SERVICE_LOADBALANCERIP', ''),
                    "loadBalancerSourceRanges": [],
                    "nodePort": os.getenv('SECONDARY_SERVICE_NODEPORT', ''),
                    "port": int(os.getenv('SECONDARY_SERVICE_PORT', '3306')),
                    "type": os.getenv('SECONDARY_SERVICE_TYPE', 'ClusterIP')
                },
                "sidecars": [],
                "startupProbe": {
                    "enabled": os.getenv('SECONDARY_STARTUPPROBE_ENABLED', 'True') == 'True',
                    "failureThreshold": int(os.getenv('SECONDARY_STARTUPPROBE_FAILURETHRESHOLD', '15')),
                    "initialDelaySeconds": int(os.getenv('SECONDARY_STARTUPPROBE_INITIALDELAYSECONDS', '15')),
                    "periodSeconds": int(os.getenv('SECONDARY_STARTUPPROBE_PERIODSECONDS', '10')),
                    "successThreshold": int(os.getenv('SECONDARY_STARTUPPROBE_SUCCESSTHRESHOLD', '1')),
                    "timeoutSeconds": int(os.getenv('SECONDARY_STARTUPPROBE_TIMEOUTSECONDS', '1'))
                },
                "tolerations": [],
                "updateStrategy": os.getenv('SECONDARY_UPDATESTRATEGY', 'RollingUpdate')
            },
            "serviceAccount": {
                "annotations": {},
                "create": os.getenv('SERVICEACCOUNT_CREATE', 'True') == 'True',
                "name": os.getenv('SERVICEACCOUNT_NAME', '')
            },
            "volumePermissions": {
                "enabled": os.getenv('VOLUMEPERMISSIONS_ENABLED', 'False') == 'True',
                "image": {
                    "pullPolicy": os.getenv('VOLUMEPERMISSIONS_IMAGE_PULLPOLICY', 'IfNotPresent'),
                    "pullSecrets": [],
                    "registry": os.getenv('VOLUMEPERMISSIONS_IMAGE_REGISTRY', 'docker.io'),
                    "repository": os.getenv('VOLUMEPERMISSIONS_IMAGE_REPOSITORY', 'bitnami/bitnami-shell'),
                    "tag": os.getenv('VOLUMEPERMISSIONS_IMAGE_TAG', '10-debian-10-r367')
                },
                "resources": {}
            }
        },
        skip_await = True
    )
)

pulumi.export('app', app)
