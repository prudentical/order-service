package persistence

type InvalidUpdateRequestError struct{}

func (InvalidUpdateRequestError) Error() string {
	return "The requested entity doesn't exists in the database"
}

type RecordNotFoundError struct{}

func (RecordNotFoundError) Error() string {
	return "Record not found"
}
