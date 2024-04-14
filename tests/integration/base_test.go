package integration

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
	"time"
)

type ServerTestSuite struct {
	suite.Suite
	ctx         context.Context
	cleanUpTest func()
	client      *resty.Client
}

func (s *ServerTestSuite) TestBannerPipeline() {
	type Banner struct {
		ID        int64      `json:"banner_id"`
		TagIds    []int64    `json:"tag_ids"`
		FeatureId int64      `json:"feature_id"`
		Content   string     `json:"content"`
		IsActive  bool       `json:"is_active"`
		CreatedAt *time.Time `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}
	banner := Banner{}

	//1
	r, err := s.client.R().SetBody(Banner{
		TagIds:    []int64{1, 2},
		FeatureId: 1,
		Content:   "Banner 1",
		IsActive:  true,
	}).SetResult(&banner).Post("/banner")
	s.T().Log(string(r.Body()))
	require.Error(s.T(), err)
	require.Equalf(s.T(), http.StatusUnauthorized, r.StatusCode(), "Do not authorize, 'token' is empty")

	//2
	r, err = s.client.R().SetHeader("token", "user_token").SetBody(Banner{
		TagIds:    []int64{1, 2},
		FeatureId: 1,
		Content:   "Banner 1",
		IsActive:  true,
	}).SetResult(&banner).Post("/banner")
	s.T().Log(string(r.Body()))
	require.Error(s.T(), err)
	require.Equalf(s.T(), http.StatusForbidden, r.StatusCode(), "Do not have permission, 'token' is not suitable. POST /banner")

	//3
	r, err = s.client.R().SetHeader("token", "user_token").SetBody("").SetResult(&banner).Post("/banner")
	s.T().Log(string(r.Body()))
	require.Error(s.T(), err)
	require.Equalf(s.T(), http.StatusBadRequest, r.StatusCode(), "Do not have required fields. POST /banner")

	//4
	r, err = s.client.R().SetHeader("token", "admin_token").SetBody(Banner{
		TagIds:    []int64{1, 2},
		FeatureId: 1,
		Content:   "Banner 1",
		IsActive:  true,
	}).SetResult(&banner).Post("/banner")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	require.Equalf(s.T(), http.StatusCreated, r.StatusCode(), "Valid POST /banner")
	require.NotEmpty(s.T(), banner.ID)

	banner = Banner{}

	//5
	r, err = s.client.R().SetHeader("token", "user_token").SetBody(Banner{
		TagIds:    []int64{1, 2},
		FeatureId: 2,
	}).SetResult(&banner).Get("/user_banner")
	s.T().Log(string(r.Body()))
	require.Error(s.T(), err)
	require.Equalf(s.T(), http.StatusNotFound, r.StatusCode(), "Not found banner with query fields. GET /user_banner")

	//6
	r, err = s.client.R().SetHeader("token", "user_token").SetBody(Banner{
		TagIds:    []int64{1, 2},
		FeatureId: 1,
	}).SetResult(&banner).Get("/user_banner")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	require.Equalf(s.T(), http.StatusOK, r.StatusCode(), "Valid GET /user_banner")
	require.NotEmpty(s.T(), banner.Content)

	//7
	r, err = s.client.R().SetHeader("token", "admin_token").SetBody(Banner{
		TagIds:    []int64{1, 2},
		FeatureId: 1,
	}).SetResult(&banner).Get("/user_banner")
	s.T().Log(string(r.Body()))
	require.NoError(s.T(), err)
	require.Equalf(s.T(), http.StatusOK, r.StatusCode(), "Valid GET /user_banner")
	require.NotEmpty(s.T(), banner.Content)

}

func (s *ServerTestSuite) SetupTest() {
	time.Sleep(time.Second / 2)

	s.ctx, s.cleanUpTest = context.WithTimeout(context.Background(), time.Second)

	c := resty.New()
	c.SetBaseURL(os.Getenv("AVITO_TECH_BACKEND"))

	s.client = c
}

func (s *ServerTestSuite) TearDownTest() {
	s.cleanUpTest()
}

func TestBase(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
