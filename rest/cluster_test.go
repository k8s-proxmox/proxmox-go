package rest

import "context"

func (s *TestSuite) TestGetNextID() {
	nextid, err := s.restclient.GetNextID(context.TODO())
	if err != nil {
		s.T().Errorf("failed to get next id: %v", err)
	}
	s.T().Logf("get nextID: %d", nextid)
}

func (s *TestSuite) TestGetJoinConfig() {
	c, err := s.restclient.GetJoinConfig(context.Background())
	if err != nil {
		s.T().Errorf("failed to get join config: %v", err)
	}
	s.T().Logf("get join config: %v", c)
}
