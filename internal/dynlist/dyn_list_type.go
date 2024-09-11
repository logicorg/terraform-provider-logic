package dynlist

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListTypable          = (*DynListType)(nil)
	_ function.ValidateableParameter = (*DynListType)(nil)
)

type DynListType struct {
	basetypes.ListType
}

func (t DynListType) String() string {
	return "xtypes.DynamicList"
}

func (l DynListType) ElementType() attr.Type {
	return basetypes.DynamicType{}
}

func (l DynListType) WithElementType(typ attr.Type) attr.TypeWithElementType {
	return DynListType{
		ListType: basetypes.ListType{
			ElemType: basetypes.DynamicType{},
		},
	}
}

func (t DynListType) ValueType(ctx context.Context) attr.Value {
	return DynList{}
}

func (t DynListType) Equal(o attr.Type) bool {
	_, ok := o.(DynListType)

	return ok
}

func (t DynListType) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	return DynList{
		ListValue: in,
	}, nil
}

func (t DynListType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return basetypes.NewListNull(t.ElementType()), nil
	}

	if !in.Type().Is(tftypes.List{}) {
		return nil, fmt.Errorf("can't use %s as value of List with ElementType %T", in.String(), t.ElementType())
	}

	if !in.IsKnown() {
		return basetypes.NewListUnknown(t.ElementType()), nil
	}

	if in.IsNull() {
		return basetypes.NewListNull(t.ElementType()), nil
	}

	val := []tftypes.Value{}
	err := in.As(&val)
	if err != nil {
		return nil, err
	}
	elems := make([]attr.Value, 0, len(val))
	for _, elem := range val {
		av, err := t.ElementType().ValueFromTerraform(ctx, elem)
		if err != nil {
			return nil, err
		}
		elems = append(elems, av)
	}

	listValue := basetypes.NewListValueMust(t.ElementType(), elems)

	listValuable, diags := t.ValueFromList(ctx, listValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}

// ValidateParameter implements function.ValidateableParameter.
func (t *DynListType) ValidateParameter(context.Context, function.ValidateParameterRequest, *function.ValidateParameterResponse) {
	// Anyhing is valid
}
