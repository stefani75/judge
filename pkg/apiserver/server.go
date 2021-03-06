/*
Copyright 2018 The Judge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apiserver // import "github.com/gearnode/judge/pkg/apiserver"

import (
	"github.com/gearnode/judge/pkg/apiserver/v1alpha1"
	"github.com/gearnode/judge/pkg/authorize"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/policy/resource"
	"github.com/gearnode/judge/pkg/storage/memory"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server represent a server instance. This struct store the
// server configuration.
type Server struct{}

// Register register to gRPC Service
func Register(srv *grpc.Server) {
	apiserver := &Server{}
	v1alpha1.RegisterJudgeServer(srv, apiserver)
}

var (
	store = memorystore.NewMemoryStore()
)

// Authorize implement judge.api.v1alpha1.Judge.Authorize
func (s *Server) Authorize(ctx context.Context, in *v1alpha1.AuthorizeRequest) (*v1alpha1.AuthorizeResponse, error) {
	log.Info("Receive Authorize Request to execute")

	something := orn.ORN{}
	orn.Unmarshal(in.GetSomething(), &something)
	_, err := authorize.Authorize(store, orn.ORN{}, in.GetWhat(), something, in.GetContext())

	if err != nil {
		return &v1alpha1.AuthorizeResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	return &v1alpha1.AuthorizeResponse{}, nil
}

// GetPolicy implement judge.api.v1alpha1.Judge.GetPolicy
func (s *Server) GetPolicy(ctx context.Context, in *v1alpha1.GetPolicyRequest) (*v1alpha1.Policy, error) {
	log.Info("Receive GetPolicy Request to execute")

	data, err := store.Describe("policies", in.GetOrn())
	if err != nil {
		return &v1alpha1.Policy{}, status.Error(codes.NotFound, "the policy with the orn "+in.GetOrn()+" is not found")
	}

	pol := data.(policy.Policy)

	res := v1alpha1.Policy{
		Orn:         orn.Marshal(&pol.ORN),
		Name:        pol.Name,
		Description: pol.Description,
		Document: &v1alpha1.Document{
			Version:    "v1alpha1",
			Statements: make([]*v1alpha1.Statement, len(pol.Document.Statement)),
		},
	}

	for i, v := range pol.Document.Statement {

		stm := v1alpha1.Statement{Effect: v.Effect, Actions: v.Action}

		for _, v := range v.Resource {
			r := resource.Marshal(&v)
			stm.Resources = append(stm.Resources, r)
		}
		res.Document.Statements[i] = &stm
	}

	return &res, nil
}

// ListPolicies implement judge.api.vbeta1.Judge.ListPolicies
func (s *Server) ListPolicies(ctx context.Context, in *v1alpha1.ListPoliciesRequest) (*v1alpha1.ListPoliciesResponse, error) {
	log.Info("Receive ListPoliciesRequest Request to execute")
	return &v1alpha1.ListPoliciesResponse{}, nil
}

// CreatePolicy implement judge.api.v1alpha1.Judge.CreatePolicy
func (s *Server) CreatePolicy(ctx context.Context, in *v1alpha1.CreatePolicyRequest) (*v1alpha1.Policy, error) {
	log.Info("Receive CreatePolicy Request to execute")

	res := v1alpha1.Policy{}

	pol, err := policy.NewPolicy(in.GetPolicy().GetName(), in.GetPolicy().GetDescription())
	if err != nil {
		return &v1alpha1.Policy{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetOrn() != "" {
		var id orn.ORN
		err := orn.Unmarshal(in.GetOrn(), &id)
		if err != nil {
			return &v1alpha1.Policy{}, status.Error(codes.InvalidArgument, err.Error())
		}
		pol.ORN = id
	}
	res.Orn = orn.Marshal(&pol.ORN)
	res.Name = pol.Name
	res.Description = pol.Description

	if _, err := store.Describe("policies", res.Orn); err == nil {
		return &v1alpha1.Policy{}, status.Error(codes.AlreadyExists, "the policy already exists")
	}

	for _, v := range in.GetPolicy().GetDocument().GetStatements() {
		stm, err := policy.NewStatement(v.GetEffect(), v.GetActions(), v.GetResources())
		if err != nil {
			return &v1alpha1.Policy{}, status.Error(codes.InvalidArgument, err.Error())
		}
		pol.Document.Statement = append(pol.Document.Statement, *stm)
	}

	store.Put("policies", res.Orn, *pol)

	statements := make([]*v1alpha1.Statement, len(pol.Document.Statement))
	for i, v := range pol.Document.Statement {

		stm := v1alpha1.Statement{Effect: v.Effect, Actions: v.Action}

		for _, v := range v.Resource {
			r := resource.Marshal(&v)
			stm.Resources = append(stm.Resources, r)
		}
		statements[i] = &stm
	}

	doc := v1alpha1.Document{
		Version:    pol.Document.Version,
		Statements: statements,
	}

	res.Document = &doc

	return &res, nil
}

// UpdatePolicy implement judge.api.v1alpha1.Judge.UpdatePolicy
func (s *Server) UpdatePolicy(ctx context.Context, in *v1alpha1.UpdatePolicyRequest) (*v1alpha1.Policy, error) {
	log.Info("Receive UpdatePolicy Resquest to execute")

	res := v1alpha1.Policy{}
	pol, err := policy.NewPolicy(in.GetPolicy().GetName(), in.GetPolicy().GetDescription())
	if err != nil {
		return &v1alpha1.Policy{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetPolicy().GetOrn() != "" {
		var id orn.ORN
		err := orn.Unmarshal(in.GetPolicy().GetOrn(), &id)
		if err != nil {
			return &v1alpha1.Policy{}, status.Error(codes.InvalidArgument, err.Error())
		}
		pol.ORN = id
	}

	res.Orn = orn.Marshal(&pol.ORN)
	res.Name = pol.Name
	res.Description = pol.Description

	for _, v := range in.GetPolicy().GetDocument().GetStatements() {
		stm, err := policy.NewStatement(v.GetEffect(), v.GetActions(), v.GetResources())
		if err != nil {
			return &v1alpha1.Policy{}, status.Error(codes.InvalidArgument, err.Error())
		}
		pol.Document.Statement = append(pol.Document.Statement, *stm)
	}

	store.Put("policies", res.Orn, *pol)

	statements := make([]*v1alpha1.Statement, len(pol.Document.Statement))
	for i, v := range pol.Document.Statement {

		stm := v1alpha1.Statement{Effect: v.Effect, Actions: v.Action}

		for _, v := range v.Resource {
			r := resource.Marshal(&v)
			stm.Resources = append(stm.Resources, r)
		}
		statements[i] = &stm
	}

	doc := v1alpha1.Document{
		Version:    pol.Document.Version,
		Statements: statements,
	}

	res.Document = &doc

	return &res, nil
}

// DeletePolicy implement judge.api.v1alpha1.Judge.DeletePolicy
func (s *Server) DeletePolicy(ctx context.Context, in *v1alpha1.DeletePolicyRequest) (*empty.Empty, error) {
	log.Printf("Receive DeletePolicy Resquest to execute")
	return &empty.Empty{}, nil
}
