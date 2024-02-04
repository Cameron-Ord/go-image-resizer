if [ ! -d "to_process" ]; then
	mkdir to_process
fi

echo "Removing existing build.."
rm -f main
rm -f go.mod
rm -f go.sum

echo "Init main.."
go mod init main
go get -u github.com/nfnt/resize
go mod tidy

echo "building.."
go build main

echo "chmod run.sh/clean.sh"
chmod +x run.sh
chmod +x clean.sh
