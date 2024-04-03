package service

import (
	"DiaryApp/src/data/model"
	"DiaryApp/src/data/repository"
	"DiaryApp/src/exception"
	"fmt"
)

type EntryService interface {
	CreateEntry(title string, body string, diaryId string) (string, error)
	FindAllEntryBelongingTo(diaryId string) []model.Entry
	UpdateEntry(title string, diaryId string, body string)
	FindEntry(title string, diaryId string) *model.Entry
	DeleteEntry(title string, diaryId string)
	DeleteAll(diaryId string)
}

type EntryServiceImpl struct {
	entryRepository repository.EntryRepository
}

func NewEntryService() *EntryServiceImpl {
	return &EntryServiceImpl{entryRepository: new(repository.EntryRepositoryImpl)}
}

func (e *EntryServiceImpl) UpdateEntry(title string, diaryId string, body string) {
	entry := e.FindEntry(title, diaryId)
	if entry != nil {
		newBody := entry.Body() + "\n" + body
		entry.SetBody(newBody)
		fmt.Println(entry)
		e.entryRepository.Save(entry)
	}
}

func (e *EntryServiceImpl) CreateEntry(title string, body string, diaryId string) (string, error) {
	if !e.checkIfTitleExist(title, diaryId) {
		newEntry := model.NewEntry(title, body, diaryId)
		e.entryRepository.Save(newEntry)
		return newEntry.Id(), nil
	}
	return "", exception.NewException("Title already exist")
}

func (e *EntryServiceImpl) checkIfTitleExist(title string, diaryId string) bool {
	allEntry := e.FindAllEntryBelongingTo(diaryId)
	for _, entry := range allEntry {
		if entry.Title() == title {
			return true
		}
	}
	return false
}

func (e *EntryServiceImpl) FindAllEntryBelongingTo(diaryId string) []model.Entry {
	var entries []model.Entry
	for _, entry := range *e.entryRepository.FindAll() {
		if entry.DiaryId() == diaryId {
			entries = append(entries, entry)
		}
	}
	return entries
}

func (e *EntryServiceImpl) DeleteEntry(title string, diaryId string) {
	entry := e.FindEntry(title, diaryId)
	if entry != nil {
		e.entryRepository.DeleteById(entry.Id())
	}
}

func (e *EntryServiceImpl) FindEntry(title string, diaryId string) *model.Entry {
	entries := e.FindAllEntryBelongingTo(diaryId)
	for _, entry := range entries {
		if entry.Title() == title {
			return &entry
		}
	}
	return nil
}

func (e *EntryServiceImpl) DeleteAll(diaryId string) {
	entries := e.FindAllEntryBelongingTo(diaryId)
	if entries != nil {
		for _, entry := range entries {
			e.entryRepository.DeleteById(entry.Id())
		}
	}
}
