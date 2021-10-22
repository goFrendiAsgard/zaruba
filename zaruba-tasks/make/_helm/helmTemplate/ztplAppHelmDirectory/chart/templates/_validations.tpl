{{/*
Compile all warnings into a single message, and call fail.
*/}}
{{- define "app.validateValues" -}}
{{- $messages := list -}}
{{- $messages := append $messages (include "app.validateValues.nameGiven" .) -}}
{{- $messages := append $messages (include "app.validateValues.imageRepositoryGiven" .) -}}
{{- $messages := append $messages (include "app.validateValues.noSecretsData" .) -}}
{{- $messages := append $messages (include "app.validateValues.noConfigMapsData" .) -}}
{{- $messages := append $messages (include "app.validateValues.ingressHasPortsProvided" .) -}}
{{- $messages := without $messages "" -}}
{{- $message := join "\n" $messages -}}

{{- if $message -}}
{{-   printf "\nVALUES VALIDATION:\n%s" $message | fail -}}
{{- end -}}
{{- end -}}

{{/* Validate value name is given */}}
{{- define "app.validateValues.nameGiven" -}}
{{- if not .Values.name -}}
app: no-name
    You did not specify the application name for the generic app chart. Please
    set name.
{{- end -}}
{{- end -}}

{{/* Validate that the container repository name is given */}}
{{- define "app.validateValues.imageRepositoryGiven" -}}
{{- if and (not .Values.image.repository) (has .Values.appKind (list "Deployment" "StatefulSet" "DaemonSet")) -}}
app: no-image-repository
    You did not specify the application image repository. Please
    set image.repository.
{{- end -}}
{{- end -}}

{{/* Validate that the container has command or args given */}}
{{- define "app.validateValues.containerCommandOrArgsGiven" -}}
{{- if not (or .Values.command .Values.args) -}}
app: no-command-or-args
    You did not specify command or args for the application default
    container. Please set command or args.
{{- end -}}
{{- end -}}

{{/* Validate that secrets have data or stringData given */}}
{{- define "app.validateValues.noSecretsData" -}}
{{- range $_, $values := .Values.secrets -}}
{{- if not (or $values.data $values.stringData) -}}
app: no-secrets-data-or-stringdata
    Each item of .secrets must have data or stringData field. Please
    check the input configuration.
{{- end -}}
{{- end -}}
{{- end -}}

{{/* Validate that configMaps have data given */}}
{{- define "app.validateValues.noConfigMapsData" -}}
{{- range $_, $values := .Values.configMaps -}}
{{- if not $values.data -}}
app: no-configmaps-data
    Each item of .configMaps must have data field. Please
    check the input configuration.
{{- end -}}
{{- end -}}
{{- end -}}

{{/* Validate that the enabled ingress has service ports provided */}}
{{- define "app.validateValues.ingressHasPortsProvided" -}}
{{- if and .Values.ingress.enabled (not (or .Values.service.port .Values.service.ports)) -}}
app: no-service-ports-for-ingress
    You enabled ingress, but did not specify service port or ports. Please set
    service.port or service.ports.
{{- end -}}
{{- end -}}
