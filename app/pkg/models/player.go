package models

import (
	decimal  "github.com/shopspring/decimal"
)

// Player struct is the data of players of Team
type Player struct {
	ID               int
	Name             string
	Nickname         string
	Photo            string
	Price            decimal.Decimal
	Score            decimal.Decimal
	Median			 		 decimal.Decimal
	NumMatches       int
	ScoutStats		   []PlayerStats
	Team             Team
	Position 		     Position 
	Status 			     Status		
}