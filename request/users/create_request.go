package users

type CrateRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}
