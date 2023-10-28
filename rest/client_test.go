package rest

import (
	"context"
	"fmt"
)

func (s *TestSuite) TestRateLimit() {
	s.restclient.SetMaxQPS(10)
	for i := 0; i < 13; i++ {
		_, err := s.restclient.GetVersion(context.TODO())
		if err != nil {
			s.T().Errorf("failed to ger version: %v", err)
		}
		fmt.Println(i)
	}
}
