package common

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "reflect"
)

type NonIaaS struct {
    ProjectPath         string
    ProgrammingLanguage ProgrammingLanguage
    Databases           []Database
}

type IaaS struct {
    Size string
}

type Blueprint struct {
    NonIaaS NonIaaS
    IaaS    IaaS
}

type blueprintForYaml struct {
    ProjectPath             string `yaml:"project_path,omitempty"`
    ProgrammingLanguage     string `yaml:"programming_language,omitempty"`
    DatabaseDriver          string `yaml:"database_driver,omitempty"`
    DatabaseVersion         string `yaml:"database_version,omitempty"`
    DatabaseStorage         string `yaml:"database_storage,omitempty"`
    DatabaseOriginHost      string `yaml:"database_origin_host,omitempty"`
    DatabaseOriginName      string `yaml:"database_origin_name,omitempty"`
    DatabaseOriginUsername  string `yaml:"database_origin_username,omitempty"`
    DatabaseOriginPassword  string `yaml:"database_origin_password,omitempty"`
    VMGroupSize             string `yaml:"vmgroup_size,omitempty"`
}

// TODO check if it fits into one of the artifacts templates
func (b Blueprint) finalValidation() error {
    return nil
}

func (b Blueprint) ExposeYaml(filePath string) error {
    if err := b.finalValidation(); err != nil {
        return err
    }
    bpfy := blueprintForYaml{
        ProjectPath: b.NonIaaS.ProjectPath,
        ProgrammingLanguage: string(b.NonIaaS.ProgrammingLanguage),
        VMGroupSize: b.IaaS.Size,
    }
    // TODO support more. Assume at most one database for now.
    if len(b.NonIaaS.Databases) > 0 {
        bpfy.DatabaseDriver = b.NonIaaS.Databases[0].Driver
        bpfy.DatabaseVersion = b.NonIaaS.Databases[0].Version
        bpfy.DatabaseStorage = b.NonIaaS.Databases[0].Storage
        bpfy.DatabaseOriginHost = b.NonIaaS.Databases[0].OriginHost
        bpfy.DatabaseOriginName = b.NonIaaS.Databases[0].OriginName
        bpfy.DatabaseOriginUsername = b.NonIaaS.Databases[0].OriginUsername
        bpfy.DatabaseOriginPassword = b.NonIaaS.Databases[0].OriginPassword
    }
    bpBytes, _ := yaml.Marshal(bpfy)
    return ioutil.WriteFile(filePath, bpBytes, 0644)
}

func ImportBlueprintYaml(filePath string) (Blueprint, error) {
    bp := Blueprint{}
    bpfy := blueprintForYaml{}
    bpBytes, err := ioutil.ReadFile(filePath)
    if err != nil {
        return bp, err
    }
    if err = yaml.Unmarshal(bpBytes, &bpfy); err != nil {
        return bp, err
    }

    bp.IaaS = IaaS{
        Size:  bpfy.VMGroupSize,
    }
    bp.NonIaaS = NonIaaS{
        ProjectPath: bpfy.ProjectPath,
    }

    bp.NonIaaS.ProgrammingLanguage, err =
        ParseProgrammingLanguage(bpfy.ProgrammingLanguage)
    if err != nil {
        return bp, err
    }

    // TODO support more. Assume at most one database for now.
    db := Database{
        Driver: bpfy.DatabaseDriver,
        Version: bpfy.DatabaseVersion,
        Storage: bpfy.DatabaseStorage,
        OriginHost: bpfy.DatabaseOriginHost,
        OriginName: bpfy.DatabaseOriginName,
        OriginUsername: bpfy.DatabaseOriginUsername,
        OriginPassword: bpfy.DatabaseOriginPassword,
    }
    allEmpty := true
    s := reflect.ValueOf(&db).Elem()
    for i := 0; i < s.NumField(); i++ {
        valField := s.Field(i)
        val := valField.Interface().(string)
        if (val != "") {
            allEmpty = false
        }
    }
    if !allEmpty {
        bp.NonIaaS.Databases = append(bp.NonIaaS.Databases, db)
    }

    if err = bp.finalValidation(); err != nil {
        return bp, err
    }

    return bp, nil
}
