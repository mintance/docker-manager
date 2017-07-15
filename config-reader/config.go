package config_reader

import (
	"os"
	"path/filepath"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {

}

func ReadConfigFiles(path string) ([]*Config, error) {

	config_files := []*Config{}

	config_files_path, err := findConfigs(path)

	if err != nil {
		return config_files, err
	}

	for _, config_path := range config_files_path {

		config, err := ReadConfig(config_path)

		if err != nil {

			logrus.Error(err)

			logrus.Info("Skipping config: ", config_path)

		}

		config_files = append(config_files, config)

	}

	return config_files, nil
}

func findConfigs(path string) ([]string, error) {

	config_files := []string{}

	logrus.Info("Reading config directory: ", path)

	dir_pointer, err := os.Open(path)

	if err != nil {

		return config_files, err

	}

	defer dir_pointer.Close()

	files, err := dir_pointer.Readdir(-1)

	if err != nil {

		return config_files, err

	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".yaml" {
				config_files = append(config_files, file.Name())
			}
		}
	}

	logrus.Info("Found ", len(config_files), " configs.")

	return config_files, nil
}

func ReadConfig(config_path string) (*Config, error) {

	config := Config{}

	logrus.Info("Reading config file: " + config_path)

	data, err := ioutil.ReadFile(config_path)

	if err != nil {

		logrus.Error("Config open error: ", err)

		return nil, err

	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {

		logrus.Error("Config read & unmarshal error: ", err)

		return nil, err

	}

	return &config, nil
}
//
//func copyConfigFile(src, dst string) error {
//
//	sfi, err := os.Stat(src)
//
//	if err != nil {
//		return
//	}
//
//	if !sfi.Mode().IsRegular() {
//		// cannot copy non-regular files (e.g., directories,
//		// symlinks, devices, etc.)
//		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
//	}
//	dfi, err := os.Stat(dst)
//	if err != nil {
//		if !os.IsNotExist(err) {
//			return
//		}
//	} else {
//		if !(dfi.Mode().IsRegular()) {
//			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
//		}
//		if os.SameFile(sfi, dfi) {
//			return
//		}
//	}
//	if err = os.Link(src, dst); err == nil {
//		return
//	}
//	err = copyFileContents(src, dst)
//	return
//}
//
//// copyFileContents copies the contents of the file named src to the file named
//// by dst. The file will be created if it does not already exist. If the
//// destination file exists, all it's contents will be replaced by the contents
//// of the source file.
//func copyFileContents(src, dst string) error {
//	in, err := os.Open(src)
//	if err != nil {
//		return
//	}
//	defer in.Close()
//	out, err := os.Create(dst)
//	if err != nil {
//		return
//	}
//	defer func() {
//		cerr := out.Close()
//		if err == nil {
//			err = cerr
//		}
//	}()
//	if _, err = io.Copy(out, in); err != nil {
//		return
//	}
//	err = out.Sync()
//	return
//}