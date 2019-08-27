package model

// Usuario es la representacion de un usuario
type Usuario struct {
	Mail   string `json:"mail"`
	Nombre string `json:"nombre"`
}

// Contains verifica si un usuario esta en la lista
func Contains(usuarios map[string]Usuario, mail string) bool {
	if _, exist := usuarios[mail]; exist {
		return true
	}
	return false
}
