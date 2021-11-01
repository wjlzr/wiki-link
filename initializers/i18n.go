package initializers

import (
	qorI18n "github.com/qor/i18n"
	"github.com/qor/i18n/backends/yaml"

	"wiki-link/i18n"
)

func init() {
	i18n.I18n = qorI18n.New(yaml.New("i18n/locales"))
}
