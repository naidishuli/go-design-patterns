package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

// this works fine for singe threaded application
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}
