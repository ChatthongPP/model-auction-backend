package controllers

import (
	"fmt"
	"net/http"
	"os"

	"backend-service/internal/domain"
	"backend-service/pkg/utilities/responses"

	"github.com/labstack/echo/v4"
)

func (c *Controller) UploadMedia2(ctx echo.Context) error {
	// 1. Bind & Validate
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Invalid multipart form"))
	}

	files := form.File["file"]
	if len(files) == 0 {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "File is required"))
	}

	// 2. For simplicity, allow one file only
	file := files[0]

	// 3. Call UseCase to handle file save logic
	url, err := c.uc.UploadMedia2(file)
	if err != nil {
		switch err {
		case domain.ErrUnsupportedFileType:
			return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Unsupported file type"))
		default:
			return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, "Upload failed"))
		}
	}

	// 4. Return success
	return ctx.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "File uploaded successfully", echo.Map{
		"url": url,
	}))
}

func (c *Controller) UploadMedia(ctx echo.Context) error {
	// 1. รับ topic จาก form field
	topic := ctx.FormValue("topic")
	if topic == "" {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Topic is required"))
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Invalid multipart form"))
	}

	files := form.File["file"]
	if len(files) == 0 {
		return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "File is required"))
	}

	file := files[0]

	// 2. ส่ง topic ไปให้ usecase
	url, err := c.uc.UploadMedia(file, topic)
	if err != nil {
		switch err {
		case domain.ErrUnsupportedFileType:
			return ctx.JSON(http.StatusBadRequest, responses.Error(http.StatusBadRequest, "Unsupported file type"))
		default:
			return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, "Upload failed"))
		}
	}

	return ctx.JSON(http.StatusCreated, responses.Ok(http.StatusCreated, "File uploaded successfully", echo.Map{
		"url": url,
	}))
}

func (c *Controller) GetMedia2(ctx echo.Context) error {
	filename := ctx.Param("filename")

	// 1. Call UseCase to fetch file path
	filePath, err := c.uc.GetMediaPath(filename)
	if err != nil {
		if err == domain.ErrFileNotFound {
			return ctx.JSON(http.StatusNotFound, responses.Error(http.StatusNotFound, "File not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, responses.Error(http.StatusInternalServerError, "Failed to get file"))
	}

	// 2. Serve file
	return ctx.File(filePath)
}

func (c *Controller) GetMedia(ctx echo.Context) error {
	topic := ctx.Param("topic")
	filename := ctx.Param("filename")
	path := fmt.Sprintf("./uploads/%s/%s", topic, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return ctx.JSON(http.StatusNotFound, responses.Error(http.StatusNotFound, "File not found"))
	}

	return ctx.File(path)
}
