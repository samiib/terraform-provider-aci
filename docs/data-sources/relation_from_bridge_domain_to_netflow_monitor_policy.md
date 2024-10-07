---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Generic"
layout: "aci"
page_title: "ACI: aci_relation_from_bridge_domain_to_netflow_monitor_policy"
sidebar_current: "docs-aci-data-source-aci_relation_from_bridge_domain_to_netflow_monitor_policy"
description: |-
  Data source for ACI Relation From Bridge Domain To NetFlow Monitor Policy
---

# aci_relation_from_bridge_domain_to_netflow_monitor_policy #

Data source for ACI Relation From Bridge Domain To NetFlow Monitor Policy

## API Information ##

* Class: [fvRsBDToNetflowMonitorPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsBDToNetflowMonitorPol/overview)

* Supported in ACI versions: 2.2(1k) and later.

* Distinguished Name Formats:
  - `uni/tn-{name}/BD-{name}/rsBDToNetflowMonitorPol-[{tnNetflowMonitorPolName}]-{fltType}`
  - `uni/tn-{name}/rsBDToNetflowMonitorPol-[{tnNetflowMonitorPolName}]-{fltType}`
  - `uni/tn-{name}/svcBD-{name}/rsBDToNetflowMonitorPol-[{tnNetflowMonitorPolName}]-{fltType}`

## GUI Information ##

* Location: `Generic`

## Example Usage ##

```hcl

data "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "example_bridge_domain" {
  parent_dn                   = aci_bridge_domain.example.id
  filter_type                 = "ipv4"
  netflow_monitor_policy_name = aci_netflow_monitor_policy.example.name
}

```

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_bridge_domain](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/bridge_domain) ([fvBD](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvBD/overview))
  - The distinguished name (DN) of classes below can be used but currently there is no available resource for it:
    - [fvSvcBD](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvSvcBD/overview)

* `filter_type` (fltType) - (string) The filter type of the NetFlow Monitor Policy object.
  - Valid Values: `ce`, `ipv4`, `ipv6`, `unspecified`.
* `netflow_monitor_policy_name` (tnNetflowMonitorPolName) - (string) The name of the NetFlow Monitor Policy object. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/netflow_monitor_policy) with `aci_netflow_monitor_policy.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/netflow_monitor_policy) with `data.aci_netflow_monitor_policy.example.name`.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Relation From Bridge Domain To NetFlow Monitor Policy object.
* `annotation` (annotation) - (string) The annotation of the Relation From Bridge Domain To NetFlow Monitor Policy object.

* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.

* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.