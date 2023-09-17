
resource "aci_pim_route_map_policy" "example" {
  parent_dn = aci_tenant.example.id
  name      = "test_name"
  annotations = [
    {
      key   = "test_key"
      value = "test_value"
    },
  ]
}

