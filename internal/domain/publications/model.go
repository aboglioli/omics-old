package publications

import "omics/pkg/models"

type Publication struct {
	models.Base
	Name     string
	Synopsis string
}
