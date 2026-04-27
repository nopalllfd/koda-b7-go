package model

type Education struct {
	Name  string
	Major string
}
type User struct {
	Name      string
	Photo     string
	Email     string
	Age       int
	Telp      string
	IsMarried bool
	Education []Education
}

func TampilinData() User {
	nopal := User{
		Name:      "Nopal",
		Photo:     "Link",
		Email:     "email",
		Age:       19,
		Telp:      "01010101",
		IsMarried: false,
		Education: []Education{
			{
				Name:  "Koda",
				Major: "Talent",
			},
		},
	}
	return nopal
}
