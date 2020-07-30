package keeper

import (
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/pki/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, cdc: cdc}
}

/*
	Approved Root or Intermediate certificate
*/

// Gets the entire Certificate record associated with a Subject/SubjectKeyID combination.
func (k Keeper) GetCertificates(ctx sdk.Context, subject string, subjectKeyID string) types.Certificates {
	if !k.IsCertificatePresent(ctx, subject, subjectKeyID) {
		return types.NewCertificates([]types.Certificate{})
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetApprovedCertificateKey(subject, subjectKeyID))

	var cert types.Certificates

	k.cdc.MustUnmarshalBinaryBare(bz, &cert)

	return cert
}

// Sets the entire Certificate record for a Subject/SubjectKeyID combination.
func (k Keeper) SetCertificate(ctx sdk.Context, certificate types.Certificate) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetApprovedCertificateKey(certificate.Subject, certificate.SubjectKeyID),
		k.cdc.MustMarshalBinaryBare(types.NewCertificates([]types.Certificate{certificate})))
}

// Sets the entire Certificate record for a Subject/SubjectKeyID combination.
func (k Keeper) SetCertificates(ctx sdk.Context, subject string,
	subjectKeyID string, certificates types.Certificates) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetApprovedCertificateKey(subject, subjectKeyID), k.cdc.MustMarshalBinaryBare(certificates))
}

// Check if the Certificate record associated with a Subject/SubjectKeyID combination is present in the store or not.
func (k Keeper) IsCertificatePresent(ctx sdk.Context, subject string, subjectKeyID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetApprovedCertificateKey(subject, subjectKeyID))
}

// Count total Certificates.
func (k Keeper) CountTotalCertificates(ctx sdk.Context) int {
	return k.countTotal(ctx, types.ApprovedCertificatePrefix)
}

func (k Keeper) IterateCertificates(ctx sdk.Context, prefix string,
	process func(info types.Certificates) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, append(types.ApprovedCertificatePrefix, []byte(prefix)...))
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var certificates types.Certificates

		k.cdc.MustUnmarshalBinaryBare(val, &certificates)

		if process(certificates) {
			return
		}

		iter.Next()
	}
}

// Deletes the entire Certificate record associated with a Subject/SubjectKeyID combination.
func (k Keeper) DeleteCertificates(ctx sdk.Context, subject string, subjectKeyID string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetApprovedCertificateKey(subject, subjectKeyID))
}

/*
	Proposed Root certificate
*/

// Gets the entire Proposed Certificate record associated with a Subject/SubjectKeyID combination.
func (k Keeper) GetProposedCertificate(ctx sdk.Context,
	subject string, subjectKeyID string) types.ProposedCertificate {
	if !k.IsProposedCertificatePresent(ctx, subject, subjectKeyID) {
		panic("Proposed Certificate does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetProposedCertificateKey(subject, subjectKeyID))

	var cert types.ProposedCertificate

	k.cdc.MustUnmarshalBinaryBare(bz, &cert)

	return cert
}

// Sets the entire Proposed Certificate record for a Subject/SubjectKeyID combination.
func (k Keeper) SetProposedCertificate(ctx sdk.Context, certificate types.ProposedCertificate) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetProposedCertificateKey(
		certificate.Subject, certificate.SubjectKeyID), k.cdc.MustMarshalBinaryBare(certificate))
}

// Check if the Proposed Certificate record associated with a
// Subject/SubjectKeyID combination is present in the store or not.
func (k Keeper) IsProposedCertificatePresent(ctx sdk.Context, subject string, subjectKeyID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetProposedCertificateKey(subject, subjectKeyID))
}

// Count total Proposed Certificates.
func (k Keeper) CountTotalProposedCertificates(ctx sdk.Context) int {
	return k.countTotal(ctx, types.ProposedCertificatePrefix)
}

// Iterate over all Proposed Certificates.
func (k Keeper) IterateProposedCertificates(ctx sdk.Context,
	process func(info types.ProposedCertificate) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ProposedCertificatePrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var pendingCertificate types.ProposedCertificate

		k.cdc.MustUnmarshalBinaryBare(val, &pendingCertificate)

		if process(pendingCertificate) {
			return
		}

		iter.Next()
	}
}

