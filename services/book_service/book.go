package book_service

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/model/resp_contract"
	"github.com/umarkotak/ytkidd-api/repos/book_content_repo"
	"github.com/umarkotak/ytkidd-api/repos/book_repo"
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
