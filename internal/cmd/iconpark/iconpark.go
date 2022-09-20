package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type IconDB struct {
	nameCN     string
	nameEN     string
	iconTypeCN string
	iconTypeEN string
}

func main() {
	var iconDBs []IconDB
	f, _ := os.Open("db.csv")
	reader := csv.NewReader(f)
	reader.Comma = ','
	records, _ := reader.ReadAll()
	for _, v := range records {
		if v[0] != "Name-CN" {
			iconDBs = append(iconDBs, IconDB{v[0], v[1], v[2], v[3]})
		}
	}

	var pname []string
	for _, v := range iconDBs {
		oriName, iconType, name := v.iconName()
		pname = append(pname, iconType)
		resultDir := "resource/" + iconType + "/"
		err := CreateDir(resultDir)
		if err != nil {
			log.Fatalln(err)
		}
		err = copySvg(oriName, resultDir+name)
		if err != nil {
			log.Fatalln(err)
		}
	}

	pname = Unique(pname)
	for _, p := range pname {
		dir := "iconpark/" + p + "/"
		err := CreateDir(dir)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("fyne bundle -package %s ./resource/%s/ > ./iconpark/%s/bundle.go\n", p, p, p)
		fmt.Printf("generator.exe %s\n", p)
	}
}

func (i *IconDB) iconName() (string, string, string) {
	name := "download/" + i.nameCN + "_" + i.nameEN + ".svg"
	return name, i.iconTypeEN, i.nameEN + ".svg"
}

func CreateDir(absPath string) error {
	return os.MkdirAll(path.Dir(absPath), os.ModePerm)
}

func copySvg(srcFilePath string, dstFilePath string) error {
	srcFile, err := os.ReadFile(srcFilePath)
	if err != nil {
		return err
	}

	distFile, err := os.Create(dstFilePath)
	if err != nil {
		return err
	}
	defer func(distFile *os.File) {
		_ = distFile.Close()
	}(distFile)

	newSvg := strings.Replace(string(srcFile), "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", "", -1)

	_, err = distFile.WriteString(newSvg)
	if err != nil {
		return err
	}
	return nil
}

func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
