package tmpl

import (
	"testing"

	"github.com/linode/terraform-provider-linode/v2/linode/acceptance"
)

type TemplateData struct {
	Region      string
	Unavailable string
}

func DataBasic(t testing.TB) string {
	return acceptance.ExecuteTemplate(t,
		"account_availabilities_data_basic", nil)
}

func DataFilterRegion(t testing.TB, region string) string {
	return acceptance.ExecuteTemplate(t,
		"account_availabilities_data_by_region", TemplateData{
			Region: region,
		})
}
