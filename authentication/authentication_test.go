package authentication_test

import (
	"testing"

	"github.com/gearnode/judge/authentication"
)

func ok(t *testing.T, err error) {
	if err != nil {
		t.FailNow()
	}
}

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Error(msg)
	}
}

func TestAuthenticate(t *testing.T) {
	cred, err := authentication.Authenticate("foo", "bar")
	ok(t, err)
	assert(t, cred.OrganizationID == "testOrgID", "bad org id")
	assert(t, cred.UserID == "testUserID", "bad user id")
}
