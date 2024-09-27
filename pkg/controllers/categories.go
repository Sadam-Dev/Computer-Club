package controllers

import (
	"ComputerClub/errs"
	"ComputerClub/models"
	"ComputerClub/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateCategory godoc
// @Summary      Create a new category
// @Description  Creates a new category for the club.
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      models.Category  true  "Category Details"
// @Success      201       {object}  map[string]string
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /categories [post]
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	if err := service.CreateCategory(category); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "category created successfully"})
}

// GetCategoryByID godoc
// @Summary      Get category by ID
// @Description  Retrieve a category by its ID.
// @Tags         categories
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  models.Category
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /categories/{id} [get]
func GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := service.GetCategoryByID(uint(id))
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

// GetAllCategories godoc
// @Summary      Get all categories
// @Description  Retrieve a list of all categories.
// @Tags         categories
// @Produce      json
// @Success      200  {array}   models.Category
// @Failure      500  {object}  map[string]string
// @Router       /categories [get]
func GetAllCategories(c *gin.Context) {
	categories, err := service.GetAllCategories()
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

// UpdateCategory godoc
// @Summary      Update a category
// @Description  Update an existing category by ID.
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        id        path      int             true  "Category ID"
// @Param        category  body      models.Category  true  "Updated Category Details"
// @Success      200       {object}  map[string]string
// @Failure      400       {object}  map[string]string
// @Failure      404       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		handleError(c, errs.ErrValidationFailed)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	category.ID = uint(id)

	if err := service.UpdateCategory(category); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category updated successfully"})
}

// DeleteCategory godoc
// @Summary      Delete a category
// @Description  Delete a category by its ID.
// @Tags         categories
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteCategory(uint(id)); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}
