package book_service

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
	"github.com/umarkotak/ytkidd-api/contract"
	"github.com/umarkotak/ytkidd-api/contract_resp"
	"github.com/umarkotak/ytkidd-api/datastore"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/repos/book_content_repo"
	"github.com/umarkotak/ytkidd-api/repos/book_repo"
	"github.com/umarkotak/ytkidd-api/repos/file_bucket_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_repo"
	"github.com/umarkotak/ytkidd-api/repos/user_subscription_repo"
	"github.com/umarkotak/ytkidd-api/utils"
)

func GetBooks(ctx context.Context, params contract.GetBooks) (contract_resp.GetBooks, error) {
	books, err := book_repo.GetByParams(ctx, params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.GetBooks{}, err
	}

	bookDatas := []contract_resp.Book{}
	for _, book := range books {
		var coverFileUrl string
		if book.Storage == model.STORAGE_R2 {
			coverFileUrl, _ = datastore.GetObjectUrl(ctx, book.CoverFilePath)
		} else {
			coverFileUrl = utils.GenRawFileUrl(config.Get().FileBucketPath, book.CoverFilePath)
		}

		bookData := contract_resp.Book{
			ID:           book.ID,
			Slug:         book.Slug,
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

	tagGroup := []contract_resp.TagGroup{
		{
			Name: "All Tags",
			Tags: tags,
		},
	}

	return contract_resp.GetBooks{
		TagGroup: tagGroup,
		Books:    bookDatas,
	}, nil
}

func GetBookDetail(ctx context.Context, params contract.GetBooks) (contract_resp.BookDetail, error) {
	var err error

	isSlug := utils.StringMustInt64(params.Slug) == 0

	var book model.Book
	if isSlug {
		book, err = book_repo.GetBySlug(ctx, params.Slug)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return contract_resp.BookDetail{}, err
		}
	} else {
		book, err = book_repo.GetByID(ctx, utils.StringMustInt64(params.Slug))
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return contract_resp.BookDetail{}, err
		}
	}

	if !book.IsFree() {
		if params.UserGuid == "" {
			return contract_resp.BookDetail{}, model.ErrLoginRequired
		}

		user, err := user_repo.GetByGuid(ctx, params.UserGuid)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return contract_resp.BookDetail{}, err
		}

		subs, err := user_subscription_repo.GetActiveByUserID(ctx, user.ID)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return contract_resp.BookDetail{}, err
		}

		if len(subs) <= 0 {
			return contract_resp.BookDetail{}, model.ErrSubscriptionRequired
		}
	}

	bookContents, err := book_content_repo.GetByBookID(ctx, book.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return contract_resp.BookDetail{}, err
	}

	bookContentDatas := []contract_resp.BookContent{}
	for _, bookContent := range bookContents {
		var imageFileUrl string
		if book.Storage == model.STORAGE_R2 {
			imageFileUrl, _ = datastore.GetObjectUrl(ctx, bookContent.ImageFilePath)
		} else {
			imageFileUrl = utils.GenRawFileUrl(config.Get().FileBucketPath, bookContent.ImageFilePath)
		}

		bookContentData := contract_resp.BookContent{
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
	// 	pdfUrl, _ = datastore.GetObjectUrl(ctx, book.PdfFileGuid)
	// } else {
	// 	pdfUrl = utils.GenRawFileUrl(config.Get().FileBucketPath, book.PdfFileGuid)
	// }

	var coverFileUrl string
	if book.Storage == model.STORAGE_R2 {
		coverFileUrl, _ = datastore.GetObjectUrl(ctx, book.CoverFilePath)
	} else {
		coverFileUrl = utils.GenRawFileUrl(config.Get().FileBucketPath, book.CoverFilePath)
	}
	bookDetail := contract_resp.BookDetail{
		ID:           book.ID,
		Slug:         book.Slug,
		Title:        book.Title,
		Description:  book.Description,
		CoverFileUrl: coverFileUrl,
		Tags:         book.Tags,
		Type:         book.Type,
		AccessTags:   book.AccessTags,
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

func UpdateBook(ctx context.Context, params contract.UpdateBook) error {
	book, err := book_repo.GetByID(ctx, params.ID)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	book.Slug = params.Slug
	book.Title = params.Title
	book.Description = params.Description
	book.Tags = params.Tags
	book.Type = params.Type
	book.Active = params.Active
	book.OriginalPdfUrl = params.OriginalPdfUrl
	book.AccessTags = params.AccessTags

	err = book_repo.Update(ctx, nil, book)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
