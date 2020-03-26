package FileUtils

import (
  "os"
)

func FileExist(path string) (exist bool, err error) {
  _, err = os.Lstat(path)

  if err == nil {
    return true, nil
  }

  if os.IsNotExist(err) {
    return false, nil
  }

  return exist, err
}
