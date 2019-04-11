package git

type LocalStore interface {
	Adder
	Getter
	GetAller
	Deleter
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
