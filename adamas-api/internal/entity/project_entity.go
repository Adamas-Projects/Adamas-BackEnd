package entity

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"` 
}

type Project struct {
	ID             int64              `json:"id"`
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	Content        string             `json:"content"`
	FirstOwnerID   int                `json:"owner_id"`
	FirstOwnerName string             `json:"owner_name"`
	Owners         []*User `json:"owners"`
	Categories     []*Category        `json:"categories"`
	Comments       []*Comment         `json:"comments"`
}

type Comment struct {
	CommentID int64  `json:"comment_id"`
	UserID    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	Comment   string `json:"comment"`
}


func NewProject(title, description, content string, ownerID int) *Project {
	return &Project{
		Title:        title,
		Description:  description,
		Content:      content,
		FirstOwnerID: ownerID,
	}
}
