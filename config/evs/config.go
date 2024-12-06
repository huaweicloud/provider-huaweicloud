package evs

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupEvs = "evs"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_evs_snapshot", func(r *config.Resource) {
		r.ShortGroup = shortGroupEvs

		r.References["volume_id"] = config.Reference{
			TerraformName: "huaweicloud_evs_volume",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_evs_snapshot_rollback", func(r *config.Resource) {
		r.ShortGroup = shortGroupEvs

		r.References["volume_id"] = config.Reference{
			TerraformName: "huaweicloud_evs_volume",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["snapshot_id"] = config.Reference{
			TerraformName: "huaweicloud_evs_snapshot",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_evs_volume", func(r *config.Resource) {
		r.ShortGroup = shortGroupEvs

		r.References["server_id"] = config.Reference{
			TerraformName: "huaweicloud_compute_instance",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}

		r.References["snapshot_id"] = config.Reference{
			TerraformName: "huaweicloud_evs_snapshot",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_evs_volume_transfer", func(r *config.Resource) {
		r.ShortGroup = shortGroupEvs

		r.References["volume_id"] = config.Reference{
			TerraformName: "huaweicloud_evs_volume",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_evs_volume_transfer_accepter", func(r *config.Resource) {
		r.ShortGroup = shortGroupEvs

		r.References["transfer_id"] = config.Reference{
			TerraformName: "huaweicloud_evs_volume_transfer",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})
}
