
resource "aci_epg_useg_ip_attribute" "example_epg_useg_block_statement" {
  parent_dn = aci_epg_useg_block_statement.example.id
  ip        = "131.107.1.200"
  name      = "131"
}