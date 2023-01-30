package lib

type Util interface {
	Contains(s []string, str string) bool
}

type util struct {
}

func NewUtil() Util {
	return &util{}
}

func (*util) Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
