package judge

// import (
// 	"encoding/json"
// 	"errors"
// 	"regexp"
// 	"strings"
//
// 	"github.com/boltdb/bolt"
// 	"github.com/satori/go.uuid"
// )
//
// // An Judge user is an entity that you create in Judge to represent
// // the person or service that uses it to interact with Judge. A user
// // in Judge consists of a name and credentials.
// type User struct {
// 	ID       string
// 	ORN      string
// 	Name     string
// 	Groups   map[string]bool
// 	Policies map[string]bool
// }
//
// // An Judge group is a collection of Judge users. Groups let you specify
// // permissions for multiple users, which can make it easier to manage
// // the permissions for those users.
// type Group struct {
// 	ID       string
// 	ORN      string
// 	Name     string
// 	Users    map[string]bool
// 	Policies map[string]bool
// }
//
// // Impersonation
// // use for assume role
// type Role struct {
// 	ID       string
// 	ORN      string
// 	Name     string
// 	Policies map[string]bool
// }
//
// // A policy is an entity in Judge that, when attached to an identity, defines
// // their permissions. Judge evaluates these policies when a principal, such as
// // a user, makes a request. Permissions in the policies determine whether the
// // request is allowed or denied.
// type Policy struct {
// 	ID          string
// 	ORN         string
// 	Name        string
// 	Description string
// 	Type        string
// 	Doc         Statement
// }
//
// type Statement struct {
// 	Version   string
// 	Statement []struct {
// 		Effect   string
// 		Action   []string
// 		Resource []string
// 	}
// }
//
// func (o *Organization) OpenBucket(tx *bolt.Tx) (*bolt.Bucket, error) {
// 	root := tx.Bucket([]byte(o.ID))
// 	if root == nil {
// 		return &bolt.Bucket{}, errors.New("Organization does not exits")
// 	}
// 	return root, nil
// }
//
// func (o *Organization) CreateUser(u *User) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		bck := root.Bucket([]byte("USERS"))
// 		u.ID = uuid.Must(uuid.NewV4()).String()
// 		buf, err := json.Marshal(u)
// 		if err != nil {
// 			return err
// 		}
// 		return bck.Put([]byte(u.ID), buf)
// 	})
// }
//
// func (o *Organization) UpdateUser(u *User) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		bck := root.Bucket([]byte("USERS"))
//
// 		v := bck.Get([]byte(u.ID))
// 		if v == nil {
// 			return errors.New("User does not exists.")
// 		}
// 		buf, err := json.Marshal(u)
// 		if err != nil {
// 			return err
// 		}
//
// 		return bck.Put([]byte(u.ID), []byte(buf))
// 	})
// }
//
// func (o *Organization) DeleteUser(u *User) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		bck := root.Bucket([]byte("USERS"))
// 		return bck.Delete([]byte(u.ID))
// 	})
// }
//
// func (o *Organization) DescribeUser(u *User) error {
// 	return o.Store.View(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
//
// 		bck := root.Bucket([]byte("USERS"))
// 		v := bck.Get([]byte(u.ID))
// 		if v == nil {
// 			return errors.New("User does not exists.")
// 		}
// 		return json.Unmarshal(v, u)
// 	})
// }
//
// func (o *Organization) AttachPolicyToUser(u *User, p *Policy) error {
// 	buf, err := json.Marshal(p)
// 	if err != nil {
// 		return err
// 	}
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
//
// 		bck := root.Bucket([]byte("POLICIES"))
// 		v := bck.Get([]byte(p.ID))
// 		if v == nil {
// 			return errors.New("The policy does not exists.")
// 		}
//
// 		bck = root.Bucket([]byte("USERS"))
// 		v = bck.Get([]byte(u.ID))
// 		if v == nil {
// 			return errors.New("The user does not exists.")
// 		}
// 		u.Policies[p.ID] = true
// 		return bck.Put([]byte(u.ID), buf)
// 	})
// }
//
// func (o *Organization) DetachPolicyToUser(u *User, p *Policy) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
//
// 		bck := root.Bucket([]byte("USERS"))
// 		v := bck.Get([]byte(string(u.ID)))
//
// 		if json.Unmarshal(v, u) != nil {
// 			return errors.New("Can't Unmarshal the DB user")
// 		}
//
// 		delete(u.Policies, p.ID)
//
// 		buf, err := json.Marshal(u)
// 		if err != nil {
// 			return err
// 		}
//
// 		return bck.Put([]byte(string(u.ID)), buf)
// 	})
// }
//
// func (o *Organization) CreateGroup(g *Group) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		b := root.Bucket([]byte(BucketGroup))
// 		g.ID = uuid.Must(uuid.NewV4()).String()
// 		buf, err := json.Marshal(g)
// 		if err != nil {
// 			return err
// 		}
// 		return b.Put([]byte(g.ID), buf)
// 	})
// }
//
// func (o *Organization) UpdateGroup(g *Group) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		b := root.Bucket([]byte(BucketGroup))
//
// 		v := b.Get([]byte(g.ID))
// 		if v == nil {
// 			return errors.New("The group does not exists.")
// 		}
// 		buf, err := json.Marshal(g)
// 		if err != nil {
// 			return err
// 		}
// 		return b.Put([]byte(g.ID), buf)
// 	})
// }
//
// func (o *Organization) DeleteGroup(g *Group) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		bck := root.Bucket([]byte(BucketGroup))
// 		return bck.Delete([]byte(g.ID))
// 	})
// }
//
// func (o *Organization) DescribeGroup(g *Group) error {
// 	return o.Store.View(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
//
// 		bck := root.Bucket([]byte(BucketGroup))
// 		v := bck.Get([]byte(g.ID))
// 		if v == nil {
// 			return errors.New("Group does not exists.")
// 		}
// 		return json.Unmarshal(v, g)
// 	})
// }
//
// func (o *Organization) AttachUserToGroup(g *Group, u *User) error {
// 	if err := o.DescribeGroup(g); err != nil {
// 		return err
// 	} else if err := o.DescribeUser(u); err != nil {
// 		return err
// 	}
//
// 	u.Groups[g.ID] = true
// 	g.Users[u.ID] = true
//
// 	if err := o.UpdateUser(u); err != nil {
// 		return err
// 	} else if err := o.UpdateGroup(g); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (o *Organization) DetachUserFromGroup(g *Group, u *User) error {
// 	if err := o.DescribeGroup(g); err != nil {
// 		return err
// 	} else if err := o.DescribeUser(u); err != nil {
// 		return err
// 	}
//
// 	delete(u.Groups, g.ID)
// 	delete(g.Users, u.ID)
//
// 	if err := o.UpdateUser(u); err != nil {
// 		u.Groups[g.ID] = true
// 		return err
// 	} else if err := o.UpdateGroup(g); err != nil {
// 		g.Users[u.ID] = true
// 		return err
// 	}
// 	return nil
// }
//
// func (o *Organization) AttachPolicyToGroup(g *Group, p *Policy) error {
// 	if err := o.DescribeGroup(g); err != nil {
// 		return err
// 	} else if err := o.DescribePolicy(p); err != nil {
// 		return err
// 	}
// 	g.Policies[p.ID] = true
// 	return o.UpdateGroup(g)
// }
//
// func (o *Organization) DetechPolicyFromGroup(g *Group, p *Policy) error {
// 	if err := o.DescribeGroup(g); err != nil {
// 		return err
// 	}
//
// 	delete(g.Policies, p.ID)
//
// 	if err := o.UpdateGroup(g); err != nil {
// 		return err
// 	}
// 	return nil
// }
//
// func (o *Organization) CreateRole(r *Role) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		bck := root.Bucket([]byte(BucketRole))
// 		r.ID = uuid.Must(uuid.NewV4()).String()
// 		buf, err := json.Marshal(r)
// 		if err != nil {
// 			return err
// 		}
// 		return bck.Put([]byte(r.ID), buf)
// 	})
// }
//
// func (o *Organization) UpdateRole(r *Role) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		b := root.Bucket([]byte(BucketRole))
//
// 		v := b.Get([]byte(r.ID))
// 		if v == nil {
// 			return errors.New("The role does not exists.")
// 		}
// 		buf, err := json.Marshal(r)
// 		if err != nil {
// 			return err
// 		}
// 		return b.Put([]byte(r.ID), buf)
// 	})
// }
//
// func (o *Organization) DeleteRole(r *Role) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		bck := root.Bucket([]byte(BucketRole))
// 		return bck.Delete([]byte(r.ID))
// 	})
// }
//
// func (o *Organization) DescribeRole(r *Role) error {
// 	return o.Store.View(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
//
// 		bck := root.Bucket([]byte(BucketRole))
// 		v := bck.Get([]byte(r.ID))
// 		if v == nil {
// 			return errors.New("The role does not exists.")
// 		}
// 		return json.Unmarshal(v, r)
// 	})
// }
//
// func (o *Organization) AttachPolicyToRole(r *Role, p *Policy) error {
// 	if err := o.DescribeRole(r); err != nil {
// 		return err
// 	} else if err := o.DescribePolicy(p); err != nil {
// 		return err
// 	}
// 	r.Policies[p.ID] = true
// 	return o.UpdateRole(r)
// }
//
// func (o *Organization) DetachPolicyFromRole(r *Role, p *Policy) error {
// 	if err := o.DescribeRole(r); err != nil {
// 		return err
// 	}
// 	delete(r.Policies, p.ID)
// 	return o.UpdateRole(r)
// }
//
// func (o *Organization) CreatePolicy(p *Policy) error {
// 	p.ID = uuid.Must(uuid.NewV4()).String()
// 	buf, err := json.Marshal(p)
// 	if err != nil {
// 		return err
// 	}
//
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
//
// 		bck := root.Bucket([]byte("POLICIES"))
// 		return bck.Put([]byte(p.ID), buf)
// 	})
// }
//
// func (o *Organization) UpdatePolicy(p *Policy) error {
// 	buf, err := json.Marshal(p)
// 	if err != nil {
// 		return err
// 	}
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
//
// 		bck := root.Bucket([]byte("POLICIES"))
// 		v := bck.Get([]byte(p.ID))
// 		if v == nil {
// 			return errors.New("Policy does not exsts.")
// 		}
// 		return bck.Put([]byte(p.ID), buf)
// 	})
// }
//
// func (o *Organization) DeletePolicy(p *Policy) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		bck := root.Bucket([]byte("POLICIES"))
// 		return bck.Delete([]byte(p.ID))
// 	})
// }
//
// func (o *Organization) DescribePolicy(p *Policy) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		root, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		bck := root.Bucket([]byte("POLICIES"))
// 		v := bck.Get([]byte(p.ID))
// 		if v == nil {
// 			return errors.New("policy does not exists.")
// 		}
// 		return json.Unmarshal(v, p)
// 	})
// }
//
// type Context map[string]interface{}
//
// func (o *Organization) Authorize(u *User, role *Role, resource *string, action *string, context *Context) (bool, error) {
// 	var policies []*Policy
// 	assumeRole := false
//
// 	if role.ID != "" {
// 		assumeRole = true
// 		err := o.DescribeRole(role)
// 		if err != nil {
// 			return false, errors.New("AssumeRoleError: The role " + role.ID + "does not exists.")
// 		}
//
// 		for k := range role.Policies {
// 			p := Policy{ID: k}
// 			o.DescribePolicy(&p)
// 			policies = append(policies, &p)
// 		}
// 	}
//
// 	for k := range u.Groups {
// 		g := Group{ID: k}
// 		o.DescribeGroup(&g)
// 		for k := range g.Policies {
// 			p := Policy{ID: k}
// 			if o.DescribePolicy(&p) != nil {
// 				// TODO: Log Error
// 				continue
// 			}
// 			policies = append(policies, &p)
// 		}
// 	}
//
// 	for k := range u.Policies {
// 		p := Policy{ID: k}
// 		if o.DescribePolicy(&p) != nil {
// 			// TODO: Log Error
// 			continue
// 		}
// 		policies = append(policies, &p)
// 	}
//
// 	assumeRoleAction := "judge:assumeRole"
// 	for i := range policies {
// 		if assumeRole {
// 			if isPermitedResource(policies[i], &role.ORN) &&
// 				isPermitedAction(policies[i], &assumeRoleAction) {
// 			}
// 		}
// 		if isPermitedResource(policies[i], resource) {
// 			if isPermitedAction(policies[i], action) {
// 				return true, nil
// 			}
// 		}
// 	}
// 	return false, errors.New("The user " + u.ID + " can't perform the action " + *action + " on the resource " + *resource)
// }
//
// func isPermitedResource(p *Policy, resource *string) bool {
// 	expr := strings.Replace("", "*", ".*", -1)
// 	matched, err := regexp.MatchString(expr, *resource)
// 	if err == nil && matched {
// 		return true
// 	}
// 	return false
// }
//
// func isPermitedAction(p *Policy, action *string) bool {
// 	// Use HashSet to have O(1) func
// 	for i := range []string{"todo"} {
// 		if []string{"todo"}[i] == *action {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// type Organization struct {
// 	ID    string
// 	Name  string
// 	Store *bolt.DB
// }
//
// var (
// 	BUCKETS = []string{BucketUser, BucketGroup, BucketRole, BucketPolicy}
// )
//
// const (
// 	BucketUser   = "USERS"
// 	BucketGroup  = "GROUPS"
// 	BucketRole   = "ROLES"
// 	BucketPolicy = "POLICIES"
// )
//
// func (o *Organization) CreateOrganization(org *Organization) error {
// 	org.ID = uuid.Must(uuid.NewV4()).String()
// 	buf, err := json.Marshal(org)
//
// 	tx, err := o.Store.Begin(true)
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()
//
// 	// Setup organization buckets
// 	root, err := tx.CreateBucket([]byte(org.ID))
// 	if err != nil {
// 		return err
// 	}
// 	for i := range BUCKETS {
// 		if _, err = root.CreateBucketIfNotExists([]byte(BUCKETS[i])); err != nil {
// 			return err
// 		}
// 	}
//
// 	// Store organization informations
// 	root.Put([]byte("metadata"), buf)
//
// 	return tx.Commit()
// }
//
// func (o *Organization) UpdateOrganization(org *Organization) error {
// 	buf, err := json.Marshal(org)
// 	if err != nil {
// 		return err
// 	}
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(org.ID))
// 		if b == nil {
// 			return errors.New("Organization does not exit")
// 		}
// 		return b.Put([]byte("metadata"), buf)
// 	})
// }
//
// func (o *Organization) DescribeOrganization(org *Organization) error {
// 	return o.Store.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(org.ID))
// 		if b == nil {
// 			return errors.New("Organization does not exit")
// 		}
// 		v := b.Get([]byte("metadata"))
// 		return json.Unmarshal(v, &org)
// 	})
// }
//
// func (o *Organization) DeleteOrganization(_ *Organization) error {
// 	return o.Store.Update(func(tx *bolt.Tx) error {
// 		_, err := o.OpenBucket(tx)
// 		if err != nil {
// 			return err
// 		}
// 		return tx.DeleteBucket([]byte(o.ID))
// 	})
// }
