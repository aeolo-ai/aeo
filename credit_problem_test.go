package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOutOfCreditsKeepsTheProblemBody(t *testing.T) {
	body := `{"type":"https://docs.aeolo.io/errors#not-enough-credits","title":"Not enough credits","status":402,"code":"INSUFFICIENT_CREDITS","message":"Not enough production credits. You need 3 credits and have 0.","required_credits":3,"available_credits":0}`
	var accept string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accept = r.Header.Get("Accept")
		w.Header().Set("Content-Type", "application/problem+json")
		w.WriteHeader(http.StatusPaymentRequired)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()

	_, err := callAuthenticatedJSON(server.URL, "POST", []byte(`{}`), "key")
	var credits *insufficientCreditsError
	if !errors.As(err, &credits) {
		t.Fatalf("want insufficientCreditsError, got %v", err)
	}
	if string(credits.body) != body || credits.Error() != "Not enough production credits. You need 3 credits and have 0." {
		t.Fatalf("lost the problem: %+v", credits)
	}
	if accept == "" {
		t.Fatal("client must advertise problem+json")
	}
}
