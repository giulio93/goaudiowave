
# 🎵 Go Audio Waveform Generator

Generate a waveform image from an audio file using the `goaudiowave` tool.

> This project requires **Go** and **FFmpeg** to be installed on your system.

---

## 📋 Table of Contents

- [🎵 Go Audio Waveform Generator](#-go-audio-waveform-generator)
  - [📋 Table of Contents](#-table-of-contents)
  - [✅ Prerequisites](#-prerequisites)
  - [⚙️ Installation](#️-installation)
    - [FFmpeg Installation](#ffmpeg-installation)
      - [macOS](#macos)
      - [Windows](#windows)
      - [Linux](#linux)
    - [Go Installation](#go-installation)
      - [macOS](#macos-1)
      - [Windows](#windows-1)
      - [Linux](#linux-1)
  - [🚀 Usage Instructions](#-usage-instructions)
    - [Step 1: Convert Your Audio File](#step-1-convert-your-audio-file)
    - [Step 2: Clone the Repository](#step-2-clone-the-repository)
    - [Step 3: Run the Go Program](#step-3-run-the-go-program)
  - [🧪 Example](#-example)
  - [🪪 License](#-license)

---

## ✅ Prerequisites

Make sure the following are installed:

- [Go](https://golang.org/dl/): Used to run the waveform generator.
- [FFmpeg](https://ffmpeg.org/): Used to convert audio files into the required format.

---

## ⚙️ Installation

### FFmpeg Installation

#### macOS

```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
brew install ffmpeg
```

#### Windows

1. Visit: [FFmpeg for Windows](https://ffmpeg.org/download.html#build-windows)
2. Download the latest **GPL ZIP** release.
3. Extract the ZIP (e.g., to `C:\ffmpeg`).
4. Add `C:\ffmpeg\bin` to your system's **PATH** environment variable.

#### Linux

**Debian/Ubuntu:**

```sh
sudo apt update
sudo apt install ffmpeg
```

**Fedora/CentOS/RHEL:**

```sh
sudo dnf install ffmpeg
# or
sudo yum install ffmpeg
```

---

### Go Installation

#### macOS

```sh
brew install go
```

#### Windows

1. Download the MSI installer from [https://golang.org/dl/](https://golang.org/dl/)
2. Run the installer and follow the instructions.

#### Linux

**Debian/Ubuntu:**

```sh
sudo apt update
sudo apt install golang-go
```

**Fedora/CentOS/RHEL:**

```sh
sudo dnf install golang
# or
sudo yum install golang
```

---

## 🚀 Usage Instructions

### Step 1: Convert Your Audio File

Use FFmpeg to convert your audio file into the correct WAV format:

```sh
ffmpeg -i input.mp3 -acodec pcm_s16le -ac 1 -ar 44100 soundwave.wav
```

Explanation:
- `-i input.mp3`: Your input file
- `-acodec pcm_s16le`: Use 16-bit PCM codec
- `-ac 1`: Convert to mono
- `-ar 44100`: Set sample rate to 44.1kHz
- `soundwave.wav`: Output file name

---

### Step 2: Clone the Repository

```sh
git clone https://github.com/giulio93/goaudiowave.git
cd goaudiowave
```

---

### Step 3: Run the Go Program

```sh
go run main.go -file=../soundwave.wav
```

Or if `soundwave.wav` is in the same directory:

```sh
go run main.go -file=soundwave.wav
```

---

## 🧪 Example

```sh
# Create a working directory
mkdir my-waveform-project
cd my-waveform-project

# Add your audio file (e.g., input.mp3)

# Convert it to a suitable WAV format
ffmpeg -i input.mp3 -acodec pcm_s16le -ac 1 -ar 44100 soundwave.wav

# Clone the waveform tool
git clone https://github.com/giulio93/goaudiowave.git

# Run the generator
go run goaudiowave/main.go -file=soundwave.wav
```

After running the program, a `waveform.png` file will be created in your directory.  
This image represents the audio waveform.

---

## 🪪 License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
