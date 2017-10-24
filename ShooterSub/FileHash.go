package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func computeFileHash(filePath string) string {
	fp, err := os.Open(filePath)
	if err != nil {
		logger.Error(err)
	}
	stat, err := fp.Stat()
	if err != nil {
		logger.Error(err)
	}
	size := float64(stat.Size())
	samplePositions := [4]int64{
		4 * 1024,
		int64(math.Floor(size / 3 * 2)),
		int64(math.Floor(size / 3)),
		int64(size - 8*1024)}
	var samples [4][]byte
	for i, position := range samplePositions {
		samples[i] = make([]byte, 4*1024)
		fp.ReadAt(samples[i], position)
	}
	var hash string
	for _, sample := range samples {
		if len(hash) > 0 {
			hash += ";"
		}
		hash += fmt.Sprintf("%x", md5.Sum(sample))
	}

	return hash
}

func getfilehash(fullpath string) string {
	fp, err := os.Open(fullpath)
	if err != nil {
		logger.Error(err)
		return ""
	}
	defer fp.Close()
	stats, statsErr := fp.Stat()
	if statsErr != nil {
		logger.Error(err)
		return ""
	}
	filelen := stats.Size()
	offsetary := [...]int64{4096, (filelen / 3) * 2, filelen / 3, filelen - 8192}
	buf := make([]byte, 4096)
	hashary := make([]string, 0, len(offsetary))
	for _, offset := range offsetary {
		fp.Seek(offset, 0)
		n, _ := io.ReadFull(fp, buf)
		hashary = append(hashary, fmt.Sprintf("%x", md5.Sum(buf[:n])))
	}
	return strings.Join(hashary, ";")
}
