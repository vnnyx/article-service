package repository

import (
	"github.com/vnnyx/article-service/graph/model/entity"
	"github.com/vnnyx/article-service/internal/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewAuthorRepository(database *mongo.Database) AuthorRepository {
	return &AuthorRepositoryImpl{
		Collection: database.Collection("authors"),
	}
}

func (repository *AuthorRepositoryImpl) InsertAuthor(author *entity.AuthorEntity) error {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	author.Article = make([]*entity.ArticleEntity, 0)
	_, err := repository.Collection.InsertOne(ctx, author)
	return err
}

func (repository *AuthorRepositoryImpl) FindAllAuthor() (authors []*entity.AuthorEntity, err error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	if err != nil {
		return authors, err
	}

	err = cursor.All(ctx, &authors)
	if err != nil {
		return authors, err
	}

	return authors, nil
}

func (repository *AuthorRepositoryImpl) FindAuthorByID(id string) (author *entity.AuthorEntity, err error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	err = repository.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&author)
	return author, err
}

func (repository *AuthorRepositoryImpl) FindAuthorByUsername(username string) (authors []*entity.AuthorEntity, err error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{"username": bson.M{"$regex": username, "$options": "im"}})
	if err != nil {
		return authors, err
	}

	err = cursor.All(ctx, &authors)
	if err != nil {
		return authors, err
	}

	return authors, err
}

func (repository *AuthorRepositoryImpl) UpdateAuthor(author *entity.AuthorEntity) (bool, error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": author.ID}
	update := bson.M{"$set": author}

	_, err := repository.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repository *AuthorRepositoryImpl) DeleteAuthor(id string) (bool, error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": id}
	_, err := repository.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repository *AuthorRepositoryImpl) UpdatePassword(id, newPassword string) (bool, error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": id}
	set := bson.M{"$set": bson.M{"password": newPassword}}

	_, err := repository.Collection.UpdateOne(ctx, filter, set)
	if err != nil {
		return false, err
	}
	return true, nil
}
