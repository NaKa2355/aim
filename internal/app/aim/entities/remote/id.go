package remote

type ID string

func NewID(id string) (ID, error) {
	return ID(id), nil
}
