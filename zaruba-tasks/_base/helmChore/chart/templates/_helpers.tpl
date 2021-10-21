{{/* vim: set filetype=mustache: */}}

{{/*
Return the proper App image name
*/}}
{{- define "app.image" -}}
{{ include "common.images.image" (dict "imageRoot" .Values.image "global" .Values.global) }}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "app.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "common.names.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{/*
Return the proper Docker Image Registry Secret Names
*/}}
{{- define "app.imagePullSecrets" -}}
{{- include "common.images.pullSecrets" (dict "images" (list .Values.image) "global" .Values.global) -}}
{{- end -}}

{{/*
Set App PVC
*/}}
{{- define "app.pvc" -}}
{{- .Values.persistence.existingClaim | default (include "common.names.fullname" .) -}}
{{- end -}}

{{/*
Return  the proper Storage Class
*/}}
{{- define "app.storageClass" -}}
{{- include "common.storage.class" (dict "persistence" .Values.persistence "global" .Values.global) -}}
{{- end -}}

{{/*
Choose port name based on .port .targetPort. Fallbacks to "app".

Examples:
  {{- include "app.service.defaultPortName" (dict "port" "80" "targetPort" "8080") }}
  {{- include "app.service.defaultPortName" (dict "port" "80" "targetPort" "api") }}
*/}}
{{- define "app.service.defaultPortName" -}}
{{- $targetPort := .targetPort | default .port -}}
{{- if regexMatch "^[0-9]+$" ($targetPort | toString) -}}
app
{{- else -}}
{{ $targetPort }}
{{- end -}}
{{- end -}}
