package main

import (
	"context"
	"ent/ent"
	"ent/service"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Kết nối MySQL
	client, err := ent.Open("mysql", "user:password@tcp(localhost:3306)/dbname?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	// Tự động migrate schema (chỉ dùng cho dev)
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Ví dụ tạo Translation
	choices := []map[string]interface{}{
		{
			"confidence": 0.9,
			"content":    "Xin chào",
			"extras":     map[string]interface{}{"note": "greeting"},
		},
	}
	tr, err := service.CreateTranslation(ctx, client, "Hello", "ext123", "en-vi", "owner1", choices)
	if err != nil {
		log.Fatalf("create translation failed: %v", err)
	}
	log.Printf("Created translation: %+v", tr)

	// Ví dụ tạo TranscriptionStats
	stats, err := service.CreateTranscriptionStats(ctx, client, 5000, 100, 50, "gpt-4", "owner1", 0.05, tr.ID)
	if err != nil {
		log.Fatalf("create transcription stats failed: %v", err)
	}
	log.Printf("Created transcription stats: %+v", stats)

	// Ví dụ tạo LanguageDetection
	detectedLanguages := map[string]interface{}{
		"en": 0.95,
		"vi": 0.05,
	}
	ld, err := service.CreateLanguageDetection(ctx, client, "Hello world", detectedLanguages, 2000, "gpt-4", 10, 5, 2, "owner1", 0.02)
	if err != nil {
		log.Fatalf("create language detection failed: %v", err)
	}
	log.Printf("Created language detection: %+v", ld)

	// Ví dụ đọc Translation theo ID
	readTr, err := service.GetTranslationByID(ctx, client, tr.ID)
	if err != nil {
		log.Fatalf("get translation failed: %v", err)
	}
	log.Printf("Read translation: %+v", readTr)

	log.Println("All operations completed successfully!")
}
