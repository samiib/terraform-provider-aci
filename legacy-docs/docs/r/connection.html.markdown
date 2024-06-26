---
subcategory: "L4-L7 Services"
layout: "aci"
page_title: "ACI: aci_connection"
sidebar_current: "docs-aci-resource-aci_connection"
description: |-
  Manages ACI Connection
---

# aci_connection

Manages ACI Connection

## Example Usage

```hcl
resource "aci_connection" "conn2" {
  l4_l7_service_graph_template_dn  = aci_l4_l7_service_graph_template.example.id
  name  = "conn2"
  adj_type  = "L3"
  description = "from terraform"
  annotation  = "example"
  conn_dir  = "consumer"
  conn_type  = "internal"
  direct_connect  = "yes"
  name_alias  = "example"
  unicast_route  = "yes"
  relation_vns_rs_abs_connection_conns = [
    aci_l4_l7_service_graph_template.example.term_cons_dn,
    aci_function_node.example.conn_consumer_dn
  ]
}
```

## Argument Reference

- `l4_l7_service_graph_template_dn` - (Required) Distinguished name of parent L4-L7 Service Graph Template object. Type: String.
- `name` - (Required) Name of object connection. Type: String.
    - The valid connection name format for cloud APICs is `CONX`, where X is a number starting with 0.
    - The valid connection name format for on-prem APICs is `CX`, where X is a number starting with 1.
- `adj_type` - (Optional) Connector adjacency type. Allowed values are "L2", "L3". Default value is "L2". Type: String.
- `annotation` - (Optional) Annotation for object connection. Type: String.
- `description` - (Optional) Description for object connection. Type: String.
- `conn_dir` - (Optional) Connection directory for object connection. Allowed values are "consumer", "provider". Default value is "provider". Type: String.
- `conn_type` - (Optional) Connection type of connection object. Allowed values are "external", "internal". Default value is "external". Type: String.
- `direct_connect` - (Optional) Direct connect for object connection. Allowed values are "yes" and "no". Default value is "no". Type: String.
- `name_alias` - (Optional) Name alias for object connection. Type: String.
- `unicast_route` - (Optional) Unicast route for connection object. Unicast route setting should be true for L3 connections. Allowed values are "yes" and "no". Default value is "yes". Type: String.

- `relation_vns_rs_abs_copy_connection` - (Optional) A list of relation to class vnsAConn. Cardinality - ONE_TO_M. Type: List.
- `relation_vns_rs_abs_connection_conns` - (Optional) A list of relation to class vnsAConn. Cardinality - ONE_TO_M. Type: List.

## Attribute Reference

The only attribute that this resource exports is the `id`, which is set to the
Dn of the Connection.
## Importing

An existing Connection can be [imported][docs-import] into this resource via its Dn, via the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html

```
terraform import aci_connection.example <Dn>
```
