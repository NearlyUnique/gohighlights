package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/NearlyUnique/gohighlights/src6/tfl"
	"github.com/pkg/errors"
)

func main() {
	client := tfl.NewBikeClient(http.DefaultClient)
	var action string
	if len(os.Args) >= 2 {
		action = os.Args[1]
	}
	var err error
	switch action {
	case "list":
		err = list(client)
	case "show":
		err = show(client, os.Args[1:])
	case "showAll":
		err = showAll(client)
	default:
		err = errors.Errorf("Invalid action %s", action)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}

func list(client tfl.Client) error {
	places, err := client.ViewIndex()
	if err != nil {
		return err
	}
	for _, p := range places {
		fmt.Printf("%s %s\n", standID(p.ID), p.CommonName)
	}
	return nil
}

func show(client tfl.Client, args []string) error {
	if len(args) < 1 {
		return errors.New("Missing arguments, no id")
	}
	dock, err := client.ViewDocking(args[0])
	if err != nil {
		return err
	}
	fmt.Println(dock.CommonName)
	for _, p := range dock.AdditionalProperties {
		fmt.Printf("%s : %s\n", p.Key, p.Value)
	}
	return nil
}

func showAll(client tfl.Client) error {
	t0 := time.Now()
	places, err := client.ViewIndex()
	if err != nil {
		return err
	}
	for _, p := range places {
		err := show(client, []string{standID(p.ID)})
		fmt.Printf("Cumlative time :%v\n", time.Since(t0))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Total: %v\n", time.Since(t0))
	return nil
}

func standID(id string) string {
	if strings.HasPrefix(id, "BikePoints_") {
		return id[len("BikePoints_"):]
	}
	return id
}
