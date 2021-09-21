pkill -9 "__debug_bin|dlv"
while true; do
    dlv debug --headless --log --listen=:2345 --api-version=2 --accept-multiclient --continue &
    inotifywait -e modify -e move -e create -e delete -e attrib --exclude __debug_bin -r .
    pkill -9 "__debug_bin|dlv"
done