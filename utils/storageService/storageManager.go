package storageService

import (
	"fmt"

	"github.com/Universal-Health-Chain/common-utils-golang/storageUtils"
)

type StorageServicesManager struct {
	storageServices []*StorageService
}

// CreateStorageService checks if the alternateName already exists in some selfStorage or privateStorages
func (m *StorageServicesManager) CreateStorageService(alternateName string, parameters storageUtils.StorageParameters, databaseTimeout uint64) error {
	for _, service := range m.storageServices {
		// Check if the alternateName exists in selfStorage
		if service.selfStorage.GetAlternateName() == alternateName {
			return fmt.Errorf("a storage service with alternate name %s already exists", alternateName)
		}
		// Check if the alternateName exists in any of the privateStorages
		for _, privateStorage := range service.privateStorages {
			if privateStorage.GetAlternateName() == alternateName {
				return fmt.Errorf("the alternate name %s already exists in a private storage of a storage service with alternate name %s", alternateName, service.selfStorage.GetAlternateName())
			}
		}
	}

	provider, err := storageUtils.NewStorageProvider(parameters, databaseTimeout)
	if err != nil {
		return err
	}

	newService := NewStorageService(provider)
	newService.selfStorage.SetAlternateName(alternateName)
	m.storageServices = append(m.storageServices, newService)

	return nil
}

func (m *StorageServicesManager) GetStorageServiceByAlternateName(alternateName string) (*StorageService, error) {
	for _, service := range m.storageServices {
		if service.selfStorage.GetAlternateName() == alternateName {
			return service, nil
		}
	}

	return nil, fmt.Errorf("no storage service found with alternate name %s", alternateName)
}

func (m *StorageServicesManager) GetStorageServiceByPrivateStorage(alternateName string) (*StorageService, error) {
	for _, service := range m.storageServices {
		if service.selfStorage.GetAlternateName() == alternateName {
			return service, nil
		}

		for _, privateStorage := range service.privateStorages {
			if privateStorage.GetAlternateName() == alternateName {
				return service, nil
			}
		}
	}

	return nil, fmt.Errorf("no storage service found containing private storage with alternate name %s", alternateName)
}
