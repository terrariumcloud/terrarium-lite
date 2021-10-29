package terrariummongo

import (
	"context"
	"fmt"

	"github.com/dylanrhysscott/terrarium/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// VCSBackend is a struct that implements Mongo operations for organizations
type VCSBackend struct {
	CollectionName string
	Database       string
	client         *mongo.Client
}

// Init initializes the VCS table
func (v *VCSBackend) Init() error {
	return nil
}

// Create Adds a new VCS to the VCS table
func (v *VCSBackend) Create(orgID string, orgName string, link *types.VCSOAuthClientLink) (*types.VCS, error) {
	ctx := context.TODO()
	oid, err := primitive.ObjectIDFromHex(orgID)
	if err != nil {
		return nil, err
	}
	vcsID := primitive.NewObjectID()
	link.CallbackURI = fmt.Sprintf("/oauth/github/%s/callback", vcsID.Hex())
	vcsConnection := &types.VCS{
		ID: vcsID,
		Organization: &types.ResourceLink{
			ID:   oid,
			Link: fmt.Sprintf("/v1/organizations/%s", orgName),
		},
		OAuth: link,
	}
	_, err = v.client.Database(v.Database).Collection(v.CollectionName).InsertOne(ctx, vcsConnection, options.InsertOne())
	if err != nil {
		return nil, err
	}
	return vcsConnection, nil
}

// ReadAll Returns all VCSs from the VCS table
func (v *VCSBackend) ReadAll(limit int, offset int) ([]*types.VCS, error) {
	// ctx := context.TODO()
	// limitOpt := options.Find().SetLimit(int64(limit))
	// skipOpt := options.Find().SetSkip(int64(offset))
	// cur, err := v.client.Database(v.Database).Collection(v.CollectionName).Find(ctx, bson.D{}, limitOpt, skipOpt)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

// ReadOne Returns a single VCS from the VCSs table
func (v *VCSBackend) ReadOne(id string) (*types.VCS, error) {
	ctx := context.TODO()
	vcs := &types.VCS{}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := v.client.Database(v.Database).Collection(v.CollectionName).FindOne(ctx, bson.M{"_id": oid}, options.FindOne())
	err = result.Decode(vcs)
	if err != nil {
		return nil, err
	}
	return vcs, nil
}

// Update Updates an VCS in the VCS table
func (v *VCSBackend) Update(orgID string, orgName string, link *types.VCSOAuthClientLink) (*types.VCS, error) {
	return nil, nil
}

// UpdateVCSToken Updates the VCS OAuth Token in the database
func (v *VCSBackend) UpdateVCSToken(clientID string, token *types.VCSToken) error {
	ctx := context.TODO()
	update := bson.M{
		"$set": bson.M{
			"oauth.token": bson.M{
				"access_token":             token.AccessToken,
				"expires_in":               token.ExpiresIn,
				"token_type":               token.TokenType,
				"scope":                    token.Scope,
				"refresh_token_expires_in": token.RefreshTokenExpiresIn,
				"refresh_token":            token.RefreshToken,
			},
		},
	}
	query := bson.M{
		"oauth.client_id": clientID,
	}
	upsert := options.Update().SetUpsert(false)
	_, err := v.client.Database(v.Database).Collection(v.CollectionName).UpdateOne(ctx, query, update, upsert)
	if err != nil {
		return err
	}
	return nil
}

// Delete Removes an VCS from the VCS table
func (v *VCSBackend) Delete(name string) error {
	return nil
}
