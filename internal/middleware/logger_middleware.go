package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type CustomResponseWrite struct { //custom để gán vào interface ctx.Write
	gin.ResponseWriter
	body *bytes.Buffer
}

// Khi gán ctx.Write bằng 1 biến có kdl là CustomResponseWrite thì phương thức Write này sẽ ghi đè tại phương thức Write của ctx.Write
// Mục đích để ép cứ mỗi lần gửi data thành công vào hàm này thì body của CustomResponseWrite sẽ bắt lấy được data đó để lưu trữ
func (w *CustomResponseWrite) Write(data []byte) (n int, err error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func LoggingMiddleware(httpLogger *zerolog.Logger) gin.HandlerFunc {
	// if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
	// 	panic(err)
	// }
	// fileLog, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	// if err != nil {
	// 	panic(err)
	// }

	//zerolog là 1 thư viện ghi log cực nhanh và tối ưu
	//New(fileLog): Dùng để biết nơi để ghi các dữ liệu của logger vào file nào
	//With:
	//Timestamp: Thêm thời gian vào các dòng log giúp ta biết được khi nào dòng log này xảy ra
	//Logger: Kết thúc và trả về đối tượng logger hoàn chỉnh
	return func(ctx *gin.Context) {
		start := time.Now()
		requestBody := make(map[string]any)
		var formFile []map[string]any
		contentType := ctx.GetHeader("Content-Type")
		if strings.HasPrefix(contentType, "multipart/form-data") {
			if err := ctx.Request.ParseMultipartForm(32 << 20); err == nil && ctx.Request.MultipartForm != nil {
				for key, val := range ctx.Request.MultipartForm.Value {
					if len(val) == 1 {
						requestBody[key] = val[0]
					} else {
						requestBody[key] = val
					}
				}

				for field, files := range ctx.Request.MultipartForm.File {
					for _, file := range files {
						formFile = append(formFile, map[string]any{
							"file_name":    file.Filename,
							"file_size":    formatSize(file.Size),
							"field":        field,
							"content-type": file.Header.Get("content-type"),
						})
					}
				}
			}
			if len(formFile) > 0 {
				requestBody["form_file"] = formFile
			}
		} else {
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				httpLogger.Error().Err(err).Msg("Failed to read body.")
			}

			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) //nạp lại dữ liệu về lại cho Body sau khi lấy đi
			if strings.HasPrefix(contentType, "application/json") {
				json.Unmarshal(bodyBytes, &requestBody) //giải mã mảng Bytes thành các cấu trúc như Json
			} else {
				values, _ := url.ParseQuery(string(bodyBytes))
				for key, val := range values {
					if len(val) == 1 {
						requestBody[key] = val[0]
					} else {
						requestBody[key] = val
					}
				}
			}
		}
		// ctx.MultipartForm()

		//RESPONSE
		responseWrite := &CustomResponseWrite{
			ResponseWriter: ctx.Writer,
			body:           bytes.NewBufferString(""),
		}
		ctx.Writer = responseWrite

		//next
		ctx.Next()
		duration := time.Since(start)
		responseContentType := ctx.Writer.Header().Get("content-type")
		responseBody := responseWrite.body.String()
		var response any
		if strings.HasPrefix(responseContentType, "image/") {
			response = "BINARY DATA"
		} else if strings.HasPrefix(responseContentType, "application/json") ||
			strings.HasPrefix(strings.TrimSpace(responseBody), "{") ||
			strings.HasPrefix(strings.TrimSpace(responseBody), "[") {
			if err := json.Unmarshal([]byte(responseBody), &response); err != nil {
				response = responseBody
			}
			fmt.Println(response)
		} else {
			response = responseBody
		}
		logInfo := httpLogger.Info()
		if ctx.Writer.Status() >= 500 {
			logInfo = httpLogger.Error()
		} else if ctx.Writer.Status() >= 400 {
			logInfo = httpLogger.Warn()
		}
		logInfo.
			Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("query", ctx.Request.URL.RawQuery).
			Int("status", ctx.Writer.Status()).
			Int64("duration_ms", duration.Milliseconds()).
			Str("clientIP", ctx.ClientIP()).
			Str("userAgent", ctx.Request.UserAgent()).
			Interface("request", requestBody).
			Interface("response", response).
			Msg("HTTP Request Log")
	}
}

func formatSize(size int64) string {
	if size > 1<<20 {
		return fmt.Sprintf("%.2fMB", float32(size)/1024/1024)
	} else if size > 1<<10 {
		return fmt.Sprintf("%.2fKB", float32(size)/1024)
	} else {
		return fmt.Sprintf("%dB", size)
	}
}
