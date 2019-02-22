package handlers

import (
	"google.golang.org/grpc/metadata"
	"strings"
)

type MetadataReaderWriter struct {
	md metadata.MD
}

// NewMetadataReaderWriter creates an object that implements the opentracing.TextMapReader and opentracing.TextMapWriter interfaces
func NewMetadataReaderWriter(md metadata.MD) *MetadataReaderWriter {
	return &MetadataReaderWriter{md: md}
}

func (mrw *MetadataReaderWriter) ForeachKey(handler func(string, string) error) error {
	for key, values := range mrw.md {
		for _, value := range values {
			if err := handler(key, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func (mrw *MetadataReaderWriter) Set(key, value string) {
	// headers should be lowercase
	k := strings.ToLower(key)
	mrw.md[k] = append(mrw.md[k], value)
}
