package storage

func (s *Storage) UpdateTransactionStatus(id uint64, newStatus string) error {
	return nil
}

func (s *Storage) ReadTransactionByID(id uint64) *Transaction {
	return &Transaction{}
}

func (s *Storage) ReadAllUsersTransactionsByID(userid string) []Transaction {
	return nil
}

func (s *Storage) ReadAllUsersTransactionsByEmail(email string) []Transaction {
	return nil
}
