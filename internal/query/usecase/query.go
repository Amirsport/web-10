package usecase

func (u *Usecase) FetchQuery() (string, error) {
	msg, err := u.p.SelectQuery()
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) SetQuery(msg string) error {
	isExist, err := u.p.CheckQueryExitByMsg(msg)
	if err != nil {
		return err
	}

	if isExist {
		return nil
	}

	err = u.p.InsertQuery(msg)
	if err != nil {
		return err
	}

	return nil
}
