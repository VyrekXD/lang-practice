package main

import (
	"log"
	"os"

	"github.com/vyrekxd/bot/loader/utils"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	log.Println("Starting generator of loaded commands... 🚀")
	log.Println("Reading directory... 📂")

	pwd, err := os.Getwd()
	if err != nil {
		log.Panicf("Error when getting working directory: %v", err)
	}

	paths, err := utils.GetFilesFromDirectory(pwd + "/../bot/commands")
	if err != nil {
		log.Panicf("Error when reading directory: %v", err)
	}

	log.Printf("Got %v files, validating if they are commands... 🔍", len(paths))

	cmds := map[string]([]utils.CmdData){}

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			log.Panicf("Error when reading file %v: %v", path, err)
		}

		cmdData, cmdRun, cmdPackage, err := utils.GetCommandData(data)
		if err != nil {
			log.Panicf("Error when getting command data: %v", err)
		}

		importPath, err := utils.GetImportByData(path)
		if err != nil {
			log.Panicf("Error when generating import: %v", err)
		}

		exist := cmds[importPath]
		if len(exist) == 0 {
			cmds[importPath] = []utils.CmdData{
				{
					Path:    path,
					Data:    cmdData,
					Run:     cmdRun,
					Package: cmdPackage,
				},
			}
		} else {
			cmds[importPath] = append(cmds[importPath], utils.CmdData{
				Path:    path,
				Data:    cmdData,
				Run:     cmdRun,
				Package: cmdPackage,
			})
		}

	}

	log.Printf("Validated all commands, got %v imports now generating loaded file... 📄", len(cmds))

	genContent := utils.GenerateFile(cmds)

	log.Printf("Loaded file generated, %v lines, creating it in the directory... 📨", len(genContent))

	genFile, err := os.Create(pwd + "/../bot/commands/commands.go")
	if err != nil {
		log.Panicf("Error when creating loaded file: %v", err)
	}

	_, err = genFile.WriteString(genContent)
	if err != nil {
		log.Panicf("Error when writing loaded file: %v", err)
	}

	genFile.Close()

	log.Printf("Loader finished! ✅")
}
