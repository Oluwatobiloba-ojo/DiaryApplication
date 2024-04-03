package test

import (
	"DiaryApp/src/service"
	"testing"
)

func Test_createEntry(t *testing.T) {
	entryService := service.NewEntryService()
	title := "Yesterday"
	body := "Ope good morning"
	entryService.CreateEntry(title, body, "1")
	entries := entryService.FindAllEntryBelongingTo("1")
	if len(entries) != 1 {
		t.Errorf("Error")
	}
}

func TestThatEntryHavingTheTitleAndYouWantToUpdateTheBodyItAppendTHeBody(t *testing.T) {
	services := service.NewEntryService()
	title := "Yesterday"
	body := "Ope good morning"
	services.CreateEntry(title, body, "1")
	services.UpdateEntry("Yesterday", "1", "Ope good afternoon")
	bigBody := body + "\n" + "Ope good afternoon"
	entry := services.FindEntry(title, "1")
	if entry.Body() != bigBody {
		t.Errorf("Actual %s\n	Expected %s", bigBody, entry.Body())
	}
}

func TestThatEntryWhenCreatedTwoAndDeleteOneReturnOneEntryInTheList(t *testing.T) {
	services := service.NewEntryService()
	services.CreateEntry("Tomorrow", "Ope is a fine girl", "1")
	services.CreateEntry("Yesterday", "Ope good morning", "1")
	if len(services.FindAllEntryBelongingTo("1")) != 2 {
		t.Errorf("Actual %d\n	Expected %d", 2, len(services.FindAllEntryBelongingTo("1")))
	}
	services.DeleteEntry("Tomorrow", "1")
	if len(services.FindAllEntryBelongingTo("1")) != 1 {
		t.Errorf("Actual %d\n	Expected %d", 1, len(services.FindAllEntryBelongingTo("1")))
	}
}
func TestThatEntryWhenCreatedOneForUserONeAndAnotherFOrUser2DeleteAllUser1OnlyDeleteUser1(t *testing.T) {
	services := service.NewEntryService()
	services.CreateEntry("Tomorrow", "Ope is a fine girl", "1")
	services.CreateEntry("Yesterday", "Ope good morning", "2")
	services.DeleteAll("1")
	userTwoEntry := services.FindAllEntryBelongingTo("2")
	userOneEntry := services.FindAllEntryBelongingTo("1")
	if len(userTwoEntry) != 1 {
		t.Errorf("Actual %d\n	Expected %d", 1, len(userTwoEntry))
	}
	if userOneEntry != nil {
		t.Errorf("Actual: nil\n  Expected: %s", userOneEntry)
	}
}
