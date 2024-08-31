package domain

type Token struct {
	AccessToken  string
	RefreshToken RefreshToken
}

type RefreshToken struct {
	RootID       string
	ParentID     string
	RefreshToken string
}
