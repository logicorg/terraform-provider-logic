package provider

import (
	"context"
	"terraform-provider-logic/internal/dynlist"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ function.Function = ExactlyOneNotNullFunction{}
)

func NewExactlyOneNotNullFunction() function.Function {
	return ExactlyOneNotNullFunction{}
}

type ExactlyOneNotNullFunction struct{}

// Definition implements function.Function.
func (e ExactlyOneNotNullFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return true if exactly one value in the list is not null.",
		MarkdownDescription: "Return `true` if exactly one value in the list is not `null`.",
		Parameters: []function.Parameter{
			function.ListParameter{
				Name:                "list",
				Description:         "List to check.",
				MarkdownDescription: "List to check.",
				AllowNullValue:      false,
				AllowUnknownValues:  true,
				CustomType:          dynlist.DynListType{},
			},
		},
		Return: function.BoolReturn{},
	}
}

// Metadata implements function.Function.
func (e ExactlyOneNotNullFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "exactly_one_not_null"
}

// Run implements function.Function.
func (e ExactlyOneNotNullFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var list dynlist.DynList

	resp.Error = req.Arguments.Get(ctx, &list)
	if resp.Error != nil {
		return
	}

	hasUnknown := false
	notNull := 0
	for _, v := range list.Elements() {
		if notNull > 1 {
			break
		}

		tfVal, err := v.ToTerraformValue(ctx)
		if err != nil {
			resp.Error = function.ConcatFuncErrors(
				resp.Error,
				function.NewFuncError("Unexpected error while casting dynamic value to Terraform value. Please file an issue to the provider maintainers."),
			)
		}

		if !tfVal.IsKnown() {
			hasUnknown = true
			continue
		}

		if !tfVal.IsNull() {
			notNull++
		}
	}

	if notNull > 1 {
		// we know it will be false no matter what (regardless of any unknown value)
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, false))
		return
	} else if hasUnknown {
		// The value is (truly) unknown
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, basetypes.NewBoolUnknown()))
		return
	}

	result := notNull == 1
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
