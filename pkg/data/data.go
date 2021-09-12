package data

type Driver interface {
	Connect()
	Create()
	Read()
	Update()
	Delete()
}
