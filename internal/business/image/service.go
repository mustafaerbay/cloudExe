package image

type Service interface {
	UploadImage(i *Image) error
	GetImage(id string) (*Image, error)
	UpdateImage(id string, i *Image) error
}

type imageService struct {
	store *Image
}

func (s * imageService) GetImage(id string) (*Image, error) {
	
	return s.store(),error
}