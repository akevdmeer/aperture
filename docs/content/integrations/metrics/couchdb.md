---
title: CouchDB
description: Integrating CouchDB Metrics
keywords:
  - couchdb
  - otel
  - opentelemetry
  - collector
  - metrics
---

:::info

See also [couchdbreceiver docs][receiver] in `opentelemetry-collector-contrib`
repository.

:::

:::note

The `couchdbreceiver` extension is available in the default agent image. If
you're [building][build] your own Aperture Agent, add
`integrations/otel/couchdbreceiver` to the `bundled_extensions` list to make
[the receiver][receiver] available.

:::

You can configure the [OpenTelemetry Collector][opentelemetry-collector] for
CouchDB as part of [Policy resources][policy-resources] while [applying the
policy][applying-policy]:

```yaml
policy:
  resources:
    infra_meters:
      couchdb:
        agent_group: default
        per_agent_group: true
        receivers:
          couchdb: [couchdbreceiver configuration here]
```

[build]: /reference/aperturectl/build/agent/agent.md
[receiver]:
  https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/couchdbreceiver
[opentelemetry-collector]: /reference/configuration/spec.md#telemetry-collector
[applying-policy]: /use-cases/use-cases.md
[policy-resources]: /reference/configuration/spec.md#resources
