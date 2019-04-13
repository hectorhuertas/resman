package git

type LocalStore interface {
	LocationGenerator
	Adder
	Getter
	GetAller
	Deleter
}

type LocationGenerator interface {
	GenerateLocation() string
}

type Adder interface {
	Add(Local) error
}

type Getter interface {
	Get(string) (Local, error)
}

type GetAller interface {
	GetAll() []Local
}

type Deleter interface {
	Delete(string)
}
