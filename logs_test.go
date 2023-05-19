package ctstream

import "testing"

func TestFindByName(t *testing.T) {

	log := FindByName("Argon2024")
	if log == nil {
		t.Fatalf("Failed to find Argon2024: nil\n")
	}

	if log.Name != "Argon2024" {
		t.Fatalf("Failed to find Argon2024: Name differs\n")
	}

	if log.URI != "https://ct.googleapis.com/logs/us1/argon2024/" {
		t.Fatalf("Failed to find Argon2024: URI differs\n")
	}

	if log.PubKey != "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHblsqctplMVc5ramA7vSuNxUQxcomQwGAVAdnWTAWUYr3MgDHQW0LagJ95lB7QT75Ve6JgT2EVLOFGU7L3YrwA==" {
		t.Fatalf("Failed to find Argon2024: Key differs\n")
	}
}

func TestFindByURI(t *testing.T) {

	log := FindByURI("https://ct.googleapis.com/logs/us1/argon2024/")
	if log == nil {
		t.Fatalf("Failed to find Argon2024: nil\n")
	}

	if log.Name != "Argon2024" {
		t.Fatalf("Failed to find Argon2024: Name differs\n")
	}

	if log.URI != "https://ct.googleapis.com/logs/us1/argon2024/" {
		t.Fatalf("Failed to find Argon2024: URI differs\n")
	}

	if log.PubKey != "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHblsqctplMVc5ramA7vSuNxUQxcomQwGAVAdnWTAWUYr3MgDHQW0LagJ95lB7QT75Ve6JgT2EVLOFGU7L3YrwA==" {
		t.Fatalf("Failed to find Argon2024: Key differs\n")
	}
}

func TestFindByKey(t *testing.T) {

	log := FindByPubKey("MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHblsqctplMVc5ramA7vSuNxUQxcomQwGAVAdnWTAWUYr3MgDHQW0LagJ95lB7QT75Ve6JgT2EVLOFGU7L3YrwA==")
	if log == nil {
		t.Fatalf("Failed to find Argon2024: nil\n")
	}

	if log.Name != "Argon2024" {
		t.Fatalf("Failed to find Argon2024: Name differs\n")
	}

	if log.URI != "https://ct.googleapis.com/logs/us1/argon2024/" {
		t.Fatalf("Failed to find Argon2024: URI differs\n")
	}

	if log.PubKey != "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHblsqctplMVc5ramA7vSuNxUQxcomQwGAVAdnWTAWUYr3MgDHQW0LagJ95lB7QT75Ve6JgT2EVLOFGU7L3YrwA==" {
		t.Fatalf("Failed to find Argon2024: Key differs\n")
	}
}
