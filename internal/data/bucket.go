package repository
import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	
)
type BucketRepository struct {
	bucket *s3.Client
}

func (r *BucketRepository) GetImage(id string) (*Image.Image, error)