package data

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"student-planner/domain"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"google.golang.org/api/option"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetUser(id int) (domain.UserModel, error) {
	var user domain.UserModel
	err := s.db.QueryRow("SELECT * FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Badge)
	if err != nil {
		return domain.UserModel{}, err
	}
	if user.Badge == nil {
		user.Badge = pq.StringArray{}
	}
	return user, nil
}

func (s *UserStore) Login(email string, password string) (domain.UserModel, error) {
	query := "SELECT * FROM users WHERE email=$1"
	user := s.db.QueryRow(query, email)
	var u domain.UserModel
	err := user.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Badge)
	if u.Badge == nil {
		u.Badge = pq.StringArray{} //set to empty array
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.UserModel{}, fmt.Errorf("user not found")
		}
		return domain.UserModel{}, err
	}
	if u.Password == password {
		return u, nil
	}
	return domain.UserModel{}, fmt.Errorf("password incorrect")
}

func (s *UserStore) Register(name string, email string, password string) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	_, err := s.db.Exec(query, name, email, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserStore) UpdateBadge(id int, badge string) error {
	query := "UPDATE users SET badges = array_append(badges, $1) WHERE id=$2"
	_, err := s.db.Exec(query, badge, id)
	if err != nil {
		return err
	}
	return nil
}

// Gemini struct holds the generative AI client
type Gemini struct {
	client *genai.Client
	ctx    context.Context
}

// NewGemini initializes a Gemini instance with an API key
func NewGemini(apiKey string) (*Gemini, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %v", err)
	}

	return &Gemini{
		client: client,
		ctx:    ctx,
	}, nil
}
func (s *UserStore) AskGemini(message string) (string, error) {
	if _, err := os.Stat(".env"); err == nil {
		if loadErr := godotenv.Load(); loadErr != nil {
			log.Printf("Warning: Could not load .env file: %v", loadErr)
		}
	}
	var apiKey = os.Getenv("GEMINI_API_KEY")
	gemini, err := NewGemini(apiKey)
	if err != nil {
		log.Fatalf("Error initializing Gemini: %v", err)
	}
	defer gemini.client.Close()
	model := gemini.client.GenerativeModel("gemini-1.5-pro")

	resp, err := model.GenerateContent(gemini.ctx, genai.Text(message))
	if err != nil {
		return "", err
	}

	// Extract response text properly
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		if text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			return string(text), nil // Convert TextPart to string
		}
		return "", fmt.Errorf("response format mismatch")
	}
	return "", fmt.Errorf("no response from Gemini")
}
