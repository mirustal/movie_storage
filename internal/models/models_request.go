package models

type MovieRequest struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Rating string `json:"rating,omitempty"`
	ActorIds []string `json:"actorIds,omitempty"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role,omitempty"`
}


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ParamsFilmRequest struct {
	typeSort  string
	FilmQuery  string
	ActorQuery string
}
