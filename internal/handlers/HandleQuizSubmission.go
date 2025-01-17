package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rehart-Kcalb/EduNexus-Monolith/internal/db"
)

func handleQuizSubmission(w http.ResponseWriter, r *http.Request, assignment db.Assignment, DB *db.Queries) {
	type quiz struct {
		Questions []string   `json:"questions"`
		Variant   [][]string `json:"variant"`
		Answers   []string   `json:"answers"`
	}
	type UserAnswers struct {
		Answers []string `json:"answers"`
	}
	var quiz1 quiz
	if err := json.Unmarshal(assignment.Content, &quiz1); err != nil {
		http.Error(w, "Failed to parse assignment content", http.StatusInternalServerError)
		return
	}

	var user_answer UserAnswers
	if err := json.NewDecoder(r.Body).Decode(&user_answer); err != nil {
		http.Error(w, "Failed to parse user answers", http.StatusBadRequest)
		return
	}

	if len(quiz1.Answers) != len(user_answer.Answers) {
		http.Error(w, "Mismatch in number of answers", http.StatusBadRequest)
		return
	}

	grade := 0
	for i := range quiz1.Answers {
		if quiz1.Answers[i] == user_answer.Answers[i] {
			grade++
		}
	}

	// TODO: Save grade to the database
	fmt.Fprintf(w, "Grade: %d/%d", grade, len(quiz1.Answers))
}
