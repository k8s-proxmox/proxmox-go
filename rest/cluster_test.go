package rest

func (s *TestSuite) TestGetNextID() {
	nextid, err := s.restclient.GetNextID()
	if err != nil {
		s.T().Errorf("failed to get next id: %v", err)
	}
	s.T().Logf("get nextID: %d", nextid)
}
