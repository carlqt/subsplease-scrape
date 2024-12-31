package organizer

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// func watchDir(dir string) {
// 	dirEntries, err := os.ReadDir(dir)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, f := range dirEntries {
// 		extName := filepath.Ext(f.Name())
// 		if f.IsDir() || !slices.Contains(videoTypes, extName) {
// 			continue
// 		}

// 		fileEntry := vidFile{
// 			Path: path.Join(dir, f.Name()),
// 		}

// 		destFolder := path.Join(dir, fileEntry.Title())

// 		fmt.Printf("Moving %s to %s\n", fileEntry.Path, destFolder)
// 		err = moveDir(fileEntry.Path, destFolder)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

func moveDir(src string, dest string) error {
	dest = fmt.Sprintf("%s/%s", dest, filepath.Base(src))
	err := os.Rename(src, dest)
	if err != nil {
		return fmt.Errorf("failed to move %s to %s: %w\n", src, dest, err)
	}

	return nil
}

func dirExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func Run(episodeName string, sourceDir string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}

	if sourceDir == "" {
		sourceDir = path.Join(homeDir, "Documents", "Videos")
	}

	vid := newVidFile(path.Join(sourceDir, episodeName))

	// find or create folder
	log.Println("find or create folder")
	folder := findOrCreateFolder(vid.Title(), sourceDir)

	// Move all matched to the folder
	dirEntries, _ := os.ReadDir(sourceDir)

	log.Println("moving all matched files to folder", "folderName", folder)
	for _, f := range dirEntries {
		if !f.IsDir() {
			distance := fuzzySearch(episodeName, f.Name())

			if distance <= 25 {
				fileSource := path.Join(sourceDir, f.Name())
				moveDir(fileSource, folder)
			}
		}
	}
}

func fuzzySearch(source string, target string) int {
	return levenshtein.DistanceForStrings([]rune(source), []rune(target), levenshtein.DefaultOptions)
}

func findOrCreateFolder(name string, sourceDir string) string {
	if !dirExists(sourceDir) {
		msg := fmt.Sprintf("Folder %s doesn't exist", sourceDir)
		panic(msg)
	}

	dirEntries, err := os.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	for _, f := range dirEntries {
		if f.IsDir() && f.Name() == name {
			return path.Join(sourceDir, name)
		}
	}

	folderName := path.Join(sourceDir, name)
	err = os.Mkdir(folderName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	return folderName
}
