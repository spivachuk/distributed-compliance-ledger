package rest

// nolint:goimports
import (
	"fmt"
	"net/http"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/rest"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func getAllX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		getListCertificates(restCtx,
			fmt.Sprintf("custom/%s/all_x509_root_certs", storeName), "", "")
	}
}

func getAllProposedX509RootCertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		getListCertificates(restCtx,
			fmt.Sprintf("custom/%s/all_proposed_x509_root_certs", storeName), "", "")
	}
}

func getAllX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyID := r.FormValue(rootSubjectKeyID)
		getListCertificates(restCtx, fmt.Sprintf("custom/%s/all_x509_certs", storeName), rootSubject, rootSubjectKeyID)
	}
}

func getAllSubjectX509CertsHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)
		vars := restCtx.Variables()
		subject := vars[subject]
		rootSubject := r.FormValue(rootSubject)
		rootSubjectKeyID := r.FormValue(rootSubjectKeyID)
		getListCertificates(restCtx, fmt.Sprintf("custom/%s/all_subject_x509_certs/%s",
			storeName, subject), rootSubject, rootSubjectKeyID)
	}
}

func getProposedX509RootCertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		res, height, err := restCtx.QueryStore(types.GetProposedCertificateKey(subject, subjectKeyID), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrProposedCertificateDoesNotExist(subject, subjectKeyID).Error())
			return
		}

		var proposedCertificate types.ProposedCertificate

		restCtx.Codec().MustUnmarshalBinaryBare(res, &proposedCertificate)

		restCtx.EncodeAndRespondWithHeight(proposedCertificate, height)
	}
}

func getX509CertHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		res, height, err := restCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), storeName)
		if err != nil || res == nil {
			restCtx.WriteErrorResponse(http.StatusNotFound,
				types.ErrCertificateDoesNotExist(subject, subjectKeyID).Error())
			return
		}

		var certificate types.Certificates

		cliCtx.Codec.MustUnmarshalBinaryBare(res, &certificate)

		restCtx.EncodeAndRespondWithHeight(certificate, height)
	}
}

func getX509CertChainHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		restCtx := rest.NewRestContext(w, r).WithCodec(cliCtx.Codec)

		vars := restCtx.Variables()
		subject := vars[subject]
		subjectKeyID := vars[subjectKeyID]

		chain := types.NewCertificates([]types.Certificate{})

		height, err := chainCertificates(restCtx, storeName, subject, subjectKeyID, &chain)

		if err != nil {
			restCtx.WriteErrorResponse(http.StatusNotFound, err.Error())
			return
		}

		restCtx.EncodeAndRespondWithHeight(chain, height)
	}
}

func chainCertificates(restCtx rest.RestContext, storeName string,
	subject string, subjectKeyID string, chain *types.Certificates) (int64, sdk.Error) {

	res, height, err := restCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), storeName)
	if err != nil || res == nil {
		return height, types.ErrCertificateDoesNotExist(subject, subjectKeyID)
	}

	var certificates types.Certificates

	restCtx.Codec().MustUnmarshalBinaryBare(res, &certificates)

	certificate := certificates.Items[len(certificates.Items)-1]
	chain.Items = append(chain.Items, certificate)

	if certificate.Type != "root" {
		return chainCertificates(restCtx, storeName, certificate.Issuer, certificate.AuthorityKeyID, chain)
	}

	return height, nil
}

func getListCertificates(restCtx rest.RestContext, path string, rootSubject string, rootSubjectKeyID string) {
	paginationParams, err := restCtx.ParsePaginationParams()
	if err != nil {
		return
	}

	params := types.NewListCertificatesQueryParams(paginationParams, rootSubject, rootSubjectKeyID)
	restCtx.QueryList(path, params)
}