package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/helpers"
	"time"
)

type Menu struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Name           string             `bson:"name" json:"name"`
	IsUsed         bool               `bson:"is_used" json:"is_used"`
	MenuLocationId primitive.ObjectID `bson:"menu_location_id" json:"menu_location_id"`
	Items          []MenuItem         `bson:"items" json:"items"`
	DepartmentId   primitive.ObjectID `bson:"department_id" json:"department_id"`
	Type           string             `bson:"type" json:"type"`
	CreatedBy      primitive.ObjectID `bson:"created_by" json:"created_by"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedBy      primitive.ObjectID `bson:"updated_by" json:"updated_by"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedBy      primitive.ObjectID `bson:"deleted_by" json:"deleted_by"`
	DeletedAt      *time.Time         `bson:"deleted_at" json:"deleted_at"`
	//Edge Loading
	Location   *MenuLocation `bson:"-" json:"location"`
	Department *Department   `bson:"-" json:"department"`
}
type MenuItem struct {
	ID          int64      `bson:"_id" json:"id"`
	Type        string     `bson:"type" json:"type"`
	Title       string     `bson:"title" json:"title"`
	Description string     `bson:"description" json:"description"`
	Url         string     `bson:"url" json:"url"`
	OpenNewTab  bool       `bson:"open_new_tab" json:"openNewTab"`
	Children    []MenuItem `bson:"children" json:"children"`
}

type Menus []Menu

func (u *Menu) CollectionName() string {
	return "menus"
}

func (u *Menu) Find(filter interface{}, opts ...*options.FindOptions) (Menus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Menus, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Menu
			if err = cursor.Decode(&elem); err != nil {
				return data, err
			}
			elem.Preload("menu_location")
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

func (u *Menu) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Menu) Update() error {
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

func (u *Menu) Delete() error {
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

func (u *Menu) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Menu) Create() error {
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
func (u *Menu) Preload(values ...string) {
	for _, value := range values {
		if value == "menu_location" {
			var location MenuLocation
			filter := bson.M{
				"_id": u.MenuLocationId,
			}
			if err := location.First(filter); err == nil {
				u.Location = &location
			}
		}
	}
}
