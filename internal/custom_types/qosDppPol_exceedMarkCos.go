package customTypes

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// QosDppPolExceedMarkCos custom string type.

var _ basetypes.StringTypable = QosDppPolExceedMarkCosStringType{}

type QosDppPolExceedMarkCosStringType struct {
	basetypes.StringType
}

func (t QosDppPolExceedMarkCosStringType) Equal(o attr.Type) bool {
	other, ok := o.(QosDppPolExceedMarkCosStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t QosDppPolExceedMarkCosStringType) String() string {
	return "QosDppPolExceedMarkCosStringType"
}

func (t QosDppPolExceedMarkCosStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	value := QosDppPolExceedMarkCosStringValue{
		StringValue: in,
	}

	return value, nil
}

func (t QosDppPolExceedMarkCosStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t QosDppPolExceedMarkCosStringType) ValueType(ctx context.Context) attr.Value {
	return QosDppPolExceedMarkCosStringValue{}
}

// QosDppPolExceedMarkCos custom string value.

var _ basetypes.StringValuableWithSemanticEquals = QosDppPolExceedMarkCosStringValue{}

type QosDppPolExceedMarkCosStringValue struct {
	basetypes.StringValue
}

func (v QosDppPolExceedMarkCosStringValue) Equal(o attr.Value) bool {
	other, ok := o.(QosDppPolExceedMarkCosStringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v QosDppPolExceedMarkCosStringValue) Type(ctx context.Context) attr.Type {
	return QosDppPolExceedMarkCosStringType{}
}

func (v QosDppPolExceedMarkCosStringValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(QosDppPolExceedMarkCosStringValue)

	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	priorMappedValue := QosDppPolExceedMarkCosValueMap(v.StringValue)

	newMappedValue := QosDppPolExceedMarkCosValueMap(newValue.StringValue)

	return priorMappedValue.Equal(newMappedValue), diags
}

func QosDppPolExceedMarkCosValueMap(value basetypes.StringValue) basetypes.StringValue {
	matchMap := map[string]string{
		"0xffff": "unspecified",
	}

	if val, ok := matchMap[value.ValueString()]; ok {
		return basetypes.NewStringValue(val)
	}

	return value
}

func NewQosDppPolExceedMarkCosStringNull() QosDppPolExceedMarkCosStringValue {
	return QosDppPolExceedMarkCosStringValue{
		StringValue: basetypes.NewStringNull(),
	}
}

func NewQosDppPolExceedMarkCosStringUnknown() QosDppPolExceedMarkCosStringValue {
	return QosDppPolExceedMarkCosStringValue{
		StringValue: basetypes.NewStringUnknown(),
	}
}

func NewQosDppPolExceedMarkCosStringValue(value string) QosDppPolExceedMarkCosStringValue {
	return QosDppPolExceedMarkCosStringValue{
		StringValue: basetypes.NewStringValue(value),
	}
}

func NewQosDppPolExceedMarkCosStringPointerValue(value *string) QosDppPolExceedMarkCosStringValue {
	return QosDppPolExceedMarkCosStringValue{
		StringValue: basetypes.NewStringPointerValue(value),
	}
}
