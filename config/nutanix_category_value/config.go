package nutanix_category_value

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("nutanix_category_value", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "nutanix"
		r.ShortGroup = "category_value"

		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to CategoryKey
		// object, we can build cross resource referencing.
		r.References["repository"] = config.Reference{
			Type: "github.com/thunderboltsid/provider-jet-nutanix/apis/category_key/v1alpha1.CategoryKey",
		}
	})


}
