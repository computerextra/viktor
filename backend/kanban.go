package backend

import (
	"fmt"
	"time"
	"viktor/db"

	"github.com/lucsky/cuid"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateKanban(id string, Name string) bool {
	user := a.GetUser(id)
	if user == nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage("[CreateKanban] Benutzer konnte nicht gefunden werden")
		dialog.Show()
		return false
	}
	board := db.Kanban{
		Id:        cuid.New(),
		Name:      Name,
		Posts:     []db.Post{},
		UserId:    user.Id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if !a.DB.HasKey("Kanban") {
		return a.DB.Set("Kanban", board) == nil
	}

	err := a.DB.Update("Kanban", board)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateKanban] Fehler beim Anlegen", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) CreatePost(Name, Status, Importance string, Description *string, KId string) bool {
	ap := db.Post{
		Id:          cuid.New(),
		Name:        Name,
		KanbanId:    KId,
		Description: Description,
		Status:      Status,
		Importance:  Importance,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if !a.DB.HasKey("Post") {
		return a.DB.Set("Post", ap) == nil
	}

	err := a.DB.Update("Post", ap)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateAnsprechpartner] Fehler beim Anlegen", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) GetKanbanBord(id string) *db.Kanban {
	boards, err := a.DB.Get("Kanban")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetKanbanBord] Fehler", err))
		dialog.Show()
		return nil
	}
	posts := a.GetPostsFromBoard(id)
	var board db.Kanban

	for _, x := range boards.([]db.Kanban) {
		if x.Id == id {
			board = x
		}
	}
	board.Posts = posts
	return &board
}

func (a *App) GetKanbanBoardsFromUser(id string) []db.Kanban {
	user := a.GetUser(id)
	if user == nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage("[GetKanbanBoardsFromUser] Benutzer konnte nicht gefunden werden")
		dialog.Show()
		return nil
	}
	boards, err := a.DB.Get("Kanban")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetKanbanBoardsFromUser] [GET(Kanban)] Fehler", err))
		dialog.Show()
		return nil
	}
	var res []db.Kanban
	for _, x := range boards.([]db.Kanban) {
		if x.UserId == id {
			res = append(res, x)
		}
	}
	return res
}

func (a *App) GetPostsFromBoard(id string) []db.Post {
	posts, err := a.DB.Get("Kanban")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetPostsFromBoard] Fehler", err))
		dialog.Show()
		return nil
	}
	var res []db.Post
	for _, x := range posts.([]db.Post) {
		if x.KanbanId == id {
			res = append(res, x)
		}
	}
	return res
}

func (a *App) UpdateKanban(id, Name string) bool {
	board := a.GetKanbanBord(id)

	board.Name = Name
	board.UpdatedAt = time.Now()

	err := a.DB.Update("Kanban", board)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdateKanban] Fehler beim Speichern", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) UpdatePost(id, Name, Status, Importance string, Description *string) bool {
	posts, err := a.DB.Get("Post")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdatePost] Fehler", err))
		dialog.Show()
		return false
	}

	var post db.Post
	var found bool = false
	for _, x := range posts.([]db.Post) {
		if x.Id == id {
			post = x
			found = true
		}
	}
	if !found {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdatePost] Post nicht gefunden", err))
		dialog.Show()
		return false
	}
	post.Name = Name
	post.Description = Description
	post.Importance = Importance
	post.Status = Status
	post.UpdatedAt = time.Now()

	err = a.DB.Update("Post", post)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdatePost] Fehler beim Speichern", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeletePost(id string) bool {
	err := a.DB.Delete("Post", id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeletePost] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeleteBoard(id string) bool {
	err := a.DB.Delete("Kanban", id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteBoard] Fehler", err))
		dialog.Show()
		return false
	}
	posts := a.GetPostsFromBoard(id)
	for _, post := range posts {
		err := a.DB.Delete("Post", post.Id)
		if err != nil {
			dialog := application.ErrorDialog()
			dialog.SetTitle("FEHLER!")
			dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteBoard] Fehler", err))
			dialog.Show()
			return false
		}
	}
	return true
}
