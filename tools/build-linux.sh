# Builds ticker executable for Linux and outputs as build/ticker
echo "[ticker] Building linux binary..."

# Enter src/ directory after resetting working directory to tools/
cd "${0%/*}"
cd ../src

# Run the build
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../build/ticker

# Delete current config.json in build folder
rm -rf ../build/config.json

# Copy the config to the build directory for local testing
cp ./config.json ../build/config.json

# Exit src/ directory into project root
cd ..

echo "[ticker] Done!"
