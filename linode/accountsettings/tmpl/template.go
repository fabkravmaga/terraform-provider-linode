package tmpl

import (
	"testing"

	"github.com/linode/terraform-provider-linode/v2/linode/acceptance"
)

type TemplateData struct {
	LongviewSubscription string
	BackupsEnabled       bool
	NetworkHelper        bool
}

func Basic(t testing.TB) string {
	return acceptance.ExecuteTemplate(t,
		"account_settings_basic", nil)
}

func DataBasic(t testing.TB) string {
	return acceptance.ExecuteTemplate(t,
		"account_settings_data_basic", nil)
}

func Updates(t testing.TB, longviewSubscription string, backupsEnabled, networkHelper bool) string {
	return acceptance.ExecuteTemplate(t,
		"account_settings_updates", TemplateData{
			LongviewSubscription: longviewSubscription,
			BackupsEnabled:       backupsEnabled, NetworkHelper: networkHelper,
		})
}
