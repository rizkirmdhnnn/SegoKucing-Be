package main

import (
	"fmt"

	"github.com/rizkirmdhnnn/segokucing-be/internal/config"
	"github.com/rizkirmdhnnn/segokucing-be/internal/entity"
)

func main() {
	// Load configuration
	cfg := config.NewViper()

	// Initialize logger
	logger := config.NewLogger(cfg)

	// Connect to database
	db := config.NewDatabase(cfg, logger)

	// Seed data users
	users := []entity.Users{
		{Name: "rizkirmdn", Email: "me@rizkirmdhn.cloud", Password: "$2a$10$Fggci1zPO/kgURXSBZZpkuJBL7vHyF1JKeuH/TqYEMASOgv1GiE9e"},
		{Name: "Alice", Email: "alice@example.com", Phone: "081234567890", Password: "hashedpassword", ImageUrl: "https://example.com/alice.jpg"},
		{Name: "Bob", Email: "bob@example.com", Phone: "081234567891", Password: "hashedpassword", ImageUrl: "https://example.com/bob.jpg"},
		{Name: "Charlie", Email: "charlie@example.com", Phone: "081234567892", Password: "hashedpassword", ImageUrl: "https://example.com/charlie.jpg"},
		{Name: "David", Email: "david@example.com", Phone: "081234567893", Password: "hashedpassword", ImageUrl: "https://example.com/david.jpg"},
		{Name: "Eve", Email: "eve@example.com", Phone: "081234567894", Password: "hashedpassword", ImageUrl: "https://example.com/eve.jpg"},
		{Name: "Frank", Email: "frank@example.com", Phone: "081234567895", Password: "hashedpassword", ImageUrl: "https://example.com/frank.jpg"},
		{Name: "Grace", Email: "grace@example.com", Phone: "081234567896", Password: "hashedpassword", ImageUrl: "https://example.com/grace.jpg"},
		{Name: "Hank", Email: "hank@example.com", Phone: "081234567897", Password: "hashedpassword", ImageUrl: "https://example.com/hank.jpg"},
		{Name: "Ivy", Email: "ivy@example.com", Phone: "081234567898", Password: "hashedpassword", ImageUrl: "https://example.com/ivy.jpg"},
	}

	for _, user := range users {
		db.FirstOrCreate(&user, entity.Users{Email: user.Email})
	}
	fmt.Println("✅ Users seeded successfully!")

	// Seed friends relationship
	friends := []entity.Friends{
		{UserID: 1, FriendID: 2}, {UserID: 1, FriendID: 3}, {UserID: 1, FriendID: 4}, {UserID: 1, FriendID: 5}, {UserID: 1, FriendID: 6},
		{UserID: 1, FriendID: 7}, {UserID: 1, FriendID: 8}, {UserID: 1, FriendID: 9}, {UserID: 1, FriendID: 10},
		{UserID: 2, FriendID: 3}, {UserID: 2, FriendID: 4}, {UserID: 2, FriendID: 5}, {UserID: 2, FriendID: 6}, {UserID: 2, FriendID: 7},
		{UserID: 3, FriendID: 4}, {UserID: 3, FriendID: 5}, {UserID: 3, FriendID: 6}, {UserID: 3, FriendID: 7}, {UserID: 3, FriendID: 8},
		{UserID: 4, FriendID: 5}, {UserID: 4, FriendID: 6}, {UserID: 4, FriendID: 7}, {UserID: 4, FriendID: 8}, {UserID: 4, FriendID: 9},
		{UserID: 5, FriendID: 6}, {UserID: 5, FriendID: 7}, {UserID: 5, FriendID: 8}, {UserID: 5, FriendID: 9}, {UserID: 5, FriendID: 10},
		{UserID: 6, FriendID: 7}, {UserID: 6, FriendID: 8}, {UserID: 6, FriendID: 9}, {UserID: 6, FriendID: 10}, {UserID: 6, FriendID: 1},
		{UserID: 7, FriendID: 8}, {UserID: 7, FriendID: 9}, {UserID: 7, FriendID: 10}, {UserID: 7, FriendID: 1}, {UserID: 7, FriendID: 2},
		{UserID: 8, FriendID: 9}, {UserID: 8, FriendID: 10}, {UserID: 8, FriendID: 1}, {UserID: 8, FriendID: 2}, {UserID: 8, FriendID: 3},
		{UserID: 9, FriendID: 10}, {UserID: 9, FriendID: 1}, {UserID: 9, FriendID: 2}, {UserID: 9, FriendID: 3}, {UserID: 9, FriendID: 4},
		{UserID: 10, FriendID: 1}, {UserID: 10, FriendID: 2}, {UserID: 10, FriendID: 3}, {UserID: 10, FriendID: 4}, {UserID: 10, FriendID: 5},
	}

	for _, friend := range friends {
		db.FirstOrCreate(&friend, entity.Friends{UserID: friend.UserID, FriendID: friend.FriendID})
	}
	fmt.Println("✅ Friends seeded successfully!")

	// Seed Posts
	posts := []entity.Posts{
		{UserID: 1, Content: "Hello, I'm learning Golang!"},
		{UserID: 1, Content: "I love coding!"},
		{UserID: 1, Content: "Just finished a new project!"},
		{UserID: 1, Content: "Excited about the new tech trends!"},
		{UserID: 1, Content: "Reading a great book on software architecture!"},
		{UserID: 2, Content: "Good morning, world!"},
		{UserID: 2, Content: "Enjoying the weekend!"},
		{UserID: 2, Content: "Working on a new feature!"},
		{UserID: 2, Content: "Learning something new every day!"},
		{UserID: 2, Content: "Thinking about the future of AI!"},
		{UserID: 3, Content: "Excited about new projects!"},
		{UserID: 3, Content: "Just finished a great book!"},
		{UserID: 3, Content: "Trying a new recipe today!"},
		{UserID: 3, Content: "Loving the new tech trends!"},
		{UserID: 3, Content: "Planning a trip to the mountains!"},
		{UserID: 4, Content: "Just finished reading a great book!"},
		{UserID: 4, Content: "Working on a new project!"},
		{UserID: 4, Content: "Learning new things every day!"},
		{UserID: 4, Content: "Thinking about the future!"},
		{UserID: 4, Content: "Excited about the weekend!"},
		{UserID: 5, Content: "Trying a new recipe today!"},
		{UserID: 5, Content: "Working on a new feature!"},
		{UserID: 5, Content: "Learning something new every day!"},
		{UserID: 5, Content: "Thinking about the future of AI!"},
		{UserID: 5, Content: "Excited about the new tech trends!"},
		{UserID: 6, Content: "Loving the new tech trends!"},
		{UserID: 6, Content: "Working on a new project!"},
		{UserID: 6, Content: "Learning new things every day!"},
		{UserID: 6, Content: "Thinking about the future!"},
		{UserID: 6, Content: "Excited about the weekend!"},
		{UserID: 7, Content: "Planning a trip to the mountains!"},
		{UserID: 7, Content: "Working on a new feature!"},
		{UserID: 7, Content: "Learning something new every day!"},
		{UserID: 7, Content: "Thinking about the future of AI!"},
		{UserID: 7, Content: "Excited about the new tech trends!"},
		{UserID: 8, Content: "Just finished a great book!"},
		{UserID: 8, Content: "Working on a new project!"},
		{UserID: 8, Content: "Learning new things every day!"},
		{UserID: 8, Content: "Thinking about the future!"},
		{UserID: 8, Content: "Excited about the weekend!"},
		{UserID: 9, Content: "Trying a new recipe today!"},
		{UserID: 9, Content: "Working on a new feature!"},
		{UserID: 9, Content: "Learning something new every day!"},
		{UserID: 9, Content: "Thinking about the future of AI!"},
		{UserID: 9, Content: "Excited about the new tech trends!"},
		{UserID: 10, Content: "Loving the new tech trends!"},
		{UserID: 10, Content: "Working on a new project!"},
		{UserID: 10, Content: "Learning new things every day!"},
		{UserID: 10, Content: "Thinking about the future!"},
		{UserID: 10, Content: "Excited about the weekend!"},
	}

	for _, post := range posts {
		db.FirstOrCreate(&post, entity.Posts{UserID: post.UserID, Content: post.Content})
	}
	fmt.Println("✅ Posts seeded successfully!")

	// Seed Tags
	tags := []entity.Tags{
		{Tag: "learning"}, {Tag: "morning"}, {Tag: "projects"}, {Tag: "books"}, {Tag: "food"}, {Tag: "tech"},
		{Tag: "weekend"}, {Tag: "future"}, {Tag: "AI"}, {Tag: "mountains"}, {Tag: "architecture"},
	}

	for _, tag := range tags {
		db.FirstOrCreate(&tag, entity.Tags{Tag: tag.Tag})
	}
	fmt.Println("✅ Tags seeded successfully!")

	// Assign tags to posts
	postTags := []entity.PostTags{
		{PostID: 1, TagID: 1}, {PostID: 2, TagID: 1}, {PostID: 3, TagID: 3}, {PostID: 4, TagID: 6}, {PostID: 5, TagID: 11},
		{PostID: 6, TagID: 2}, {PostID: 7, TagID: 7}, {PostID: 8, TagID: 3}, {PostID: 9, TagID: 1}, {PostID: 10, TagID: 9},
		{PostID: 11, TagID: 3}, {PostID: 12, TagID: 4}, {PostID: 13, TagID: 5}, {PostID: 14, TagID: 6}, {PostID: 15, TagID: 10},
		{PostID: 16, TagID: 4}, {PostID: 17, TagID: 3}, {PostID: 18, TagID: 1}, {PostID: 19, TagID: 8}, {PostID: 20, TagID: 7},
		{PostID: 21, TagID: 5}, {PostID: 22, TagID: 3}, {PostID: 23, TagID: 1}, {PostID: 24, TagID: 9}, {PostID: 25, TagID: 6},
		{PostID: 26, TagID: 6}, {PostID: 27, TagID: 3}, {PostID: 28, TagID: 1}, {PostID: 29, TagID: 8}, {PostID: 30, TagID: 7},
		{PostID: 31, TagID: 10}, {PostID: 32, TagID: 3}, {PostID: 33, TagID: 1}, {PostID: 34, TagID: 9}, {PostID: 35, TagID: 6},
		{PostID: 36, TagID: 4}, {PostID: 37, TagID: 3}, {PostID: 38, TagID: 1}, {PostID: 39, TagID: 8}, {PostID: 40, TagID: 7},
		{PostID: 41, TagID: 5}, {PostID: 42, TagID: 3}, {PostID: 43, TagID: 1}, {PostID: 44, TagID: 9}, {PostID: 45, TagID: 6},
		{PostID: 46, TagID: 6}, {PostID: 47, TagID: 3}, {PostID: 48, TagID: 1}, {PostID: 49, TagID: 8}, {PostID: 50, TagID: 7},
	}

	for _, postTag := range postTags {
		db.FirstOrCreate(&postTag, entity.PostTags{PostID: postTag.PostID, TagID: postTag.TagID})
	}
	fmt.Println("✅ PostTags seeded successfully!")

	// Seed Comments
	comments := []entity.Comments{
		{UserID: 2, PostID: 1, Comment: "That's awesome!"},
		{UserID: 3, PostID: 1, Comment: "Keep it up!"},
		{UserID: 4, PostID: 1, Comment: "Great job!"},
		{UserID: 5, PostID: 1, Comment: "I'm also learning Golang!"},
		{UserID: 6, PostID: 2, Comment: "I love coding too!"},
		{UserID: 7, PostID: 2, Comment: "What are you working on?"},
		{UserID: 8, PostID: 2, Comment: "Keep it up!"},
		{UserID: 9, PostID: 2, Comment: "Great job!"},
		{UserID: 10, PostID: 3, Comment: "What project did you finish?"},
		{UserID: 1, PostID: 3, Comment: "Tell us more about it!"},
		{UserID: 2, PostID: 3, Comment: "I'm excited to see it!"},
		{UserID: 3, PostID: 3, Comment: "Great job!"},
		{UserID: 4, PostID: 4, Comment: "What tech trends are you excited about?"},
		{UserID: 5, PostID: 4, Comment: "I'm also excited about AI!"},
		{UserID: 6, PostID: 4, Comment: "Great job!"},
		{UserID: 7, PostID: 4, Comment: "Keep it up!"},
		{UserID: 8, PostID: 5, Comment: "What book are you reading?"},
		{UserID: 9, PostID: 5, Comment: "I love reading too!"},
		{UserID: 10, PostID: 5, Comment: "Great job!"},
		{UserID: 1, PostID: 5, Comment: "Keep it up!"},
		{UserID: 2, PostID: 6, Comment: "Good morning to you too!"},
		{UserID: 3, PostID: 6, Comment: "Have a great day!"},
		{UserID: 4, PostID: 6, Comment: "Great job!"},
		{UserID: 5, PostID: 6, Comment: "Keep it up!"},
		{UserID: 6, PostID: 7, Comment: "Enjoy your weekend!"},
		{UserID: 7, PostID: 7, Comment: "What are you planning to do?"},
		{UserID: 8, PostID: 7, Comment: "Great job!"},
		{UserID: 9, PostID: 7, Comment: "Keep it up!"},
		{UserID: 10, PostID: 8, Comment: "What feature are you working on?"},
		{UserID: 1, PostID: 8, Comment: "I'm excited to see it!"},
		{UserID: 2, PostID: 8, Comment: "Great job!"},
		{UserID: 3, PostID: 8, Comment: "Keep it up!"},
		{UserID: 4, PostID: 9, Comment: "What are you learning today?"},
		{UserID: 5, PostID: 9, Comment: "I'm also learning new things!"},
		{UserID: 6, PostID: 9, Comment: "Great job!"},
		{UserID: 7, PostID: 9, Comment: "Keep it up!"},
		{UserID: 8, PostID: 10, Comment: "What do you think about the future?"},
		{UserID: 9, PostID: 10, Comment: "I'm also thinking about it!"},
		{UserID: 10, PostID: 10, Comment: "Great job!"},
		{UserID: 1, PostID: 10, Comment: "Keep it up!"},
	}

	for _, comment := range comments {
		db.FirstOrCreate(&comment, entity.Comments{UserID: comment.UserID, PostID: comment.PostID, Comment: comment.Comment})
	}
	fmt.Println("✅ Comments seeded successfully!")
}
