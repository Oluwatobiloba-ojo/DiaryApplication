package test

import (
	"DiaryApp/src/data/model"
	"DiaryApp/src/data/repository"
	"testing"
)

func TestThatTheSaveOfADiaryAndDiaryListBecomeOne(t *testing.T) {
	var repo repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	diary := model.NewDiary("Tobi123", "1234567")
	repo.Save(diary)
	if repo.Count() != 1 {
		t.Errorf("Actual %d\n   Expected %d", 1, repo.Count())
	}
}

func TestIfDiaryAlreadyExistInTheListOfDiaryDiaryWillNotBeSave(t *testing.T) {
	var repo repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	diary := model.NewDiary("Tobi123", "1234567")
	repo.Save(diary)
	if repo.Count() != 1 {
		t.Errorf("Actual %d\n   Expected %d", 1, repo.Count())
	}
	repo.Save(diary)
	if repo.Count() != 1 {
		t.Errorf("Actual %d\n   Expected %d", 1, repo.Count())
	}
}

func TestThatADiaryWhenItIsSaveWeCanFindTheEntryByTheId(t *testing.T) {
	var repo repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	diary := model.NewDiary("Ope123", "wieidwudiwd")
	repo.Save(diary)
	newDiary := repo.FindById(diary.Id())
	if diary.Username() != newDiary.Username() {
		t.Errorf("Actual: %s\n	Expected: %s", diary.Username(), newDiary.Username())
	}
}

func TestThatTheDiaryAfterItHasBeenSavedAlreadyAndWantToEditTheUsername(t *testing.T) {
	var repo repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	diary := model.NewDiary("Emmanuella", "bwhuhahbgfalh")
	repo.Save(diary)
	diary.SetUsername("Delighted")
	repo.Save(diary)
	findDiary := repo.FindById(diary.Id())
	if findDiary.Username() != diary.Username() {
		t.Errorf("Actual: %s\n	Expected: %s", diary.Username(), findDiary.Username())
	}
	if repo.Count() != 1 {
		t.Errorf("Actual %d\n   Expected %d", 1, repo.Count())
	}
}

func TestTestThatDiaryCanBeDeletedByTheId(t *testing.T) {
	var repo repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	diary := model.NewDiary("Ope123", "wieidwudiwd")
	repo.Save(diary)
	if repo.Count() != 1 {
		t.Errorf("Actual %d\n   Expected %d", 1, repo.Count())
	}
	repo.DeleteById(diary.Id())
	if repo.Count() != 0 {
		t.Errorf("Actual %d\n   Expected %d", 0, repo.Count())
	}
}

func TestThatAllDiaryCanBeDeletedFromTheSlice(t *testing.T) {
	var repo repository.DiaryRepository = new(repository.DiaryRepositoryImpl)
	diary := model.NewDiary("Ope123", "wieidwudiwd")
	repo.Save(diary)

	repo.DeleteAll()
	if repo.Count() != 0 {
		t.Errorf("Actual %d\n   Expected %d", 0, repo.Count())
	}

}
