# 📢 echo-wise

A minimalist, high-efficiency terminal-based quote management system. Designed for thinkers, writers, and developers who want to capture and organize wisdom without leaving the command line.

Built with **Go** for speed and a lightweight footprint.

---

## ✨ Features

- **Keyboard-Centric**: Navigate, edit, and manage quotes entirely via shortcuts.
- **Real-Time Filtering**: Use the `/` operator to instantly find quotes by keyword or author.
- **Form-Based Entry**: A structured "Add" menu with intuitive field jumping.
- **Dynamic Reloading**: Refresh your collection on the fly without restarting the app.
- **Visual Feedback**: Clear cursor tracking to see exactly which quote you are managing.

### 🚀 Real-time Data Sync

Check out how the sync works in the demo below:

https://github.com/user-attachments/assets/d60d640c-ad06-4678-acad-7559ae1910ec


---

## 🏗️ Architecture & Data

echo-wise follows a t1 (Single-Tier) DBMS application architecture.

In this architecture, the User Interface (TUI), the application logic (Go), and the Database (SQLite3) all reside on a single local machine. This ensures:

- Zero Latency: Local file I/O means instant quote retrieval.
- Offline First: No internet connection or external server required.
- Data Portability: Your entire database is a single file, making it easy to sync

## Technical Stack

- Language: Go 1.26.0+
- Database Engine: SQLite3
- Architecture: t1 (Single-Tier)

## ⌨️ Shortcuts & Controls

### 🏠 Main Menu

The central hub for navigating the application.

| Key | Action                    |
| :-- | :------------------------ |
| `a` | **Add** a new quote       |
| `l` | Open **List** view        |
| `r` | **Reload** quote database |
| `q` | **Quit** application      |

### 📜 List View (Management)

The cursor indicates the currently selected quote for all actions.

| Key        | Action                        |
| :--------- | :---------------------------- |
| `j` / `k`  | Move cursor **Down** / **Up** |
| `/`        | **Filter** quotes by text     |
| `e`        | **Edit** selected quote       |
| `Ctrl + d` | **Delete** selected quote     |
| `?`        | Toggle **More Keys** menu     |
| `Esc`      | Back to **Main Menu**         |

### ➕ Add/Edit Menu

Use these controls when filling out quote details.

| Key           | Action                     |
| :------------ | :------------------------- |
| `Tab`         | Move to **Next** field     |
| `Shift + Tab` | Move to **Previous** field |
| `Enter`       | **Submit** / Continue      |
| `Esc`         | **Cancel** and go back     |

---

## 🚀 Installation

### Prerequisites

- [Go](https://go.dev/doc/install) (1.26.0 or higher
- GCC (Required for sqlite3 driver compilation)

### Build from Source

1. **Clone the repository:**

   ```bash
   git clone https://github.com/DevSatyamCollab/echo-wise.git
   cd echo-wise
   ```

2. Build the binary

   ```bash
   go build -o echo-wise .
   ```

3. (Optional) Move to your path:
   ```bash
   sudo mv echo-wise /usr/local/bin/
   ```
