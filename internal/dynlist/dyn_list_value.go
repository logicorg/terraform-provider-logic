package dynlist

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.ListValuable = (*DynList)(nil)
)

type DynList struct {
	basetypes.ListValue
}

func (v DynList) Type(_ context.Context) attr.Type {
	return DynListType{
		ListType: basetypes.ListType{
			ElemType: basetypes.DynamicType{},
		},
	}
}

func (v DynList) Equal(o attr.Value) bool {
	other, ok := o.(DynList)

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

// castNullOrUnknownValues checks if the input value is an unknown/nil DynamicPseudoType
// and casts it to a known value that matches the concrete value type.
// If in is not an unknown/null DynamicPseudoType, the output value will not be modified.
//
// Parameters:
// - t: the target type to cast to.
// - in: the input value to check and cast.
// - out: a pointer to the output value where the result will be stored.
func castNullOrUnknownValues(t tftypes.Type, in tftypes.Value, out *tftypes.Value) {
	// If the value is an unknown/nil DynamicPseudoType, we need to append a unknown/nil that matches the concrete value type
	if in.Type().Is(tftypes.DynamicPseudoType) {
		if in.IsNull() {
			*out = tftypes.NewValue(t, nil)
		} else if !in.IsKnown() {
			*out = tftypes.NewValue(t, tftypes.UnknownValue)
		}
	}
}

func (v DynList) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	listType := tftypes.List{ElementType: v.ElementType(ctx).TerraformType(ctx)}

	if v.ListValue.IsNull() {
		return tftypes.NewValue(tftypes.List{ElementType: tftypes.DynamicPseudoType}, nil), nil
	}

	if v.ListValue.IsUnknown() {
		return tftypes.NewValue(tftypes.List{ElementType: tftypes.DynamicPseudoType}, tftypes.UnknownValue), nil
	}

	var elemTfType tftypes.Type = tftypes.DynamicPseudoType

	// Since the element type is dynamic, the final list element type will be determined by the value.
	for _, elem := range v.Elements() {
		val, err := elem.ToTerraformValue(ctx)
		// Find the first non-dynamic value and use that as the type
		if err == nil && !val.Type().Is(tftypes.DynamicPseudoType) {
			elemTfType = val.Type()
			break
		}
	}

	vals := make([]tftypes.Value, 0, len(v.Elements()))

	for _, elem := range v.Elements() {
		val, err := elem.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(listType, tftypes.UnknownValue), err
		}

		castNullOrUnknownValues(elemTfType, val, &val)

		vals = append(vals, val)
	}

	if err := tftypes.ValidateValue(listType, vals); err != nil {
		return tftypes.NewValue(listType, tftypes.UnknownValue), err
	}

	return tftypes.NewValue(listType, vals), nil
}

func NewDynListNull() DynList {
	return DynList{
		ListValue: basetypes.NewListNull(basetypes.DynamicType{}),
	}
}

func NewDynListUnknown() DynList {
	return DynList{
		ListValue: basetypes.NewListUnknown(basetypes.DynamicType{}),
	}
}

func NewDynListValue(elements []attr.Value) (DynList, diag.Diagnostics) {
	listValue, diags := basetypes.NewListValue(basetypes.DynamicType{}, elements)
	if diags.HasError() {
		return NewDynListUnknown(), diags
	}

	return DynList{
		ListValue: listValue,
	}, nil
}
