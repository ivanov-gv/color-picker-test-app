package model

type User struct {
	Id int
}

func (u User) ToDto() UserDto {
	return UserDto{Id: u.Id}
}

type UserDto struct {
	Id int `json:"id" validate:"required"`
}

func (d UserDto) FromDto() User {
	return User{Id: d.Id}
}
