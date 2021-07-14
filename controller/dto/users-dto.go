package dto

type UserUpdateDTO struct {
	ID 				uint64 `json:"id,,string,omitempty" form:"id"`
	Name 			string `json:"name" form:"name" binding:"required"`
	Email 			string `json:"email" form:"email" binding:"required,email" validate:"email"`
	Password 		string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}

// type UserUpdateDTO struct {
// 	ID 				uint64 `json:"id" form:"id"`
// 	Password 		string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
// }