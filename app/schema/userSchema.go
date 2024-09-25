package schema

type UserUpdateRequest struct {
	Id string `json:"origin_url" binding:"required,min=6,max=12"`
	Pw string `json:"short_path" binding:"required,min=8,max=16"`
}
