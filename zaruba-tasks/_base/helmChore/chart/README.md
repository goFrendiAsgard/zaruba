# Generic application Helm chart

![Version: 0.3.0](https://img.shields.io/badge/Version-0.3.0-informational?style=flat-square)

Generic application Helm chart

This chart is a humble effort to provide a generic chart which can be a ready-to-go solution for deployments lacking their own Helm charts.

## Introduction

The app chart allows to perform basically a generic application deployment. At least you need to specify `name` and the `image.repository` to perform the deployment. Evidently such an installation doesn't bring a lot value, since there will only a Deployment created without Service, ConfigMaps etc.

The charts allows you to manage the following resources:

  - Deployment
  - Service
  - Secrets
  - ConfigMaps
  - Ingress
  - HorizontalPodAutoscaler
  - PersistentVolumeClaim
  - ServiceAccount

## TL;DR

```console
$ helm repo add dysnix https://dysnix.github.io/charts
$ helm install my-release dysnix/app
```

## Examples

### Install a generic app example

```yaml
name: myapp
image:
  repository: nginx

podAnnotations:
  hello/world: "true"

checksums:
  - /secrets.yaml

secrets:
  env:
    stringData:
      HELLO: world

configMaps:
  my.conf:
    data: |
      hello: world
      foo: coo

envFrom:
  - secretRef:
      name: '{{ template "app.fullname" . }}-env'

# containerPorts:
#   - name: web
#     containerPort: 8090

service:
  port: 8080
  targetPort: 80

  # ports:
  #   - name: web
  #     port: 9000
  #     targetPort: web

ingress:
  enabled: true
  hostname: my.tld
  certManager: true
  tls: true

  nginx:
    serverSnippet: |
      ## Redirect mobile
      set $agentflag 0;
      if ($http_user_agent ~* "(Mobile)" ){
        set $agentflag 1;
      }
      if ( $agentflag = 1 ) {
        return 301 https://m.example.com;
      }

```

## Source Code

* <https://github.com/dysnix/charts>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| https://charts.bitnami.com/bitnami | common | 1.x.x |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| args | list | `[]` |  |
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `10` |  |
| autoscaling.minReplicas | int | `1` |  |
| checksums | list | `[]` |  |
| command | list | `[]` |  |
| commonAnnotations | object | `{}` |  |
| commonLabels | object | `{}` |  |
| configMaps | object | `{}` |  |
| containerPorts | list | `[]` |  |
| containerSecurityContext.enabled | bool | `false` |  |
| containerSecurityContext.runAsUser | int | `1001` |  |
| dnsPolicy | string | `"ClusterFirst"` |  |
| env | object | `{}` |  |
| envFrom | list | `[]` |  |
| hostAliases | list | `[]` |  |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.registry | string | `""` |  |
| image.repository | string | `""` |  |
| image.tag | string | `"latest"` |  |
| ingress.annotations | object | `{}` |  |
| ingress.apiVersion | string | `nil` |  |
| ingress.certManager | bool | `false` |  |
| ingress.enabled | bool | `false` |  |
| ingress.hostname | string | `"app.local"` |  |
| ingress.nginx.configurationSnippet | string | `nil` |  |
| ingress.nginx.serverSnippet | string | `nil` |  |
| ingress.path | string | `"/"` |  |
| ingress.pathType | string | `"ImplementationSpecific"` |  |
| ingress.tls | bool | `false` |  |
| initContainers | list | `[]` |  |
| kind | string | `"Deployment"` |  |
| livenessProbe.enabled | bool | `false` |  |
| livenessProbe.failureThreshold | int | `6` |  |
| livenessProbe.initialDelaySeconds | int | `60` |  |
| livenessProbe.periodSeconds | int | `10` |  |
| livenessProbe.successThreshold | int | `1` |  |
| livenessProbe.timeoutSeconds | int | `5` |  |
| name | string | `nil` |  |
| nodeAffinityPreset.key | string | `""` |  |
| nodeAffinityPreset.type | string | `""` |  |
| nodeAffinityPreset.values | list | `[]` |  |
| nodeSelector | object | `{}` |  |
| persistence.accessMode | string | `"ReadWriteOnce"` |  |
| persistence.enabled | bool | `false` |  |
| persistence.mountPath | string | `"/data"` |  |
| persistence.size | string | `"10Gi"` |  |
| podAffinityPreset | string | `""` |  |
| podAnnotations | object | `{}` |  |
| podAntiAffinityPreset | string | `"soft"` |  |
| podLabels | object | `{}` |  |
| podSecurityContext.enabled | bool | `false` |  |
| podSecurityContext.fsGroup | int | `1001` |  |
| priorityClassName | string | `""` |  |
| readinessProbe.enabled | bool | `false` |  |
| readinessProbe.failureThreshold | int | `6` |  |
| readinessProbe.initialDelaySeconds | int | `60` |  |
| readinessProbe.periodSeconds | int | `10` |  |
| readinessProbe.successThreshold | int | `1` |  |
| readinessProbe.timeoutSeconds | int | `5` |  |
| resources | object | `{}` |  |
| secrets | object | `{}` |  |
| service.annotations | object | `{}` |  |
| service.labels | object | `{}` |  |
| service.ports | list | `[]` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `false` |  |
| serviceAccount.name | string | `nil` |  |
| tests.httpChecks.default | bool | `false` |  |
| tolerations | list | `[]` |  |
| topologySpreadConstraints | list | `[]` |  |
| updateStrategy.type | string | `"RollingUpdate"` |  |
| volumeMounts | list | `[]` |  |
| volumes | list | `[]` |  |


Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. For example,

```console
helm install my-release -f values.yaml dysnix/app
```

> **Tip**: You can use the default [values.yaml](values.yaml)

## Configuration and installation details

### Ingress

This chart provides support for Ingress resources. If an Ingress controller, such as [nginx-ingress](https://kubeapps.com/charts/stable/nginx-ingress) or [traefik](https://kubeapps.com/charts/stable/traefik), that Ingress controller can be used to serve app.

To enable Ingress integration, set `ingress.enabled` to `true`. The `ingress.hostname` property can be used to set the host name. The `ingress.tls` parameter can be used to add the TLS configuration for this host. It is also possible to have more than one host, with a separate TLS configuration for each host.

### Environment variables

In case you want to add extra environment variables (useful for advanced operations like custom init scripts), you can use the `env` property.

```yaml
env:
  - name: LOG_LEVEL
    value: error
```

Alternatively, you can use a ConfigMap or a Secret with the environment variables. To do so, use the `extraEnvVarsCM` or the `extraEnvVarsSecret` values.

### Pod affinity

This chart allows you to set your custom affinity using the `affinity` parameter. Find more information about Pod affinity in the [kubernetes documentation](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity).

As an alternative, use one of the preset configurations for pod affinity, pod anti-affinity, and node affinity available at the [bitnami/common](https://github.com/bitnami/charts/tree/master/bitnami/common#affinities) chart. To do so, set the `podAffinityPreset`, `podAntiAffinityPreset`, or `nodeAffinityPreset` parameters.
