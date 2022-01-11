package mongo

import (

)

type OrderTicketModel struct {
	FirstName      string `bson:firstName`
	LastName       string `bson:lastName`
	Email          string `bson:email`
	NumberOfTicket uint32 `bson:numberOfTicket`
}