package main

import (
	"bufio"
	"fmt"
	"github.com/Masterminds/semver"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 1 {
		panic(fmt.Errorf("argument invalid"))
	}
	if os.Args[1] == "" {
		var versions []*semver.Version
		buf := bufio.NewScanner(os.Stdin)
		for buf.Scan() {
			v, err := semver.NewVersion(buf.Text())
			if err == nil {
				versions = append(versions, v)
			}
		}
		sort.Sort(semver.Collection(versions))
		for _, v := range versions {
			fmt.Println(v.String())
		}
	}
	if os.Args[1] == "latest" {
		var versions []*semver.Version
		buf := bufio.NewScanner(os.Stdin)
		for buf.Scan() {
			v, err := semver.NewVersion(buf.Text())
			if err == nil {
				versions = append(versions, v)
			}
		}
		sort.Sort(semver.Collection(versions))
		fmt.Print(versions[len(versions)-1].String())
	} else {
		target, err := semver.NewConstraint(os.Args[1])
		if err != nil {
			panic(err.Error())
		}
		var versions []*semver.Version
		buf := bufio.NewScanner(os.Stdin)
		for buf.Scan() {
			v, err := semver.NewVersion(buf.Text())
			if err == nil && target.Check(v) {
				versions = append(versions, v)
			}
		}
		sort.Sort(semver.Collection(versions))
		for _, v := range versions {
			fmt.Println(v.String())
		}
	}

	os.Exit(0)
}
