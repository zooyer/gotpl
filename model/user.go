package model

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zooyer/gotpl/common/consts"
	"github.com/zooyer/gotpl/dto"
	"github.com/zooyer/miskit/micro"
)

type User struct {
	micro.ModelBase
	dto.UserCreateReq
	micro.ModelExtra
}

func (User) TableName() string {
	return consts.TableNameUser
}

func (u User) getID() string {
	return fmt.Sprintf("%d", u.ID)
}

func (u User) CreateModel(now time.Time) micro.ModelExtra {
	var uid = u.getID()

	return micro.ModelExtra{
		CreatedID: uid,
		UpdatedID: uid,
		DeletedID: "",
		CreatedBy: u.Username,
		UpdatedBy: u.Username,
		DeletedBy: "",
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}
}

func (u User) UpdateModel(model micro.Model, now time.Time) micro.Model {
	model.UpdatedID = u.getID()
	model.UpdatedBy = u.Username
	model.UpdatedAt = now.Unix()

	return model
}

func (u User) Update(update micro.Update, now time.Time) micro.Update {
	if update == nil {
		return update
	}

	update["updated_id"] = u.getID()
	update["updated_by"] = u.Username
	update["updated_at"] = now.Unix()

	return update
}

func (u User) DeleteModel(now time.Time) micro.Update {
	return micro.Update{
		"deleted_id": u.getID(),
		"deleted_by": u.Username,
		"deleted_at": now.Unix(),
	}
}

func (u User) MarshalJSON() ([]byte, error) {
	type user User

	u.Password = ""

	return json.Marshal(user(u))
}

func init() {
	gob.Register(new(User))
}
