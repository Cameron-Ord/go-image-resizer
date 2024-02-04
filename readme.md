
# Simple Script for Batch Processing JPG Images into Various Sizes

## REQUIREMENTS
1. **Linux**
2. **Golang**

## Setup on linux
Either git clone it, or download the zip (code -> Download zip)
```bash
git clone https://github.com/Cameron-Ord/go-image-resizer ~/Documents/
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

The script processes JPG images located in the `/to_process/` folder. The `run.sh` script creates both `processed` and `to_process` directories. Ensure that images are present in `/to_process/` for effective processing.
