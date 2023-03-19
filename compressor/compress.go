package compressor

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
)

const base64Str = `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAD5AAAAYCCAYAAABHwJQMAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAACxMAAAsTAQCanBgAAA0UaVRYdFhNTDpjb20uYWRvYmUueG1wAAAAAAA8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIiB4OnhtcHRrPSJYTVAgQ29yZSA1LjQuMCI+CiAgIDxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+CiAgICAgIDxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSIiCiAgICAgICAgICAgIHhtbG5zOmV4aWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20vZXhpZi8xLjAvIgogICAgICAgICAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29tL3RpZmYvMS4wLyIKICAgICAgICAgICAgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIgogICAgICAgICAgICB4bWxuczpleGlmRVg9Imh0dHA6Ly9jaXBhLmpwL2V4aWYvMS4wLyIKICAgICAgICAgICAgeG1sbnM6cGhvdG9zaG9wPSJodHRwOi8vbnMuYWRvYmUuY29tL3Bob3Rvc2hvcC8xLjAvIj4KICAgICAgICAgPGV4aWY6U2h1dHRlclNwZWVkVmFsdWU+MTA4MTUvOTI5PC9leGlmOlNodXR0ZXJTcGVlZFZhbHVlPgogICAgICAgICA8ZXhpZjpDb2xvclNwYWNlPjE8L2V4aWY6Q29sb3JTcGFjZT4KICAgICAgICAgPGV4aWY6Rm9jYWxMZW5ndGg8SFDmj548kZbgWVQNypjVbocU/nDmTfo8+Hj9Dtxqc+HZCAGvUPbwInDMQBP3umVowVaK7Vu45W7rnb4gSqjD0A/0G2vfgCK7pV2JvYO0MYBn+2Tfm9Q0hjJADOSnZIEKpQe2vigS9nnuNDQ+hM7wLMZys0PeDxVERr9FS/lNotojLHZZO7IGAPeUNlOZCAqz8jyDSsdqU88rQ0aX2xhPzljXvwe8ScIJ0GgcvOZgLyyr2w0znYmgHd8dB768hoeHsoF4CDjKB0F2E7dOqrc2p3f0FBWm/1FJ3UdOfbwUrq16ADoECGBznbK1p10jAfq2fcC5HX0pq1Z99wHrcdC8uZYB5n+5cPGmS/wAXUuPwoKjwGOj5r9ATrbpAp1Wp07JjPQ34Bto1/NXfzN9/+MbCDnd39utbv3h5GJ6+Hhv+ld/8avIs/DAy7Ppp1c/dny6KFJkjaR/MeY3N2ytMotpL7zr3/zlr6dvXwTOn/Js5FkUtenk6F3gvzTtPYhOeSV/5Ly0YHEgeSjf/vLr6dnzR+lviMzmSilOPNjK5t`

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func CompressText() {

	var buf bytes.Buffer

	gzpWriter := gzip.NewWriter(&buf)

	_, err := gzpWriter.Write([]byte(base64Str))
	panicError(err)

	err = gzpWriter.Close()
	panicError(err)

	fmt.Printf("Original string length: %d\n", len(base64Str))
	fmt.Printf("Compressed data length: %d\n", buf.Len())

	println(buf.String())

	err = os.WriteFile("assets/test.txt", buf.Bytes(), 0644)
	panicError(err)
}
