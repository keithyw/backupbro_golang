package mongo

import (
	"gopkg.in/mgo.v2"
	"backupBro/pkg"
	"gopkg.in/mgo.v2/bson"
)

type UserService struct {
	collection *mgo.Collection
	hash root.Hash
}

func NewUserService(session *Session, dbName string, collectionName string, hash root.Hash) *UserService {
	collection := session.GetCollection(dbName, collectionName)
	collection.EnsureIndex(userModelIndex())
	return &UserService{collection, hash}
}

func (p *UserService) CreateUser(u *root.User) (error) {
	user := newUserModel(u)
	hashedPassword, err := p.hash.Generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return p.collection.Insert(&user)
}

func (p *UserService) GetByEmail(email string) (*root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"email": email}).One(&model)
	return model.toRootUser(), err
}

func (p *UserService) GetByUsername(username string) (*root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"username": username}).One(&model)
	return model.toRootUser(), err
}

func (p *UserService) Login(c root.Credentials) (root.User, error) {
	model := userModel{}
	err := p.collection.Find(bson.M{"email": c.Email}).One(&model)
	if err != nil {
		return root.User{}, err
	}
	err = p.hash.Compare(model.Password, c.Password)
	if err != nil {
		return root.User{}, err
	}
	return root.User{
		Id: model.Id.Hex(),
		Email: "-",
		Username: model.Username,
		Password: "-",
	}, err

}

