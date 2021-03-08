package auth

type Auth struct {
	Token string `json:"token,omitempty" pg:",pk"`
}
