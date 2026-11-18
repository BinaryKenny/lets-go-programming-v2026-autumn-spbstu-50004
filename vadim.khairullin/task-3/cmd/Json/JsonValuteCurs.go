package Json

import (
	"cmp"
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
)

func EncodeToJson(curs []JsonValute, path string) {
	direction := filepath.Dir(path)

	err := os.MkdirAll(direction, 0755)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	slices.SortFunc(curs, func(i, j JsonValute) int {
		return cmp.Compare(j.Value, i.Value)
	})

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(curs)
	if err != nil {
		panic(err)
	}
}
