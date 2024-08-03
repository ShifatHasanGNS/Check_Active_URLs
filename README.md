# Check Active URLs
It can check a list of http/https protocol based URLs and from them generates a new list of active-URLs.

**You can make your own URL-List, for example - `your_urls.csv`. For Reference and Syntax - Check `urls.csv`.**

The `your_urls.csv` is the Input-File for the program and while running it, you should provide a name for the Output-File, for example - `your-active-urls.txt`.

#

## Installation Process 

1. Go [Here](https://github.com/ShifatHasanGNS/Check_Active_URLs/releases/tag/v1.0.0) and Download as per your Operating System.
2. Download `urls.csv` as your Input-File URL-List or make your own CSV File. But, make sure to maintain the exact same format as the `urls.csv`.
3. Open your Terminal in the same directory where the downloaded executable-file located.
4. Run-Command-Format: `<Program> <Input> <Output>`

**Examples:**

- Linux
```bash
./check-active-urls-linux-amd64 urls.csv active-urls.txt
```
- Mac
```bash
./check-active-urls-darwin-arm64 urls.csv active-urls.txt
```
- Windows
```bash
.\check-active-urls-windows-amd64 urls.csv active-urls.txt
```
#


### How to use without downloading the executable-file?
0. Make sure you have `Go` installed in your System. If not, then install it from [Here](https://golang.org/doc/install).
1. Open your Terminal in a Folder where you want to Keep the Repository.
2. Clone the Repository:
```bash
git clone https://github.com/ShifatHasanGNS/Check_Active_URLs.git
```
2. Go inside the Repository.
3. Check the `Makefile` where you will find command for the Built-Process.
4. I would suggest to Cross-Compile first and then use the one which is suitable for yout operating system.
5. Command for Cross-Compilation:
```bash
make cross-compile
```
6. The `bin` folder will be generated where you will find the Executable-Files.
7. Now You are Good to Go.

# 