package collections

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Permission struct {
	ID          string `bson:"_id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Subject     string `bson:"subject" json:"subject"`
	Action      string `bson:"action" json:"action"`
	Description string `bson:"description" json:"description"`
}

type Permissions []Permission

func (u *Permission) CollectionName() string {
	return "permissions"
}

func (u *Permission) First(filter bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	if result := DB().Collection(u.CollectionName()).FindOne(ctx, filter); result.Err() != nil {
		return result.Err()
	} else {
		return result.Decode(&u)
	}
}

func (u *Permission) Find(filter interface{}, opts ...*options.FindOptions) (Permissions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), QueryTimeOut)
	defer cancel()
	data := make(Permissions, 0)

	/* Lấy danh sách bản ghi */
	if cursor, err := DB().Collection(u.CollectionName()).Find(ctx, filter, opts...); err == nil {
		for cursor.Next(ctx) {
			var elem Permission
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
