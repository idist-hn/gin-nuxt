package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/helpers"
)

type Template struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
	Slug string             `bson:"slug" json:"slug"`
	Type string             `bson:"type" json:"type"`
}

type Templates []Template

func (u *Template) CollectionName() string {
	return "templates"
}

func (u *Template) Find(filter interface{}, opts ...*options.FindOptions) (Templates, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Templates, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Template
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

func (u *Template) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Template) Update() error {
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

func (u *Template) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if _, err := DB().Collection(u.CollectionName()).DeleteOne(ctx, bson.M{"_id": u.ID}, options.Delete()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Template) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Template) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.Slug = helpers.MakeSlug(u.Name)
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}
