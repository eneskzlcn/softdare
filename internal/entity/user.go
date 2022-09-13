package entity

import "time"

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/*UserSessionData keeps the values that provide user identify.
The fields can be change for usage like just with id or password.
Do not touch the name of the struct due to dependent parts. If you
change the struct name be sure to reformat related parts.

*/
type UserSessionData struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
