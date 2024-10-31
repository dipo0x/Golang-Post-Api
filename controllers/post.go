package controller

import (
	"context"
	// "net/http"
	"time"
	// "fmt"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"go-rest-api/config"
	"go-rest-api/models"
	"github.com/google/uuid"
)

// func GetPosts(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	collection := config.MongoDatabase.Collection("posts")

// 	cursor, err := collection.Find(ctx, bson.M{}, options.Find())
// 	if err != nil {
// 		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
// 		log.Println("Error fetching posts:", err)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var posts []map[string]interface{}
// 	if err = cursor.All(ctx, &posts); err != nil {
// 		http.Error(w, "Failed to decode posts", http.StatusInternalServerError)
// 		log.Println("Error decoding posts:", err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintf("%+v", posts)))
// }

func CreatePost(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	post.ID = uuid.New()
	post.CreatedAt = time.Now()
	post.UpdatedAt = post.CreatedAt

	collection := config.MongoDatabase.Collection("posts")
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create post"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": 201,
		"success": true,
		"data": post,
	})
}