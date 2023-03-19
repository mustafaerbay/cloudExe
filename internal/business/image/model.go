package image

import "mime/multipart"

type Image struct {
	Image *multipart.FileHeader
}