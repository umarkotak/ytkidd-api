package book_service

import (
	"context"
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
	"time"

	"github.com/gen2brain/go-fitz"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/config"
	"github.com/umarkotak/ytkidd-api/datastore"
	"github.com/umarkotak/ytkidd-api/model"
	"github.com/umarkotak/ytkidd-api/model/contract"
	"github.com/umarkotak/ytkidd-api/repos/book_content_repo"
	"github.com/umarkotak/ytkidd-api/repos/book_repo"
	"github.com/umarkotak/ytkidd-api/repos/file_bucket_repo"
	"github.com/umarkotak/ytkidd-api/utils"
	"github.com/umarkotak/ytkidd-api/utils/random"
)

func InsertFromPdf(ctx context.Context, params contract.InsertFromPdf) error {
	tempDir, err := os.MkdirTemp("", "pdf-images-")
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}
	defer os.RemoveAll(tempDir)

	tempPDFPath := fmt.Sprintf("%s/uploaded.pdf", tempDir)
	err = os.WriteFile(tempPDFPath, params.PdfBytes, 0644)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	doc, err := fitz.New(tempPDFPath)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}
	defer doc.Close()

	successFilePaths := []string{}
	err = datastore.Transaction(ctx, datastore.Get().Db, func(tx *sqlx.Tx) error {
		bookType := "default"
		if params.BookType != "" {
			bookType = params.BookType
		}
		book := model.Book{
			Title:          params.Title,
			Description:    params.Description,
			CoverFileGuid:  "",
			Tags:           pq.StringArray{},
			Type:           bookType,
			PdfFileGuid:    "",
			Active:         true,
			OriginalPdfUrl: params.OriginalPdfUrl,
		}
		book.ID, err = book_repo.Insert(ctx, tx, book)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		err = utils.CreateFolderIfNotExists(fmt.Sprintf("%s/book_contents", config.Get().FileBucketPath))
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		coverFileGuid := ""
		for pageNum := 0; pageNum < doc.NumPage(); pageNum++ {
			img, err := doc.Image(pageNum)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return err
			}

			var extension, httpContentType string
			if params.ImgFormat == "jpeg" {
				extension = "jpeg"
				httpContentType = "image/jpeg"
			} else {
				extension = "png"
				httpContentType = "image/png"
			}

			guid := random.MustGenUUIDTimes(2)
			if params.CustomImageSlug != "" {
				guid = fmt.Sprintf("%v-%v", params.CustomImageSlug, pageNum)
			}
			filePath := fmt.Sprintf("%s/book_contents/%v_%s.%s", config.Get().FileBucketPath, time.Now().UnixMilli(), guid, extension)
			file, err := os.Create(filePath)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return err
			}
			defer file.Close()

			if params.ImgFormat == "jpeg" {
				err = jpeg.Encode(file, img, &jpeg.Options{
					Quality: 90,
				})
			} else {
				err = png.Encode(file, img)
			}
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return err
			}
			successFilePaths = append(successFilePaths, filePath)

			fileBucket := model.FileBucket{
				Guid:            guid,
				Name:            fmt.Sprintf("book %v - page %v", book.ID, pageNum+1),
				BaseType:        "image",
				Extension:       extension,
				HttpContentType: httpContentType,
				Metadata:        model.FileBucketMetadata{},
				Data:            []byte{},
				ExactPath:       filePath,
			}
			_, fileGuid, err := file_bucket_repo.Insert(ctx, tx, fileBucket)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return err
			}

			bookContent := model.BookContent{
				BookID:        book.ID,
				PageNumber:    int64(pageNum + 1),
				ImageFileGuid: fileGuid,
				Description:   "",
				Metadata:      model.BookContentMetadata{},
			}
			_, err = book_content_repo.Insert(ctx, tx, bookContent)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return err
			}

			if pageNum == 0 {
				coverFileGuid = fileGuid
			}

			logrus.WithContext(ctx).Infof("success inserting image %v/%v", pageNum+1, doc.NumPage())
		}

		book.CoverFileGuid = coverFileGuid
		err = book_repo.Update(ctx, tx, book)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		// clean up successfully uploaded file
		go func() {
			for _, successFilePath := range successFilePaths {
				err = os.Remove(successFilePath)
				if err != nil {
					logrus.WithContext(ctx).Error(err)
				}
			}
		}()

		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
