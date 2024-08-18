package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	helpers2 "idist-core/app/helpers"
	"time"
)

type Partner struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Logo        string             `bson:"logo" json:"logo"`
	Banner      string             `bson:"banner" json:"banner"`
	Description string             `bson:"description" json:"description"`
	Slug        string             `bson:"slug" json:"slug"`
	Url         string             `bson:"url" json:"url"`
	IsHighlight bool               `bson:"is_highlight" json:"is_highlight"`
	IsActive    bool               `bson:"is_active" json:"is_active"`
	Priority    int                `bson:"priority" json:"priority"`
	CreatedBy   primitive.ObjectID `bson:"created_by" json:"created_by"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedBy   primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedBy   primitive.ObjectID `bson:"deleted_by" json:"deleted_by"`
	DeletedAt   *time.Time         `bson:"deleted_at" json:"deleted_at"`

	// Edge loading
	CountArticles int64 `bson:"-" json:"count_articles"`
}

type Partners []Partner

func (u *Partner) CollectionName() string {
	return "partners"
}

func (u *Partner) Find(filter interface{}, opts ...*options.FindOptions) (Partners, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Partners, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Partner
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

func (u *Partner) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Partner) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.UpdatedAt = helpers2.Now()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Partner) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.DeletedAt = helpers2.PNow()
	if _, err := DB().Collection(u.CollectionName()).UpdateOne(ctx, bson.M{"_id": u.ID}, bson.M{
		"$set": u,
	}, options.Update()); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Partner) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Partner) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.Slug = helpers2.MakeSlug(u.Name)
	u.CreatedAt = helpers2.Now()
	u.UpdatedAt = helpers2.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}
