package main

import (
	"fmt"
	"log"

	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=rizkirmdhn password=rizkirmdhn dbname=segokucing port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Seed data users
	users := []entity.Users{
		{Name: "Alice", Email: "alice@example.com", Phone: "081234567890", Password: "hashedpassword", ImageUrl: "https://example.com/alice.jpg"},
		{Name: "Bob", Email: "bob@example.com", Phone: "081234567891", Password: "hashedpassword", ImageUrl: "https://example.com/bob.jpg"},
		{Name: "Charlie", Email: "charlie@example.com", Phone: "081234567892", Password: "hashedpassword", ImageUrl: "https://example.com/charlie.jpg"},
	}

	// Insert users ke database
	for _, user := range users {
		db.FirstOrCreate(&user, entity.Users{Email: user.Email})
	}

	fmt.Println("✅ Users seeded successfully!")

	// Seed friends relationship
	friends := []entity.Friends{
		{UserID: 1, FriendID: 2}, // Alice berteman dengan Bob
		{UserID: 1, FriendID: 3}, // Alice berteman dengan Charlie
		{UserID: 2, FriendID: 3}, // Bob berteman dengan Charlie
	}

	// Insert friends ke database
	for _, friend := range friends {
		db.FirstOrCreate(&friend, entity.Friends{UserID: friend.UserID, FriendID: friend.FriendID})
	}

	fmt.Println("✅ Friends seeded successfully!")

	// Seed Posts
	posts := []entity.Posts{
		{UserID: 1, Content: "Hello, this is Alice!"},
		{UserID: 2, Content: "Hello, this is Bob!"},
		{UserID: 3, Content: "Hello, this is Charlie!"},
	}

	// Insert posts ke database
	for _, post := range posts {
		db.FirstOrCreate(&post, entity.Posts{UserID: post.UserID, Content: post.Content})
	}

	fmt.Println("✅ Posts seeded successfully!")

	// Seed Tags
	tags := []entity.Tags{
		{Tag: "greeting"},
		{Tag: "introduction"},
		{Tag: "hello"},
	}

	// Insert tags ke database
	for _, tag := range tags {
		db.FirstOrCreate(&tag, entity.Tags{Tag: tag.Tag})
	}

	fmt.Println("✅ Tags seeded successfully!")

	// Assign tags to posts
	postTags := []entity.PostTags{
		{PostID: 1, TagID: 1}, // Post Alice diberi tag greeting
		{PostID: 2, TagID: 2}, // Post Bob diberi tag introduction
		{PostID: 3, TagID: 3}, // Post Charlie diberi tag hello
	}

	// Insert postTags ke database
	for _, postTag := range postTags {
		db.FirstOrCreate(&postTag, entity.PostTags{PostID: postTag.PostID, TagID: postTag.TagID})
	}

	fmt.Println("✅ PostTags seeded successfully!")

	// Seed Comments
	comments := []entity.Comments{
		{UserID: 1, PostID: 2, Comment: "Hello Bob!"},
		{UserID: 2, PostID: 3, Comment: "Hello Charlie!"},
		{UserID: 3, PostID: 1, Comment: "Hello Alice!"},
	}

	// Insert comments ke database
	for _, comment := range comments {
		db.FirstOrCreate(&comment, entity.Comments{UserID: comment.UserID, PostID: comment.PostID, Comment: comment.Comment})
	}

	fmt.Println("✅ Comments seeded successfully!")

}
