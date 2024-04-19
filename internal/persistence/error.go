package persistence

type InvalidUpdateRequest struct{}

func (InvalidUpdateRequest) Error() string {
	return "The requested entity doesn't exists in the database"
}
