package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/vyrekxd/bot/loader/constants"
)

type CmdData struct {
	Path    string
	Data    string
	Run     string
	Package string
}

func GetFilesFromDirectory(dir string) ([]string, error) {
	files := []string{}

	err := filepath.WalkDir(dir, func(path string, data fs.DirEntry, err error) error {
		if !data.IsDir() && constants.FileNameRegex.FindString(data.Name()) != "commands.go" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

func GetCommandData(file []byte) (string, string, string, error) {
	contents := string(file)

	dataVar := constants.DataRegex.FindString(contents)
	if dataVar == "" || len(strings.Split(dataVar, "\x20")) == 0 {
		return "", "", "", errors.New("cant find data variable or cant split it")
	}

	runFunc := constants.RunRegex.FindString(contents)
	if runFunc == "" || len(strings.Split(runFunc, "\x20")) == 0 {
		return "", "", "", errors.New("cant find run function or cant split it")
	}

	packageStr := constants.PackageRegex.FindString(contents)
	if runFunc == "" || len(strings.Split(runFunc, "\x20")) == 0 {
		return "", "", "", errors.New("cant find package name or cant split it")
	}

	data := strings.Split(dataVar, "\x20")[1]
	run := strings.Replace(strings.Split(runFunc, "\x20")[1], "(", "", 1)
	packageName := strings.Split(packageStr, "\x20")[1]

	return data, run, packageName, nil
}

func GetImportByData(path string) (string, error) {
	intoDir := strings.Split(path, "commands")
	if len(intoDir) < 2 {
		return "", errors.New("cant split commands dir")
	}

	fileName := constants.FileNameRegex.FindString(intoDir[1])
	dir := strings.TrimRight(strings.Split(intoDir[1], fileName)[0], "/")

	return constants.DefaultImport + "/commands" + dir, nil
}

func GenerateFile(cmds map[string]([]CmdData)) string {
	lines := []string{
		"package commands\n",
		"import (",
		"	\"github.com/bwmarrin/discordgo\"\n",
	}
	allImports := []string{}

	for k := range cmds {
		allImports = append(allImports, fmt.Sprintf("	\"%v\"", k))
	}

	lines = append(lines, allImports...)
	lines = append(lines,
		")\n",
		"type CommandRun func(*discordgo.Session, *discordgo.InteractionCreate)\n",
		"type Command struct {",
		"	Data *discordgo.ApplicationCommand",
		"	Run  CommandRun",
		"}\n",
		"var Commands = map[string]*Command{",
	)

	for _, k := range cmds {
		for _, cmd := range k {

			//info.PingData.Name: {Data: info.PingData, Run: info.PingRun},
			lines = append(lines, fmt.Sprintf("	%v: {Data: %v, Run: %v},",
				fmt.Sprintf("%v.%v.Name", cmd.Package, cmd.Data),
				fmt.Sprintf("%v.%v", cmd.Package, cmd.Data),
				fmt.Sprintf("%v.%v", cmd.Package, cmd.Run),
			),
			)
		}
	}

	lines = append(lines, "}")

	return strings.Join(lines, "\n")
}
