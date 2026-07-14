# File Sync Tool

Simple Go application that compares two folders and determines which files should be copied during synchronization.

The project does not modify files. It only builds a synchronization plan.

---

## Preview

```
Source : Documents

Destination : Backup

Files to copy

+ invoice.pdf
+ notes.txt

Already synchronized

✓ photo.jpg
✓ music.mp3
```

---

## Project Layout

```
main.go
scanner.go
manifest.go
compare.go
syncer.go
output.go
sample_files.go
```

---

## Running

```bash
go run .
```

---

## Next Ideas

- SHA256 comparison
- Recursive folders
- File deletion
- Copy progress
- Command-line flags
