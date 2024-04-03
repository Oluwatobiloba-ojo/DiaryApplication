package service

import (
	"DiaryApp/src/data/model"
	"DiaryApp/src/data/repository"
	"DiaryApp/src/exception"
)

type DiaryService interface {
	CreateDiary(username string, password string) (*model.Diary, error)
	Unlock(username string, password string)
	Lock(username string)
	CreateEntry(username string, title string, body string) (string, error)
	UpdateEntry(username string, title string, body string)
	DeleteEntry(username string, title string)
	DeleteAllEntry(username string)
	FindEntry(username string, title string) *model.Entry
	FindAllEntry(username string) []model.Entry
	FindDiaryByUsername(username string) *model.Diary
	DeleteDiaryBy(username string)
	CheckIfDiaryIsLock(username string) bool
}

type DiaryServiceImpl struct {
	diaryRepository repository.DiaryRepository
	entryService    EntryService
}

func NewDiaryServiceImpl() *DiaryServiceImpl {
	return &DiaryServiceImpl{
		diaryRepository: new(repository.DiaryRepositoryImpl),
		entryService:    NewEntryService()}
}

func (receiver *DiaryServiceImpl) CreateDiary(username string, password string) (*model.Diary, error) {
	diary := receiver.FindDiaryByUsername(username)
	if diary == nil {
		newDiary := model.NewDiary(username, password)
		receiver.diaryRepository.Save(newDiary)
		return newDiary, nil
	}
	return nil, exception.NewException("Account created already not exist")
}

func (receiver *DiaryServiceImpl) Unlock(username string, password string) {
	diary := receiver.FindDiaryByUsername(username)
	if diary != nil && diary.IsLocked() && diary.Password() == password {
		diary.SetIsLocked(false)
		receiver.diaryRepository.Save(diary)
	}
}

func (receiver *DiaryServiceImpl) Lock(username string) {
	diary := receiver.FindDiaryByUsername(username)
	if diary != nil && !diary.IsLocked() {
		diary.SetIsLocked(true)
		receiver.diaryRepository.Save(diary)
	}
}

func (receiver *DiaryServiceImpl) CreateEntry(username string, title string, body string) (string, error) {
	diary := receiver.FindDiaryByUsername(username)
	if diary != nil {
		if receiver.CheckIfDiaryIsLock(username) != false {
			return "", exception.NewException("Diary is locked")
		}
		return receiver.entryService.CreateEntry(title, body, diary.Id())
	}
	return "", exception.NewException("Diary not yet created")
}

func (receiver *DiaryServiceImpl) UpdateEntry(username string, title string, body string) {
	diary := receiver.FindDiaryByUsername(username)
	if diary != nil {
		receiver.entryService.UpdateEntry(title, diary.Id(), body)
	}
}

func (receiver *DiaryServiceImpl) DeleteEntry(username string, title string) {
	diary := receiver.FindDiaryByUsername(username)
	if diary != nil {
		receiver.entryService.DeleteEntry(title, diary.Id())
	}
}

func (receiver *DiaryServiceImpl) DeleteAllEntry(username string) {
	foundDiary := receiver.FindDiaryByUsername(username)
	if foundDiary != nil {
		receiver.entryService.DeleteAll(foundDiary.Id())
	}
}

func (receiver *DiaryServiceImpl) FindEntry(username string, title string) *model.Entry {
	diary := receiver.FindDiaryByUsername(username)
	if diary != nil {
		return receiver.entryService.FindEntry(title, diary.Id())
	}
	return nil
}

func (receiver *DiaryServiceImpl) FindAllEntry(username string) []model.Entry {
	foundDiary := receiver.FindDiaryByUsername(username)
	if foundDiary != nil {
		return receiver.entryService.FindAllEntryBelongingTo(foundDiary.Id())
	}
	return nil
}

func (receiver *DiaryServiceImpl) FindDiaryByUsername(username string) *model.Diary {
	diaries := *receiver.diaryRepository.FindAll()
	for _, diary := range diaries {
		if diary.Username() == username {
			return &diary
		}
	}
	return nil
}

func (receiver *DiaryServiceImpl) DeleteDiaryBy(username string) {
	foundDiary := receiver.FindDiaryByUsername(username)
	if foundDiary != nil {
		receiver.diaryRepository.DeleteById(foundDiary.Id())
	}
}

func (receiver *DiaryServiceImpl) CheckIfDiaryIsLock(username string) bool {
	diary := receiver.FindDiaryByUsername(username)
	if diary.IsLocked() == true {
		return false
	} else {
		return true
	}
}
