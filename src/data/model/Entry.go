package model

import "time"

type Entry struct {
	id          string
	title       string
	body        string
	dateCreated string
	diaryId     string
}

func (e *Entry) DiaryId() string {
	return e.diaryId
}

func (e *Entry) SetDiaryId(diaryId string) {
	e.diaryId = diaryId
}

func (e *Entry) Id() string {
	return e.id
}

func (e *Entry) SetId(id string) {
	e.id = id
}

func (e *Entry) Title() string {
	return e.title
}

func (e *Entry) SetTitle(title string) {
	e.title = title
}

func (e *Entry) Body() string {
	return e.body
}

func (e *Entry) SetBody(body string) {
	e.body = body
}

func (e *Entry) DateCreated() string {
	return e.dateCreated
}

func (e *Entry) SetDateCreated(dateCreated string) {
	e.dateCreated = dateCreated
}

func NewEntry(title string, body string, diaryId string) *Entry {
	times := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), time.Now().Hour(), time.Now().Minute(),
		time.Now().Second(), time.Now().Nanosecond(), time.Local)
	entry := &Entry{title: title, body: body, diaryId: diaryId, dateCreated: times.Format("02-Jan-2006")}
	return entry
}
