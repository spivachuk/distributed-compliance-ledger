package keeper

// nolint:goimports
import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryAllProposedX509RootCerts = "all_proposed_x509_root_certs"
	QueryProposedX509RootCert     = "proposed_x509_root_cert"
	QueryAllX509RootCerts         = "all_x509_root_certs"
	QueryAllX509Certs             = "all_x509_certs"
	QueryX509Cert                 = "x509_cert"
	QueryAllSubjectX509Certs      = "all_subject_x509_certs"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryAllProposedX509RootCerts:
			return queryAllProposedX509RootCerts(ctx, req, keeper)
		case QueryAllX509RootCerts:
			return queryAllX509RootCerts(ctx, req, keeper)
		case QueryProposedX509RootCert:
			return queryProposedX509RootCert(ctx, path[1:], keeper)
		case QueryAllX509Certs:
			return queryAllX509Certs(ctx, req, keeper)
		case QueryAllSubjectX509Certs:
			return queryAllSubjectX509Certs(ctx, path[1:], req, keeper)
		case QueryX509Cert:
			return queryX509Cert(ctx, path[1:], keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown pki query endpoint")
		}
	}
}

func queryAllProposedX509RootCerts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var params types.ListCertificatesQueryParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Failed to parse request params: %s", err))
	}

	result := types.NewListProposedCertificates()

	skipped := 0

	keeper.IterateProposedCertificates(ctx, func(certificate types.ProposedCertificate) (stop bool) {
		result.Total++

		if skipped < params.Skip {
			skipped++
			return false
		}

		if len(result.Items) < params.Take || params.Take == 0 {
			result.Items = append(result.Items, certificate)
			return false
		}

		return false
	})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}

func queryProposedX509RootCert(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	subject := path[0]
	subjectKeyID := path[1]

	if !keeper.IsProposedCertificatePresent(ctx, subject, subjectKeyID) {
		return nil, types.ErrProposedCertificateDoesNotExist(subject, subjectKeyID)
	}

	certificate := keeper.GetProposedCertificate(ctx, subject, subjectKeyID)

	res = codec.MustMarshalJSONIndent(keeper.cdc, certificate)

	return res, nil
}

func queryX509Cert(ctx sdk.Context, path []string, keeper Keeper) (res []byte, err sdk.Error) {
	subject := path[0]
	subjectKeyID := path[1]

	if !keeper.IsCertificatePresent(ctx, subject, subjectKeyID) {
		return nil, types.ErrCertificateDoesNotExist(subject, subjectKeyID)
	}

	certificate := keeper.GetCertificates(ctx, subject, subjectKeyID)

	res = codec.MustMarshalJSONIndent(keeper.cdc, certificate)

	return res, nil
}

func queryAllX509RootCerts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	return queryX509Certs(ctx, req, keeper, types.RootCertificate, "")
}

func queryAllX509Certs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	return queryX509Certs(ctx, req, keeper, "", "")
}

func queryAllSubjectX509Certs(ctx sdk.Context, path []string,
	req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	subject := path[0]
	return queryX509Certs(ctx, req, keeper, "", subject)
}

func queryX509Certs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper,
	certificateType types.CertificateType, iteratorPrefix string) (res []byte, err sdk.Error) {
	var params types.ListCertificatesQueryParams
	if err := keeper.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Failed to parse request params: %s", err))
	}

	result := types.NewListCertificates()

	skipped := 0

	keeper.IterateCertificates(ctx, iteratorPrefix,
		func(certificates types.Certificates) (stop bool) {
			for _, certificate := range certificates.Items {
				// filter by certificate type (Root/Any)
				if len(certificateType) != 0 && certificate.Type != certificateType {
					return false
				}

				// filter by root subject
				if len(params.RootSubject) > 0 && params.RootSubject != certificate.RootSubject {
					return false
				}

				// filter by root subject ky id
				if len(params.RootSubjectKeyID) > 0 && params.RootSubjectKeyID != certificate.RootSubjectKeyID {
					return false
				}

				result.Total++

				if skipped < params.Skip {
					skipped++
					return false
				}

				if len(result.Items) < params.Take || params.Take == 0 {
					result.Items = append(result.Items, certificate)
					return false
				}
			}
			return false
		})

	res = codec.MustMarshalJSONIndent(keeper.cdc, result)

	return res, nil
}
