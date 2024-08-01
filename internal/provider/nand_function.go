package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = NandFunction{}
)

func NewNandFunction() function.Function {
	return NandFunction{}
}

type NandFunction struct{}

// Definition implements function.Function.
func (x NandFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return the NAND between two values.",
		MarkdownDescription: "Return the NAND between two values.",
		Parameters: []function.Parameter{
			function.BoolParameter{
				Name:               "first",
				Description:        "The first value.",
				AllowNullValue:     true,
				AllowUnknownValues: false,
			},
			function.BoolParameter{
				Name:               "second",
				Description:        "The second value.",
				AllowNullValue:     true,
				AllowUnknownValues: false,
			},
		},
		Return: function.BoolReturn{},
	}
}

// Metadata implements function.Function.
func (x NandFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "nand"
}

// Run implements function.Function.
func (x NandFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var first, second bool

	resp.Error = req.Arguments.Get(ctx, &first, &second)
	if resp.Error != nil {
		return
	}

	result := !(first && second)

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
