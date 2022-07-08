package common

import (
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/AndreeJait/GO-ANDREE-UTILITIES/logs"
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/util/andreerror"
	"github.com/labstack/echo/v4"
)

func SystemResponse(ec echo.Context, data interface{}, err error) error {
	if err == nil {
		successResponse := andreerror.New(andreerror.SUCCESS, nil)
		return ec.JSON(successResponse.GetHTTPStatus(), &BaseResponseDto{
			Code:       successResponse.GetCode(),
			Message:    successResponse.GetMessage(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}

	if s, ok := err.(andreerror.ErrorStandard); ok {
		return ec.JSON(s.GetHTTPStatus(), &BaseResponseDto{
			Code:       s.GetCode(),
			Message:    s.GetMessage(),
			Errors:     s.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	} else {
		errResponse := andreerror.New(andreerror.SYSTEM_ERROR, err)
		return ec.JSON(errResponse.GetHTTPStatus(), &BaseResponseDto{
			Code:       errResponse.GetCode(),
			Message:    errResponse.GetMessage(),
			Errors:     errResponse.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
}

func SaveFile(files []*multipart.FileHeader, log logs.Logger, subfolder string) ([]string, error) {
	result := make([]string, 0)
	for _, file := range files {

		// Source
		src, err := file.Open()

		if err != nil {
			log.Errorf("[Failed to read File] path : %v", err)
			return nil, err
		}

		defer src.Close()

		split := strings.Split(file.Filename, ".")

		// Destination
		fileName := "assets/" + subfolder + time.Now().Format("2006-01-02:15-04:05") + "-" + RandStringRunes(10) + "." + split[len(split)-1]
		dst, err := os.Create(fileName)

		if err != nil {
			log.Errorf("[Failed to create file] path : %v", err)
			return []string{}, err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return []string{}, err
		}
		log.Infof("[Upload File] path : %s", fileName)
		result = append(result, fileName)
	}

	return result, nil
}

func DeleteFile(paths []string, log logs.Logger) {

	for _, path := range paths {
		e := os.Remove(path)
		if e != nil {
			log.Errorf("[Failed to delete File] path: %s", path)
		} else {
			log.Info("[Delete File] path : %s", path)
		}
	}
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
