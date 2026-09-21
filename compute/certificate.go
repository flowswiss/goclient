package compute

import (
	"context"

	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/core"
)

type CertificateService struct {
	client *core.Client
}

func NewCertificateService(client *core.Client) *CertificateService {
	return &CertificateService{client: client}
}

func (r CertificateService) List(ctx context.Context, cursor core.Cursor) (list common.List[Certificate], err error) {
	list.Pagination, err = r.client.List(ctx, getCertificatesPath(), cursor, &list.Items)
	return
}

type CertificateGetReq struct {
	ID uint `json:"-"`
}

func (r CertificateService) Get(ctx context.Context, req CertificateGetReq) (certificate Certificate, err error) {
	err = r.client.Get(ctx, getSpecificCertificatePath(req.ID), &certificate)
	return
}

type CertificateCreateReq struct {
	Name        string `json:"name"`
	LocationID  int    `json:"location_id"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

func (r CertificateService) Create(ctx context.Context, req CertificateCreateReq) (certificate Certificate, err error) {
	err = r.client.Create(ctx, getCertificatesPath(), req, &certificate)
	return
}

type CertificateDeleteReq struct {
	ID uint `json:"-"`
}

func (r CertificateService) Delete(ctx context.Context, req CertificateDeleteReq) (err error) {
	err = r.client.Delete(ctx, getSpecificCertificatePath(req.ID))
	return
}

const certificatesSegment = "/v4/compute/certificates"

func getCertificatesPath() string {
	return certificatesSegment
}

func getSpecificCertificatePath(certificateID uint) string {
	return core.Join(certificatesSegment, certificateID)
}
