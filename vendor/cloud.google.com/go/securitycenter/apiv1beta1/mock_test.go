// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package securitycenter

import (
	emptypb "github.com/golang/protobuf/ptypes/empty"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1beta1"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockSecurityCenterServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	securitycenterpb.SecurityCenterServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockSecurityCenterServer) CreateSource(ctx context.Context, req *securitycenterpb.CreateSourceRequest) (*securitycenterpb.Source, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.Source), nil
}

func (s *mockSecurityCenterServer) CreateFinding(ctx context.Context, req *securitycenterpb.CreateFindingRequest) (*securitycenterpb.Finding, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.Finding), nil
}

func (s *mockSecurityCenterServer) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest) (*iampb.Policy, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*iampb.Policy), nil
}

func (s *mockSecurityCenterServer) GetOrganizationSettings(ctx context.Context, req *securitycenterpb.GetOrganizationSettingsRequest) (*securitycenterpb.OrganizationSettings, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.OrganizationSettings), nil
}

func (s *mockSecurityCenterServer) GetSource(ctx context.Context, req *securitycenterpb.GetSourceRequest) (*securitycenterpb.Source, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.Source), nil
}

func (s *mockSecurityCenterServer) GroupAssets(ctx context.Context, req *securitycenterpb.GroupAssetsRequest) (*securitycenterpb.GroupAssetsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.GroupAssetsResponse), nil
}

func (s *mockSecurityCenterServer) GroupFindings(ctx context.Context, req *securitycenterpb.GroupFindingsRequest) (*securitycenterpb.GroupFindingsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.GroupFindingsResponse), nil
}

func (s *mockSecurityCenterServer) ListAssets(ctx context.Context, req *securitycenterpb.ListAssetsRequest) (*securitycenterpb.ListAssetsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.ListAssetsResponse), nil
}

func (s *mockSecurityCenterServer) ListFindings(ctx context.Context, req *securitycenterpb.ListFindingsRequest) (*securitycenterpb.ListFindingsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.ListFindingsResponse), nil
}

func (s *mockSecurityCenterServer) ListSources(ctx context.Context, req *securitycenterpb.ListSourcesRequest) (*securitycenterpb.ListSourcesResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.ListSourcesResponse), nil
}

func (s *mockSecurityCenterServer) RunAssetDiscovery(ctx context.Context, req *securitycenterpb.RunAssetDiscoveryRequest) (*longrunningpb.Operation, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*longrunningpb.Operation), nil
}

func (s *mockSecurityCenterServer) SetFindingState(ctx context.Context, req *securitycenterpb.SetFindingStateRequest) (*securitycenterpb.Finding, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.Finding), nil
}

func (s *mockSecurityCenterServer) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest) (*iampb.Policy, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*iampb.Policy), nil
}

func (s *mockSecurityCenterServer) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest) (*iampb.TestIamPermissionsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*iampb.TestIamPermissionsResponse), nil
}

func (s *mockSecurityCenterServer) UpdateFinding(ctx context.Context, req *securitycenterpb.UpdateFindingRequest) (*securitycenterpb.Finding, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.Finding), nil
}

func (s *mockSecurityCenterServer) UpdateOrganizationSettings(ctx context.Context, req *securitycenterpb.UpdateOrganizationSettingsRequest) (*securitycenterpb.OrganizationSettings, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.OrganizationSettings), nil
}

func (s *mockSecurityCenterServer) UpdateSource(ctx context.Context, req *securitycenterpb.UpdateSourceRequest) (*securitycenterpb.Source, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.Source), nil
}

