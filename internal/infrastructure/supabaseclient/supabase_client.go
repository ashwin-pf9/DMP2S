// Create Connection To the SUPABASE project
package supabaseclient

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
)

// func InitSupabaseClient() *supabase.Client {
// 	// godotenv.Load() function loads Environment Variables into processe's address space
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatalf("Error loading .env file: %v", err)
// 	}

// 	// Supabase project details
// 	SUPABASE_URL := os.Getenv("SUPABASE_URL")
// 	ANON_KEY := os.Getenv("SUPABASE_KEY")

// 	// Create Supabase client
// 	return supabase.CreateClient(SUPABASE_URL, ANON_KEY)

// }
func InitSupabaseClient() *supabase.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("ANON_KEY")

	return supabase.CreateClient(url, key)
}
