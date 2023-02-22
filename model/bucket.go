package model

type Bucket struct {
	capacity float32
	contents float32
}

func (b *Bucket) PourIn(addedVolume float32) {
	volumePresent := b.contents + addedVolume
	b.contents = b.contrainedToCapacity(volumePresent)
}

func (b *Bucket) contrainedToCapacity(volumePlacedIn float32) float32 {
	if volumePlacedIn > b.capacity {
		return b.capacity
	}

	return volumePlacedIn
}

// func (b *Bucket) PourIn(addedVolume float32) {
// 	if b.contents+addedVolume <= b.capacity {
// 		b.contents = b.capacity
// 		return
// 	}

// 	b.contents += addedVolume
// }
