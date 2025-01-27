package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"koriebruh/management/controller"
	"koriebruh/management/domain"
	"koriebruh/management/repository"
	"koriebruh/management/routes"
	"koriebruh/management/service"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupTestDB() *gorm.DB {
	dsn := "root:korie123@tcp(127.0.0.1:3306)/go_management_test?charset=utf8mb4&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ { // Coba ulang hingga 5 kali
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}

		log.Printf("Percobaan %d: Gagal menghubungkan ke database. Coba lagi dalam 5 detik...\n", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(errors.New("gagal terhubung ke database"))
	}
	// Auto Migrate
	err = db.AutoMigrate(
		&domain.Admin{},
		&domain.Category{},
		&domain.Supplier{},
		&domain.Item{},
	)
	if err != nil {
		panic(errors.New("Failed Migrated"))
	}

	return db
}

func setupRouter() *fiber.App {

	db := setupTestDB()
	validate := validator.New()

	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(db, authRepository, validate)
	authController := controller.NewAuthController(authService)

	app := fiber.New()
	routes.SetupAuthRoutes(app, authController)

	return app
}

func TestRegisterControllerSuccess(t *testing.T) {

	router := setupRouter()

	// PAYLOAD REQUEST
	payload := map[string]string{
		"username": "fatlem",
		"password": "fatlem224",
		"email":    "fatlem@example.com",
	}
	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := router.Test(req)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "CREATED", responseBody["status"])
	assert.Equal(t, "SUCCESS CREATED", responseBody["data"])

}

func TestRegisterControllerFail(t *testing.T) {

	router := setupRouter()

	// PAYLOAD REQUEST
	payload := map[string]string{
		"username": "fat",
		"password": "",
		"email":    "87y869sdasda",
	}
	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := router.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestLoginSuccess(t *testing.T) {

	router := setupRouter()

	// STEP 1: Register a new user
	registerPayload := map[string]string{
		"username": "admin",
		"password": "admin123",
		"email":    "admin@example.com",
	}
	registerBytes, _ := json.Marshal(registerPayload)
	registerReq := httptest.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(registerBytes))
	registerReq.Header.Set("Content-Type", "application/json")

	registerResp, _ := router.Test(registerReq)
	assert.Equal(t, http.StatusCreated, registerResp.StatusCode)

	// PAYLOAD REQUEST
	payload := map[string]string{
		"username": "admin",
		"password": "admin123",
	}
	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := router.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)

	data := responseBody["data"].(map[string]interface{})
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "SUCCESS", responseBody["status"])
	assert.NotEmpty(t, data["token"])
}

func TestLoginFail(t *testing.T) {

	router := setupRouter()

	// PAYLOAD REQUEST
	payload := map[string]string{
		"username": "adminAHHH",
		"password": "",
	}
	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := router.Test(req)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestLogoutSuccess(t *testing.T) {
	router := setupRouter()

	// STEP 1: Login (Simulasi)
	loginPayload := map[string]string{
		"username": "admin",
		"password": "admin123",
	}
	loginBytes, _ := json.Marshal(loginPayload)
	reqLogin := httptest.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(loginBytes))
	reqLogin.Header.Set("Content-Type", "application/json")

	respLogin, _ := router.Test(reqLogin)
	assert.Equal(t, http.StatusOK, respLogin.StatusCode)

	var loginResponse map[string]interface{}
	json.NewDecoder(respLogin.Body).Decode(&loginResponse)
	token := loginResponse["data"].(map[string]interface{})["token"].(string)

	// STEP 2: Logout
	reqLogout := httptest.NewRequest(http.MethodPost, "/api/auth/logout", nil)
	reqLogout.Header.Set("Authorization", "Bearer "+token) // Tambahkan spasi setelah "Bearer"
	respLogout, _ := router.Test(reqLogout)

	var responseBody map[string]interface{}
	json.NewDecoder(respLogout.Body).Decode(&responseBody)

	assert.Equal(t, http.StatusOK, respLogout.StatusCode)
	assert.Equal(t, "OK", responseBody["status"])

}

// FAIL NOI INCUDE TOKEN
func TestLogoutFail(t *testing.T) {
	router := setupRouter()

	reqLogout := httptest.NewRequest(http.MethodPost, "/api/auth/logout", nil)
	reqLogout.Header.Set("Authorization", "Bearer ") // Tambahkan spasi setelah "Bearer"
	respLogout, _ := router.Test(reqLogout)

	assert.Equal(t, http.StatusUnauthorized, respLogout.StatusCode)

}
