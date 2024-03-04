package handler

func handleErr(err error) error {
	if err == nil {
		return nil
	}

	// code := coderr.ErrorCode(err)
	// TODO: finish error handling
	return err
}
