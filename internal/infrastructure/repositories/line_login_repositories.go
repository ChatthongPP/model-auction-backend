package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"backend-service/internal/domain"
	"backend-service/internal/infrastructure/database/models"

	"gorm.io/gorm"
)

func (r *Repo) CreateLineUser(lineUser *domain.Line) error {
	log.Printf("CreateLineUser: %+v", lineUser)
	dbRepo := models.LoginModel{
		UID:      lineUser.UID,
		DisplayName: lineUser.DisplayName,
		Picture:  lineUser.PictureURL,
	}

	if err := r.db.Create(&dbRepo).Error; err != nil {
		return err
	}

	// Populate the ID back into the lineUser
	lineUser.Id = dbRepo.Id

	return nil
}

func (r *Repo) FindLineUserByUID(lineUID string) (*domain.Line, error) {
	var dbRepo models.LoginModel
	if err := r.db.Where("uid = ? AND is_delete is NULL ", lineUID).First(&dbRepo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	lineUser := domain.Line{
		Id:          dbRepo.Id,
		UID:         dbRepo.UID,
		DisplayName: dbRepo.DisplayName,
		PictureURL:  dbRepo.Picture,
	}
	return &lineUser, nil
}
func (r *Repo) VerifyIDToken(idToken string) (*domain.Line, error) {
	endpoint := "https://api.line.me/oauth2/v2.1/verify"
	data := url.Values{}
	data.Set("id_token", idToken)
	data.Set("client_id", os.Getenv("CHANNEL_ID"))

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to verify ID token failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to verify ID token: %s", body)
	}

	var verifiedData struct {
		Sub     string `json:"sub"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	err = json.Unmarshal(body, &verifiedData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	profile := domain.Line{
		UID:         verifiedData.Sub,
		DisplayName: verifiedData.Name,
		PictureURL:  verifiedData.Picture,
	}

	return &profile, nil
}

func (r *Repo) UpdateLineUser(lineUser domain.Line) error {
	log.Printf("UpdateLineUser: %+v", lineUser)
	dbRepo := models.LoginModel{
		UID:      lineUser.UID,
		DisplayName: lineUser.DisplayName,
		Picture:  lineUser.PictureURL,
	}

	if err := r.db.Model(&models.LoginModel{}).Where("uid = ?", lineUser.UID).Updates(dbRepo).Error; err != nil {
		return fmt.Errorf("failed to update line user: %w", err)
	}

	return nil
}