package book_service

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/datastore"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/book_content_repo"
	"github.com/umarkotak/ytkidd-api/repos/book_repo"
	"github.com/umarkotak/ytkidd-api/repos/file_bucket_repo"
	"github.com/umarkotak/ytkidd-api/utils"
)

func GetBooks(ctx context.Context, params contract.GetBooks) (resp_contract.GetBooks, error) {
	books, err := book_repo.GetByParams(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.GetBooks{}, err
	}

	bookDatas := []resp_contract.Book{}
	for _, book := range books {
		bookData := resp_contract.Book{
			ID:           book.ID,
			Title:        book.Title,
			Description:  book.Description,
			CoverFileUrl: utils.GenRawFileUrl(book.CoverFilePath),
			Tags:         book.Tags,
			Type:         book.Type,
		}
		bookDatas = append(bookDatas, bookData)
	}

	return resp_contract.GetBooks{
		Books: bookDatas,
	}, nil
}

func GetBookDetail(ctx context.Context, params contract.GetBooks) (resp_contract.BookDetail, error) {
	book, err := book_repo.GetByID(ctx, params.BookID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.BookDetail{}, err
	}

	bookContents, err := book_content_repo.GetByBookID(ctx, book.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.BookDetail{}, err
	}

	bookContentDatas := []resp_contract.BookContent{}
	for _, bookContent := range bookContents {
		bookContentData := resp_contract.BookContent{
			ID:           bookContent.ID,
			BookID:       bookContent.BookID,
			PageNumber:   bookContent.PageNumber,
			ImageFileUrl: utils.GenRawFileUrl(bookContent.ImageFilePath),
			Description:  bookContent.Description,
		}

		bookContentDatas = append(bookContentDatas, bookContentData)
	}

	bookDetail := resp_contract.BookDetail{
		ID:           book.ID,
		Title:        book.Title,
		Description:  book.Description,
		CoverFileUrl: utils.GenRawFileUrl(book.CoverFilePath),
		Tags:         book.Tags,
		Type:         book.Type,
		Contents:     bookContentDatas,
	}

	return bookDetail, nil
}

func DeleteBook(ctx context.Context, params contract.DeleteBook) error {
	book, err := book_repo.GetByID(ctx, params.BookID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	bookContents, err := book_content_repo.GetByBookID(ctx, book.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	fileBucketGuids := []string{}
	for _, bookContent := range bookContents {
		fileBucketGuids = append(fileBucketGuids, bookContent.ImageFileGuid)
	}

	fileBuckets, err := file_bucket_repo.GetByGuids(ctx, fileBucketGuids)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	err = datastore.Transaction(ctx, datastore.Get().Db, func(tx *sqlx.Tx) error {
		err = book_content_repo.DeleteByBookID(ctx, tx, book.ID)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		err = book_repo.Delete(ctx, tx, book.ID)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		err = file_bucket_repo.DeleteByGuids(ctx, tx, fileBucketGuids)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		for _, fileBucket := range fileBuckets {
			err = utils.DeleteFileIfExists(fileBucket.ExactPath)
			if err != nil {
				logrus.WithContext(ctx).WithFields(logrus.Fields{
					"book_id":          book.ID,
					"file_bucket_guid": fileBucket.Guid,
					"file_path":        fileBucket.ExactPath,
				}).Error(err)
				return err
			}
		}

		return nil
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