// Deletes the Proposed Certificate from the store.
func (k Keeper) DeleteProposedCertificate(ctx sdk.Context, subject string, subjectKeyID string) {
	if !k.IsProposedCertificatePresent(ctx, subject, subjectKeyID) {
		panic("Proposed Certificate does not exist")
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetProposedCertificateKey(subject, subjectKeyID))
}

func (k Keeper) countTotal(ctx sdk.Context, prefix []byte) int {
	store := ctx.KVStore(k.storeKey)
	res := 0

	iter := sdk.KVStorePrefixIterator(store, prefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		res++
	}

	return res
}

/*
	A record contains the list of direct child certificates referring to Subject/SubjectKeyID of parent certificate
*/

// Gets the Child Certificates for a record associated with a combination Subject/SubjectKeyID.
func (k Keeper) GetChildCertificates(ctx sdk.Context, subject string, subjectKeyID string) types.ChildCertificates {
	if !k.IsChildCertificatesPresent(ctx, subject, subjectKeyID) {
		return types.NewChildCertificates(subject, subjectKeyID)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetChildCertificatesKey(subject, subjectKeyID))

	var childCertificates types.ChildCertificates

	k.cdc.MustUnmarshalBinaryBare(bz, &childCertificates)

	return childCertificates
}

// Sets the entire Child Certificate record associated with a combination Subject/SubjectKeyID.
func (k Keeper) SetChildCertificatesList(ctx sdk.Context, childCertificates types.ChildCertificates) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetChildCertificatesKey(childCertificates.Subject,
		childCertificates.SubjectKeyID), k.cdc.MustMarshalBinaryBare(childCertificates))
}

// Adds a Child Certificate for a record associated with a combination Subject/SubjectKeyID.
func (k Keeper) AddChildCertificate(ctx sdk.Context, subject string,
	subjectKeyID string, childCertificate types.Certificate) {
	store := ctx.KVStore(k.storeKey)

	childIdentifier := types.NewCertificateIdentifier(
		childCertificate.Subject, childCertificate.SubjectKeyID)

	certificateChain := k.GetChildCertificates(ctx, subject, subjectKeyID)
	certificateChain.ChildCertificates = append(certificateChain.ChildCertificates, childIdentifier)

	store.Set(types.GetChildCertificatesKey(subject, subjectKeyID), k.cdc.MustMarshalBinaryBare(certificateChain))
}

// Checks if the the list of Child Certificates for a combination Subject/SubjectKeyID is present in the store or not.
func (k Keeper) IsChildCertificatesPresent(ctx sdk.Context, subject string, subjectKeyID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetChildCertificatesKey(subject, subjectKeyID))
}

// Iterate over all Child Certificates records.
func (k Keeper) IterateChildCertificatesRecords(ctx sdk.Context,
	process func(info types.ChildCertificates) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ChildCertificatesPrefix)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		val := iter.Value()

		var childCertificatesList types.ChildCertificates

		k.cdc.MustUnmarshalBinaryBare(val, &childCertificatesList)

		if process(childCertificatesList) {
			return
		}

		iter.Next()
	}
}

/*
	Combination of Issuer : Serial Number must be unique
	Helper collection to track uniqueness
*/
// Sets existence flag for combination of Issuer/Serial Number.
func (k Keeper) AddCertificateExistenceFlag(ctx sdk.Context, issuer string, serialNumber string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetCertificateByIssuerSerialNumberKey(issuer, serialNumber), k.cdc.MustMarshalBinaryBare(true))
}

// Check if the certificate for combination of Issuer/Serial Number is present.
func (k Keeper) IsCertificateExists(ctx sdk.Context, issuer string, serialNumber string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetCertificateByIssuerSerialNumberKey(issuer, serialNumber))
}

// Deletes the Existence Flag.
func (k Keeper) DeleteExistenceFlag(ctx sdk.Context, issuer string, serialNumber string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetCertificateByIssuerSerialNumberKey(issuer, serialNumber))
}