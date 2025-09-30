#!/bin/sh

echo "No Green Tea----------------------------------------------"
for _ in $(seq 0 4); do
  go test -gcflags="-N -l" -bench . -benchmem | grep "ns/op"
done

echo ""
echo "Green Tea 🍵----------------------------------------------"
for _ in $(seq 0 4); do
  GOEXPERIMENT=greenteagc go test -gcflags="-N -l" -bench . -benchmem | grep "ns/op"
done
