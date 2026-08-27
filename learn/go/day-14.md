# Day 14: Errors, Collections & I/O — JSON Encoding and Decoding

**Project:** Study Buddy — JSON API for sessions and courses

## 1. Marshal Structs

Convert structs to JSON bytes:

```go
type StudySession struct {
	ID        string    `json:"id"`
	Subject   string    `json:"subject"`
	Minutes   int       `json:"minutes"`
	Notes     string    `json:"notes,omitempty"`
	StartedAt time.Time `json:"started_at"`
	Completed bool      `json:"completed"`
}

session := StudySession{
	ID: "s1", Subject: "Go", Minutes: 60,
	StartedAt: time.Now(), Completed: true,
}

data, err := json.Marshal(session)
// {"id":"s1","subject":"Go","minutes":60,"started_at":"...","completed":true}
// Notes omitted — empty string + omitempty tag

pretty, err := json.MarshalIndent(session, "", "  ")
```

**Struct tags** control JSON field names and omission behavior.

## 2. Unmarshal JSON

Parse JSON into structs:

```go
input := `{"id":"s2","subject":"Math","minutes":45,"completed":false}`

var session StudySession
err := json.Unmarshal([]byte(input), &session)
if err != nil {
	return fmt.Errorf("parse session JSON: %w", err)
}
fmt.Println(session.Subject) // Math
```

Unknown JSON fields are **silently ignored** by default — clients can send extra data safely.

Handle malformed JSON:

```go
if err := json.Unmarshal(data, &session); err != nil {
	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return fmt.Errorf("invalid JSON at offset %d: %w", syntaxErr.Offset, err)
	}
	return fmt.Errorf("unmarshal session: %w", err)
}
```

## 3. Use json.Encoder/Decoder

Stream JSON for HTTP-sized payloads:

```go
// Encode to writer (e.g., HTTP response)
func WriteSessions(w io.Writer, sessions []StudySession) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(sessions)
}

// Decode from reader (e.g., HTTP request body)
func ReadSession(r io.Reader) (StudySession, error) {
	var session StudySession
	dec := json.NewDecoder(r)
	if err := dec.Decode(&session); err != nil {
		return StudySession{}, fmt.Errorf("decode session: %w", err)
	}
	return session, nil
}
```

Encoders/decoders process one JSON value per call — ideal for streaming APIs.

## 4. Model Optional Fields

JSON fields that may be absent need special handling:

**Pointer fields** — `nil` means absent, non-nil means present:

```go
type Course struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Credits     *int    `json:"credits,omitempty"`
}

desc := "Intro to Go"
course := Course{
	ID: "go-101", Name: "Go Fundamentals",
	Description: &desc,
}
// credits omitted entirely from JSON
```

**omitempty tag** — omit zero values:

```go
type SessionSummary struct {
	Subject string `json:"subject"`
	Minutes int    `json:"minutes,omitempty"` // 0 omitted
	Notes   string `json:"notes,omitempty"`   // "" omitted
}
```

| Approach | When to use |
|---|---|
| `omitempty` | Zero value means "not set" is acceptable |
| Pointer field | Distinguish "not sent" from "sent as zero/null" |

## Study Buddy JSON Example

Full session export:

```json
{
  "id": "s1",
  "subject": "Go",
  "minutes": 90,
  "started_at": "2026-08-21T14:00:00Z",
  "completed": true
}
```

## Summary

JSON is Study Buddy's interchange format for configs, API payloads, and exported reports. Struct tags, optional pointers, and streaming encoders prepare for HTTP handlers in later days.
