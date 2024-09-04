package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/helpers"
	"time"
)

type Article struct {
	ID                primitive.ObjectID   `bson:"_id" json:"id"`
	GroupId           primitive.ObjectID   `bson:"group_id" json:"group_id"`
	Title             string               `bson:"title" json:"title"`
	SeoTitle          string               `bson:"seo_title" json:"seo_title"`
	Slug              string               `bson:"slug" json:"slug"`
	Description       string               `bson:"description" json:"description"`
	SeoDescription    string               `bson:"seo_description" json:"seo_description"`
	SeoKeywords       []string             `bson:"seo_keywords" json:"seo_keywords"`
	Thumbnail         string               `bson:"thumbnail" json:"thumbnail"`
	Content           string               `bson:"content" json:"content"`
	Note              string               `bson:"note" json:"note"`
	Status            string               `bson:"status" json:"status"`
	IsHighlight       bool                 `bson:"is_highlight" json:"is_highlight"`
	IsHot             bool                 `bson:"is_hot" json:"is_hot"`
	CategoryId        primitive.ObjectID   `bson:"category_id" json:"-"`
	TagIds            []primitive.ObjectID `bson:"tag_ids" json:"tag_ids"`
	RelatedArticleIds []primitive.ObjectID `bson:"related_article_ids" json:"related_article_ids"`
	ViewUpdatedAt     *time.Time           `bson:"view_updated_at" json:"-"`
	PublishedBy       *primitive.ObjectID  `bson:"published_by" json:"-"`
	PublishedAt       *time.Time           `bson:"published_at" json:"published_at"`
	CreatedBy         primitive.ObjectID   `bson:"created_by" json:"-"`
	CreatedAt         time.Time            `bson:"created_at" json:"created_at"`
	UpdatedBy         primitive.ObjectID   `bson:"updated_by" json:"-"`
	UpdatedAt         time.Time            `bson:"updated_at" json:"updated_at"`
	DeletedBy         *primitive.ObjectID  `bson:"deleted_by" json:"-"`
	DeletedAt         *time.Time           `bson:"deleted_at" json:"-"`

	// Status
	View int64 `bson:"count_views" json:"count_views"`

	// Edge Loading
	Tags            []Tag     `bson:"-" json:"tags"`
	Publisher       *User     `bson:"-" json:"publisher"`
	Category        Category  `bson:"-" json:"category"`
	RelatedArticles []Article `bson:"-" json:"related_articles"`
}

type Articles []Article

func (u *Article) CollectionName() string {
	return "articles"
}

func (u *Article) Find(filter interface{}, opts ...*options.FindOptions) (Articles, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Articles, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Article
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

func (u *Article) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Article) Update() error {
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

func (u *Article) Delete() error {
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

func (u *Article) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Article) Create() error {
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

func (u *Article) Preload(values ...string) {
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
		if value == "related_articles" {
			filter := bson.M{
				"deleted_at": nil,
				"_id":        bson.M{"$in": u.RelatedArticleIds},
			}
			if relatedArticles, err := u.Find(filter, options.Find().SetProjection(bson.M{"content": 0})); err == nil {
				u.RelatedArticles = relatedArticles
			}
		}
		if value == "tags" {
			filter := bson.M{
				"deleted_at": nil,
				"_id":        bson.M{"$in": u.TagIds},
			}
			var tag Tag
			if tags, err := tag.Find(filter); err == nil {
				u.Tags = tags
			}
		}
	}
}
