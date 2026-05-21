package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func ValidateFileAndUpload(fileHead *multipart.FileHeader) (error, string) {

	if fileHead.Size > 5<<20 {
		return errors.New("File too large.(Max 5MB)"), ""
	}
	extFile := strings.ToLower(filepath.Ext(fileHead.Filename))
	if !ExtValidate[extFile] {
		return errors.New("Invalid file."), ""
	}

	//Muốn kiểm tra mime của 1 *multipart.FileHeader ta cần mử file đấy ra, cho 1 buffer đọc 512 byte, rồi cho http.DetectContentType đọc buffer đó
	//để xác định mime của file
	file, err := fileHead.Open()
	if err != nil {
		return errors.New("Failed to open file."), ""
	}
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return errors.New("Failed to read file."), ""
	}
	defer file.Close()
	mimetype := http.DetectContentType(buffer)
	if !MimeValidate[mimetype] {
		return errors.New("Invalid file."), ""
	}
	fileName := fmt.Sprintf("%s-%s", uuid.New().String(), fileHead.Filename)
	savePath := filepath.Join("./output/", fileName)
	out, err := os.Create(savePath)
	if err != nil {
		return err, ""
	}
	defer out.Close()
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return err, ""
	}
	io.Copy(out, file)
	return nil, fileName
}
