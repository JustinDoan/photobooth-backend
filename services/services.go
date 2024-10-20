package services

type Services struct {
	Forge *Forge
}

func NewServices() *Services {
	return &Services{
		Forge: NewForge(),
	}
}
