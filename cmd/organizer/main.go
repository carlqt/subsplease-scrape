package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

var videoTypes = []string{".mkv", ".mp4"}

type vidFile struct {
	Path string
}

func newVidFile(srcFile string) vidFile {
	p, err := filepath.Abs(srcFile)
	if err != nil {
		panic(err)
	}

	return vidFile{
		Path: p,
	}
}

func (v vidFile) translationGroup() string {
	r, err := regexp.Compile(`\[([^\]]*)\]`)
	if err != nil {
		return ""
	}

	return r.FindString(v.Path)
}

func (v vidFile) Title() string {
	s := strings.Replace(v.fileName(), v.translationGroup(), "", 1)
	s = strings.Replace(s, v.endSection(), "", 1)
	s = strings.TrimSpace(s)

	return s
}

func (v vidFile) endSection() string {
	r, err := regexp.Compile(`.*(-.*)`)
	if err != nil {
		return ""
	}

	matches := r.FindStringSubmatch(v.Path)

	lastIndex := len(matches) - 1
	return matches[lastIndex]
}

func (v vidFile) fileName() string {
	return filepath.Base(v.Path)
}

func watchDir(dir string) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range dirEntries {
		extName := filepath.Ext(f.Name())
		if f.IsDir() || !slices.Contains(videoTypes, extName) {
			continue
		}

		fileEntry := vidFile{
			Path: path.Join(dir, f.Name()),
		}

		destFolder := path.Join(dir, fileEntry.Title())

		fmt.Printf("Moving %s to %s\n", fileEntry.Path, destFolder)
		err = moveDir(fileEntry.Path, destFolder)
		if err != nil {
			log.Fatal(err)
		}
	}
}

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

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
	}

	srcDir := path.Join(homeDir, "Documents", "Videos")

	input := "[EMBER] Isekai Shikkaku - 12.mkv"
	vid := newVidFile(path.Join(srcDir, input))

	// find or create folder
	log.Println("find or create folder")
	folder := findOrCreateFolder(vid.Title(), srcDir)

	// Move all matched to the folder
	dirEntries, _ := os.ReadDir(srcDir)

	log.Println("moving all matched files to folder", "folderName", folder)
	for _, f := range dirEntries {
		if !f.IsDir() {
			distance := fuzzy(input, f.Name())

			if distance <= 25 {
				fileSource := path.Join(srcDir, f.Name())
				moveDir(fileSource, folder)
			}
		}
	}
}

func fuzzy(source string, target string) int {
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
