package service

import (
	"sync"
	"yemeksepeti-keyvalue/internal/repository"
	"yemeksepeti-keyvalue/pkg/util"
)

type StorageService struct {
	StorageRepository *repository.Repository
	mu                sync.Mutex
}

func NewConverterService(s *repository.Repository) StorageService {
	return StorageService{StorageRepository: s}
}

/*
if key used value won't save, if not key will be saved
*/
func (s *StorageService) SetIfKeyNotUsedBefore(key string, value string) bool {
	if s.StorageRepository.IsKeyUsedBefore(key) || key == "" || value == "" {
		s.StorageRepository.SetPair(key, value)
		return true
	}
	return s.StorageRepository.SetPair(key, value)
}
func (s *StorageService) GetValue(key string) string {
	return s.StorageRepository.GetValue(key)
}
func (s *StorageService) GetAll() map[string]string {
	return s.StorageRepository.GetAll()
}
func (s *StorageService) Flush() error {
	return s.StorageRepository.Flush()
}
func (s *StorageService) Delete(key string) error {
	return s.StorageRepository.Delete(key)
}

/*
if there is any json file save it to memory
 */
func (s *StorageService) FileOps() bool {
	if util.IsThereAnyFile() {
		data := util.OpenFileAndReturnDataFromFile()
		s.mu.Lock()
		s.StorageRepository.SaveToMemory(data)
		s.mu.Unlock()
	}
	return false
}
/*
save data to file from memory
 */
func (s *StorageService) SaveMemoryToFile() bool {
	s.mu.Lock()
	util.DeleteAllFiles()
	s.StorageRepository.SaveToFile()
	s.mu.Unlock()
	return false
}
