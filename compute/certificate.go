package compute

import (
	"context"

	"github.com/flowswiss/goclient"
	"github.com/flowswiss/goclient/common"
)

type Certificate struct {
	ID       int                `json:"id"`
	Name     string             `json:"name"`
	Location common.Location    `json:"location"`
	Type     string             `json:"type"`
	Details  CertificateDetails `json:"certificate"`
}

type CertificateDetails struct {
	Subject   map[string]string `json:"subject"`
	Issuer    map[string]string `json:"issuer"`
	ValidFrom common.Time       `json:"valid_from"`
	ValidTo   common.Time       `json:"valid_to"`
	Serial    string            `json:"serial"`
}

type CertificateList struct {
	Items      []Certificate
	Pagination goclient.Pagination
}

type CertificateCreate struct {
	Name        string `json:"name"`
	LocationID  int    `json:"location_id"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"private_key"`
}

type CertificateService struct {
	client goclient.Client
}

func NewCertificateService(client goclient.Client) CertificateService {
	return CertificateService{client: client}
}

func (r CertificateService) List(ctx context.Context, cursor goclient.Cursor) (list CertificateList, err error) {
	list.Pagination, err = r.client.List(ctx, getCertificatesPath(), cursor, &list.Items)
	return
}

func (r CertificateService) Get(ctx context.Context, id int) (certificate Certificate, err error) {
	err = r.client.Get(ctx, getSpecificCertificatePath(id), &certificate)
	return
}

func (r CertificateService) Create(ctx context.Context, body CertificateCreate) (certificate Certificate, err error) {
	err = r.client.Create(ctx, getCertificatesPath(), body, &certificate)
	return
}

func (r CertificateService) Delete(ctx context.Context, id int) (err error) {
	err = r.client.Delete(ctx, getSpecificCertificatePath(id))
	return
}

const certificatesSegment = "/v4/compute/certificates"

func getCertificatesPath() string {
	return certificatesSegment
}

func getSpecificCertificatePath(certificateID int) string {
	return goclient.Join(certificatesSegment, certificateID)
}