func (s *mockSecurityCenterServer) UpdateSecurityMarks(ctx context.Context, req *securitycenterpb.UpdateSecurityMarksRequest) (*securitycenterpb.SecurityMarks, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*securitycenterpb.SecurityMarks), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockSecurityCenter mockSecurityCenterServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	securitycenterpb.RegisterSecurityCenterServer(serv, &mockSecurityCenter)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestSecurityCenterCreateSource(t *testing.T) {
	var name string = "name3373707"
	var displayName string = "displayName1615086568"
	var description string = "description-1724546052"
	var expectedResponse = &securitycenterpb.Source{
		Name:        name,
		DisplayName: displayName,
		Description: description,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var source *securitycenterpb.Source = &securitycenterpb.Source{}
	var request = &securitycenterpb.CreateSourceRequest{
		Parent: formattedParent,
		Source: source,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateSource(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterCreateSourceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var source *securitycenterpb.Source = &securitycenterpb.Source{}
	var request = &securitycenterpb.CreateSourceRequest{
		Parent: formattedParent,
		Source: source,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateSource(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterCreateFinding(t *testing.T) {
	var name string = "name3373707"
	var parent2 string = "parent21175163357"
	var resourceName string = "resourceName979421212"
	var category string = "category50511102"
	var externalUri string = "externalUri-1385596168"
	var expectedResponse = &securitycenterpb.Finding{
		Name:         name,
		Parent:       parent2,
		ResourceName: resourceName,
		Category:     category,
		ExternalUri:  externalUri,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var findingId string = "findingId728776081"
	var finding *securitycenterpb.Finding = &securitycenterpb.Finding{}
	var request = &securitycenterpb.CreateFindingRequest{
		Parent:    formattedParent,
		FindingId: findingId,
		Finding:   finding,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateFinding(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterCreateFindingError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var findingId string = "findingId728776081"
	var finding *securitycenterpb.Finding = &securitycenterpb.Finding{}
	var request = &securitycenterpb.CreateFindingRequest{
		Parent:    formattedParent,
		FindingId: findingId,
		Finding:   finding,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateFinding(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterGetIamPolicy(t *testing.T) {
	var version int32 = 351608024
	var etag []byte = []byte("21")
	var expectedResponse = &iampb.Policy{
		Version: version,
		Etag:    etag,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedResource string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var request = &iampb.GetIamPolicyRequest{
		Resource: formattedResource,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetIamPolicy(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterGetIamPolicyError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedResource string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var request = &iampb.GetIamPolicyRequest{
		Resource: formattedResource,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetIamPolicy(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterGetOrganizationSettings(t *testing.T) {
	var name2 string = "name2-1052831874"
	var enableAssetDiscovery bool = false
	var expectedResponse = &securitycenterpb.OrganizationSettings{
		Name:                 name2,
		EnableAssetDiscovery: enableAssetDiscovery,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("organizations/%s/organizationSettings", "[ORGANIZATION]")
	var request = &securitycenterpb.GetOrganizationSettingsRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetOrganizationSettings(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterGetOrganizationSettingsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("organizations/%s/organizationSettings", "[ORGANIZATION]")
	var request = &securitycenterpb.GetOrganizationSettingsRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetOrganizationSettings(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterGetSource(t *testing.T) {
	var name2 string = "name2-1052831874"
	var displayName string = "displayName1615086568"
	var description string = "description-1724546052"
	var expectedResponse = &securitycenterpb.Source{
		Name:        name2,
		DisplayName: displayName,
		Description: description,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var request = &securitycenterpb.GetSourceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetSource(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterGetSourceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var request = &securitycenterpb.GetSourceRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetSource(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterGroupAssets(t *testing.T) {
	var nextPageToken string = ""
	var groupByResultsElement *securitycenterpb.GroupResult = &securitycenterpb.GroupResult{}
	var groupByResults = []*securitycenterpb.GroupResult{groupByResultsElement}
	var expectedResponse = &securitycenterpb.GroupAssetsResponse{
		NextPageToken:  nextPageToken,
		GroupByResults: groupByResults,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var groupBy string = "groupBy506361367"
	var request = &securitycenterpb.GroupAssetsRequest{
		Parent:  formattedParent,
		GroupBy: groupBy,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GroupAssets(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.GroupByResults[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterGroupAssetsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var groupBy string = "groupBy506361367"
	var request = &securitycenterpb.GroupAssetsRequest{
		Parent:  formattedParent,
		GroupBy: groupBy,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GroupAssets(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterGroupFindings(t *testing.T) {
	var nextPageToken string = ""
	var groupByResultsElement *securitycenterpb.GroupResult = &securitycenterpb.GroupResult{}
	var groupByResults = []*securitycenterpb.GroupResult{groupByResultsElement}
	var expectedResponse = &securitycenterpb.GroupFindingsResponse{
		NextPageToken:  nextPageToken,
		GroupByResults: groupByResults,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var groupBy string = "groupBy506361367"
	var request = &securitycenterpb.GroupFindingsRequest{
		Parent:  formattedParent,
		GroupBy: groupBy,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GroupFindings(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.GroupByResults[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterGroupFindingsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var groupBy string = "groupBy506361367"
	var request = &securitycenterpb.GroupFindingsRequest{
		Parent:  formattedParent,
		GroupBy: groupBy,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GroupFindings(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterListAssets(t *testing.T) {
	var nextPageToken string = ""
	var totalSize int32 = 705419236
	var listAssetsResultsElement *securitycenterpb.ListAssetsResponse_ListAssetsResult = &securitycenterpb.ListAssetsResponse_ListAssetsResult{}
	var listAssetsResults = []*securitycenterpb.ListAssetsResponse_ListAssetsResult{listAssetsResultsElement}
	var expectedResponse = &securitycenterpb.ListAssetsResponse{
		NextPageToken:     nextPageToken,
		TotalSize:         totalSize,
		ListAssetsResults: listAssetsResults,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var request = &securitycenterpb.ListAssetsRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListAssets(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.ListAssetsResults[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterListAssetsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var request = &securitycenterpb.ListAssetsRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListAssets(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterListFindings(t *testing.T) {
	var nextPageToken string = ""
	var totalSize int32 = 705419236
	var findingsElement *securitycenterpb.Finding = &securitycenterpb.Finding{}
	var findings = []*securitycenterpb.Finding{findingsElement}
	var expectedResponse = &securitycenterpb.ListFindingsResponse{
		NextPageToken: nextPageToken,
		TotalSize:     totalSize,
		Findings:      findings,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var request = &securitycenterpb.ListFindingsRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListFindings(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Findings[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterListFindingsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var request = &securitycenterpb.ListFindingsRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListFindings(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterListSources(t *testing.T) {
	var nextPageToken string = ""
	var sourcesElement *securitycenterpb.Source = &securitycenterpb.Source{}
	var sources = []*securitycenterpb.Source{sourcesElement}
	var expectedResponse = &securitycenterpb.ListSourcesResponse{
		NextPageToken: nextPageToken,
		Sources:       sources,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var request = &securitycenterpb.ListSourcesRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListSources(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Sources[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterListSourcesError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var request = &securitycenterpb.ListSourcesRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListSources(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterRunAssetDiscovery(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	any, err := ptypes.MarshalAny(expectedResponse)
	if err != nil {
		t.Fatal(err)
	}
	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], &longrunningpb.Operation{
		Name:   "longrunning-test",
		Done:   true,
		Result: &longrunningpb.Operation_Response{Response: any},
	})

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var request = &securitycenterpb.RunAssetDiscoveryRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	respLRO, err := c.RunAssetDiscovery(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	err = respLRO.Wait(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestSecurityCenterRunAssetDiscoveryError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = nil
	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], &longrunningpb.Operation{
		Name: "longrunning-test",
		Done: true,
		Result: &longrunningpb.Operation_Error{
			Error: &status.Status{
				Code:    int32(errCode),
				Message: "test error",
			},
		},
	})

	var formattedParent string = fmt.Sprintf("organizations/%s", "[ORGANIZATION]")
	var request = &securitycenterpb.RunAssetDiscoveryRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	respLRO, err := c.RunAssetDiscovery(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	err = respLRO.Wait(context.Background())

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestSecurityCenterSetFindingState(t *testing.T) {
	var name2 string = "name2-1052831874"
	var parent string = "parent-995424086"
	var resourceName string = "resourceName979421212"
	var category string = "category50511102"
	var externalUri string = "externalUri-1385596168"
	var expectedResponse = &securitycenterpb.Finding{
		Name:         name2,
		Parent:       parent,
		ResourceName: resourceName,
		Category:     category,
		ExternalUri:  externalUri,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("organizations/%s/sources/%s/findings/%s", "[ORGANIZATION]", "[SOURCE]", "[FINDING]")
	var state securitycenterpb.Finding_State = securitycenterpb.Finding_STATE_UNSPECIFIED
	var startTime *timestamppb.Timestamp = &timestamppb.Timestamp{}
	var request = &securitycenterpb.SetFindingStateRequest{
		Name:      formattedName,
		State:     state,
		StartTime: startTime,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SetFindingState(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterSetFindingStateError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("organizations/%s/sources/%s/findings/%s", "[ORGANIZATION]", "[SOURCE]", "[FINDING]")
	var state securitycenterpb.Finding_State = securitycenterpb.Finding_STATE_UNSPECIFIED
	var startTime *timestamppb.Timestamp = &timestamppb.Timestamp{}
	var request = &securitycenterpb.SetFindingStateRequest{
		Name:      formattedName,
		State:     state,
		StartTime: startTime,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SetFindingState(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterSetIamPolicy(t *testing.T) {
	var version int32 = 351608024
	var etag []byte = []byte("21")
	var expectedResponse = &iampb.Policy{
		Version: version,
		Etag:    etag,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedResource string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var policy *iampb.Policy = &iampb.Policy{}
	var request = &iampb.SetIamPolicyRequest{
		Resource: formattedResource,
		Policy:   policy,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SetIamPolicy(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterSetIamPolicyError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedResource string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var policy *iampb.Policy = &iampb.Policy{}
	var request = &iampb.SetIamPolicyRequest{
		Resource: formattedResource,
		Policy:   policy,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SetIamPolicy(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterTestIamPermissions(t *testing.T) {
	var expectedResponse *iampb.TestIamPermissionsResponse = &iampb.TestIamPermissionsResponse{}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var formattedResource string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var permissions []string = nil
	var request = &iampb.TestIamPermissionsRequest{
		Resource:    formattedResource,
		Permissions: permissions,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.TestIamPermissions(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterTestIamPermissionsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var formattedResource string = fmt.Sprintf("organizations/%s/sources/%s", "[ORGANIZATION]", "[SOURCE]")
	var permissions []string = nil
	var request = &iampb.TestIamPermissionsRequest{
		Resource:    formattedResource,
		Permissions: permissions,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.TestIamPermissions(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterUpdateFinding(t *testing.T) {
	var name string = "name3373707"
	var parent string = "parent-995424086"
	var resourceName string = "resourceName979421212"
	var category string = "category50511102"
	var externalUri string = "externalUri-1385596168"
	var expectedResponse = &securitycenterpb.Finding{
		Name:         name,
		Parent:       parent,
		ResourceName: resourceName,
		Category:     category,
		ExternalUri:  externalUri,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var finding *securitycenterpb.Finding = &securitycenterpb.Finding{}
	var request = &securitycenterpb.UpdateFindingRequest{
		Finding: finding,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateFinding(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterUpdateFindingError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var finding *securitycenterpb.Finding = &securitycenterpb.Finding{}
	var request = &securitycenterpb.UpdateFindingRequest{
		Finding: finding,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateFinding(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterUpdateOrganizationSettings(t *testing.T) {
	var name string = "name3373707"
	var enableAssetDiscovery bool = false
	var expectedResponse = &securitycenterpb.OrganizationSettings{
		Name:                 name,
		EnableAssetDiscovery: enableAssetDiscovery,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var organizationSettings *securitycenterpb.OrganizationSettings = &securitycenterpb.OrganizationSettings{}
	var request = &securitycenterpb.UpdateOrganizationSettingsRequest{
		OrganizationSettings: organizationSettings,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateOrganizationSettings(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterUpdateOrganizationSettingsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var organizationSettings *securitycenterpb.OrganizationSettings = &securitycenterpb.OrganizationSettings{}
	var request = &securitycenterpb.UpdateOrganizationSettingsRequest{
		OrganizationSettings: organizationSettings,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateOrganizationSettings(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterUpdateSource(t *testing.T) {
	var name string = "name3373707"
	var displayName string = "displayName1615086568"
	var description string = "description-1724546052"
	var expectedResponse = &securitycenterpb.Source{
		Name:        name,
		DisplayName: displayName,
		Description: description,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var source *securitycenterpb.Source = &securitycenterpb.Source{}
	var request = &securitycenterpb.UpdateSourceRequest{
		Source: source,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateSource(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterUpdateSourceError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var source *securitycenterpb.Source = &securitycenterpb.Source{}
	var request = &securitycenterpb.UpdateSourceRequest{
		Source: source,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateSource(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestSecurityCenterUpdateSecurityMarks(t *testing.T) {
	var name string = "name3373707"
	var expectedResponse = &securitycenterpb.SecurityMarks{
		Name: name,
	}

	mockSecurityCenter.err = nil
	mockSecurityCenter.reqs = nil

	mockSecurityCenter.resps = append(mockSecurityCenter.resps[:0], expectedResponse)

	var securityMarks *securitycenterpb.SecurityMarks = &securitycenterpb.SecurityMarks{}
	var request = &securitycenterpb.UpdateSecurityMarksRequest{
		SecurityMarks: securityMarks,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateSecurityMarks(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockSecurityCenter.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestSecurityCenterUpdateSecurityMarksError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockSecurityCenter.err = gstatus.Error(errCode, "test error")

	var securityMarks *securitycenterpb.SecurityMarks = &securitycenterpb.SecurityMarks{}
	var request = &securitycenterpb.UpdateSecurityMarksRequest{
		SecurityMarks: securityMarks,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.UpdateSecurityMarks(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
