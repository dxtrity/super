package initialise

import (
  "fmt"
  "os"
  "io/ioutil"
)

const SUPER_INIT_DIRNAME string = ".super"
const SUPER_REPO_FILENAME string = "repo.super"

func InitEmptyRepo(folder string) error {
  err := os.Mkdir(fmt.Sprintf("%s%s",folder,SUPER_INIT_DIRNAME),0750)
  if err != nil {
    return error(err)
  }
  
  file, err := os.Create(fmt.Sprintf("./%s/%s",SUPER_INIT_DIRNAME,SUPER_REPO_FILENAME))
  if err != nil {
    return error(err)
  }
  defer file.Close()

  directory_files, err := ioutil.ReadDir(".")
  if err != nil {
    return error(err)
  }

  for _, f := range directory_files {
    _, err := file.WriteString(fmt.Sprintf("%s\n",string(f.Name())))
    if err != nil {
      return error(err)
    }
  }

  return nil
}
