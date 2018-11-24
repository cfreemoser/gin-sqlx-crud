package forms

// CreatePerson is the data needed to create a new person entity.
type CreatePerson struct {
	FirstName *string `form:"firstName" json:"firstName"`
	LastName  *string `form:"lastName" json:"lastName"`
	Address   *string `form:"address" json:"address"`
	Age       *int    `form:"age" json:"age"`
}