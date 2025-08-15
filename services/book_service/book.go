package book_service

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
	"github.com/umarkotak/ytkidd-api/datastore"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/book_content_repo"
	"github.com/umarkotak/ytkidd-api/repos/book_repo"
	"github.com/umarkotak/ytkidd-api/repos/file_bucket_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_subscription_repo"
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
		var coverFileUrl string
		if book.Storage == model.STORAGE_R2 {
			coverFileUrl, _ = datastore.GetPresignedUrl(ctx, book.CoverFilePath, 1*time.Minute)
		} else {
			coverFileUrl = utils.GenRawFileUrl(config.Get().FileBucketPath, book.CoverFilePath)
		}

		bookData := resp_contract.Book{
			ID:           book.ID,
			Title:        book.Title,
			Description:  book.Description,
			CoverFileUrl: coverFileUrl,
			Tags:         book.Tags,
			Type:         book.Type,
			IsFree:       book.IsFree(),
		}
		bookDatas = append(bookDatas, bookData)
	}

	tags, err := book_repo.GetTags(ctx)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
	}

	return resp_contract.GetBooks{
		Tags:  tags,
		Books: bookDatas,
	}, nil
}

func GetBookDetail(ctx context.Context, params contract.GetBooks) (resp_contract.BookDetail, error) {
	book, err := book_repo.GetByID(ctx, params.BookID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.BookDetail{}, err
	}

	if !book.IsFree() {
		if params.UserGuid == "" {
			return resp_contract.BookDetail{}, model.ErrLoginRequired
		}

		user, err := user_repo.GetByGuid(ctx, params.UserGuid)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return resp_contract.BookDetail{}, err
		}

		subs, err := user_subscription_repo.GetActiveByUserID(ctx, user.ID)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return resp_contract.BookDetail{}, err
		}

		if len(subs) <= 0 {
			return resp_contract.BookDetail{}, model.ErrSubscriptionRequired
		}
	}

	bookContents, err := book_content_repo.GetByBookID(ctx, book.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return resp_contract.BookDetail{}, err
	}

	bookContentDatas := []resp_contract.BookContent{}
	for _, bookContent := range bookContents {
		var imageFileUrl string
		if book.Storage == model.STORAGE_R2 {
			imageFileUrl, _ = datastore.GetPresignedUrl(ctx, bookContent.ImageFilePath, 1*time.Minute)
		} else {
			imageFileUrl = utils.GenRawFileUrl(config.Get().FileBucketPath, bookContent.ImageFilePath)
		}

		bookContentData := resp_contract.BookContent{
			ID:           bookContent.ID,
			BookID:       bookContent.BookID,
			PageNumber:   bookContent.PageNumber,
			ImageFileUrl: imageFileUrl,
			Description:  bookContent.Description,
		}

		bookContentDatas = append(bookContentDatas, bookContentData)
	}

	var pdfUrl string
	// TODO: implement logic
	// if book.Storage == model.STORAGE_R2 {
	// 	pdfUrl, _ = datastore.GetPresignedUrl(ctx, book.PdfFileGuid, 1*time.Minute)
	// } else {
	// 	pdfUrl = utils.GenRawFileUrl(config.Get().FileBucketPath, book.PdfFileGuid)
	// }

	var coverFileUrl string
	if book.Storage == model.STORAGE_R2 {
		coverFileUrl, _ = datastore.GetPresignedUrl(ctx, book.CoverFilePath, 1*time.Minute)
	} else {
		coverFileUrl = utils.GenRawFileUrl(config.Get().FileBucketPath, book.CoverFilePath)
	}
	bookDetail := resp_contract.BookDetail{
		ID:           book.ID,
		Title:        book.Title,
		Description:  book.Description,
		CoverFileUrl: coverFileUrl,
		Tags:         book.Tags,
		Type:         book.Type,
		Contents:     bookContentDatas,
		PdfUrl:       pdfUrl,
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

	fileBucketGuids := make([]string, 0, len(bookContents)+1)
	fileBucketGuids = append(fileBucketGuids, book.PdfFileGuid)
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

		if book.Storage == model.STORAGE_R2 {
			keys := make([]string, 0, len(fileBuckets))
			for _, fileBucket := range fileBuckets {
				keys = append(keys, fileBucket.ExactPath)
			}
			err = datastore.DeleteByKeys(ctx, keys)

		} else {
			err = utils.DeleteFolder(fmt.Sprintf("%s/books/%s", config.Get().FileBucketPath, book.Slug))
		}
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
