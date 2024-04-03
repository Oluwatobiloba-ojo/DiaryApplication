package repository

import (
	"DiaryApp/src/data/model"
	"strconv"
)

type DiaryRepository interface {
	Save(diary *model.Diary)
	FindById(id string) *model.Diary
	DeleteById(id string)
	FindAll() *[]model.Diary
	DeleteAll()
	Count() int
}

type DiaryRepositoryImpl struct {
	diaries []model.Diary
	count   int
}

func (e *DiaryRepositoryImpl) Save(diary *model.Diary) {
	oldDiary := e.FindById(diary.Id())
	if oldDiary == nil {
		id := e.generateId()
		diary.SetId(strconv.Itoa(id))
		e.diaries = append(e.diaries, *diary)
	} else {
		e.DeleteById(oldDiary.Id())
		e.diaries = append(e.diaries, *diary)
	}
}

func (e *DiaryRepositoryImpl) FindById(id string) *model.Diary {
	diaries := e.FindAll()
	if diaries != nil {
		for _, entry := range *diaries {
			if entry.Id() == id {
				return &entry
			}
		}
	}
	return nil
}

func (e *DiaryRepositoryImpl) DeleteById(id string) {
	allDiary := e.FindAll()
	for index, entry := range *allDiary {
		if entry.Id() == id {
			e.diaries = append(e.diaries[:index], e.diaries[index+1:]...)
		}
	}

}

func (e *DiaryRepositoryImpl) FindAll() *[]model.Diary {
	return &e.diaries
}

func (e *DiaryRepositoryImpl) DeleteAll() {
	e.diaries = append(e.diaries[:0], e.diaries[len(e.diaries):]...)
}

func (e *DiaryRepositoryImpl) Count() int {
	return len(e.diaries)
}

func (e *DiaryRepositoryImpl) generateId() int {
	e.count += 1
	return e.count
}
