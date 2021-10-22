{{/* vim: set filetype=mustache: */}}

{{/*
Ingress annotations
*/}}
{{- define "app.ingress.annotations" -}}
{{- if .Values.ingress.certManager }}
kubernetes.io/tls-acme: "true"
{{- end }}

{{- with .Values.ingress.nginx.configurationSnippet }}
nginx.ingress.kubernetes.io/configuration-snippet: |
{{- include "common.tplvalues.render" (dict "value" . "context" $) | nindent 2 }}
{{- end }}

{{- with .Values.ingress.nginx.serverSnippet }}
nginx.ingress.kubernetes.io/server-snippet: |
{{- include "common.tplvalues.render" (dict "value" . "context" $) | nindent 2 }}
{{- end }}
{{- end -}}
