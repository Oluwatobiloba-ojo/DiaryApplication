package test

import (
	"DiaryApp/src/data/model"
	"DiaryApp/src/data/repository"
	"testing"
)

func TestForTheSaveOfEntryAndTheEntryListBecomesOne(t *testing.T) {
	var entryRepository repository.EntryRepository = new(repository.EntryRepositoryImpl)
	entry := model.NewEntry("Today", "I am a boy", "1")
	entryRepository.Save(entry)
	entries := *entryRepository.FindAll()
	if len(entries) != 1 {
		t.Errorf("Actual :%d\n Expected: %d", 1, len(entries))
	}
}

func TestThatTheEntryIfItAlreadyExistInTheListDoesNotSaveThatEntry(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	entry := model.NewEntry("Today", "I am a boy", "1")
	repo.Save(entry)
	entries := repo.FindAll()
	if len(*entries) != 1 {
		t.Errorf("Actual :%d\n Expected: %d", 1, len(*entries))
	}
}

func TestThatTheEntryIfItSaveInTheListWeCanGetTheEntryFromTheList(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	entry := model.NewEntry("Yesterday", "I am going to a man house today", "1")
	repo.Save(entry)
	oneEntry := repo.FindById(entry.Id())
	if oneEntry.Title() != entry.Title() {
		t.Errorf("Actual : %s\n Expected %s ", entry.Title(), oneEntry.Title())
	}
}

func TestThatEntryRepositoryCanSaveThreeEntryAndCountOfEntriesIsThree(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	firstEntry := model.NewEntry("School", "Ope is my all", "1")
	secondEntry := model.NewEntry("Market", "I went to the market yesterday", "1")
	thirdEntry := model.NewEntry("Lagos", "Lagos is a peaceful and calm place to stay", "1")
	repo.Save(firstEntry)
	repo.Save(secondEntry)
	repo.Save(thirdEntry)
	count := repo.Count()
	if count != 3 {
		t.Errorf("Actual :%d\n  Expected : %d", 3, count)
	}
}

func TestThatEntryWhenTheirIsNoEntryInTheSliceItReturnAMessageNoEntryFound(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	repo.DeleteById("1")
}

func TestThatWhenTheEntryAlreadyExistInTheSliceAndTriesToDeleteItReturnSuccessfullyDeleted(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	entry := model.NewEntry("Yesterday", "Lagos is peaceful", "1")
	repo.Save(entry)
	if repo.Count() != 1 {
		t.Errorf("Actual: %d\n Expected: %d ", 1, repo.Count())
	}
	repo.DeleteById(entry.Id())
	if repo.Count() != 0 {
		t.Errorf("Actual: %d\n Expected %d", 0, repo.Count())
	}
}

func TestThatWhenWeAddThreeEntryAndWeWantToDeleteAllOfTheEntriesUsedDeleteAll(t *testing.T) {
	var repo repository.EntryRepository = new(repository.EntryRepositoryImpl)
	entry := model.NewEntry("Today", "Lagos is nice and cool", "1")
	firstEntry := model.NewEntry("Tomorrow", "Ope is so beautiful and cool", "1")
	secondEntry := model.NewEntry("Next Tomorrow", "Delighted is nice and super cool", "1")
	repo.Save(entry)
	repo.Save(firstEntry)
	repo.Save(secondEntry)
	if repo.Count() != 3 {
		t.Errorf("Actual %d\n  Expected %d", 3, repo.Count())
	}
	repo.DeleteAll()
	if repo.Count() != 0 {
		t.Errorf("Actual %d\n  Expected %d", 0, repo.Count())
	}
}
