package models

type MovieResponse struct{
	Title string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"releaseDate"`
	Rating string `json:"rating"`
	Actors []string `json:"actors"`
}


type MovieResponseActor struct{
	Id string `json:"id,omitempty"`
	Title string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"releaseDate"`
	Rating string `json:"rating"`
	Actors []Actor `json:"actors"`
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

type MovieDetail struct {
    Title       string  `json:"title"`
    ReleaseDate string  `json:"release_date"`
    Rating      float64 `json:"rating"`
}

type ActorWithMovies struct {
    ID     int            `json:"id"`
    Name   string         `json:"name"`
    Movies []MovieDetail `json:"movies"`
}