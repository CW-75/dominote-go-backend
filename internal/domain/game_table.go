package domain

// Player representa a un jugador básico.
// Ajusta esta estructura según los datos reales del jugador que necesites.
type Player struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

// GameTable representa la entidad de una mesa de juego.
type GameTable struct {
	UID                     string   `json:"uid"`
	Name                    string   `json:"name"`
	Players                 []Player `json:"players"`
	// Punteros para campos opcionales, lo que permite que sean omitidos en el JSON si son nil
	BestScore               *int     `json:"best_score,omitempty"`
	PlayerWithMostVictories *string  `json:"player_with_most_victories,omitempty"`
}
