package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/helpers"
	"time"
)

type MenuLocation struct {
	ID          primitive.ObjectID  `bson:"_id" json:"id"`
	Name        string              `bson:"name" json:"name"`
	Alias       string              `bson:"alias" json:"alias"`
	IsUsed      bool                `bson:"is_used" json:"is_used"`
	CanNested   bool                `bson:"can_nested" json:"can_nested"`
	MenuId      *primitive.ObjectID `bson:"menu_id" json:"menu_id"`
	Language    string              `bson:"language" json:"language"`
	Description string              `bson:"description" json:"description"`

	CreatedBy primitive.ObjectID `bson:"created_by" json:"created_by"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedBy primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedBy primitive.ObjectID `bson:"deleted_by" json:"-"`
	DeletedAt *time.Time         `bson:"deleted_at" json:"-"`

	// Edge Loading
	Menu *Menu `bson:"-" json:"menu"`
}

type MenuLocations []MenuLocation

func (u *MenuLocation) CollectionName() string {
	return "menu_locations"
}

func (u *MenuLocation) Find(filter interface{}, opts ...*options.FindOptions) (MenuLocations, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(MenuLocations, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem MenuLocation
			if err = cursor.Decode(&elem); err != nil {
				return data, err
			}
			data = append(data, elem)
		}
		if err = cursor.Err(); err != nil {
			return data, err
		}
		return data, cursor.Close(ctx)
	} else {
		return data, err
	}
}

func (u *MenuLocation) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *MenuLocation) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *MenuLocation) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.DeletedAt = helpers.PNow()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *MenuLocation) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *MenuLocation) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.CreatedAt = helpers.Now()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}
