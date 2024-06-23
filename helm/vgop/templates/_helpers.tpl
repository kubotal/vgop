

{{/*
Expand the name of the chart.
*/}}
{{- define "vgop.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "vgop.fullname" -}}
{{- if .Values.fullNameOverride }}
{{- .Values.fullNameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "vgop.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "vgop.selectorLabels" -}}
app.kubernetes.io/name: {{ include "vgop.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
app.kubernetes.io/component: node
{{- end }}


{{/*
Common labels
*/}}
{{- define "vgop.labels" -}}
helm.sh/chart: {{ include "vgop.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}



{{/*
Create the name of the deployment
*/}}
{{- define "vgop.deploymentName" -}}
{{- default (printf "%s" (include "vgop.fullname" .)) .Values.deploymentName }}
{{- end }}


{{/*
Create the name of the configuration configmap
*/}}
{{- define "vgop.configmapName" -}}
{{- default (printf "%s" (include "vgop.fullname" .)) .Values.configmapName }}
{{- end }}


{{/*
Create the name of the service account to use
*/}}
{{- define "vgop.serviceAccountName" -}}
{{- default (printf "%s" (include "vgop.fullname" .)) .Values.serviceAccountName }}
{{- end }}

{{/*
Create the name of the associated role
*/}}
{{- define "vgop.roleName" -}}
{{- default (printf "%s" (include "vgop.fullname" .)) .Values.roleName }}
{{- end }}


{{/*
Create the name of the associated clusterRole
*/}}
{{- define "vgop.clusterRoleName" -}}
{{- default (printf "%s" (include "vgop.fullname" .)) .Values.clusterRoleName }}
{{- end }}
