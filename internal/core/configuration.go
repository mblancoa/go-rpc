package core

import (
	"github.com/mblancoa/go-rpc/conf"
	"github.com/mblancoa/go-rpc/internal/core/ports"
	"github.com/rs/zerolog/log"
)

type coreConfiguration struct {
	FileStorage struct {
		StorageDire string `yaml:"storage-dir"`
		FilePattern string `yaml:"file-pattern"`
	} `yaml:"file-storage"`
}

var PersistenceContext *persistenceContext = &persistenceContext{}
var Context *coreContext = &coreContext{}

type coreContext struct {
	InfoFileService InfoFileService
}
type persistenceContext struct {
	InfoFileRepository ports.InfoFileRepository
}

// SetupCoreConfiguration creates the core context. Adapter contexts have to be created before
func SetupCoreConfiguration() {
	log.Info().Msg("Initializing core configuration")
	var c coreConfiguration
	conf.LoadYamlConfiguration(conf.GetConfigFile(), &c)
	p := PersistenceContext
	Context.InfoFileService = NewInfoFileService(c.FileStorage.StorageDire, c.FileStorage.FilePattern, p.InfoFileRepository)
}
