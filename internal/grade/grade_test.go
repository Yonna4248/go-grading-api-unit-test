package grade

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCalculateGradeA(t *testing.T) {
	_, grade := CalculateGrade(80, 70, 90)
	assert.Equal(t, "A", grade)
}

func TestCalculateGradeB(t *testing.T) {
	_, grade := CalculateGrade(70, 70, 70)
	assert.Equal(t, "B", grade)
}

func TestCalculateGradeC(t *testing.T) {
	_, grade := CalculateGrade(60, 60, 60)
	assert.Equal(t, "C", grade)
}

func TestCalculateGradeD(t *testing.T) {
	_, grade := CalculateGrade(50, 50, 50)
	assert.Equal(t, "D", grade)
}

func TestCalculateGradeF(t *testing.T) {
	_, grade := CalculateGrade(30, 30, 30)
	assert.Equal(t, "F", grade)
}

func TestInvalidScore(t *testing.T) {
	_, grade := CalculateGrade(-1, 0, 0)
	assert.Equal(t, "Invalid", grade)
}

func TestBoundaryScore(t *testing.T) {
	_, grade := CalculateGrade(100, 100, 100)
	assert.Equal(t, "A", grade)
}
func TestCalculateGrade_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		homework float64
		midterm  float64
		final    float64
		expected string
	}{
		{"Grade A", 80, 70, 90, "A"},
		{"Grade B", 70, 70, 70, "B"},
		{"Grade C", 60, 60, 60, "C"},
		{"Grade D", 50, 50, 50, "D"},
		{"Grade F", 30, 30, 30, "F"},
		{"Invalid Score", -1, 0, 0, "Invalid"},
		{"Boundary Score", 100, 100, 100, "A"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, grade := CalculateGrade(tt.homework, tt.midterm, tt.final)
			assert.Equal(t, tt.expected, grade)
		})
	}
}
func TestCheckGrade(t *testing.T) {
	mockRepo := &MockRepository{}
	service := &GradeService{Repo: mockRepo}

	res, err := service.CheckGrade("65001")

	assert.NoError(t, err)
	assert.Equal(t, "65001", res.StudentID)
	assert.Equal(t, "A", res.Grade)
}

func TestGetGradeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{}
	handler := NewHandler(mockService)

	router := gin.Default()
	router.GET("/grade/:studentId", handler.GetGradeHandler)

	req, _ := http.NewRequest("GET", "/grade/65001", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
