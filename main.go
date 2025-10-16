package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const rootEnvVar = "ITO_ROOT"

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	root, err := defaultRoot()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ito: %v\n", err)
		return 1
	}

	if len(args) == 0 {
		if err := listEntries(root); err != nil {
			fmt.Fprintf(os.Stderr, "ito: %v\n", err)
			return 1
		}
		return 0
	}

	if args[0] == "--help" || args[0] == "-h" {
		printUsage()
		return 0
	}

	switch args[0] {
	case "list":
		if len(args) > 1 {
			fmt.Fprintln(os.Stderr, "ito: list は引数を取りません")
			return 1
		}
		if err := listEntries(root); err != nil {
			fmt.Fprintf(os.Stderr, "ito: %v\n", err)
			return 1
		}
	default:
		if len(args) > 1 {
			fmt.Fprintf(os.Stderr, "ito: 余計な引数があります: %q\n", args[1:])
			return 1
		}
		path, err := resolveEntry(root, args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "ito: %v\n", err)
			return 1
		}
		fmt.Println(path)
	}
	return 0
}

func printUsage() {
	fmt.Println(`ito は ~/.ito 以下のエントリへ素早くアクセスするためのコマンドです。環境変数 ITO_ROOT でルートディレクトリを変更できます。

使い方:
  ito list        ルートディレクトリ内のエントリを列挙します
  ito <entry>     指定したエントリへの絶対パスを出力します`)
}

func defaultRoot() (string, error) {
	if override := os.Getenv(rootEnvVar); override != "" {
		path, err := expandHome(override)
		if err != nil {
			return "", fmt.Errorf("環境変数 %s を解決できません: %w", rootEnvVar, err)
		}
		return filepath.Clean(path), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("ホームディレクトリを取得できません: %w", err)
	}
	return filepath.Join(home, ".ito"), nil
}

func expandHome(path string) (string, error) {
	if path == "~" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return home, nil
	}
	if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "~\\") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[2:]), nil
	}
	return path, nil
}

func listEntries(root string) error {
	dirEntries, err := os.ReadDir(root)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%s が存在しません", root)
	}
	if err != nil {
		return fmt.Errorf("%s を読み取れません: %w", root, err)
	}

	names := make([]string, 0, len(dirEntries))
	for _, entry := range dirEntries {
		name := entry.Name()
		if strings.HasPrefix(name, ".") {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}

func resolveEntry(root, name string) (string, error) {
	path := filepath.Join(root, name)
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return path, nil
		}
		return "", fmt.Errorf("%s を確認できません: %w", path, err)
	}
	return path, nil
}
