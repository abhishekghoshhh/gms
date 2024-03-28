package model

import "math"

type Batch struct {
	Indices []int
	Count   int
	Size    int
}

func NewBatch(totalResults, batchSize, maxBatchCount int) *Batch {
	batchCount := int(math.Ceil(float64(totalResults) / float64(batchSize)))
	if batchCount > maxBatchCount {
		batchSize = int(math.Ceil(float64(totalResults) / float64(maxBatchCount)))
		batchCount = maxBatchCount
	}
	var startingIndices []int
	for i := 0; i < batchCount; i++ {
		startingIndices = append(startingIndices, i*batchSize+1)
	}
	return &Batch{
		Indices: startingIndices,
		Count:   batchCount,
		Size:    batchSize,
	}
}
