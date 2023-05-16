package repository

import (
	"context"

	"github.com/ahmadaidin/ginplating/domain/dto"
	"github.com/ahmadaidin/ginplating/domain/entity"
	"github.com/ahmadaidin/ginplating/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	bookCollection = "books"
)

type BookRepository struct {
	db         *mongo.Database
	collection string
}

func NewBookRepository(db *mongo.Database) *BookRepository {
	return &BookRepository{
		db:         db,
		collection: bookCollection,
	}
}

func (r *BookRepository) FindAll(ctx context.Context, options ...dto.FindAllBookOptions) (books []entity.Book, total int64, err errors.Error) {
	option := dto.MergeFindAllBookOptions(options...)

	cursor, ierr := r.db.Collection(r.collection).Find(ctx, findAllBookOptionToFindAllFilter(option))
	if ierr != nil {
		return books, total, errors.NewInternalError(ierr, "error in BookRepository.FindAll when calling db.Collection.Find")
	}

	if ierr := cursor.All(ctx, &books); ierr != nil {
		return books, total, errors.NewInternalError(ierr, "error in BookRepository.FindAll when calling cursor.All")
	}

	total, ierr = r.db.Collection(r.collection).CountDocuments(ctx, findAllBookOptionToFindAllFilter(option))
	if ierr != nil {
		return books, total, errors.NewInternalError(ierr, "error in BookRepository.FindAll when calling db.Collection.CountDocuments")
	}

	return books, total, err
}

func findAllBookOptionToFindAllFilter(option dto.FindAllBookOptions) (filter bson.M) {
	filter = bson.M{}

	if option.Title != "" {
		filter["title"] = option.Title
	}

	if option.Author != "" {
		filter["author"] = option.Author
	}

	return filter
}
