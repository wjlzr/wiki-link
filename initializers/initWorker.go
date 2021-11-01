package initializers

import (
	"github.com/spf13/viper"

	"wiki-link/common"
	"wiki-link/service/workers/sidekiqs"
)

func initWorkers() {
	viper.UnmarshalKey("workers", &common.AllWorkers)
	common.AllWorkerIs["AddressCheck"] = sidekiqs.CreateAddressCheck
	common.AllWorkerIs["OKLinkBlockHeightCheck"] = sidekiqs.CreateOKLinkBlockHeightCheck
	common.AllWorkerIs["ChainInfoBlockHeightCheck"] = sidekiqs.CreateChainInfoBlockHeightCheck
	common.AllWorkerIs["SendEmail"] = sidekiqs.CreateSendEmail
}
