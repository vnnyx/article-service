package repository

import (
	"fmt"

	"github.com/vnnyx/article-service/graph/model/entity"
	"github.com/vnnyx/article-service/internal/infrastructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArticleRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewArticleRepository(database *mongo.Database) ArticleRepository {
	return &ArticleRepositoryImpl{
		Collection: database.Collection("authors"),
	}
}

func (repository *ArticleRepositoryImpl) CreateArticle(authorID string, article *entity.ArticleEntity) error {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	filter := bson.M{"_id": authorID}

	set := bson.M{"$push": bson.M{"articles": article}}

	_, err := repository.Collection.UpdateOne(ctx, filter, set)
	return err
}

func (repository *ArticleRepositoryImpl) FindArticleByID(articleID string) (author *entity.AuthorEntity, err error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	filter := bson.M{"articles._id": articleID}

	opts := options.FindOne().SetProjection(bson.M{
		"_id":      1,
		"username": 1,
		"email":    1,
		"passwrod": 1,
		"articles": bson.M{"$elemMatch": bson.M{"_id": articleID}},
	})

	err = repository.Collection.FindOne(ctx, filter, opts).Decode(&author)
	if err != nil {
		return author, err
	}

	return author, nil
}

func (repository *ArticleRepositoryImpl) FindArticleByName(name string) (authors []*entity.AuthorEntity, err error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{"articles.name": bson.M{"$regex": name, "$options": "im"}})
	if err != nil {
		return authors, err
	}

	err = cursor.All(ctx, &authors)
	if err != nil {
		return authors, err
	}

	return authors, err
}

func (repository *ArticleRepositoryImpl) UpdateArticle(article *entity.ArticleEntity) (bool, error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	fmt.Println(article)
	filter := bson.M{
		"articles._id": article.ID,
	}
	set := bson.M{"$set": bson.M{"articles.$": article}}

	_, err := repository.Collection.UpdateOne(ctx, filter, set)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repository *ArticleRepositoryImpl) DeleteArticle(articleID string) (bool, error) {
	ctx, cancel := infrastructure.NewMongoContext()
	defer cancel()

	filter := bson.M{
		"articles._id": articleID,
	}

	pull := bson.M{
		"$pull": bson.M{
			"articles": bson.M{"_id": articleID},
		}}

	_, err := repository.Collection.UpdateOne(ctx, filter, pull)
	if err != nil {
		return false, err
	}
	return true, nil
}
