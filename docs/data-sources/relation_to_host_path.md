---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Generic"
layout: "aci"
page_title: "ACI: aci_relation_to_host_path"
sidebar_current: "docs-aci-data-source-aci_relation_to_host_path"
description: |-
  Data source for Relation To Host Path
---

# aci_relation_to_host_path #

Data source for Relation To Host Path

## API Information ##

* Class: [infraRsHPathAtt](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/infraRsHPathAtt/overview)

* Supported in ACI versions: 1.1(1j) and later.

* Distinguished Name Format: `uni/infra/hpaths-{name}/rsHPathAtt-[{tDn}]`

## GUI Information ##

* Location: `Generic`

## Example Usage ##

```hcl

data "aci_relation_to_host_path" "example_access_interface_override" {
  parent_dn = aci_access_interface_override.example.id
  target_dn = "topology/pod-1/paths-101/pathep-[eth1/1]"
}

```

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_access_interface_override](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/access_interface_override) ([infraHPathS](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/infraHPathS/overview))
* `target_dn` (tDn) - (string) The distinguished name of the target.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Relation To Host Path object.
* `annotation` (annotation) - (string) The annotation of the Relation To Host Path object.

* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.

* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
  * `key` (key) - (string) The key used to uniquely identify this configuration object.
  * `value` (value) - (string) The value of the property.