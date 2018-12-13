mkdir release
gox -arch="amd64" -os="darwin linux windows" -ldflags="-w -s" --output="release/{{.Dir}}_{{.OS}}_{{.Arch}}"