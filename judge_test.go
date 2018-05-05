package judge_test

//
// import (
// 	"github.com/boltdb/bolt"
// 	"github.com/gearnode/judge"
// 	"testing"
// 	"time"
// )
//
// type Metadata struct {
// 	Store *bolt.DB
// }
//
// var (
// 	ctx = Metadata{}
// )
//
// func init() {
// 	db, err := bolt.Open("state_test.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
// 	if err != nil {
// 		panic(err)
// 	}
// 	ctx.Store = db
// }
//
// func TestCreateOrganization(t *testing.T) {
// 	o := judge.Organization{
// 		Name:  "test company",
// 		Store: ctx.Store,
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
//
// 	if o.ID == "" {
// 		t.Error("create organization failed.")
// 	}
// }
//
// func TestDescribeOrganization(t *testing.T) {
// 	o := judge.Organization{
// 		Name:  "test company",
// 		Store: ctx.Store,
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
//
// 	expect := judge.Organization{ID: o.ID}
//
// 	if o.DescribeOrganization(&expect) != nil {
// 		t.Error("describe organization failed.")
// 	}
//
// 	if expect.Name != o.Name {
// 		t.Error("describe organization failed.")
// 	}
//
// 	if expect.ID != o.ID {
// 		t.Error("describe organization failed.")
// 	}
//
// 	expect = judge.Organization{}
// 	if o.DescribeOrganization(&expect) == nil {
// 		t.Error("describe organization failed.")
// 	}
// }
//
// func TestUpdateOrganization(t *testing.T) {
// 	o := judge.Organization{
// 		Name:  "test company",
// 		Store: ctx.Store,
// 	}
// 	if o.UpdateOrganization(&o) == nil {
// 		t.Error("update organization failed.")
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
// 	o.Name = "test company 2"
// 	if o.UpdateOrganization(&o) != nil {
// 		t.Error("update organization failed.")
// 	}
// 	expect := judge.Organization{ID: o.ID}
// 	if o.DescribeOrganization(&expect) != nil {
// 		t.Error("describe organization failed.")
// 	}
// 	if expect.Name != "test company 2" {
// 		t.Error("update organization failed.")
// 	}
// }
//
// func TestDeleteOrganization(t *testing.T) {
// 	o := judge.Organization{
// 		ID:    "notexistingid",
// 		Store: ctx.Store,
// 	}
// 	if o.DeleteOrganization(&o) == nil {
// 		t.Error("delete organization failed")
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
// 	if o.DeleteOrganization(&o) != nil {
// 		t.Error("delete organization failed")
// 	}
// 	if o.DeleteOrganization(&o) == nil {
// 		t.Error("delete organization failed")
// 	}
// }
//
// func TestCreatePolicy(t *testing.T) {
// 	o := judge.Organization{
// 		ID:    "notexistingid",
// 		Store: ctx.Store,
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
// 	p := judge.Policy{
// 		Name:        "allow all",
// 		Description: "Give all permissons",
// 		Type:        "managed",
// 		Doc:         judge.Statement{},
// 	}
// 	if o.CreatePolicy(&p) != nil {
// 		t.Error("create policy failed.")
// 	}
//
// 	if p.ID == "" {
// 		t.Error("create policy failed.")
// 	}
// }
//
// func TestDescribePolicy(t *testing.T) {
// 	o := judge.Organization{
// 		ID:    "notexistingid",
// 		Store: ctx.Store,
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
// 	p := judge.Policy{
// 		Name:        "allow all",
// 		Description: "Give all permissons",
// 		Type:        "managed",
// 		Doc:         judge.Statement{},
// 	}
// 	if o.CreatePolicy(&p) != nil {
// 		t.Error("create policy failed.")
// 	}
//
// 	if p.ID == "" {
// 		t.Error("create policy failed.")
// 	}
//
// 	expect := judge.Policy{ID: p.ID}
// 	if o.DescribePolicy(&expect) != nil {
// 		t.Error("describe policy failed.")
// 	}
// 	if expect.Name != p.Name ||
// 		expect.ID != p.ID ||
// 		expect.Type != p.Type ||
// 		expect.Description != p.Description {
// 		t.Error("describe policy failed.")
// 	}
// }
//
// func TestCreateUser(t *testing.T) {
// 	o := judge.Organization{
// 		ID:    "notexistingid",
// 		Store: ctx.Store,
// 	}
// 	if o.CreateUser(&judge.User{}) == nil {
// 		t.Error("expected function return an error.")
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
// 	u := judge.User{Name: "gearnode"}
// 	if o.CreateUser(&u) != nil {
// 		t.Error("create user failed.")
// 	}
//
// 	if u.ID == "" {
// 		t.Error("create user failed.")
// 	}
//
// 	if o.DescribeUser(&u) != nil {
// 		t.Error("describe user failed.")
// 	}
// }
//
// func TestUpdateUser(t *testing.T) {
// 	o := judge.Organization{
// 		ID:    "notexistingid",
// 		Store: ctx.Store,
// 	}
// 	if o.UpdateUser(&judge.User{}) == nil {
// 		t.Error("expected function return an error.")
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
// 	u := judge.User{Name: "gearnode"}
// 	if o.CreateUser(&u) != nil {
// 		t.Error("create user failed.")
// 	}
//
// 	u.Name = "new name"
//
// 	if o.UpdateUser(&u) != nil {
// 		t.Error("update user failed.")
// 	}
//
// 	expect := judge.User{ID: u.ID}
// 	if o.DescribeUser(&expect) != nil {
// 		t.Error("describe user failed.")
// 	}
//
// 	if expect.Name != "new name" {
// 		t.Error("update user failed.")
// 	}
//
// 	expect.ID = "notexistingid"
// 	if o.UpdateUser(&expect) == nil {
// 		t.Error("expect error was not nil.")
// 	}
//
// }
//
// func TestDescribeUser(t *testing.T) {
// 	o := judge.Organization{
// 		ID:    "notexistingid",
// 		Store: ctx.Store,
// 	}
// 	if o.DescribeUser(&judge.User{}) == nil {
// 		t.Error("expected function return an error.")
// 	}
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
// 	u := judge.User{Name: "gearnode"}
// 	if o.CreateUser(&u) != nil {
// 		t.Error("create user failed.")
// 	}
//
// 	expect := judge.User{ID: u.ID}
// 	if o.DescribeUser(&expect) != nil {
// 		t.Error("describe user failed.")
// 	}
//
// 	expect.ID = "notexistingid"
// 	if o.DescribeUser(&expect) == nil {
// 		t.Error("expect error was not nil.")
// 	}
// }
//
// func TestDeleteUser(t *testing.T) {
// 	o := judge.Organization{
// 		Name:  "test organization",
// 		Store: ctx.Store,
// 	}
// 	if o.DeleteUser(&judge.User{}) == nil {
// 		t.Error("expected function return an error.")
// 	}
//
// 	if o.CreateOrganization(&o) != nil {
// 		t.Error("create organization failed.")
// 	}
// 	u := judge.User{Name: "gearnode"}
// 	if o.CreateUser(&u) != nil {
// 		t.Error("create user failed.")
// 	}
//
// 	if o.DeleteUser(&u) != nil {
// 		t.Error("delete user failed.")
// 	}
// 	if o.DeleteUser(&u) != nil {
// 		t.Error("delete user failed.")
// 	}
// }
//
// func TestCreateGroup(t *testing.T) {
// 	o := judge.Organization{Name: "Some Org", Store: ctx.Store}
// 	if o.CreateGroup(&judge.Group{}) == nil {
// 		t.Error("does returns error, but no error has been returned")
// 	}
// 	o.CreateOrganization(&o)
//
// 	g := judge.Group{Name: "foo"}
// 	if o.CreateGroup(&g) != nil {
// 		t.Error("does create group, but error has been returned")
// 	}
// 	if g.ID == "" {
// 		t.Error("does generate an ID, but no ID has been generated")
// 	}
// }
