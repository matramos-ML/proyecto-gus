package model

// Usuario es la representacion de un usuario
type Usuario struct {
	Mail   string `json:"mail,omitempty"`
	Nombre string `json:"nombre"`
}

// Contains verifica si un usuario esta en la lista
func Contains(usuarios map[string]Usuario, mail string) bool {
	_, exist := usuarios[mail]
	return exist
}
