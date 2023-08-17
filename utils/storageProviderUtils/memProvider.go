package storageProviderUtils

import (
	storageProvider "github.com/trustbloc/edv/pkg/edvprovider/memedvprovider"
)

var hostProviderMem *storageProvider.MemEDVProvider

var SetHostProviderMem = func(newProvider *storageProvider.MemEDVProvider) (errMsg string) {
	// TODO: if it is not nil then it already has been initialized, should it return error?
	hostProviderMem = newProvider
	return ""
}

// this function can be mocked in the tests to return edvMock.Provider
var GetHostProviderMem = func() *storageProvider.MemEDVProvider {
	return hostProviderMem
}
