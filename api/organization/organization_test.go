package organization

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/dylanrhysscott/terrarium/pkg/types/mock_types"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var orgs []*types.Organization

func mockHandlerWithOrgStore(mockStore *mock_types.MockOrganizationStore, method string, path string) (*httptest.ResponseRecorder, *http.Request, http.Handler) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, nil)
	// TODO: Fix nill mocks for tests
	api := NewOrganizationAPI(nil, "", mockStore, nil, &types.TerrariumAPIResponseWriter{}, &types.TerrariumAPIErrorHandler{})
	return rr, req, api.ListOrganizationsHandler()
}

func marshalToServerResponse(body *bytes.Buffer) (*types.TerrariumServerResponse, error) {
	resp, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	tr := &types.TerrariumServerResponse{}
	err = json.Unmarshal(resp, tr)
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func marshalToOrgList(body *bytes.Buffer) ([]*types.Organization, int, error) {
	resp, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, 0, err
	}
	data := &types.TerrariumDataResponse{}
	err = json.Unmarshal(resp, data)
	if err != nil {
		return nil, 0, err
	}
	orgsData, err := json.Marshal(data.Data)
	if err != nil {
		return nil, 0, err
	}
	orgs := []*types.Organization{}
	err = json.Unmarshal(orgsData, &orgs)
	if err != nil {
		return nil, 0, err
	}
	log.Printf("%v", orgs)
	return orgs, data.Code, nil
}

func TestListOrganizationsErrorHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	expected := "Internal Server Error - Some err"
	mockOrganizationStore := mock_types.NewMockOrganizationStore(ctrl)
	mockOrganizationStore.EXPECT().ReadAll(gomock.Eq(0), gomock.Eq(0)).Return(nil, errors.New("Some err")).Times(1)
	rr, req, h := mockHandlerWithOrgStore(mockOrganizationStore, "GET", "/")
	h.ServeHTTP(rr, req)
	if rr.Result().StatusCode != http.StatusInternalServerError {
		t.Fatalf("Expected status code to equal %d got %d", http.StatusInternalServerError, rr.Result().StatusCode)
	}
	tr, err := marshalToServerResponse(rr.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	if tr.Code != http.StatusInternalServerError {
		t.Fatalf("Expected response body property code to equal %d got %d", http.StatusInternalServerError, tr.Code)
	}
	if tr.Message != expected {
		t.Fatalf("Expected response body property message to equal %s got %s", expected, tr.Message)
	}
}

func TestListOrganizationsSuccessHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOrganizationStore := mock_types.NewMockOrganizationStore(ctrl)
	mockOrganizationStore.EXPECT().ReadAll(gomock.Eq(0), gomock.Eq(0)).Return(orgs, nil).Times(1)
	rr, req, h := mockHandlerWithOrgStore(mockOrganizationStore, "GET", "/")
	h.ServeHTTP(rr, req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Fatalf("Expected status code to equal %d got %d", http.StatusOK, rr.Result().StatusCode)
	}
	orgs, code, err := marshalToOrgList(rr.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	if code != http.StatusOK {
		t.Fatalf("Expected response body property code to equal %d got %d", http.StatusOK, code)
	}
	if len(orgs) != 2 {
		t.Fatalf("Expected 2 organization to be in response body")
	}
	testOrg := orgs[0]
	if testOrg.ID != orgs[0].ID {
		t.Fatal("Expected ID for org to be 1")
	}
	if testOrg.Name != "Test Org" {
		t.Fatal("Expected name for org to be Test Org")
	}
	if testOrg.Email != "test@test.com" {
		t.Fatal("Expected email for org to be test@test.com")
	}
	if testOrg.CreatedOn != "2020-01-01T00:00:00Z" {
		t.Fatal("Expected CreateOn to be 2020-01-01T00:00:00Z")
	}
}

func TestMain(m *testing.M) {
	orgs = []*types.Organization{
		{
			ID:        primitive.NewObjectID(),
			Name:      "Test Org",
			Email:     "test@test.com",
			CreatedOn: "2020-01-01T00:00:00Z",
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Test Org2",
			Email:     "test@test.com",
			CreatedOn: "2020-01-01T00:00:00Z",
		},
	}
	m.Run()
}
