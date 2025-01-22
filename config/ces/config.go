package ces

import (
	"github.com/crossplane/upjet/pkg/config"
)

const shortGroupCes = "ces"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("huaweicloud_ces_alarm_template", func(r *config.Resource) {
		r.ShortGroup = shortGroupCes
	})

	p.AddResourceConfigurator("huaweicloud_ces_alarmrule", func(r *config.Resource) {
		r.ShortGroup = shortGroupCes

		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"metric.dimensions"},
		}
	})

	p.AddResourceConfigurator("huaweicloud_ces_dashboard", func(r *config.Resource) {
		r.ShortGroup = shortGroupCes
	})

	p.AddResourceConfigurator("huaweicloud_ces_dashboard_widget", func(r *config.Resource) {
		r.ShortGroup = shortGroupCes

		r.References["dashboard_id"] = config.Reference{
			TerraformName: "huaweicloud_ces_dashboard",
			Extractor:     `github.com/crossplane/upjet/pkg/resource.ExtractResourceID()`,
		}
	})

	p.AddResourceConfigurator("huaweicloud_ces_event_report", func(r *config.Resource) {
		r.ShortGroup = shortGroupCes
	})

	p.AddResourceConfigurator("huaweicloud_ces_one_click_alarm", func(r *config.Resource) {
		r.ShortGroup = shortGroupCes
	})

	p.AddResourceConfigurator("huaweicloud_ces_resource_group", func(r *config.Resource) {
		r.ShortGroup = shortGroupCes

		r.LateInitializer = config.LateInitializer{
			ConditionalIgnoredFields: []string{"type", "tags", "associated_eps_ids", "resources"},
		}
	})
}
