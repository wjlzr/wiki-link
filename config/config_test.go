package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	ReadConfig("D:\\gitlab\\wiki-link\\config\\config.toml")
	fmt.Println(Conf().GetOkLinkApiKey())
}
