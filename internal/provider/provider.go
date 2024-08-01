// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure LogicProvider satisfies various provider interfaces.
var _ provider.Provider = &LogicProvider{}
var _ provider.ProviderWithFunctions = &LogicProvider{}

// LogicProvider defines the provider implementation.
type LogicProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// LogicProviderModel describes the provider data model.
type LogicProviderModel struct{}

func (p *LogicProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "logic"
	resp.Version = p.version
}

func (p *LogicProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provider of logic utility functions.",
	}
}

func (p *LogicProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data LogicProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *LogicProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *LogicProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *LogicProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewExactlyOneTrueFunction,
		NewXorFunction,
		NewXnorFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &LogicProvider{
			version: version,
		}
	}
}
