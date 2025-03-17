// Create Connection To the SUPABASE project
package supabaseclient

import (

	// "log"
	// "github.com/joho/godotenv"

	"os"

	"github.com/nedpals/supabase-go"
)

func InitSupabaseClient() *supabase.Client {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("ANON_KEY")

	return supabase.CreateClient(url, key)
}
