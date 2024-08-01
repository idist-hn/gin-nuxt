package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/helpers"
	"time"
)

type Description struct {
	ID                    primitive.ObjectID   `bson:"_id" json:"id"`
	Title                 string               `bson:"title" json:"title"`
	SubTitle              string               `bson:"sub_title" json:"sub_title"`
	Slug                  string               `bson:"slug" json:"slug"`
	Description           string               `bson:"description" json:"description"`
	Thumbnail             string               `bson:"thumbnail" json:"thumbnail"`
	Content               string               `bson:"content" json:"content"`
	Libraries             []BlockImage         `bson:"libraries" json:"libraries"`
	Meta                  MetaDescription      `bson:"meta" json:"meta"`
	Note                  string               `bson:"note" json:"note"`
	Status                string               `bson:"status" json:"status"`
	IsHighlight           bool                 `bson:"is_highlight" json:"is_highlight"`
	IsHot                 bool                 `bson:"is_hot" json:"is_hot"`
	CategoryId            primitive.ObjectID   `bson:"category_id" json:"-"`
	TagIds                []primitive.ObjectID `bson:"tag_ids" json:"-"`
	RelatedDescriptionIds []primitive.ObjectID `bson:"related_Description_ids" json:"related_Description_ids"`
	ViewUpdatedAt         *time.Time           `bson:"view_updated_at" json:"-"`
	PublishedBy           *primitive.ObjectID  `bson:"published_by" json:"-"`
	PublishedAt           *time.Time           `bson:"published_at" json:"published_at"`
	CreatedBy             primitive.ObjectID   `bson:"created_by" json:"-"`
	CreatedAt             time.Time            `bson:"created_at" json:"created_at"`
	UpdatedBy             primitive.ObjectID   `bson:"updated_by" json:"-"`
	UpdatedAt             time.Time            `bson:"updated_at" json:"updated_at"`
	DeletedBy             *primitive.ObjectID  `bson:"deleted_by" json:"-"`
	DeletedAt             *time.Time           `bson:"deleted_at" json:"-"`

	// Status
	View int64 `bson:"count_views" json:"count_views"`

	// Edge Loading
	Tags                []Tag         `bson:"-" json:"tags"`
	Publisher           *User         `bson:"-" json:"publisher"`
	Category            Category      `bson:"-" json:"category"`
	RelatedDescriptions []Description `bson:"-" json:"related_Descriptions"`
}

type MetaDescription struct {
	HighlightBlock HighlightBlock `bson:"highlight_block" json:"highlight_block"`
}
type HighlightBlock struct {
	Title       string          `bson:"title" json:"title"`
	Description string          `bson:"description" json:"description"`
	Items       []HighlightItem `bson:"items" json:"items"`
}
type HighlightItem struct {
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Icon        string `bson:"icon" json:"icon"`
}
type BlockImage struct {
	Url    string `bson:"url" json:"url"`
	Name   string `bson:"name" json:"name"`
	Uid    string `bson:"uid" json:"uid"`
	Status string `bson:"status" json:"status"`
}
type Descriptions []Description

func (u *Description) CollectionName() string {
	return "Descriptions"
}

func (u *Description) Find(filter interface{}, opts ...*options.FindOptions) (Descriptions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Descriptions, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Description
			if err = cursor.Decode(&elem); err != nil {
				return data, err
			}
			elem.Preload("category")
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

func (u *Description) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Description) Update() error {
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

func (u *Description) Delete() error {
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

func (u *Description) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Description) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	if u.Slug == "" {
		u.Slug = helpers.MakeSlug(u.Title)
	}
	u.CreatedAt = helpers.Now()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Description) Preload(values ...string) {
	for _, value := range values {
		if value == "category" {
			var category Category
			filter := bson.M{
				"_id": u.CategoryId,
			}
			if err := category.First(filter); err == nil {
				u.Category = category
			}
		}
		if value == "related_Descriptions" {
			filter := bson.M{
				"deleted_at": nil,
				"_id":        bson.M{"$in": u.RelatedDescriptionIds},
			}
			if relatedDescriptions, err := u.Find(filter); err == nil {
				u.RelatedDescriptions = relatedDescriptions
			}
		}
	}
}
