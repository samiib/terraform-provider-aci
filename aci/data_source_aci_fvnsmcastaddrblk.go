package aci

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAciAbstractionofIPAddressBlock() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceAciAbstractionofIPAddressBlockRead,
		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"multicast_pool_dn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"from": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"to": {
				Type:     schema.TypeString,
				Required: true,
			},
		})),
	}
}

func dataSourceAciAbstractionofIPAddressBlockRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	aciClient := m.(*client.Client)
	from := d.Get("from").(string)
	to := d.Get("to").(string)
	MulticastAddressPoolDn := d.Get("multicast_pool_dn").(string)
	rn := fmt.Sprintf(models.RnfvnsMcastAddrBlk, from, to)
	dn := fmt.Sprintf("%s/%s", MulticastAddressPoolDn, rn)

	fvnsMcastAddrBlk, err := getRemoteAbstractionofIPAddressBlock(aciClient, dn)
	if err != nil {
		return nil
	}

	d.SetId(dn)

	_, err = setAbstractionofIPAddressBlockAttributes(fvnsMcastAddrBlk, d)
	if err != nil {
		return nil
	}

	return nil
}
