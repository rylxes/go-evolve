package books

import (
	"github.com/gin-gonic/gin"
	"go-postgres/domain"
	"go-postgres/utils"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func (h Handler) paginate(c *gin.Context) *gorm.DB {
	pagination := utils.GeneratePaginationFromRequest(c)
	offset := (pagination.Page - 1) * pagination.Limit
	return h.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
}

// find book by author (one)
func (h Handler) findByIsbn(c *gin.Context, isbn string) {
	var book domain.Book
	if err := h.DB.Where("isbn = ?", isbn).Find(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func (h Handler) findByDate(c *gin.Context, dateFrom string, dateTo string) {
	//var book domain.Book
	var books []domain.Book

	if dateFrom != "" && dateTo != "" {
		if err := h.paginate(c).Where("created_at BETWEEN ? AND ?", dateFrom, dateTo).Find(&books).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}
	}

	//filter := h.DB
	//if dateFrom != "" {
	//	date, _ := time.Parse(dateFrom, dateFrom)
	//	log.Println(date)
	//	filter = filter.Where("created_at >= ?", date)
	//}
	//
	//if dateTo != "" {
	//	date, _ := time.Parse(dateTo, dateTo)
	//	log.Println(date)
	//	filter = filter.Where("created_at <= ?", date)
	//}
	//
	//if err := filter.Find(&books).Error; err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func (h Handler) FindBook(c *gin.Context) {
	isbn := c.Query("isbn")
	if isbn != "" {
		h.findByIsbn(c, isbn)
		return
	}

	dateFrom := c.Query("dateFrom")
	dateTo := c.Query("dateTo")
	if dateFrom != "" || dateTo != "" {
		h.findByDate(c, dateFrom, dateTo)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
}
