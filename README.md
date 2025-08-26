# SuperGo Downloader
An ultra-fast, multithreaded command-line tool designed for efficient and reliable file downloads.

It accelerates the process by splitting files into manageable chunks and downloading them concurrently, ensuring optimal speed and the ability to resume interrupted transfers.

This will either max out your storage or your network :satisfied:

## Key Features
- :white_check_mark: Multithreaded Downloads: Concurrently fetches multiple parts of a file, drastically reducing download times.

- :white_check_mark: Chunk-based Downloading: Divides files into smaller, manageable chunks for efficient handling and reconstructs after all chunks have been downloaded.

- :white_check_mark: Cross-Platform: Built with Go, it works seamlessly on Windows, macOS, and Linux.

## To Do:

- :pushpin: Resume Capability: Automatically detects and continues interrupted downloads from the last successful chunk, saving time and bandwidth.

- :pushpin: Real-time Progress: A clean, live progress bar in your terminal shows download status, speed, and estimated time remaining.


## Installation
You can get started with the SuperGo downloader using a few simple methods.

From a Release Binary
Grab the latest pre-compiled binary from the Releases page. This is the fastest way to get started.

From Source (using Go)
If you have Go installed on your machine, you can build the tool yourself:

    git clone https://github.com/ahhossain/SuperGo.git
    cd SuperGo
    go build -o SuperGo.exe .\cmd\SuperGo\main.go

## Usage
The CLI is designed to be straightforward and intuitive.

Basic Download
To download a file, simply provide the URL and the path to directory where you want to save it:

    SuperGo.exe --url "https://releases.ubuntu.com/25.04/ubuntu-25.04-desktop-amd64.iso" --path "C:\temp\"


## Contributing
I welcome contributions from everyone! If you find a bug, have a feature idea, or want to improve the code, please feel free to open an issue or a pull request
