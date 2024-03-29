import pulumi
from pulumi_kubernetes.helm.v3 import Chart, ChartOpts, LocalChartOpts
import os

app = Chart(
    'ztpl-deployment-name', 
    config=LocalChartOpts(
        path = './chart',
        namespace = os.getenv('NAMESPACE', 'default'),
        values = {
            "advancedConfiguration": os.getenv('ADVANCEDCONFIGURATION', ''),
            "affinity": {},
            "args": [],
            "auth": {
                "erlangCookie": os.getenv('AUTH_ERLANGCOOKIE', ''),
                "existingErlangSecret": os.getenv('AUTH_EXISTINGERLANGSECRET', ''),
                "existingPasswordSecret": os.getenv('AUTH_EXISTINGPASSWORDSECRET', ''),
                "password": os.getenv('AUTH_PASSWORD', ''),
                "tls": {
                    "autoGenerated": os.getenv('AUTH_TLS_AUTOGENERATED', 'False') == 'True',
                    "caCertificate": os.getenv('AUTH_TLS_CACERTIFICATE', ''),
                    "enabled": os.getenv('AUTH_TLS_ENABLED', 'False') == 'True',
                    "existingSecret": os.getenv('AUTH_TLS_EXISTINGSECRET', ''),
                    "existingSecretFullChain": os.getenv('AUTH_TLS_EXISTINGSECRETFULLCHAIN', 'False') == 'True',
                    "failIfNoPeerCert": os.getenv('AUTH_TLS_FAILIFNOPEERCERT', 'True') == 'True',
                    "serverCertificate": os.getenv('AUTH_TLS_SERVERCERTIFICATE', ''),
                    "serverKey": os.getenv('AUTH_TLS_SERVERKEY', ''),
                    "sslOptionsVerify": os.getenv('AUTH_TLS_SSLOPTIONSVERIFY', 'verify_peer')
                },
                "username": os.getenv('AUTH_USERNAME', 'user')
            },
            "clusterDomain": os.getenv('CLUSTERDOMAIN', 'cluster.local'),
            "clustering": {
                "addressType": os.getenv('CLUSTERING_ADDRESSTYPE', 'hostname'),
                "enabled": os.getenv('CLUSTERING_ENABLED', 'True') == 'True',
                "forceBoot": os.getenv('CLUSTERING_FORCEBOOT', 'False') == 'True',
                "partitionHandling": os.getenv('CLUSTERING_PARTITIONHANDLING', 'autoheal'),
                "rebalance": os.getenv('CLUSTERING_REBALANCE', 'False') == 'True'
            },
            "command": [],
            "commonAnnotations": {},
            "communityPlugins": os.getenv('COMMUNITYPLUGINS', ''),
            "containerSecurityContext": {},
            "customLivenessProbe": {},
            "customReadinessProbe": {},
            "customStartupProbe": {},
            "diagnosticMode": {
                "args": [
                    os.getenv('DIAGNOSTICMODE_ARGS_0', 'infinity')
                ],
                "command": [
                    os.getenv('DIAGNOSTICMODE_COMMAND_0', 'sleep')
                ],
                "enabled": os.getenv('DIAGNOSTICMODE_ENABLED', 'False') == 'True'
            },
            "extraContainerPorts": [],
            "extraDeploy": [],
            "extraEnvVars": [],
            "extraEnvVarsCM": os.getenv('EXTRAENVVARSCM', ''),
            "extraEnvVarsSecret": os.getenv('EXTRAENVVARSSECRET', ''),
            "extraPlugins": os.getenv('EXTRAPLUGINS', 'rabbitmq_auth_backend_ldap'),
            "extraSecrets": {},
            "extraSecretsPrependReleaseName": os.getenv('EXTRASECRETSPREPENDRELEASENAME', 'False') == 'True',
            "extraVolumeMounts": [],
            "extraVolumes": [],
            "fullnameOverride": os.getenv('FULLNAMEOVERRIDE', ''),
            "global": {
                "imagePullSecrets": [],
                "imageRegistry": os.getenv('GLOBAL_IMAGEREGISTRY', ''),
                "storageClass": os.getenv('GLOBAL_STORAGECLASS', '')
            },
            "hostAliases": [],
            "image": {
                "debug": os.getenv('IMAGE_DEBUG', 'False') == 'True',
                "pullPolicy": os.getenv('IMAGE_PULLPOLICY', 'IfNotPresent'),
                "pullSecrets": [],
                "registry": os.getenv('IMAGE_REGISTRY', 'docker.io'),
                "repository": os.getenv('IMAGE_REPOSITORY', 'bitnami/rabbitmq'),
                "tag": os.getenv('IMAGE_TAG', '3.9.14-debian-10-r0')
            },
            "ingress": {
                "annotations": {},
                "enabled": os.getenv('INGRESS_ENABLED', 'False') == 'True',
                "extraHosts": [],
                "extraRules": [],
                "extraTls": [],
                "hostname": os.getenv('INGRESS_HOSTNAME', 'rabbitmq.local'),
                "ingressClassName": os.getenv('INGRESS_INGRESSCLASSNAME', ''),
                "path": os.getenv('INGRESS_PATH', '/'),
                "pathType": os.getenv('INGRESS_PATHTYPE', 'ImplementationSpecific'),
                "secrets": [],
                "selfSigned": os.getenv('INGRESS_SELFSIGNED', 'False') == 'True',
                "tls": os.getenv('INGRESS_TLS', 'False') == 'True'
            },
            "initContainers": [],
            "kubeVersion": os.getenv('KUBEVERSION', ''),
            "ldap": {
                "enabled": os.getenv('LDAP_ENABLED', 'False') == 'True',
                "port": int(os.getenv('LDAP_PORT', '389')),
                "servers": [],
                "tls": {
                    "enabled": os.getenv('LDAP_TLS_ENABLED', 'False') == 'True'
                },
                "user_dn_pattern": os.getenv('LDAP_USER_DN_PATTERN', 'cn=${username},dc=example,dc=org')
            },
            "livenessProbe": {
                "enabled": os.getenv('LIVENESSPROBE_ENABLED', 'True') == 'True',
                "failureThreshold": int(os.getenv('LIVENESSPROBE_FAILURETHRESHOLD', '6')),
                "initialDelaySeconds": int(os.getenv('LIVENESSPROBE_INITIALDELAYSECONDS', '120')),
                "periodSeconds": int(os.getenv('LIVENESSPROBE_PERIODSECONDS', '30')),
                "successThreshold": int(os.getenv('LIVENESSPROBE_SUCCESSTHRESHOLD', '1')),
                "timeoutSeconds": int(os.getenv('LIVENESSPROBE_TIMEOUTSECONDS', '20'))
            },
            "loadDefinition": {
                "enabled": os.getenv('LOADDEFINITION_ENABLED', 'False') == 'True',
                "existingSecret": os.getenv('LOADDEFINITION_EXISTINGSECRET', '')
            },
            "logs": os.getenv('LOGS', '-'),
            "maxAvailableSchedulers": os.getenv('MAXAVAILABLESCHEDULERS', ''),
            "memoryHighWatermark": {
                "enabled": os.getenv('MEMORYHIGHWATERMARK_ENABLED', 'False') == 'True',
                "type": os.getenv('MEMORYHIGHWATERMARK_TYPE', 'relative'),
                "value": os.getenv('MEMORYHIGHWATERMARK_VALUE', '0.4')
            },
            "metrics": {
                "enabled": os.getenv('METRICS_ENABLED', 'False') == 'True',
                "plugins": os.getenv('METRICS_PLUGINS', 'rabbitmq_prometheus'),
                "podAnnotations": {
                    "prometheus.io/scrape": os.getenv('METRICS_PODANNOTATIONS_PROMETHEUS_IO_SCRAPE', 'true')
                },
                "prometheusRule": {
                    "additionalLabels": {},
                    "enabled": os.getenv('METRICS_PROMETHEUSRULE_ENABLED', 'False') == 'True',
                    "namespace": os.getenv('METRICS_PROMETHEUSRULE_NAMESPACE', ''),
                    "rules": []
                },
                "serviceMonitor": {
                    "additionalLabels": {},
                    "enabled": os.getenv('METRICS_SERVICEMONITOR_ENABLED', 'False') == 'True',
                    "honorLabels": os.getenv('METRICS_SERVICEMONITOR_HONORLABELS', 'False') == 'True',
                    "interval": os.getenv('METRICS_SERVICEMONITOR_INTERVAL', '30s'),
                    "metricRelabelings": [],
                    "namespace": os.getenv('METRICS_SERVICEMONITOR_NAMESPACE', ''),
                    "path": os.getenv('METRICS_SERVICEMONITOR_PATH', ''),
                    "podTargetLabels": {},
                    "relabelings": [],
                    "relabellings": [],
                    "scrapeTimeout": os.getenv('METRICS_SERVICEMONITOR_SCRAPETIMEOUT', ''),
                    "targetLabels": {}
                }
            },
            "nameOverride": os.getenv('NAMEOVERRIDE', ''),
            "networkPolicy": {
                "additionalRules": [],
                "allowExternal": os.getenv('NETWORKPOLICY_ALLOWEXTERNAL', 'True') == 'True',
                "enabled": os.getenv('NETWORKPOLICY_ENABLED', 'False') == 'True'
            },
            "nodeAffinityPreset": {
                "key": os.getenv('NODEAFFINITYPRESET_KEY', ''),
                "type": os.getenv('NODEAFFINITYPRESET_TYPE', ''),
                "values": []
            },
            "nodeSelector": {},
            "onlineSchedulers": os.getenv('ONLINESCHEDULERS', ''),
            "pdb": {
                "create": os.getenv('PDB_CREATE', 'False') == 'True',
                "maxUnavailable": os.getenv('PDB_MAXUNAVAILABLE', ''),
                "minAvailable": int(os.getenv('PDB_MINAVAILABLE', '1'))
            },
            "persistence": {
                "accessMode": os.getenv('PERSISTENCE_ACCESSMODE', 'ReadWriteOnce'),
                "annotations": {},
                "enabled": os.getenv('PERSISTENCE_ENABLED', 'True') == 'True',
                "existingClaim": os.getenv('PERSISTENCE_EXISTINGCLAIM', ''),
                "mountPath": os.getenv('PERSISTENCE_MOUNTPATH', '/bitnami/rabbitmq/mnesia'),
                "selector": {},
                "size": os.getenv('PERSISTENCE_SIZE', '8Gi'),
                "storageClass": os.getenv('PERSISTENCE_STORAGECLASS', ''),
                "subPath": os.getenv('PERSISTENCE_SUBPATH', ''),
                "volumes": []
            },
            "plugins": os.getenv('PLUGINS', 'rabbitmq_management rabbitmq_peer_discovery_k8s'),
            "podAffinityPreset": os.getenv('PODAFFINITYPRESET', ''),
            "podAnnotations": {},
            "podAntiAffinityPreset": os.getenv('PODANTIAFFINITYPRESET', 'soft'),
            "podLabels": {},
            "podManagementPolicy": os.getenv('PODMANAGEMENTPOLICY', 'OrderedReady'),
            "podSecurityContext": {
                "enabled": os.getenv('PODSECURITYCONTEXT_ENABLED', 'True') == 'True',
                "fsGroup": int(os.getenv('PODSECURITYCONTEXT_FSGROUP', '1001')),
                "runAsUser": int(os.getenv('PODSECURITYCONTEXT_RUNASUSER', '1001'))
            },
            "priorityClassName": os.getenv('PRIORITYCLASSNAME', ''),
            "rbac": {
                "create": os.getenv('RBAC_CREATE', 'True') == 'True'
            },
            "readinessProbe": {
                "enabled": os.getenv('READINESSPROBE_ENABLED', 'True') == 'True',
                "failureThreshold": int(os.getenv('READINESSPROBE_FAILURETHRESHOLD', '3')),
                "initialDelaySeconds": int(os.getenv('READINESSPROBE_INITIALDELAYSECONDS', '10')),
                "periodSeconds": int(os.getenv('READINESSPROBE_PERIODSECONDS', '30')),
                "successThreshold": int(os.getenv('READINESSPROBE_SUCCESSTHRESHOLD', '1')),
                "timeoutSeconds": int(os.getenv('READINESSPROBE_TIMEOUTSECONDS', '20'))
            },
            "replicaCount": int(os.getenv('REPLICACOUNT', '1')),
            "resources": {
                "limits": {},
                "requests": {}
            },
            "schedulerName": os.getenv('SCHEDULERNAME', ''),
            "service": {
                "annotations": {},
                "annotationsHeadless": {},
                "distNodePort": os.getenv('SERVICE_DISTNODEPORT', ''),
                "distPort": int(os.getenv('SERVICE_DISTPORT', '25672')),
                "distPortEnabled": os.getenv('SERVICE_DISTPORTENABLED', 'True') == 'True',
                "distPortName": os.getenv('SERVICE_DISTPORTNAME', 'dist'),
                "epmdNodePort": os.getenv('SERVICE_EPMDNODEPORT', ''),
                "epmdPortEnabled": os.getenv('SERVICE_EPMDPORTENABLED', 'True') == 'True',
                "epmdPortName": os.getenv('SERVICE_EPMDPORTNAME', 'epmd'),
                "externalIPs": [],
                "externalTrafficPolicy": os.getenv('SERVICE_EXTERNALTRAFFICPOLICY', 'Cluster'),
                "extraPorts": [],
                "labels": {},
                "loadBalancerIP": os.getenv('SERVICE_LOADBALANCERIP', ''),
                "loadBalancerSourceRanges": [],
                "managerNodePort": os.getenv('SERVICE_MANAGERNODEPORT', ''),
                "managerPort": int(os.getenv('SERVICE_MANAGERPORT', '15672')),
                "managerPortEnabled": os.getenv('SERVICE_MANAGERPORTENABLED', 'True') == 'True',
                "managerPortName": os.getenv('SERVICE_MANAGERPORTNAME', 'http-stats'),
                "metricsNodePort": os.getenv('SERVICE_METRICSNODEPORT', ''),
                "metricsPort": int(os.getenv('SERVICE_METRICSPORT', '9419')),
                "metricsPortName": os.getenv('SERVICE_METRICSPORTNAME', 'metrics'),
                "nodePort": os.getenv('SERVICE_NODEPORT', ''),
                "port": int(os.getenv('SERVICE_PORT', '5672')),
                "portEnabled": os.getenv('SERVICE_PORTENABLED', 'True') == 'True',
                "portName": os.getenv('SERVICE_PORTNAME', 'amqp'),
                "tlsNodePort": os.getenv('SERVICE_TLSNODEPORT', ''),
                "tlsPort": int(os.getenv('SERVICE_TLSPORT', '5671')),
                "tlsPortName": os.getenv('SERVICE_TLSPORTNAME', 'amqp-ssl'),
                "type": os.getenv('SERVICE_TYPE', 'ClusterIP')
            },
            "serviceAccount": {
                "automountServiceAccountToken": os.getenv('SERVICEACCOUNT_AUTOMOUNTSERVICEACCOUNTTOKEN', 'True') == 'True',
                "create": os.getenv('SERVICEACCOUNT_CREATE', 'True') == 'True',
                "name": os.getenv('SERVICEACCOUNT_NAME', '')
            },
            "sidecars": [],
            "statefulsetLabels": {},
            "terminationGracePeriodSeconds": int(os.getenv('TERMINATIONGRACEPERIODSECONDS', '120')),
            "tolerations": [],
            "topologySpreadConstraints": [],
            "ulimitNofiles": int(os.getenv('ULIMITNOFILES', '65536')),
            "updateStrategyType": os.getenv('UPDATESTRATEGYTYPE', 'RollingUpdate'),
            "volumePermissions": {
                "enabled": os.getenv('VOLUMEPERMISSIONS_ENABLED', 'False') == 'True',
                "image": {
                    "pullPolicy": os.getenv('VOLUMEPERMISSIONS_IMAGE_PULLPOLICY', 'IfNotPresent'),
                    "pullSecrets": [],
                    "registry": os.getenv('VOLUMEPERMISSIONS_IMAGE_REGISTRY', 'docker.io'),
                    "repository": os.getenv('VOLUMEPERMISSIONS_IMAGE_REPOSITORY', 'bitnami/bitnami-shell'),
                    "tag": os.getenv('VOLUMEPERMISSIONS_IMAGE_TAG', '10-debian-10-r373')
                },
                "resources": {
                    "limits": {},
                    "requests": {}
                }
            }
        },
        skip_await = True
    )
)

pulumi.export('app', app)
