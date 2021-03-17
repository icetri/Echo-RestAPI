package service

import (
	"bytes"
	"echo-restapi/pkg"
	"echo-restapi/types"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
)

func UpdateFileJSON(data []types.User, filedir string) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", " ")
	err := enc.Encode(data)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Ошибка в Encode файла UpdateFileJSON"))
		return err
	}
	file, err := os.OpenFile(filedir, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Ошибка в открытие или создания os.OpenFile"))
		return err
	}
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Ошибка в очистки файла UpdateFileJSON"))
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Ошибка в очистки файла UpdateFileJSON"))
	}
	defer file.Close()
	_, err = file.Write(buf.Bytes())
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Ошибка в записи UpdateFileJSON"))
		return err
	}
	return err
}

func (s *Service) CheckJsonFile(filedir string) (int, []types.User, error) {
	file, err := os.OpenFile(filedir, os.O_RDONLY, 0755)
	if err != nil {
		return 1, s.users.Users, nil
	}
	defer file.Close()
	b, _ := ioutil.ReadFile("json.json")
	err = json.Unmarshal(b, &s.users.Users)
	if err != nil {
		return 0, nil, errors.Wrap(err, "Ошибка декодирования сервиса CheckJsonFile")
	}
	for i, val := range s.users.Users {
		s.max = s.users.Users[i].Id
		if s.max < val.Id {
			s.max = val.Id
		}
	}
	s.max += 1
	return s.max, s.users.Users, err
}
