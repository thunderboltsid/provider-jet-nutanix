/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/thunderboltsid/provider-jet-nutanix/config/nutanix_virtual_machine"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/thunderboltsid/provider-jet-nutanix/config/null"
	"github.com/thunderboltsid/provider-jet-nutanix/config/nutanix_category_key"
	"github.com/thunderboltsid/provider-jet-nutanix/config/nutanix_category_value"
	"github.com/thunderboltsid/provider-jet-nutanix/config/nutanix_image"
)

const (
	resourcePrefix = "nutanix"
	modulePath     = "github.com/thunderboltsid/provider-jet-nutanix"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		//   r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		tjconfig.WithIncludeList([]string{
			"nutanix_category_key$",
			"nutanix_category_value$",
			"nutanix_image$",
			"nutanix_virtual_machine$",
		}))

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		null.Configure,
		nutanix_category_key.Configure,
		nutanix_category_value.Configure,
		nutanix_image.Configure,
		nutanix_virtual_machine.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
