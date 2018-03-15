package models

func GetContact(Id uint) *Contact {

	contact := &Contact{}
	GetConn().First(contact, Id)
	return contact
}

func GetUserContact(userId interface{}) *[]Contact {

	var contacts []Contact

	GetConn().Table("contacts").Where("user_id = ?", userId).Find(&contacts)
	return &contacts
}