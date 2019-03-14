## tutorial (based on helm official repo)

- Install helm (see the docs)

- Create a new helm chart

```shell
helm create bobba-api-chart
```

- install the chart

```shell
helm install ./bobba-api-chart
```

- Retrieve the deployed chart and see the actual template

```shell
helm get manifest <release-name>
```

- Delete a release

```shell
helm delete <release-name>
```

- Debug mode

```shell
helm install --debug --dry-run ./bobba-api-chart
```

> This will send the chart to the Tiller server, which will render the templates. But instead of installing the chart, it will return the rendered template to you so you can see the output:

> Using --dry-run will make it easier to test your code, but it won’t ensure that Kubernetes itself will accept the templates you generate. It’s best not to assume that your chart will install just because --dry-run works.

## Release naming convention

**Hard coding a release name is a bad practice**
**Name should be unique to release**

## Templating (based on Go templating)

- String: {{ .<Namesapce1>.<Namespace2> }}

## Overriding a custom value

```shell
helm install --debug --dry-run --set favoriteBubbleTea=ginger ./bobba-api-chart
```

## Override a default value

```shell
helm install stable/drupal --set image=my-registry/drupal:0.1.0 --set livenessProbe.exec.command=[cat,docroot/CHANGELOG.txt] --set livenessProbe.httpGet=null
```

## Multiline string

> The |- marker in YAML takes a multi-line string. This can be a useful technique for embedding big blocks of data inside of your manifests, as exemplified here. (use in loop for e.g)

## Tuple

> tuple is a list-like collection of fixed size, but with arbitrary data types. e.g

```yaml
sizes: |-
    {{- range tuple "small" "medium" "large" }}
    - {{ . }}
    {{- end }}
```

## Templates

Add an other templates example

use to include a partial configuration

```yaml
metadata:
  name: {{ .Release.Name }}-configmap
  {{- template "bobba-api-chart.drink" }}
```

```yaml
{{- define "bobba-api-chart.drink" }}
  labels:
    generator: helm
    drink: coconut
    date: {{ now | htmlDate }}
    chart: {{ .Chart.Name }}
{{- end }}
```

## Include

### Preferable to use include over template

> It is considered preferable to use include over template in Helm templates simply so that the output formatting can be handled better for YAML documents.

include can be used to include a partial configuration values like so.

```yaml
labels: {{- include "bobba-api-chart.drink" . | nindent 4 }}
```

```yaml
{{- define "bobba-api-chart.drink" -}}
generator: helm
drink: coconut
date: {{ now | htmlDate }}
chart: {{ .Chart.Name }}
{{- end -}}
```

With special values file

```shell
helm install --debug --dry-run -f bobba-helm-chart/values-api.yaml ./bobba-helm-chart
```

### Upgrade

```shell
helm upgrade -f bobba-helm-chart/values-api.yaml broken-pronghorn bobba-helm-chart
```

### Rollback

```shell
helm history <release>

helm rollback broken-pronghorn <number_release>
```