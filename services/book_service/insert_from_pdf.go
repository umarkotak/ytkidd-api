package book_service

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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
	err := utils.CreateFolderIfNotExists(fmt.Sprintf("%s/book_pdfs", config.Get().FileBucketPath))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	pdfDir := fmt.Sprintf("%s/book_pdfs", config.Get().FileBucketPath)

	pdfFilePath := fmt.Sprintf("%s/uploaded-%v.pdf", pdfDir, params.Slug)
	err = os.WriteFile(pdfFilePath, params.PdfBytes, 0644)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	successFilePaths := []string{}
	err = datastore.Transaction(ctx, datastore.Get().Db, func(tx *sqlx.Tx) error {
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
			PdfFileGuid:    pdfFilePath,
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

		outputDir := fmt.Sprintf("%s/book_contents", config.Get().FileBucketPath)
		baseName := fmt.Sprintf("%d_%s", time.Now().UnixMilli(), params.Slug)
		outputPattern := filepath.Join(outputDir, fmt.Sprintf("%s-%%d.jpeg", baseName))
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

		pattern := filepath.Join(outputDir, fmt.Sprintf("%s-*.jpeg", baseName))
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

		for idx, filePath := range matches {
			guid := random.MustGenUUIDTimes(2)
			fileBucket := model.FileBucket{
				Guid:            guid,
				Name:            fmt.Sprintf("book %v - page %v", book.ID, idx),
				BaseType:        "image",
				Extension:       "jpeg",
				HttpContentType: "image/jpeg",
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
		// clean up successfully uploaded file
		for _, successFilePath := range successFilePaths {
			err = os.Remove(successFilePath)
			if err != nil {
				logrus.WithContext(context.Background()).Error(err)
			}
		}

		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
