# learn_go_right_click_copy_file
(educational purposes only) learning some Go basics, so with this executable, the user can right-click with his mouse on a file and copy the full path of the file, then when it is pasted, it will have a file-link


````markdown
# learn_go_right_click_copy_file

> A small Go utility to right‑click a file on Linux and copy its full path or `file://` link to the clipboard.

## 🎯 What this is

This project provides a lightweight command‑line tool written in Go that does the following:  
- Accepts a file path, converts it to an absolute path.  
- Encodes it as a `file://` URL (so you can paste links, not just raw paths).  
- Is designed primarily for **Linux desktop environments** (e.g., Ubuntu with GNOME/Nautilus) and integrates with the file manager right‑click menu.

## 🧩 Features

- Converts one or more selected files into `file://` URLs.  
- Works from the terminal or via a right‑click “Scripts” menu.  
- Minimal external dependencies.  
- Useful for workflows where you need to paste clickable file‑links into documents, chats, or other systems.

## 🖥️ Supported environment

- Operates on Linux (tested on Ubuntu with GNOME / Nautilus).  
- Requires `xclip` (or equivalent clipboard tool) if integrating with clipboard via script.  
- The binary itself is platform‑agnostic, but the “right‑click” script and context‑menu integration assume a Linux desktop file‑manager environment.

## 🚀 Getting Started

### 1. Build the Go executable

```bash
go build -o copy_and_filelink copy_and_filelink.go
````

### 2. Move the binary to your `PATH`

```bash
sudo mv copy_and_filelink /usr/local/bin/
sudo chmod +x /usr/local/bin/copy_and_filelink
```

### 3. Set up the Nautilus “Scripts” entry (for Ubuntu/GNOME)

```bash
mkdir -p ~/.local/share/nautilus/scripts
```

Create a file called `~/.local/share/nautilus/scripts/copy_as_filelink` with this content:

```bash
#!/bin/bash
tmpfile=$(mktemp)

for file in "$@"; do
  /usr/local/bin/copy_and_filelink "$file" "$tmpfile"
  xclip -selection clipboard "$tmpfile"
done

rm -f "$tmpfile"
```

Then make it executable:

```bash
chmod +x ~/.local/share/nautilus/scripts/copy_as_filelink
```

Restart Nautilus:

```bash
nautilus -q && nautilus &
```

Now, in your file manager: **right‑click any file → Scripts → copy_as_filelink**.
After you click, the `file://…` link should be in your clipboard.

### 4. Verify clipboard tool

Make sure you have `xclip` installed (or swap in another clipboard tool).
On Ubuntu:

```bash
sudo apt update && sudo apt install xclip
```

## 📋 Usage Examples

**From terminal:**

```bash
copy_and_filelink /home/user/Documents/example.txt /tmp/output.txt
cat /tmp/output.txt
# Output: [example.txt]file:///home/user/Documents/example.txt
```

**Via file‑manager:**

* Right‑click `example.txt` → Scripts → *copy_as_filelink*
* Paste somewhere (Ctrl+V) → you should see `file:///home/user/Documents/example.txt`

## 🛠 Customisation

* Modify the Go program to add CLI flags (e.g., choose plain text path vs `file://` link).
* Replace `xclip` with `xsel` or another clipboard utility if preferred.
* Extend the script to handle multiple files and combine output into a single clipboard entry.
* Integrate with other file‑managers (e.g., Dolphin, Thunar) by creating equivalent service‑menu or custom action entries.

## 🧠 Why I built this

As part of my learning journey with Go and Linux desktop automation, I wanted a small but useful utility: right‑click a file and instantly get a `file://` link to paste somewhere. It helped me understand:

* Working with file paths in Go
* URL‑encoding for file URLs
* Building binaries and installing in Linux
* Integrating with desktop file‑manager context menus

## 🚧 Limitations

* Currently tailored to Linux desktop environments (especially GNOME/Nautilus).
* Doesn’t (yet) support Windows or macOS context‑menu installation.
* Assumes availability of a clipboard tool (like `xclip`).
* The script version uses a temporary file; you might prefer a direct‑clipboard version.

## 📝 License

Distributed under the [MIT License](LICENSE). Feel free to fork, modify and share.

---

Thank you for checking out the project!
If you have suggestions or issues, please open an issue on GitHub.
Pull requests are welcome 😊

```

---

