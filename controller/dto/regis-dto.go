package dto

type RegisDTO struct {
	Name 		string `json:"name" form:"name" binding:"required"`
	Email 		string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password 	string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}