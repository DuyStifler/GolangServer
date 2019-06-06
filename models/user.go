package models

type User struct {
	userID  string `json:"user_id"`
	gold    int    `json:"gold"`
	cash    int    `json:"cash"`
	stamina int    `json:"stamina"`
	level   int    `json:"level"`
	exp     int    `json:"exp"`
}

func (u *User) Exp() int {
	return u.exp
}

func (u *User) SetExp(exp int) {
	u.exp = exp
}

func (u *User) Level() int {
	return u.level
}

func (u *User) SetLevel(level int) {
	u.level = level
}

func (u *User) Cash() int {
	return u.cash
}

func (u *User) SetCash(cash int) {
	u.cash = cash
}

func (u *User) Gold() int {
	return u.gold
}

func (u *User) SetGold(gold int) {
	u.gold = gold
}

func (u *User) UserID() string {
	return u.userID
}

func (u *User) SetUserID(userID string) {
	u.userID = userID
}

func (u *User) Stamina() int {
	return u.stamina
}

func (u *User) SetStamina(stamina int) {
	u.stamina = stamina
}
