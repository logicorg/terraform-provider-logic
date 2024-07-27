package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ function.Function = ExactlyOneTrueFunction{}
)

func NewExactlyOneTrueFunction() function.Function {
	return ExactlyOneTrueFunction{}
}

type ExactlyOneTrueFunction struct{}

// Definition implements function.Function.
func (e ExactlyOneTrueFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return true if exactly one value in the list is true.",
		MarkdownDescription: "Return `true` if exactly one value in the list is `true`.",
		Parameters: []function.Parameter{
			function.ListParameter{
				Name:               "list",
				Description:        "The list to check.",
				ElementType:        types.BoolType,
				AllowNullValue:     true,
				AllowUnknownValues: false,
			},
		},
		Return: function.BoolReturn{},
	}
}

// Metadata implements function.Function.
func (e ExactlyOneTrueFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "exactly_one_true"
}

// Run implements function.Function.
func (e ExactlyOneTrueFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var list []types.Bool

	resp.Error = req.Arguments.Get(ctx, &list)
	if resp.Error != nil {
		return
	}

	trueCount := 0
	for _, v := range list {
		if trueCount > 1 {
			break
		}

		if v.ValueBool() {
			trueCount++
		}
	}
	result := trueCount == 1

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
