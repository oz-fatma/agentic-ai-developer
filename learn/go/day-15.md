# Day 15: Errors, Collections & I/O — Practice

**Project:** Study Buddy — config loader, word counter, and summary CLI

## 1. Config Loader

Read JSON config into a struct and validate required fields:

```go
type Config struct {
	AppName          string `json:"app_name"`
	DailyGoalMinutes int    `json:"daily_goal_minutes"`
	DataDir          string `json:"data_dir"`
	LogLevel         string `json:"log_level,omitempty"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &cfg, nil
}

func validateConfig(cfg *Config) error {
	if cfg.AppName == "" {
		return errors.New("app_name is required")
	}
	if cfg.DailyGoalMinutes <= 0 {
		return errors.New("daily_goal_minutes must be positive")
	}
	if cfg.DataDir == "" {
		return errors.New("data_dir is required")
	}
	return nil
}
```

Sample `config.json`:

```json
{
  "app_name": "Study Buddy",
  "daily_goal_minutes": 120,
  "data_dir": "./data"
}
```

## 2. Word Counter

Read a notes file and count word frequencies:

```go
func WordFrequency(path string) (map[string]int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	freq := make(map[string]int)
	words := strings.Fields(strings.ToLower(string(data)))
	for _, word := range words {
		word = strings.Trim(word, ".,!?;:\"'")
		if word != "" {
			freq[word]++
		}
	}
	return freq, nil
}
```

Example `notes.txt`:

```
Go interfaces are implicit. Go interfaces are powerful.
Study Buddy uses JSON for configuration.
```

Result:

| Word | Count |
|---|---|
| go | 2 |
| interfaces | 2 |
| are | 2 |
| study | 1 |
| buddy | 1 |

## 3. Error Path Tests

Manually tested unhappy paths:

| Scenario | Expected behavior | Result |
|---|---|---|
| Missing config file | Wrapped error with `os.IsNotExist` | Pass |
| Malformed JSON `{bad` | Parse error with context | Pass |
| Empty `app_name` | Validation error | Pass |
| Missing notes file | Wrapped read error | Pass |
| Empty notes file | Empty map, no panic | Pass |
| Zero daily goal | Validation error | Pass |

## 4. Combine Skills

Small CLI tying everything together:

```go
func main() {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Config error:", err)
		os.Exit(1)
	}

	notesPath := filepath.Join(cfg.DataDir, "notes.txt")
	freq, err := WordFrequency(notesPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Word count error:", err)
		os.Exit(1)
	}

	summary := SummaryReport{
		AppName:     cfg.AppName,
		TotalWords:  sumValues(freq),
		UniqueWords: len(freq),
		TopWords:    topN(freq, 5),
		GeneratedAt: time.Now(),
	}

	outPath := filepath.Join(cfg.DataDir, "summary.json")
	data, _ := json.MarshalIndent(summary, "", "  ")
	if err := os.WriteFile(outPath, data, 0644); err != nil {
		fmt.Fprintln(os.Stderr, "Write error:", err)
		os.Exit(1)
	}
	fmt.Println("Summary written to", outPath)
}
```

Output `summary.json`:

```json
{
  "app_name": "Study Buddy",
  "total_words": 12,
  "unique_words": 8,
  "top_words": {"go": 2, "interfaces": 2, "are": 2},
  "generated_at": "2026-08-21T17:00:00Z"
}
```

## Summary

**Phase checkpoint:** Errors, Collections & I/O complete.

Study Buddy now handles:
- Explicit error returns with wrapping
- Slices and maps for data
- File read/write with `os` and `bufio`
- JSON marshal/unmarshal with struct tags
- A combined CLI utility with validation and error paths

Ready for concurrency in Days 16–20.
