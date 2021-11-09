rm -f bin/*downloader*
echo What is the name of the project?
read project_name
GOOS=windows GOARCH=386 go build -o bin/$project_name-downloader-32.exe main.go
GOOS=windows GOARCH=amd64 go build -o bin/$project_name-downloader-64.exe main.go
go build -o bin/$project_name-downloader-macos main.go
