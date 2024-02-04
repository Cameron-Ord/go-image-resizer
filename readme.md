
# Simple Script for Batch Processing JPG Images into Various Sizes
This program processes JPG images located in the `/to_process/` folder. The `run.sh` script creates both `processed` and `to_process` directories just in case, but as long as you run build.sh first you'll pretty much be good to go. Ensure that images are present in `/to_process/` for effective processing.

The output for each JPG file is 5 files each with different sizes(original,/2,/3,/4,/5). The image sizes are in the filenames after processing.

If you don't have go, then go get it :I.

## REQUIREMENTS
1. **Linux**
2. **Golang**

## Setup on linux
Either git clone it, or download the zip (code -> Download zip)
```bash
git clone https://github.com/Cameron-Ord/go-image-resizer && cd go-image-resizer
```
```bash
git clone https://github.com/Cameron-Ord/go-image-resizer && cd go-image-resizer/main && chmod +x build.sh
```
```bash
git clone https://github.com/Cameron-Ord/go-image-resizer && cd go-image-resizer/main && chmod +x build.sh && ./build.sh
```

```bash
unzip jpg_resizer.zip
```

## Running the Script

1. **Make scripts executable:**
   ```bash
   chmod +x build.sh
   ```

2. **Execute the scripts:**
   ```bash
   ./build.sh
   ```
   ```bash
   ./run.sh
   ```


