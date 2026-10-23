package main

import (
	"errors"
	"net/http"
	"net/mail"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	db *gorm.DB
}

func newUserHandler(db *gorm.DB) *userHandler {
	return &userHandler{db: db}
}

// list godoc
// @Summary Daftar semua user
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func (h *userHandler) list(c *gin.Context) {
	users := make([]User, 0)
	if err := h.db.Order("user_id ASC").Find(&users).Error; err != nil {
		respondInternalError(c)
		return
	}
	c.JSON(http.StatusOK, users)
}

// get godoc
// @Summary Ambil user berdasarkan ID
// @Tags users
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{userId} [get]
func (h *userHandler) get(c *gin.Context) {
	user, ok := h.findByID(c)
	if !ok {
		return
	}
	c.JSON(http.StatusOK, user)
}

// create godoc
// @Summary Buat user baru
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserRequest true "Data user"
// @Success 201 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func (h *userHandler) create(c *gin.Context) {
	request, ok := bindUserRequest(c)
	if !ok {
		return
	}

	exists, err := h.emailExists(request.Email, 0)
	if err != nil {
		respondInternalError(c)
		return
	}
	if exists {
		respondEmailConflict(c)
		return
	}

	user := User{FirstName: request.FirstName, LastName: request.LastName, Email: request.Email}
	if err := h.db.Create(&user).Error; err != nil {
		respondWriteError(c, err)
		return
	}

	c.Header("Location", "/api/users/"+strconv.FormatUint(uint64(user.UserID), 10))
	c.JSON(http.StatusCreated, user)
}

// update godoc
// @Summary Perbarui user
// @Tags users
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param user body UserRequest true "Data user"
// @Success 200 {object} User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{userId} [put]
func (h *userHandler) update(c *gin.Context) {
	request, ok := bindUserRequest(c)
	if !ok {
		return
	}

	user, ok := h.findByID(c)
	if !ok {
		return
	}

	exists, err := h.emailExists(request.Email, user.UserID)
	if err != nil {
		respondInternalError(c)
		return
	}
	if exists {
		respondEmailConflict(c)
		return
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	if err := h.db.Save(&user).Error; err != nil {
		respondWriteError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

// delete godoc
// @Summary Hapus user
// @Tags users
// @Param userId path int true "User ID"
// @Success 204
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users/{userId} [delete]
func (h *userHandler) delete(c *gin.Context) {
	user, ok := h.findByID(c)
	if !ok {
		return
	}

	if err := h.db.Delete(&user).Error; err != nil {
		respondInternalError(c)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *userHandler) findByID(c *gin.Context) (User, bool) {
	var user User
	if err := h.db.First(&user, c.Param("userId")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			respondInternalError(c)
		}
		return User{}, false
	}
	return user, true
}

func (h *userHandler) emailExists(email string, exceptUserID uint) (bool, error) {
	var count int64
	err := h.db.Model(&User{}).
		Where("email = ? AND user_id <> ?", email, exceptUserID).
		Count(&count).Error
	return count > 0, err
}

func bindUserRequest(c *gin.Context) (UserRequest, bool) {
	var request UserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user data",
			"details": err.Error(),
		})
		return UserRequest{}, false
	}

	request.FirstName = strings.TrimSpace(request.FirstName)
	request.LastName = strings.TrimSpace(request.LastName)
	request.Email = strings.ToLower(strings.TrimSpace(request.Email))

	address, err := mail.ParseAddress(request.Email)
	if request.FirstName == "" || request.LastName == "" || err != nil || address.Address != request.Email {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
		return UserRequest{}, false
	}
	return request, true
}

func respondWriteError(c *gin.Context, err error) {
	if strings.Contains(strings.ToLower(err.Error()), "unique") {
		respondEmailConflict(c)
		return
	}
	respondInternalError(c)
}

func respondEmailConflict(c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{"message": "Email already exists"})
}

func respondInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
}
