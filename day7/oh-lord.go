package day7

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func ohLord(r *bufio.Reader) {
  folders = []folder{}
  files = []file{}

	fOuter := folder{name: "/", path: "/"}

	folders = append(folders, fOuter)

	currentFolder := fOuter

	// Create a line buffer for lines read but not worked on yet
	bufferedLine := ""
reader:
	for {
		line := ""

		if bufferedLine == "" {
			b, _, err := r.ReadLine()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}

			line = string(b)
		} else {
			line = bufferedLine
			bufferedLine = ""
		}

		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd ..") {
				idx := slices.IndexFunc(folders, func(f folder) bool {
					upOnePath := removeLastSection.ReplaceAllString(currentFolder.path, "")
					if upOnePath == "" {
						upOnePath = "/"
					}

					return f.path == upOnePath
				})

				currentFolder = folders[idx]
			} else if strings.HasPrefix(line, "$ cd") {
				folderName := strings.Replace(line, "$ cd ", "", 1)
				folderPath := ""
				if currentFolder.path != "/" {
					folderPath = fmt.Sprint(currentFolder.path, "/", folderName)
				} else if folderName == "/" {
					folderPath = "/"
				} else {
					folderPath = fmt.Sprint("/", folderName)
				}

				idx := slices.IndexFunc(folders, func(f folder) bool {
					return f.path == folderPath
				})

				if idx >= 0 {
					currentFolder = folders[idx]
				} else {
					f := folder{
						name:         folderName,
						parentFolder: currentFolder.name,
						path:         folderPath,
					}
					// log.Printf("folder: %+v\n", f)
					folders = append(folders, f)
					currentFolder = f
				}
			} else if strings.HasPrefix(line, "$ ls") {
				// read lines until we hit a $
				for {
					b, _, err := r.ReadLine()
					if err == io.EOF {
						break
					}
					if err != nil {
						break
					}

					line = string(b)
					if strings.HasPrefix(line, "$") {
						bufferedLine = line
						continue reader
					}

					if strings.HasPrefix(line, "dir ") {
						folderName := strings.Replace(line, "dir ", "", 1)
						folderPath := ""
						if currentFolder.path != "/" {
							folderPath = fmt.Sprint(currentFolder.path, "/", folderName)
						} else if folderName == "/" {
							folderPath = "/"
						} else {
							folderPath = fmt.Sprint("/", folderName)
						}

						idx := slices.IndexFunc(folders, func(f folder) bool {
							return f.path == folderPath
						})

						if idx < 0 {
							f := folder{
								name:         folderName,
								parentFolder: currentFolder.name,
								path:         folderPath,
							}
							// log.Printf("folder: %+v\n", f)
							folders = append(folders, f)
						}
					} else {
						parts := strings.SplitN(line, " ", 2)
						size, err := strconv.Atoi(parts[0])
						if err != nil {
							log.Fatal("Invalid strconv value: ", parts[0])
						}
						name := parts[1]

						f := file{
							name:       name,
							size:       size,
							folderPath: currentFolder.path,
						}

						files = append(files, f)
					}
				}
			}
		}
	}
}
