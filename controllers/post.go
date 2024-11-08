package controller

import (
	"context"
	"go-rest-api/config"
	"go-rest-api/models"
	"go-rest-api/types"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPosts(c *fiber.Ctx) error {
    collection := config.MongoDatabase.Collection("posts")

    var posts []models.Post

    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": 400,
			"success": false,
			"error": "Failed to retrieve posts"})
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var post models.Post
        if err := cursor.Decode(&post); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": 400,
			"success": false,
			"error": "Failed to decode post"})
        }
        posts = append(posts, post)
    }

    if err := cursor.Err(); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": 400,
			"success": false,
			"error": "Cursor error"})
    }

    var iPosts []types.IPost
    for _, post := range posts {
        iPost := types.IPost{
            ID: post.ID,
            Title:   post.Title,
            Content: post.Content,
            Author:  post.Author,
        }
        iPosts = append(iPosts, iPost)
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status": 200,
        "success": true,
        "data": iPosts,
    })
}

func CreatePost(c *fiber.Ctx) error {
	var post models.Post

	if err := c.BodyParser(&post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": 400,
			"success": false,
			"error": "Invalid request payload"})
	}
	
	post.ID = uuid.New()
	post.CreatedAt = time.Now()
	post.UpdatedAt = post.CreatedAt

	collection := config.MongoDatabase.Collection("posts")
	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": 400,
			"success": false,
			"error": "Failed to create post"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": 201,
		"success": true,
		"data": post,
	})
}

func EditPost(c *fiber.Ctx) error {
    postID := c.Params("id")
    var updateData types.IPost

    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  400,
            "success": false,
            "error":   "Invalid request payload",
        })
    }

    collection := config.MongoDatabase.Collection("posts")
    var existingPost models.Post
    objID, err := uuid.Parse(postID)
    if err != nil {
        println(err, postID)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  400,
            "success": false,
            "error": "Invalid post ID",
        })
    }

    filter := bson.M{"_id": objID}
    if err := collection.FindOne(context.Background(), filter).Decode(&existingPost); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status":  404,
            "success": false,
            "error":  "Post not found",
        })
    }

    if updateData.Title != "" {
        existingPost.Title = updateData.Title
    }
    if updateData.Content != "" {
        existingPost.Content = updateData.Content
    }
    if updateData.Author != "" {
        existingPost.Author = updateData.Author
    }
    existingPost.UpdatedAt = time.Now()

    update := bson.M{
        "$set": bson.M{
            "title":      existingPost.Title,
            "content":    existingPost.Content,
            "author":     existingPost.Author,
            "updated_at": existingPost.UpdatedAt,
        },
    }

    _, err = collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status":  500,
            "success": false,
            "error":   "Failed to update post",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status":  200,
        "success": true,
        "data": existingPost,
    })
}

func DeletePost(c *fiber.Ctx) error {
    postID := c.Params("id")

    collection := config.MongoDatabase.Collection("posts")
    var existingPost models.Post
    objID, err := uuid.Parse(postID)
    if err != nil {
        println(err, postID)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status":  400,
            "success": false,
            "error": "Invalid post ID",
        })
    }

    filter := bson.M{"_id": objID}
    if err := collection.FindOne(context.Background(), filter).Decode(&existingPost); err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "status": 404,
            "success": false,
            "error": "Post not found",
        })
    }

    _, err = collection.DeleteOne(context.Background(), filter)
    if err != nil {
        return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
            "status": 500,
            "success": false,
            "error": "Something went wrong",
        })
    }

    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "status": 200,
        "success": true,
        "error": "Post successfully deleted",
    })
}
