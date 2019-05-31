package forms

type Login struct {
	Username string `form:"username"`
	Pwd string `form:"pwd"`
}
