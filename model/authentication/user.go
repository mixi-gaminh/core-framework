package model

type (
	// UserProfile - UserProfile
	UserProfile struct {
		ID          string                 `bson:"_id" json:"id"`
		UserID      string                 `bson:"user_id" json:"user_id"`
		Username    string                 `bson:"username" json:"username" validate:"required,min=4,max=20"`
		Password    string                 `bson:"password" json:"password,omitempty" validate:"required,min=8"`
		Firstname   string                 `bson:"first_name" json:"first_name,omitempty" validate:"omitempty,min=2,max=30"`
		Lastname    string                 `bson:"last_name" json:"last_name,omitempty" validate:"omitempty,min=2,max=30"`
		Email       string                 `bson:"email" json:"email" validate:"email,min=2,max=320"`
		PhoneNumber string                 `bson:"phone_number" json:"phone_number" validate:"number,min=9,max=11"`
		Address     string                 `bson:"address" json:"address,omitempty" validate:"omitempty,min=0,max=50"`
		DateOfBirth string                 `bson:"date_of_birth" json:"date_of_birth,omitempty" validate:"omitempty,number,max=10"`
		Status      string                 `bson:"status" json:"status,omitempty" validate:"omitempty,number,max=20"`
		CheckExist  []string               `bson:"check_exist" json:"check_exist" validate:"required"`
		UserInfo    map[string]interface{} `bson:"user_info" json:"user_info"`
	}
)
