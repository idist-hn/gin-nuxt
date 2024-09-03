package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"idist-core/app/helpers"
	"sync"
	"time"
)

type Department struct {
	ID                     primitive.ObjectID  `bson:"_id" json:"id"`
	Priority               int64               `bson:"priority" json:"priority"`
	Name                   string              `bson:"name" json:"name"`
	Slug                   string              `bson:"slug" json:"slug"`
	Description            string              `bson:"description" json:"description"`
	Address                string              `bson:"address" json:"address"`
	Phone                  string              `bson:"phone" json:"phone"`
	Email                  string              `bson:"email" json:"email"`
	Icon                   string              `bson:"icon" json:"icon"`
	CustomTemplate         bool                `bson:"custom_template" json:"custom_template"`
	IsActive               bool                `bson:"is_active" json:"is_active"`
	ShowSidebar            bool                `bson:"show_sidebar" json:"show_sidebar"`
	ParentId               *primitive.ObjectID `bson:"parent_id" json:"parent_id"`
	Color                  string              `bson:"color" json:"color"`
	Extras                 []ExtraDepartment   `bson:"extras" json:"extras"`
	Content                string              `bson:"content" json:"content"`
	PcDepartmentStyles     []DepartmentStyle   `bson:"pc_department_styles" json:"pc_department_styles"`
	MobileDepartmentStyles []DepartmentStyle   `bson:"mobile_department_styles" json:"mobile_department_styles"`
	MenuItems              []MenuItem          `bson:"menu_items" json:"menu_items"`
	CreatedAt              time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt              time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedAt              *time.Time          `bson:"deleted_at" json:"-"`

	// Edge Loading
	Users      *Users      `bson:"-" json:"users"`
	Parent     *Department `bson:"-" json:"parent"`
	CountUsers int64       `bson:"-" json:"count_users"`
}
type ExtraDepartment struct {
	ID    int64  `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Value string `json:"value" bson:"value"`
}
type DepartmentStyle struct {
	ID    int64  `bson:"_id" json:"id"`
	Key   string `bson:"key" json:"key"`
	Value string `json:"value" bson:"value"`
}

type Departments []Department

func (u *Department) CollectionName() string {
	return "departments"
}

func (u *Department) Find(filter interface{}, opts ...*options.FindOptions) (Departments, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Departments, 0)
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Department
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

func (u *Department) First(filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Department) Update() error {
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

func (u *Department) Delete() error {
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

func (u *Department) Count(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if total, err := DB().Collection(u.CollectionName()).CountDocuments(ctx, filter, options.Count()); err != nil {
		return 0, err
	} else {
		return total, nil
	}
}

func (u *Department) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	u.ID = primitive.NewObjectID()
	u.Slug = helpers.MakeSlug(u.Name)
	u.CreatedAt = helpers.Now()
	u.UpdatedAt = helpers.Now()
	if _, err := DB().Collection(u.CollectionName()).InsertOne(ctx, u); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *Department) Preload(properties ...string) {
	var wg sync.WaitGroup
	for _, property := range properties {
		if property == "count_user" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				entry := User{}
				entries := Users{}
				filter := bson.M{
					"deleted_at":    nil,
					"department_id": u.ID,
				}
				opts := options.Find()
				entries, _ = entry.Find(filter, opts)
				u.CountUsers = int64(len(entries))
				u.Users = &entries
			}()
		}
		if property == "parent" {
			wg.Add(1)
			go func() {
				defer wg.Done()
				entry := Department{}
				_ = entry.First(bson.M{"_id": u.ParentId, "deleted_at": nil})
				u.Parent = &entry
			}()
		}
	}
	wg.Wait()

}
