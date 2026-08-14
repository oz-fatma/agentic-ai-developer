# Day 1: Linux Fundamentals — Basic Navigation and File Management

**Environment:** macOS terminal (Unix/Linux-compatible) · zsh

## 1. Connect to a Linux Environment

macOS terminal works for Day 1 — same core commands as Linux.

```bash
echo $SHELL    # /bin/zsh
uname -a         # Darwin (macOS) — Unix-based, same CLI commands
```

**Result:** Shell is `zsh`. macOS is Unix-based — `ls`, `cd`, `mkdir`, etc. work the same as on Linux servers.

## 2. Navigate the File System

```bash
pwd                          # where am I?
ls                           # list files
ls -la                       # list with details (hidden files too)
cd ~                         # go to home directory
cd ..                        # go up one level
pwd                          # confirm new location
```

**Practice log:**

| Command | Result |
|---|---|
| `pwd` | `/Users/fatmaoz/developer/agentic-ai-developer/learn/devops` |
| `ls -la` | Shows `day-1.md`, `.`, `..` with permissions and dates |
| `cd ~/developer` | Goes to home developer folder (or parent paths) |

**Flags:**
- `-l` = long format (permissions, owner, size, date)
- `-a` = show hidden files (names starting with `.`)

## 3. Create and Remove Directories

```bash
mkdir devops-practice          # create directory
mkdir devops-practice/logs     # create nested directory
ls devops-practice             # output: logs
rmdir devops-practice/logs     # remove empty directory
rmdir devops-practice          # remove empty directory
```

**Note:** `rmdir` only removes **empty** directories. Use `rm -r dirname` for non-empty folders (careful — no undo).

## 4. Create, Copy, Move, and Delete Files

```bash
mkdir devops-practice && cd devops-practice

touch readme.txt
echo "Hello DevOps" > readme.txt

cp readme.txt readme-backup.txt   # copy → two files
mv readme-backup.txt backup.txt   # rename (same folder = rename)
ls -la
# backup.txt  readme.txt

rm backup.txt
rm readme.txt
cd .. && rmdir devops-practice
```

**Result:** All steps completed successfully. `mv` in the same folder = **rename**.

## Command Cheat Sheet

| Command | Purpose |
|---|---|
| `pwd` | Print working directory |
| `ls` / `ls -la` | List files |
| `cd` | Change directory |
| `mkdir` | Create directory |
| `rmdir` | Remove empty directory |
| `touch` | Create empty file |
| `cp` | Copy |
| `mv` | Move or rename |
| `rm` | Delete file |

## Notes

- **CLI** = text interface to the OS — DevOps lives here on servers.
- **`pwd`** before and after `cd` — always know where you are.
- **`rm` is permanent** — double-check before deleting.
- **`mv` same folder** = rename; different folder = move.
- Next (Day 2): more file commands — `cat`, `less`, `grep`.
