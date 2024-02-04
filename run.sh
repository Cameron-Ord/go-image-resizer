echo "Running bash script.."

if [ ! -d "to_process" ]; then
	mkdir to_process
	echo "to_process dir created"
fi

if [ ! -d "processed" ]; then
	mkdir processed
	echo "processed dir created"
fi

echo "Executing jpg resizing tool.."
./main
