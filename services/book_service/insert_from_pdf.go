package book_service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
	var err error

	if params.Storage == "" {
		params.Storage = model.STORAGE_LOCAL
	}

	bookDir := fmt.Sprintf("%s/books/%s", config.Get().FileBucketPath, params.Slug)

	err = utils.CreateFolderIfNotExists(bookDir)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}
	defer func() {
		if err != nil {
			utils.DeleteFolder(bookDir)
			if params.Storage == model.STORAGE_R2 {
				datastore.DeleteObjectsByPrefix(ctx, fmt.Sprintf("books/%s", params.Slug))
			}
		}
		if params.Storage == model.STORAGE_R2 {
			utils.DeleteFolder(bookDir)
		}
	}()

	pdfFilePath := fmt.Sprintf("%s/book.pdf", bookDir)
	err = os.WriteFile(pdfFilePath, params.PdfBytes, 0644)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	outputPattern := filepath.Join(bookDir, "%d.jpeg")
	gsArgs := []string{
		"-dNOPAUSE",     //
		"-dBATCH",       //
		"-dSAFER",       //
		"-sDEVICE=jpeg", //
		"-dJPEGQ=90",    //
		"-r225",         // Resolution in DPI
		fmt.Sprintf("-sOutputFile=%s", outputPattern), //
		pdfFilePath, //
	}

	cmd := exec.Command("/opt/homebrew/bin/gs", gsArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"cmd_output": output,
		}).Error(err)
		return err
	}

	pattern := filepath.Join(bookDir, "*.jpeg")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	if len(matches) == 0 {
		err = fmt.Errorf("empty pages")
		logrus.WithContext(ctx).Error(err)
		return err
	}

	bookObjectKey := fmt.Sprintf("books/%s/book.pdf", params.Slug)
	if params.Storage == model.STORAGE_R2 {
		err = datastore.UploadFileToR2(ctx, pdfFilePath, bookObjectKey, false)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}
	}

	successFilePaths := []string{}
	err = datastore.Transaction(ctx, datastore.Get().Db, func(tx *sqlx.Tx) error {
		pdfFileGuid := random.MustGenUUIDTimes(2)
		bookType := "default"
		if params.BookType != "" {
			bookType = params.BookType
		}
		book := model.Book{
			Slug:           params.Slug,
			Title:          params.Title,
			Description:    params.Description,
			CoverFileGuid:  "",
			Tags:           pq.StringArray{},
			Type:           bookType,
			PdfFileGuid:    pdfFileGuid,
			Active:         true,
			OriginalPdfUrl: params.OriginalPdfUrl,
			AccessTags:     pq.StringArray{model.ACCESS_TAG_FREE, model.ACCESS_TAG_BASIC},
			Storage:        params.Storage,
		}
		book.ID, err = book_repo.Insert(ctx, tx, book)
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}
		_, _, err = file_bucket_repo.Insert(ctx, tx, model.FileBucket{
			Guid:            pdfFileGuid,
			Name:            params.Slug,
			BaseType:        "application",
			Extension:       "pdf",
			HttpContentType: "application/pdf",
			Metadata:        model.FileBucketMetadata{},
			Data:            []byte{},
			ExactPath:       bookObjectKey,
			Storage:         params.Storage,
			SizeKb:          0,
		})
		if err != nil {
			logrus.WithContext(ctx).Error(err)
			return err
		}

		coverFileGuid := ""
		for idx, filePath := range matches {
			bookContentObjectKey := fmt.Sprintf("books/%s/%v.jpeg", params.Slug, idx+1)
			if params.Storage == model.STORAGE_R2 {
				err = datastore.UploadFileToR2(ctx, filePath, bookContentObjectKey, false)
				if err != nil {
					logrus.WithContext(ctx).Error(err)
					return err
				}
			}

			guid := random.MustGenUUIDTimes(2)
			fileBucket := model.FileBucket{
				Guid:            guid,
				Name:            fmt.Sprintf("book %v - page %v", book.ID, idx+1),
				BaseType:        "image",
				Extension:       "jpeg",
				HttpContentType: "image/jpeg",
				Metadata:        model.FileBucketMetadata{},
				Data:            []byte{},
				ExactPath:       bookContentObjectKey,
				Storage:         params.Storage,
				SizeKb:          0,
			}
			_, fileGuid, err := file_bucket_repo.Insert(ctx, tx, fileBucket)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return err
			}

			bookContent := model.BookContent{
				BookID:        book.ID,
				PageNumber:    int64(idx + 1),
				ImageFileGuid: fileGuid,
				Description:   "",
				Metadata:      model.BookContentMetadata{},
			}
			_, err = book_content_repo.Insert(ctx, tx, bookContent)
			if err != nil {
				logrus.WithContext(ctx).Error(err)
				return err
			}

			if idx == 0 {
				coverFileGuid = fileGuid
			}

			successFilePaths = append(successFilePaths, filePath)
			logrus.WithContext(ctx).Infof("success inserting image %v/%v", idx+1, len(matches))
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
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
