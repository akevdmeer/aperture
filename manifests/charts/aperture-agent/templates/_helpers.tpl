{{/*
Create the name of the service account to use
*/}}
{{- define "agent-operator.serviceAccountName" -}}
{{- default ( print (include "common.names.fullname" .) "-operator" ) .Values.operator.serviceAccount.name }}
{{- end }}

{{/*
Compile all warnings into a single message.
*/}}
{{- define "agent-operator.validateValues" -}}
{{- $messages := list -}}
{{- $messages := without $messages "" -}}
{{- $message := join "\n" $messages -}}

{{- if $message -}}
{{-   printf "\nVALUES VALIDATION:\n%s" $message -}}
{{- end -}}
{{- end -}}


{{/*
  Azure Agent helper template
*/}}
{{- define "azure.image.helper" -}}
{{- $azureImage := "" -}}
{{- if  .operator -}}
    {{- $azureImage = (printf "%s/%s@%s" .context.Values.global.azure.images.operator.registry  .context.Values.global.azure.images.operator.image .context.Values.global.azure.images.operator.digest) -}}
{{- else -}}
    {{- $azureImage = (printf "%s/%s@%s" .context.Values.global.azure.images.agent.registry  .context.Values.global.azure.images.agent.image .context.Values.global.azure.images.agent.digest) -}}
{{- end -}}
{{ print $azureImage }}
{{- end -}}

{{/*
   Agent helper template
*/}}
{{- define "agent.image.helper" -}}
    {{- $tag := get .image "tag" -}}
    {{- $newImage := .image -}}
    {{- if (not $tag) -}}
        {{- $tag = trimPrefix "v" .context.Chart.AppVersion -}}
    {{- end -}}
    {{- $_ := set $newImage "tag" $tag -}}
    {{ print (include "common.images.image" (dict "imageRoot" $newImage "global" .context.Values.global)) }}
{{- end -}}


{{/*
Get image tag for operator.
{{ include "agent-operator.image" ( dict "image" .Values.path.to.the.image "context" $.context $) }}
*/}}
{{- define "agent-operator.image" -}}
{{- $globalAzure := get .context.Values.global "azure" -}}
{{- if not (empty $globalAzure) -}}
    {{ print (include "azure.image.helper" (dict "operator" .operator "context"  .context)) }}
{{- else -}}
    {{ print (include "agent.image.helper" (dict "image" .image "context"  .context)) }}
{{- end -}}
{{- end -}}
{{/*
Create the endpoint of the etcd for Aperture Agent
{{ include "agent.etcd.endpoints" ( dict "etcd" .Values.path.to.the.etcd "context" $.context $) }}
*/}}
{{- define "agent.etcd.endpoints" -}}
{{- $endpoints := list -}}
{{ $endpoints = without .etcd.endpoints "" }}
{{- if empty $endpoints -}}
    {{- fail "Value for etcd endpoints of Agent cannot be empty." -}}
{{- end -}}
{{ print $endpoints }}
{{- end -}}

{{/*
Create the address of the Prometheus for Aperture Agent
{{ include "agent.prometheus.address" ( dict "prometheus" .Values.path.to.the.prometheus "context" $.context $) }}
*/}}
{{- define "agent.prometheus.address" -}}
{{- if .prometheus.address -}}
    {{ print .prometheus.address }}
{{- else -}}
    {{- fail "Value for prometheus address of Agent cannot be empty." -}}
{{- end -}}
{{- end -}}

{{/*
Fetch the value of the API Key secret for Aperture Agent
{{ include "agent.apiSecret.value" ( dict "agent" .Values.path.to.the.agent $) }}
*/}}
{{- define "agent.apiSecret.value" -}}
{{- if .agent.secrets.fluxNinjaExtension.create -}}
    {{- if .agent.secrets.fluxNinjaExtension.value -}}
        {{ print .agent.secrets.fluxNinjaExtension.value }}
    {{- else -}}
        {{- fail "Value of API Key for Agent cannot be empty when .Values.agent.secrets.fluxNinjaExtension.create is set to true." -}}
    {{- end -}}
{{- else -}}
    {{ print "" }}
{{- end -}}
{{- end -}}

{{/*
Fetch the Name of the API Key secret for Aperture Agent
{{ include "agent.apiSecret.name" ( dict "agent" .Values.path.to.the.agent "context" $.context $ ) }}
*/}}
{{- define "agent.apiSecret.name" -}}
{{- if .agent.secrets.fluxNinjaExtension.secretKeyRef.name -}}
    {{ print .agent.secrets.fluxNinjaExtension.secretKeyRef.name }}
{{- else -}}
    {{ print "%s-agent-apikey" .context.Release.Name }}
{{- end -}}
{{- end -}}

{{/*
Fetch the Key of the API Key secret for Aperture Agent
{{ include "agent.apiSecret.key" ( dict "agent" .Values.path.to.the.agent $ ) }}
*/}}
{{- define "agent.apiSecret.key" -}}
{{- if .agent.secrets.fluxNinjaExtension.secretKeyRef.key -}}
    {{ print .agent.secrets.fluxNinjaExtension.secretKeyRef.key }}
{{- else -}}
    {{ print "apiKey" }}
{{- end -}}
{{- end -}}

{{/*
Fetch the server port of the Aperture Agent
{{ include "agent.server.port" ( dict "agent" .Values.path.to.the.agent $ ) }}
*/}}
{{- define "agent.server.port" -}}
{{- if and .agent.config .agent.config.server .agent.config.server.listener .agent.config.server.listener.addr -}}
    {{ print (split ":" .agent.config.server.listener.addr)._1 }}
{{- else -}}
    {{ print "8080" }}
{{- end -}}
{{- end -}}

{{/*
Fetch the OTEL port of the Aperture Agent
{{ include "agent.otel.port" ( dict "agent" .Values.path.to.the.agent portName string defaultPort string $ ) }}
*/}}
{{- define "agent.otel.port" -}}
{{- if and .agent.config .agent.config.otel .agent.config.otel.ports (hasKey .agent.config.otel.ports .portName) -}}
    {{ print (get .agent.config.otel.ports .portName) }}
{{- else -}}
    {{ print .defaultPort }}
{{- end -}}
{{- end -}}

{{/*
Fetch the Distcache port of the Aperture Agent
{{ include "agent.dist_cache.port" ( dict "agent" .Values.path.to.the.agent portName string defaultPort string $ ) }}
*/}}
{{- define "agent.dist_cache.port" -}}
{{- if and .agent.config .agent.config.dist_cache (hasKey .agent.config.dist_cache .portName) -}}
    {{ print (split ":" .agent.config.dist_cache)._0 }}
{{- else -}}
    {{ print .defaultPort }}
{{- end -}}
{{- end -}}

{{/*
Add the pod labels when global azure field is set
*/}}
{{- define "agent.podlabels" -}}
{{- $globalAzure := get .context.Values.global "azure" -}}
{{- $podLabels := "" -}}
{{- if not (empty $globalAzure) -}}
    {{- $podLabels = (printf "%s: %s" "azure-extensions-usage-release-identifier" .context.Release.Name ) -}}
    {{ print $podLabels | toYaml | nindent 4 | replace "'" ""}}
{{- else -}}
  {{- if not (empty .podlabels) -}}
    {{- $podLabels = .podlabels | toYaml | nindent 4 -}}
    {{ print $podLabels }}
  {{- end -}}
{{- end -}}
{{- end -}}
