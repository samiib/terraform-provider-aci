---
subcategory: "Contract"
layout: "aci"
page_title: "ACI: contract_interface"
sidebar_current: "docs-aci-resource-contract_interface"
description: |-
  Manages ACI Contract Interface
---

# aci_contract_interface #

Manages ACI Contract Interface

## API Information ##

* `Class` - fvRsConsIf

* `Distinguished Name Formats`
  - outdefcont-{id}/rsoutDefContToOut-[{tDn}]/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/epp/fv-[{epgPKey}]/ac-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/epp/rtd-[{epgPKey}]/ac-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/epp/sec-[{epgPKey}]/ac-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/ldev-[{priKey}]-ctx-[{ctxDn}]-bd-[{bdDn}]/rsconsIf-{tnVzCPIfName}
  - uni/svcdefcont/tensvcdef-{name}/ldevdef-[{lDevDn}]/sepg-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/LDevInst-[{priKey}]-ctx-{ctxName}/G-{graphRn}-N-{nodeRn}-C-{connRn}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/LDevInst-[{priKey}]-ctx-{ctxName}/bd-[{bdDn}]/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/LDevInst-[{priKey}]-ctx-{ctxName}/epgDn-[{shEpgDn}]/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/Tnlepg-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/acAnyToEp-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/acEpToAny-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/acEpToEp-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/acEpToExt-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/acExtToEp-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/acIpToIp-{name}/acl3outcont/outdef-{name}/instPdef-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/ap-{name}/epg-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/ap-{name}/esg-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/cloudapp-{name}/cloudepg-{name}/rsconsIf-{tnVzCPIfName}
  - uni/tn-{name}/cloudapp-{name}/cloudextepg-{name}/rsconsIf-{tnVzCPIfName}
  - Too many DN formats to display, see model documentation for all possible parents.

## GUI Information ##

* `Location` - Tenants -> ... determine for muliple DN formats ...

## Example Usage ##

```hcl

resource "aci_contract_interface" "example" {
  parent_dn               = aci_application_epg.example.id
  contract_interface_name = "test_contract_interface"
  annotations = [
    {
      key = "test_annotation"
    },
  ]
}

```

## Schema

### Required

* `parent_dn` - (string) The distinquised name (DN) of the parent object, possible resources:
  - `aci_application_epg` (class: fvAEPg)
  - The distinquised name (DN) of classes below can be used but currently there is no available resource for it:
    - `cloudEPg`
    - `cloudExtEPg`
    - `cloudISvcEPg`
    - `cloudSvcEPg`
    - `dhcpCRelPg`
    - `dhcpPRelPg`
    - `fvESg`
    - `fvTnlEPg`
    - `infraCEPg`
    - `infraPEPg`
    - `l2extInstP`
    - `l3extInstP`
    - `l3extInstPDef`
    - `mgmtInB`
    - `vnsEPpInfo`
    - `vnsREPpInfo`
    - `vnsSDEPpInfo`
    - `vnsSHEPpInfo` 
* `contract_interface_name` - (string) The contract interface name.

### Read-Only

* `id` - (string) The distinquised name (DN) of the Contract Interface object.

### Optional

* `annotation` - (string) The annotation of the Contract Interface object.
  - Default: `orchestrator:terraform`
* `priority` - (string) The contract interface priority.
  - Default: `unspecified`
  - Valid Values: `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.

* `annotations` - (list) A list of Annotation objects (tagAnnotation) which can also be configured using the `aci_annotation` resource.
  #### Required
  
  * `key` - (string) The key or password used to uniquely identify this configuration object.

  #### Optional
  
  * `value` - (string) The value of the property.

## Importing ##

An existing Contract Interface can be [imported](https://www.terraform.io/docs/import/index.html) into this resource via its distinquised name (DN), via the following command:

```
terraform import aci_contract_interface.example uni/tn-{name}/cloudapp-{name}/cloudsvcepg-{name}/rsconsIf-{tnVzCPIfName}
```

Starting in Terraform version 1.5, an existing BFD Multihop Node Policy can be imported 
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/cloudapp-{name}/cloudsvcepg-{name}/rsconsIf-{tnVzCPIfName}"
  to = aci_contract_interface.example
}
```