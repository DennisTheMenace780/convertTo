package utils

func Chunk(str string, size int) []string {
	chunkSize := size
	if len(str) == 0 {
		return nil
	}
	if chunkSize >= len(str) {
		return []string{str}
	}
	var chunks []string = make([]string, 0, (len(str)-1)/chunkSize+1)
	currentLen := 0
	currentStart := 0
	for i := range str {
		if currentLen == chunkSize {
			chunks = append(chunks, str[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, str[currentStart:])
	return chunks
}
