package test

import (
	"DiaryApp/src/service"
	"testing"
)

func TestUserCanCreateADiary(t *testing.T) {
	var diaryService service.DiaryService = service.NewDiaryServiceImpl()
	diaryService.CreateDiary("Opeoluwa", "password")
	diary := diaryService.FindDiaryByUsername("Opeoluwa")
	if diary == nil {
		t.Errorf("Expected : {id : 1, username: opeoluwa, password: password}  Actual: null")
	}
}

func TestUserCanUnlockTheirDiaryAndWhenTheDiaryIsGottenIsOpened(t *testing.T) {
	var diaryService service.DiaryService = service.NewDiaryServiceImpl()
	diaryService.CreateDiary("Opeoluwa", "password")
	diaryService.Lock("Opeoluwa")
	diary := diaryService.FindDiaryByUsername("Opeoluwa")
	if diary.IsLocked() != true {
		t.Errorf("Expected: true  Actual: false")
	}
	diaryService.Unlock("Opeoluwa", "password")
	diary = diaryService.FindDiaryByUsername("Opeoluwa")
	if diary.IsLocked() != false {
		t.Errorf("Expected: false Actual : true")
	}
}

func TestThatCreateEntryFromDiary(t *testing.T) {
	var diaryService service.DiaryService = service.NewDiaryServiceImpl()
	diaryService.CreateDiary("Ope", "password")
	diaryService.CreateEntry("Ope", "Yesterday",
		"I had an hectic day in class today")
	entry := diaryService.FindEntry("Ope", "Yesterday")
	if entry == nil {
		t.Errorf("Expected {id : 1, title : Yesterday," +
			"body : I had an hectic day in class today, diaryId : 1}, Actual: nil")
	}
}

func TestThatUserCanUpdateAnEntry(t *testing.T) {
	var diaryService service.DiaryService = service.NewDiaryServiceImpl()
	diaryService.CreateDiary("ope", "password")
	diaryService.CreateEntry("ope", "Today", "I went to church today to pray to God")
	diaryService.UpdateEntry("ope", "Today", "And I saw Ope on the way")
	entry := diaryService.FindEntry("ope", "Today")
	updatedBody := "I went to church today to pray to God\nAnd I saw Ope on the way"
	if entry.Body() != updatedBody {
		t.Errorf("Expected : {%s\n} Actual: {%s}", entry.Body(), updatedBody)
	}
}

func TestThatUserCanDeleteAnEntry(t *testing.T) {
	var diaryService service.DiaryService = service.NewDiaryServiceImpl()
	diaryService.CreateDiary("ope", "password")
	diaryService.CreateEntry("ope", "Today", "I went to church today to pray to God")
	diaryService.CreateEntry("ope", "Tomorrow", "I went to church today to pray to God")
	diaryService.DeleteEntry("ope", "Today")
	entries := diaryService.FindAllEntry("ope")
	if len(entries) != 1 {
		t.Errorf("Expected :%d\n Actual %d", len(entries), 1)
	}
}

func TestThatUserCanDeleteAllEntriesIN(t *testing.T) {
	var diaryService service.DiaryService = service.NewDiaryServiceImpl()
	diaryService.CreateDiary("ope", "password")
	diaryService.CreateEntry("ope", "Today", "I went to church today to pray to God")
	diaryService.CreateEntry("ope", "Tomorrow", "I went to church today to pray to God")
	diaryService.DeleteAllEntry("ope")
	lengthOfEntry := len(diaryService.FindAllEntry("ope"))
	if lengthOfEntry != 0 {
		t.Errorf("Expected :%d Actual :%d", lengthOfEntry, 0)
	}
}

func TestThatUserCanDeleteDiary(t *testing.T) {
	var diaryService service.DiaryService = service.NewDiaryServiceImpl()
	diaryService.CreateDiary("ope", "password")
	diaryService.CreateEntry("ope", "Today", "I went to church today to pray to God")
	diaryService.DeleteDiaryBy("ope")
	if diaryService.FindDiaryByUsername("ope") != nil {
		t.Errorf("Expected:Diary, Actual null")
	}

}
