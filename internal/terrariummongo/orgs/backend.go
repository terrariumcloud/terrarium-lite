package orgs

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// OrganizationBackend is a struct that implements Mongo operations for organizations
type OrganizationBackend struct {
	CollectionName string
	Database       string
	Client         *mongo.Client
}

// Init initializes the Organizations table
func (o *OrganizationBackend) Init() error {
	collection := o.Client.Database(o.Database).Collection(o.CollectionName)
	// Ensures unique email and name combination
	_, err := collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "name", Value: 1},
			{Key: "email", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}
	return nil
}

// Create Adds a new organization to the organizations table
func (o *OrganizationBackend) Create(name string, email string) (*Organization, error) {
	org := &Organization{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Email:     email,
		CreatedOn: time.Now().UTC().String(),
	}
	ctx := context.TODO()
	_, err := o.Client.Database(o.Database).Collection(o.CollectionName).InsertOne(ctx, org, options.InsertOne())
	if err != nil {
		return nil, err
	}
	return org, nil
}

// ReadAll Returns all organizations from the organizations table
func (o *OrganizationBackend) ReadAll(limit int, offset int) ([]*Organization, error) {
	ctx := context.TODO()
	limitOpt := options.Find().SetLimit(int64(limit))
	skipOpt := options.Find().SetSkip(int64(offset))
	cur, err := o.Client.Database(o.Database).Collection(o.CollectionName).Find(ctx, bson.D{}, limitOpt, skipOpt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var organizationList []*Organization = []*Organization{}
	for cur.Next(ctx) {
		result := &Organization{}
		err := cur.Decode(result)
		if err != nil {
			return nil, err
		}
		organizationList = append(organizationList, result)
	}
	return organizationList, nil
}

// ReadOne Returns a single organization from the organizations table
func (o *OrganizationBackend) ReadOne(orgName string) (*Organization, error) {
	ctx := context.TODO()
	org := &Organization{}
	result := o.Client.Database(o.Database).Collection(o.CollectionName).FindOne(ctx, bson.M{"name": orgName}, options.FindOne())
	err := result.Decode(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// Update Updates an organization in the organization table
func (o *OrganizationBackend) Update(name string, email string) (*Organization, error) {
	ctx := context.TODO()
	update := bson.M{}
	if email != "" {
		update["email"] = email
	}
	upsert := options.Update().SetUpsert(false)
	_, err := o.Client.Database(o.Database).Collection(o.CollectionName).UpdateOne(ctx, bson.M{"name": name}, bson.M{"$set": update}, upsert)
	if err != nil {
		return nil, err
	}
	updatedOrg, err := o.ReadOne(name)
	if err != nil {
		return nil, err
	}
	return updatedOrg, nil
}

// Delete Removes an organization from the organization table
func (o *OrganizationBackend) Delete(name string) error {
	ctx := context.TODO()
	_, err := o.Client.Database(o.Database).Collection(o.CollectionName).DeleteOne(ctx, bson.M{"name": name}, options.Delete())
	if err != nil {
		return err
	}
	return nil
}

// GetBackendType Returns the type of backend used
func (o *OrganizationBackend) GetBackendType() string {
	return "mongo"
}
