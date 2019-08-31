package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"k8s.io/kubectl/pkg/cmd/util/editor"
)

type Spec struct {
	Container string `json:"container"`
	Image     string `json:"Image"`
}

type Resource struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	Spec Spec   `json:"spec"`
}

func main() {
	obj := Resource{
		Kind: "pod",
		Name: "backend",
		Spec: Spec{
			Container: "backend",
			Image:     "docker.io/xxx/backend:latest",
		},
	}
	objInBytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	edit := editor.Editor{Args: []string{"vi"}, Shell: false}
	newContent, _, err := edit.LaunchTempFile("", "", bytes.NewBufferString(string(objInBytes)))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	newObj := Resource{}
	err = json.Unmarshal(newContent, &newObj)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%v", newObj)
}
