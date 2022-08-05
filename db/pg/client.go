package pg

import (
	"errors"
	"os"

	postgrest_go "github.com/nedpals/postgrest-go/pkg"
	"github.com/nedpals/supabase-go"
)

func CreateSupabaseClient() (*postgrest_go.Client, error) {
	endpoint := os.Getenv("SUPABASE_ENDPOINT")
	if endpoint == "" {
		return nil, errors.New(`Could not create supabase client. No environmental value exists for "SUPABASE_ENDPOINT"`)
	}

	anonKey := os.Getenv("SUPABASE_ANON_KEY")
	if anonKey == "" {
		return nil, errors.New(`Could not create supabase client. No environmental value exists for "SUPABASE_ANON_KEY"`)
	}

	return supabase.CreateClient(endpoint, anonKey).DB, nil
}
