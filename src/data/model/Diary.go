package model

type Diary struct {
	username string
	password string
	isLocked bool
	entries  []Entry
	id       string
}

func (d *Diary) Id() string {
	return d.id
}

func (d *Diary) SetId(id string) {
	d.id = id
}

func (d *Diary) Username() string {
	return d.username
}

func (d *Diary) SetUsername(username string) {
	d.username = username
}

func (d *Diary) Password() string {
	return d.password
}

func (d *Diary) SetPassword(password string) {
	d.password = password
}

func (d *Diary) IsLocked() bool {
	return d.isLocked
}

func (d *Diary) SetIsLocked(isLocked bool) {
	d.isLocked = isLocked
}

func (d *Diary) Entries() []Entry {
	return d.entries
}

func (d *Diary) SetEntries(entries []Entry) {
	d.entries = entries
}

func NewDiary(username string, password string) *Diary {
	return &Diary{username: username, password: password, isLocked: false}
}
