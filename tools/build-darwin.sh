# Builds ticker executable for darwin (OSX) and outputs as build/ticker
echo "[ticker] Building darwin binary..."

# Enter src/ directory after resetting working directory to tools/

cd "${0%/*}"
cd ../src

# Run the build
CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o ../build/ticker

# Delete current config.json in build folder
rm -rf ../build/config.json

# Copy the config to the build directory
cp ./config.json ../build/config.json

# Exit src/ directory into project root
cd ..

echo "[ticker] Done!"