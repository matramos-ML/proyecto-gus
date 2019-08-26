package model

// Usuario es la representacion de un usuario
type Usuario struct {
	Mail   string `json:"mail"`
	Nombre string `json:"nombre"`
}

// Contains verifica si un usuario esta en la lista
func Contains(lista []Usuario, mail string) bool {
	for _, user := range lista {
		if user.Mail == mail {
			return true
		}
	}
	return false
}
