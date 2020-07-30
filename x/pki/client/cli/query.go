package cli

//nolint:goimports
import (
	"fmt"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/pagination"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the pki module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	complianceQueryCmd.AddCommand(client.GetCommands(
		GetCmdGetAllProposedX509RootCerts(storeKey, cdc),
		GetCmdProposedX509RootCert(storeKey, cdc),
		GetCmdGetAllX509RootCerts(storeKey, cdc),
		GetCmdX509Cert(storeKey, cdc),
		GetCmdX509CertChain(storeKey, cdc),
		GetCmdGetAllX509Certs(storeKey, cdc),
		GetCmdGetAllSubjectX509Certs(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdGetAllProposedX509RootCerts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-proposed-x509-root-certs",
		Short: "Gets all proposed but not approved root certificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllCertificates(cdc, fmt.Sprintf("custom/%s/all_proposed_x509_root_certs", queryRoute))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdProposedX509RootCert(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposed-x509-root-cert",
		Short: "Gets a proposed but not approved root certificate with the given combination of subject and subject-key-id",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			res, height, err := cliCtx.QueryStore(types.GetProposedCertificateKey(subject, subjectKeyID), queryRoute)
			if err != nil || res == nil {
				return types.ErrProposedCertificateDoesNotExist(subject, subjectKeyID)
			}

			var proposedCertificate types.ProposedCertificate
			cdc.MustUnmarshalBinaryBare(res, &proposedCertificate)

			return cliCtx.EncodeAndPrintWithHeight(proposedCertificate, height)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

func GetCmdGetAllX509RootCerts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-x509-root-certs",
		Short: "Gets all approved root certificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllCertificates(cdc, fmt.Sprintf("custom/%s/all_x509_root_certs", queryRoute))
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdX509Cert(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "x509-cert",
		Short: "Gets a certificates (either root, intermediate or leaf) " +
			"by the given combination of subject and subject-key-id",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			res, height, err := cliCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), queryRoute)
			if err != nil || res == nil {
				return types.ErrCertificateDoesNotExist(subject, subjectKeyID)
			}

			var certificate types.Certificates
			cdc.MustUnmarshalBinaryBare(res, &certificate)

			return cliCtx.EncodeAndPrintWithHeight(certificate, height)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

func GetCmdX509CertChain(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use: "x509-cert-chain",
		Short: "Gets a certificate chain (including root, intermediate and certificate itself) " +
			"by the given combination of subject and subject-key-id",
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			subject := viper.GetString(FlagSubject)
			subjectKeyID := viper.GetString(FlagSubjectKeyID)

			chain := types.NewCertificates([]types.Certificate{})

			height, err := chainCertificates(cliCtx, queryRoute, subject, subjectKeyID, &chain)
			if err != nil {
				return types.ErrCertificateDoesNotExist(subject, subjectKeyID)
			}

			return cliCtx.EncodeAndPrintWithHeight(chain, height)
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagSubjectKeyID, FlagSubjectKeyIDShortcut, "", "Certificate's subject key id (hex)")

	_ = cmd.MarkFlagRequired(FlagSubject)
	_ = cmd.MarkFlagRequired(FlagSubjectKeyID)

	return cmd
}

func GetCmdGetAllX509Certs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-x509-certs",
		Short: "Gets all certificates (root, intermediate and leaf)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getAllCertificates(cdc, fmt.Sprintf("custom/%s/all_x509_certs", queryRoute))
		},
	}

	cmd.Flags().StringP(FlagRootSubject, FlagRootSubjectShortcut, "",
		"filter certificates by `Subject` of root certificate "+
			"(only the certificates started with the given root certificate are returned)")
	cmd.Flags().StringP(FlagRootSubjectKeyID, FlagRootSubjectKeyIDShortcut, "",
		"filter certificates by `Subject Key Id` of root certificate "+
			"(only the certificates started with the given root certificate are returned)")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	return cmd
}

func GetCmdGetAllSubjectX509Certs(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-subject-x509-certs",
		Short: "Gets all certificates (root, intermediate and leaf) associated with subject",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			subject := viper.GetString(FlagSubject)
			return getAllCertificates(cdc, fmt.Sprintf("custom/%s/all_subject_x509_certs/%s", queryRoute, subject))
		},
	}

	cmd.Flags().StringP(FlagSubject, FlagSubjectShortcut, "", "Certificate's subject")
	cmd.Flags().StringP(FlagRootSubject, FlagRootSubjectShortcut, "",
		"filter certificates by `Subject` of root certificate "+
			"(only the certificates started with the given root certificate are returned)")
	cmd.Flags().StringP(FlagRootSubjectKeyID, FlagRootSubjectKeyIDShortcut, "",
		"filter certificates by `Subject Key Id` of root certificate "+
			"(only the certificates started with the given root certificate are returned)")
	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of certificates to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of certificates to take")

	_ = cmd.MarkFlagRequired(FlagSubject)

	return cmd
}

func chainCertificates(cliCtx cli.CliContext, queryRoute string,
	subject string, subjectKeyID string, chain *types.Certificates) (int64, sdk.Error) {

	res, height, err := cliCtx.QueryStore(types.GetApprovedCertificateKey(subject, subjectKeyID), queryRoute)
	if err != nil || res == nil {
		return height, types.ErrCertificateDoesNotExist(subject, subjectKeyID)
	}

	var certificates types.Certificates

	cliCtx.Codec().MustUnmarshalBinaryBare(res, &certificates)

	certificate := certificates.Items[len(certificates.Items)-1]
	chain.Items = append(chain.Items, certificate)

	if certificate.Type != "root" {
		return chainCertificates(cliCtx, queryRoute, certificate.Issuer, certificate.AuthorityKeyID, chain)
	}

	return height, nil
}

func getAllCertificates(cdc *codec.Codec, route string) error {
	cliCtx := cli.NewCLIContext().WithCodec(cdc)

	rootSubject := viper.GetString(FlagRootSubject)
	rootSubjectKeyID := viper.GetString(FlagRootSubjectKeyID)

	paginationParams := pagination.ParsePaginationParamsFromFlags()
	params := types.NewListCertificatesQueryParams(paginationParams, rootSubject, rootSubjectKeyID)

	return cliCtx.QueryList(route, params)
}