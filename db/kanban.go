package db

import (
	"gorm.io/gorm"
)

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

func (d Database) CreateKanban(id uint, name string) error {
	user, err := d.GetUser(id)
	if err != nil {
		return err
	}
	return d.db.Omit("Post.*").Create(&Kanban{
		Name: name,
		User: *user,
	}).Error
}

func (d Database) CreatePost(kanbanId uint, name string, desc *string, importance int, statusId uint) error {
	var s Status
	d.db.First(&s, statusId)
	return d.db.Create(&Post{
		KanbanId:    kanbanId,
		Name:        name,
		Description: desc,
		Importance:  importance,
		Status:      s,
	}).Error
}

func (d Database) CreateStatus(name string) error {
	return d.db.Create(&Status{
		Name: name,
	}).Error
}

// READ
func (d Database) GetKanbanBoardsFromUser(id uint) ([]Kanban, error) {
	user, err := d.GetUser(id)
	if err != nil {
		return nil, err
	}
	var k []Kanban
	err = d.db.Model(&Kanban{}).Preload("Post").Preload("Status").Where(
		Kanban{
			User: *user,
		},
	).Find(&k).Order("Name asc").Error
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (d Database) GetKanbanBord(id uint) (*Kanban, error) {
	var k Kanban
	err := d.db.Model(&Kanban{}).Preload("Post").Preload("Status").First(&k, id).Error
	if err != nil {
		return nil, err
	}
	return &k, err
}

// UPDATE

func (d Database) UpdateKanban(id uint, newName string) error {
	kanban, err := d.GetKanbanBord(id)
	if err != nil {
		return err
	}
	kanban.Name = newName
	return d.db.Save(&kanban).Error
}

func (d Database) UpdatePost(id uint, name string, desc *string, statusID uint, Importance int) error {
	var post Post
	d.db.First(&post, id)
	var s Status
	d.db.First(&s, statusID)
	post.Name = name
	post.Description = desc
	post.Status = s
	post.Importance = Importance
	return d.db.Save(&post).Error
}

func (d Database) UpdateStatus(id uint, name string) error {
	var s Status
	d.db.First(&s, id)
	s.Name = name
	return d.db.Save(&s).Error
}

// DELETE

func (d Database) DeletePost(id uint) error {
	return d.db.Delete(Post{}, id).Error
}

func (d Database) DeleteBoard(id uint) error {
	return d.db.Delete(Kanban{}, id).Error
}

func (d Database) DeleteStatus(id uint) error {
	return d.db.Delete(Post{}, id).Error
}
