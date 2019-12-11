package models

// Formation is the schema of game team
type Formation struct {
	ID        	int
	Goalkeeper	int
	Attackers 	int
	Defenders 	int
}