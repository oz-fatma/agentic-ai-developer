# Day 13: Errors, Collections & I/O — Reading and Writing Files

**Project:** Study Buddy — persist notes and load session logs

## 1. Read Files

**Whole file** — small config or notes:

```go
func ReadNotes(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read notes %q: %w", path, err)
	}
	return string(data), nil
}
```

**Line by line** — large session logs:

```go
func ReadSessionLog(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open log %q: %w", path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan log %q: %w", path, err)
	}
	return lines, nil
}
```

## 2. Write Files

**Whole file** — export summary:

```go
func WriteReport(path string, content []byte) error {
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		return fmt.Errorf("write report %q: %w", path, err)
	}
	return nil
}
```

**Buffered writer** — append session entries efficiently:

```go
func AppendSessionEntry(path, entry string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open log %q: %w", path, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(entry + "\n"); err != nil {
		return fmt.Errorf("write entry: %w", err)
	}
	return writer.Flush()
}
```

## 3. Work with io

Abstract over sources and destinations with interfaces:

```go
func CountLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}
```

Works with files, strings, network — anything implementing `io.Reader`:

```go
// From file
file, _ := os.Open("notes.txt")
count, _ := CountLines(file)

// From string
count, _ = CountLines(strings.NewReader("line1\nline2\n"))
```

Write utility:

```go
func WriteGreeting(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, "Hello, %s! Ready to study?\n", name)
	return err
}
```

## 4. Handle File Errors

Check for missing files specifically:

```go
func LoadNotes(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("notes file %q not found: %w", path, err)
		}
		return "", fmt.Errorf("read notes %q: %w", path, err)
	}
	return string(data), nil
}
```

| Error type | Handling |
|---|---|
| File not found | `os.IsNotExist(err)` — suggest creating file |
| Permission denied | Wrap with context, check file mode |
| Partial write | Check `Write` return value and `Flush` error |

Example session log format written by Study Buddy:

```
2026-08-21 | Go | 60 min | completed
2026-08-21 | Math | 45 min | completed
```

## Summary

| Operation | Package/API | Study Buddy use |
|---|---|---|
| Read whole file | `os.ReadFile` | Load notes |
| Read lines | `bufio.Scanner` | Parse session log |
| Write file | `os.WriteFile` | Export weekly report |
| Abstract I/O | `io.Reader` / `io.Writer` | Testable utilities |

File I/O connects Study Buddy to the filesystem — notes, logs, and exported reports.
