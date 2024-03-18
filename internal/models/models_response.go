package models

type MovieResponse struct{
	Title string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"releaseDate"`
	Rating string `json:"rating"`
}

type ActorResponse struct{
	Name   string `json:"name"`
	Gender string `json:"gender"`
	BirthDate string `json:"birthDate"`
}
type UserResponse struct{
	UserId string
	RefreshToken string
}