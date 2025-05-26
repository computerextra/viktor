package db

import "gorm.io/gorm"

// TODO: https://marmelab.com/blog/2023/07/28/create-a-kanban-board-in-react-admin.html

type Kanban struct {
	gorm.Model
	Name  string
	User  User
	Posts []Post
}

type Post struct {
	gorm.Model
	KanbanId    uint
	Name        string
	Description *string
	Status      Status
	Importance  int
}

type Status struct {
	gorm.Model
	Name string
}

// CREATE

func (d Database) CreateKanban(id uint, name string) {
	user := d.GetUser(id)
	d.db.Omit("Post.*").Create(&Kanban{
		Name: name,
		User: user,
	})
}

func (d Database) CreatePost(kanbanId uint, name string, desc *string, importance int, statusId uint) {
	var s Status
	d.db.First(&s, statusId)
	d.db.Create(&Post{
		KanbanId:    kanbanId,
		Name:        name,
		Description: desc,
		Importance:  importance,
		Status:      s,
	})
}

func (d Database) CreateStatus(name string) {
	d.db.Create(&Status{
		Name: name,
	})
}

// READ
func (d Database) GetKanbanBoardsFromUser(id uint) []Kanban {
	user := d.GetUser(id)
	var k []Kanban
	d.db.Model(&Kanban{}).Preload("Post").Preload("Status").Where(
		Kanban{
			User: user,
		},
	).Find(&k).Order("Name asc")
	return k
}

func (d Database) GetKanbanBord(id uint) *Kanban {
	var k Kanban
	d.db.Model(&Kanban{}).Preload("Post").Preload("Status").First(&k, id)
	return &k
}

// UPDATE

func (d Database) UpdateKanban(id uint, newName string) {
	kanban := d.GetKanbanBord(id)
	kanban.Name = newName
	d.db.Save(&kanban)
}

func (d Database) UpdatePost(id uint, name string, desc *string, statusID uint, Importance int) {
	var post Post
	d.db.First(&post, id)
	var s Status
	d.db.First(&s, statusID)
	post.Name = name
	post.Description = desc
	post.Status = s
	post.Importance = Importance
	d.db.Save(&post)
}

func (d Database) UpdateStatus(id uint, name string) {
	var s Status
	d.db.First(&s, id)
	s.Name = name
	d.db.Save(&s)
}

// DELETE

// TODO: NYI
