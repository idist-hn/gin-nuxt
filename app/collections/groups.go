package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/helpers"
)

type Group struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Alias         string             `bson:"alias" json:"alias"`
	Name          string             `bson:"name" json:"name"`
	CountArticles int64              `bson:"-" json:"count_articles"`
}

type Groups []Group

func (u *Group) CollectionName() string {
	return "groups"
}

func (u *Group) Find(filter interface{}, opts ...*options.FindOptions) (Groups, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Groups, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Group
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

func (u *Group) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Group) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Group) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Group) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Group) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	if u.Alias == "" {
		u.Alias = helpers.MakeAlias(u.Name)
	}
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}
