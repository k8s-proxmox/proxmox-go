package rest

func (s *TestSuite) TestGetVersion() {
	ver, err := s.restclient.GetVersion()
	if err != nil {
		s.T().Errorf("failed to ger version: %v", err)
	}
	s.T().Logf("get version: %v", *ver)
}
