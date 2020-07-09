package roles

type Role struct {
	Code        string
	Name        string
	Permissions []Permission
}

type Permission struct {
	Permission string
	Module     Module
}

type Module struct {
	Code string
	Name string
}
