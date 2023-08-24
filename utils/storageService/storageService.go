package storageService

import (
	"fmt"

	"github.com/Universal-Health-Chain/common-utils-golang/storageUtils"
	// "github.com/trustbloc/edge-core/pkg/log"
	"github.com/trustbloc/edv/pkg/restapi/edv/models"

	ariesStorage "github.com/hyperledger/aries-framework-go/spi/storage"
)

// var logger = log.New("edv-rest")

// StorageService is used to contain several private storages and the self storage.
// It wraps an Aries storage provider with additional functionality.
// It contains the stores for "documents" and "config" (to set the vaults created).
type StorageService struct {
	storageProvider ariesStorage.Provider
	storageType     string
	selfStorage     PrivateStorage
	privateStorages []PrivateStorage
}

func NewStorageService(storageProvider ariesStorage.Provider) *StorageService {
	return &StorageService{
		storageProvider: storageProvider,
	}
}

// Get the alternate name for this storage service
func (s *StorageService) GetAlternateName() string {
	return s.selfStorage.GetAlternateName()
}

// Set the alternate name for this storage service
func (s *StorageService) SetAlternateName(alternateName string) bool {
	return s.selfStorage.SetAlternateName(alternateName)
}

// ** Methods for private storages **

func (s *StorageService) findPrivateStorageByAlternateName(alternateName string) (*PrivateStorage, error) {
	for _, privateStorage := range s.privateStorages {
		if privateStorage.GetAlternateName() == alternateName {
			return &privateStorage, nil
		}
	}
	return nil, fmt.Errorf("private storage with alternate name %s not found", alternateName)
}

func (s *StorageService) CreatePrivateStorage(privateStorage PrivateStorage) error {
	_, err := s.findPrivateStorageByAlternateName(privateStorage.GetAlternateName())
	if err == nil { // Means a storage with that alternateName was found
		return fmt.Errorf("private storage with alternate name %s already exists", privateStorage.GetAlternateName())
	}

	s.privateStorages = append(s.privateStorages, privateStorage)
	return nil
}

func (s *StorageService) CreateNewVaultInPrivateStorage(storageAlternateName string, vaultID string, dataVaultConfiguration *models.DataVaultConfiguration) error {
	privateStorage, err := s.findPrivateStorageByAlternateName(storageAlternateName)
	if err != nil {
		return err
	}

	return privateStorage.CreateNewVault(vaultID, dataVaultConfiguration)
}

func (s *StorageService) VaultExistsInPrivateStorage(storageAlternateName, vaultID string) (bool, error) {
	privateStorage, err := s.findPrivateStorageByAlternateName(storageAlternateName)
	if err != nil {
		return false, err
	}

	return privateStorage.VaultExists(vaultID)
}

func (s *StorageService) PutInPrivateStorage(storageAlternateName, vaultID string, documents ...storageUtils.EncryptedDocument) error {
	privateStorage, err := s.findPrivateStorageByAlternateName(storageAlternateName)
	if err != nil {
		return err
	}

	return privateStorage.Put(vaultID, documents...)
}

func (s *StorageService) GetFromPrivateStorage(storageAlternateName, vaultID, documentID string) ([]byte, error) {
	privateStorage, err := s.findPrivateStorageByAlternateName(storageAlternateName)
	if err != nil {
		return nil, err
	}

	return privateStorage.Get(vaultID, documentID)
}

func (s *StorageService) DeleteFromPrivateStorage(storageAlternateName, vaultID, documentID string) error {
	privateStorage, err := s.findPrivateStorageByAlternateName(storageAlternateName)
	if err != nil {
		return err
	}

	return privateStorage.Delete(vaultID, documentID)
}

func (s *StorageService) QueryPrivateStorage(storageAlternateName, vaultID string, query storageUtils.Query) ([]storageUtils.EncryptedDocument, error) {
	privateStorage, err := s.findPrivateStorageByAlternateName(storageAlternateName)
	if err != nil {
		return nil, err
	}

	return privateStorage.Query(vaultID, query)
}
