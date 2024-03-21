package utils

func Assert(err error) error {
	if err != nil {
		return err // Return nil pool and the error encountered
	}
	return nil
}