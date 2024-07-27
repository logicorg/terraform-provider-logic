package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = XorFunction{}
)

func NewXorFunction() function.Function {
	return XorFunction{}
}

type XorFunction struct{}

// Definition implements function.Function.
func (x XorFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return the XOR between two values.",
		MarkdownDescription: "Return the XOR between two values.",
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
func (x XorFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "xor"
}

// Run implements function.Function.
func (x XorFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var first, second bool

	resp.Error = req.Arguments.Get(ctx, &first, &second)
	if resp.Error != nil {
		return
	}

	result := first != second

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
