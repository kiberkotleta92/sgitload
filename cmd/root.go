/*
Copyright © 2020 Kirill Denisov <kirill.denisov700@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "sgitload [link to github file] > [destination]",
	Short: "Downloads single file from github and prints to stdout",
	Long: `Downloads single file from github and prints to stdout

 by kirilldenisov`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Pass me URL")
		}
		return LoadGitHub(args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func LoadGitHub(url string) error {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(errors.New("Wrong URL"))
		}
	}()
	newurl := ConstructURL(url)
	data, err := Load(newurl)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ConstructURL(inp string) string {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(errors.New("Wrong URL"))
		}
	}()
	var i int
	if !strings.Contains(inp, "http") {
		i = -2
	}
	s := strings.Split(inp, "/")
	res := []string{"https://cdn.jsdelivr.net/gh", s[i+3], s[i+4], s[i+7]}
	out := strings.Join(res, "/")
	return out
}

func Load(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
