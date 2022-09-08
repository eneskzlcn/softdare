package home

const DomainName = "home"

type UserSessionData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type SessionData struct {
	IsLoggedIn bool            `json:"is_logged_in"`
	User       UserSessionData `json:"user"`
}

type homeData struct {
	Session SessionData
}
