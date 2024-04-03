package repository

import (
	"DiaryApp/src/data/model"
	"strconv"
)

type EntryRepository interface {
	Save(entry *model.Entry)
	FindById(id string) *model.Entry
	DeleteById(id string)
	DeleteAll()
	FindAll() *[]model.Entry
	Count() int
}

type EntryRepositoryImpl struct {
	entries []model.Entry
	count   int
}

func (e *EntryRepositoryImpl) Save(entry *model.Entry) {
	oldEntry := e.FindById(entry.Id())
	if oldEntry == nil {
		id := e.generateId()
		entry.SetId(strconv.Itoa(id))
		e.entries = append(e.entries, *entry)
	} else {
		e.DeleteById(oldEntry.Id())
		e.entries = append(e.entries, *entry)
	}
}

func (e *EntryRepositoryImpl) generateId() int {
	e.count += 1
	return e.count
}

func (e *EntryRepositoryImpl) FindById(id string) *model.Entry {
	entries := e.FindAll()
	if entries != nil {
		for _, entry := range *entries {
			if entry.Id() == id {
				return &entry
			}
		}
	}
	return nil
}

func (e *EntryRepositoryImpl) FindAll() *[]model.Entry {
	return &e.entries
}

func (e *EntryRepositoryImpl) Count() int {
	entries := e.entries
	return len(entries)
}

func (e *EntryRepositoryImpl) DeleteById(id string) {
	allEntry := e.FindAll()
	for index, entry := range *allEntry {
		if entry.Id() == id {
			e.entries = append(e.entries[:index], e.entries[index+1:]...)
		}
	}
}

func (e *EntryRepositoryImpl) DeleteAll() {
	e.entries = append(e.entries[:0], e.entries[len(e.entries):]...)
}
