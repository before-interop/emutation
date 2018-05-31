package basichttpbinding_emutation

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://interop-fibre.fr/"

// NewEmutation creates an initializes a Emutation.
func NewEmutation(cli *soap.Client) Emutation {
	return &emutation{cli}
}

// Emutation was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type Emutation interface {
	// ConsultationFibres was auto-generated from WSDL.
	ConsultationFibres(parameters *ConsultationFibres) (*ConsultationFibresResponse, error)

	// MiseAJourRouteOptique was auto-generated from WSDL.
	MiseAJourRouteOptique(parameters *MiseAJourRouteOptique) (*MiseAJourRouteOptiqueResponse, error)

	// RecherchePBO was auto-generated from WSDL.
	RecherchePBO(parameters *RecherchePBO) (*RecherchePBOResponse, error)
}

// DateTime in WSDL format.
type DateTime string

// Char was auto-generated from WSDL.
type Char int

// Duration was auto-generated from WSDL.
type Duration Duration

// Guid was auto-generated from WSDL.
type Guid string

// AdressePBOReponseType was auto-generated from WSDL.
type AdressePBOReponseType struct {
	AdressePBO            *string                       `xml:"AdressePBO,omitempty" json:"AdressePBO,omitempty" yaml:"AdressePBO,omitempty"`
	ReferenceGeographique *CoordonneesGeographiquesType `xml:"ReferenceGeographique,omitempty" json:"ReferenceGeographique,omitempty" yaml:"ReferenceGeographique,omitempty"`
	ReferenceHexacle      *string                       `xml:"ReferenceHexacle,omitempty" json:"ReferenceHexacle,omitempty" yaml:"ReferenceHexacle,omitempty"`
	ReferenceHexacleVoie  *ReferenceHexacleVoieType     `xml:"ReferenceHexacleVoie,omitempty" json:"ReferenceHexacleVoie,omitempty" yaml:"ReferenceHexacleVoie,omitempty"`
	ReferenceRivoli       *ReferenceRivoliType          `xml:"ReferenceRivoli,omitempty" json:"ReferenceRivoli,omitempty" yaml:"ReferenceRivoli,omitempty"`
}

// ArrayOfFibreType was auto-generated from WSDL.
type ArrayOfFibreType struct {
	FibreType []*FibreType `xml:"FibreType,omitempty" json:"FibreType,omitempty" yaml:"FibreType,omitempty"`
}

// ArrayOfPboType was auto-generated from WSDL.
type ArrayOfPboType struct {
	PboType []*PboType `xml:"PboType,omitempty" json:"PboType,omitempty" yaml:"PboType,omitempty"`
}

// ArrayOfRouteOptique was auto-generated from WSDL.
type ArrayOfRouteOptique struct {
	RouteOptique []*RouteOptique `xml:"RouteOptique,omitempty" json:"RouteOptique,omitempty" yaml:"RouteOptique,omitempty"`
}

// CodeRetourType was auto-generated from WSDL.
type CodeRetourType struct {
	CodeErreur    *string `xml:"CodeErreur,omitempty" json:"CodeErreur,omitempty" yaml:"CodeErreur,omitempty"`
	CodeRetour    *int    `xml:"CodeRetour,omitempty" json:"CodeRetour,omitempty" yaml:"CodeRetour,omitempty"`
	LibelleErreur *string `xml:"LibelleErreur,omitempty" json:"LibelleErreur,omitempty" yaml:"LibelleErreur,omitempty"`
}

// ConsultationFibres was auto-generated from WSDL.
type ConsultationFibres struct {
	Demande *ConsultationFibresDemande `xml:"demande,omitempty" json:"demande,omitempty" yaml:"demande,omitempty"`
}

// ConsultationFibresDemande was auto-generated from WSDL.
type ConsultationFibresDemande struct {
	Entete                          *EnteteRequeteType `xml:"Entete,omitempty" json:"Entete,omitempty" yaml:"Entete,omitempty"`
	ReferenceCommandePriseInterneOC *string            `xml:"ReferenceCommandePriseInterneOC,omitempty" json:"ReferenceCommandePriseInterneOC,omitempty" yaml:"ReferenceCommandePriseInterneOC,omitempty"`
	ReferencePBO                    *string            `xml:"ReferencePBO,omitempty" json:"ReferencePBO,omitempty" yaml:"ReferencePBO,omitempty"`
	ReferencePM                     *string            `xml:"ReferencePM,omitempty" json:"ReferencePM,omitempty" yaml:"ReferencePM,omitempty"`
	ReferencePrestationPrise        *string            `xml:"ReferencePrestationPrise,omitempty" json:"ReferencePrestationPrise,omitempty" yaml:"ReferencePrestationPrise,omitempty"`
}

// ConsultationFibresPBOReponse was auto-generated from WSDL.
type ConsultationFibresPBOReponse struct {
	CodeRetour *CodeRetourType    `xml:"CodeRetour,omitempty" json:"CodeRetour,omitempty" yaml:"CodeRetour,omitempty"`
	Entete     *EnteteReponseType `xml:"Entete,omitempty" json:"Entete,omitempty" yaml:"Entete,omitempty"`
	Fibre      *ListeFibreType    `xml:"Fibre,omitempty" json:"Fibre,omitempty" yaml:"Fibre,omitempty"`
	PboType    *PboType           `xml:"PboType,omitempty" json:"PboType,omitempty" yaml:"PboType,omitempty"`
}

// ConsultationFibresResponse was auto-generated from WSDL.
type ConsultationFibresResponse struct {
	ConsultationFibresResult *ConsultationFibresPBOReponse `xml:"ConsultationFibresResult,omitempty" json:"ConsultationFibresResult,omitempty" yaml:"ConsultationFibresResult,omitempty"`
}

// CoordonneesGeographiquesType was auto-generated from WSDL.
type CoordonneesGeographiquesType struct {
	CoordonneesX   *string `xml:"CoordonneesX,omitempty" json:"CoordonneesX,omitempty" yaml:"CoordonneesX,omitempty"`
	CoordonneesY   *string `xml:"CoordonneesY,omitempty" json:"CoordonneesY,omitempty" yaml:"CoordonneesY,omitempty"`
	TypeProjection *string `xml:"TypeProjection,omitempty" json:"TypeProjection,omitempty" yaml:"TypeProjection,omitempty"`
}

// EnteteReponseType was auto-generated from WSDL.
type EnteteReponseType struct {
	HorodatageReponse   *DateTime                `xml:"HorodatageReponse,omitempty" json:"HorodatageReponse,omitempty" yaml:"HorodatageReponse,omitempty"`
	HorodatageRequete   *DateTime                `xml:"HorodatageRequete,omitempty" json:"HorodatageRequete,omitempty" yaml:"HorodatageRequete,omitempty"`
	IdentidiantReponse  *int                     `xml:"IdentidiantReponse,omitempty" json:"IdentidiantReponse,omitempty" yaml:"IdentidiantReponse,omitempty"`
	OperateurCommercial *OperateurCommercialType `xml:"OperateurCommercial,omitempty" json:"OperateurCommercial,omitempty" yaml:"OperateurCommercial,omitempty"`
	VersionWS           *string                  `xml:"VersionWS,omitempty" json:"VersionWS,omitempty" yaml:"VersionWS,omitempty"`
}

// EnteteRequeteType was auto-generated from WSDL.
type EnteteRequeteType struct {
	HorodatageRequete   *DateTime                `xml:"HorodatageRequete,omitempty" json:"HorodatageRequete,omitempty" yaml:"HorodatageRequete,omitempty"`
	OperateurCommercial *OperateurCommercialType `xml:"OperateurCommercial,omitempty" json:"OperateurCommercial,omitempty" yaml:"OperateurCommercial,omitempty"`
	VersionWS           *string                  `xml:"VersionWS,omitempty" json:"VersionWS,omitempty" yaml:"VersionWS,omitempty"`
}

// EtatFibreType was auto-generated from WSDL.
type EtatFibreType struct {
	EtatFibre *string `xml:"EtatFibre,omitempty" json:"EtatFibre,omitempty" yaml:"EtatFibre,omitempty"`
}

// FibreType was auto-generated from WSDL.
type FibreType struct {
	EtatFibre           *EtatFibreType `xml:"EtatFibre,omitempty" json:"EtatFibre,omitempty" yaml:"EtatFibre,omitempty"`
	IdentifiantFibre    *string        `xml:"IdentifiantFibre,omitempty" json:"IdentifiantFibre,omitempty" yaml:"IdentifiantFibre,omitempty"`
	InformationFibrePBO *string        `xml:"InformationFibrePBO,omitempty" json:"InformationFibrePBO,omitempty" yaml:"InformationFibrePBO,omitempty"`
	InformationTubePBO  *string        `xml:"InformationTubePBO,omitempty" json:"InformationTubePBO,omitempty" yaml:"InformationTubePBO,omitempty"`
	ReferenceCablePBO   *string        `xml:"ReferenceCablePBO,omitempty" json:"ReferenceCablePBO,omitempty" yaml:"ReferenceCablePBO,omitempty"`
	ReferencePrise      *string        `xml:"ReferencePrise,omitempty" json:"ReferencePrise,omitempty" yaml:"ReferencePrise,omitempty"`
}

// ListeFibreType was auto-generated from WSDL.
type ListeFibreType struct {
	Fibres *ArrayOfFibreType `xml:"Fibres,omitempty" json:"Fibres,omitempty" yaml:"Fibres,omitempty"`
}

// ListePboType was auto-generated from WSDL.
type ListePboType struct {
	Pbo *ArrayOfPboType `xml:"Pbo,omitempty" json:"Pbo,omitempty" yaml:"Pbo,omitempty"`
}

// ListeRoutesOptiques was auto-generated from WSDL.
type ListeRoutesOptiques struct {
	RoutesOptiques *ArrayOfRouteOptique `xml:"RoutesOptiques,omitempty" json:"RoutesOptiques,omitempty" yaml:"RoutesOptiques,omitempty"`
}

// MiseAJourRouteOptique was auto-generated from WSDL.
type MiseAJourRouteOptique struct {
	Demande *MiseAJourRouteOptiqueDemande `xml:"demande,omitempty" json:"demande,omitempty" yaml:"demande,omitempty"`
}

// MiseAJourRouteOptiqueDemande was auto-generated from WSDL.
type MiseAJourRouteOptiqueDemande struct {
	Batiment                        *string                      `xml:"Batiment,omitempty" json:"Batiment,omitempty" yaml:"Batiment,omitempty"`
	BatimentTerrain                 *string                      `xml:"BatimentTerrain,omitempty" json:"BatimentTerrain,omitempty" yaml:"BatimentTerrain,omitempty"`
	CodeRetour                      *CodeRetourType              `xml:"CodeRetour,omitempty" json:"CodeRetour,omitempty" yaml:"CodeRetour,omitempty"`
	Entete                          *EnteteRequeteType           `xml:"Entete,omitempty" json:"Entete,omitempty" yaml:"Entete,omitempty"`
	Escalier                        *string                      `xml:"Escalier,omitempty" json:"Escalier,omitempty" yaml:"Escalier,omitempty"`
	EscalierTerrain                 *string                      `xml:"EscalierTerrain,omitempty" json:"EscalierTerrain,omitempty" yaml:"EscalierTerrain,omitempty"`
	Etage                           *string                      `xml:"Etage,omitempty" json:"Etage,omitempty" yaml:"Etage,omitempty"`
	EtageTerrain                    *string                      `xml:"EtageTerrain,omitempty" json:"EtageTerrain,omitempty" yaml:"EtageTerrain,omitempty"`
	IdentifiantFibre                *string                      `xml:"IdentifiantFibre,omitempty" json:"IdentifiantFibre,omitempty" yaml:"IdentifiantFibre,omitempty"`
	MotifMutation                   *string                      `xml:"MotifMutation,omitempty" json:"MotifMutation,omitempty" yaml:"MotifMutation,omitempty"`
	PorteTerrain                    *string                      `xml:"PorteTerrain,omitempty" json:"PorteTerrain,omitempty" yaml:"PorteTerrain,omitempty"`
	ReferenceAdresse                *ReferenceAdresseDemandeType `xml:"ReferenceAdresse,omitempty" json:"ReferenceAdresse,omitempty" yaml:"ReferenceAdresse,omitempty"`
	ReferenceCommandePriseInterneOC *string                      `xml:"ReferenceCommandePriseInterneOC,omitempty" json:"ReferenceCommandePriseInterneOC,omitempty" yaml:"ReferenceCommandePriseInterneOC,omitempty"`
	ReferencePrestationPrise        *string                      `xml:"ReferencePrestationPrise,omitempty" json:"ReferencePrestationPrise,omitempty" yaml:"ReferencePrestationPrise,omitempty"`
	ReferencePrise                  *string                      `xml:"ReferencePrise,omitempty" json:"ReferencePrise,omitempty" yaml:"ReferencePrise,omitempty"`
}

// MiseAJourRouteOptiqueReponse was auto-generated from WSDL.
type MiseAJourRouteOptiqueReponse struct {
	CodeRetour     *CodeRetourType      `xml:"CodeRetour,omitempty" json:"CodeRetour,omitempty" yaml:"CodeRetour,omitempty"`
	Entete         *EnteteReponseType   `xml:"Entete,omitempty" json:"Entete,omitempty" yaml:"Entete,omitempty"`
	NumeroDecharge *string              `xml:"NumeroDecharge,omitempty" json:"NumeroDecharge,omitempty" yaml:"NumeroDecharge,omitempty"`
	ReferencePBO   *string              `xml:"ReferencePBO,omitempty" json:"ReferencePBO,omitempty" yaml:"ReferencePBO,omitempty"`
	ReferencePM    *string              `xml:"ReferencePM,omitempty" json:"ReferencePM,omitempty" yaml:"ReferencePM,omitempty"`
	ReferencePMT   *string              `xml:"ReferencePMT,omitempty" json:"ReferencePMT,omitempty" yaml:"ReferencePMT,omitempty"`
	ReferencePrise *string              `xml:"ReferencePrise,omitempty" json:"ReferencePrise,omitempty" yaml:"ReferencePrise,omitempty"`
	RoutesOptiques *ListeRoutesOptiques `xml:"RoutesOptiques,omitempty" json:"RoutesOptiques,omitempty" yaml:"RoutesOptiques,omitempty"`
}

// MiseAJourRouteOptiqueResponse was auto-generated from WSDL.
type MiseAJourRouteOptiqueResponse struct {
	MiseAJourRouteOptiqueResult *MiseAJourRouteOptiqueReponse `xml:"MiseAJourRouteOptiqueResult,omitempty" json:"MiseAJourRouteOptiqueResult,omitempty" yaml:"MiseAJourRouteOptiqueResult,omitempty"`
}

// OperateurCommercialType was auto-generated from WSDL.
type OperateurCommercialType struct {
	Identifiant *string `xml:"Identifiant,omitempty" json:"Identifiant,omitempty" yaml:"Identifiant,omitempty"`
	Nom         *string `xml:"Nom,omitempty" json:"Nom,omitempty" yaml:"Nom,omitempty"`
}

// PboType was auto-generated from WSDL.
type PboType struct {
	AdressePBO              *AdressePBOReponseType     `xml:"AdressePBO,omitempty" json:"AdressePBO,omitempty" yaml:"AdressePBO,omitempty"`
	NaturePBO               *string                    `xml:"NaturePBO,omitempty" json:"NaturePBO,omitempty" yaml:"NaturePBO,omitempty"`
	NombreFibresDisponibles *int                       `xml:"NombreFibresDisponibles,omitempty" json:"NombreFibresDisponibles,omitempty" yaml:"NombreFibresDisponibles,omitempty"`
	ReferencePBO            *string                    `xml:"ReferencePBO,omitempty" json:"ReferencePBO,omitempty" yaml:"ReferencePBO,omitempty"`
	ReferencePM             *string                    `xml:"ReferencePM,omitempty" json:"ReferencePM,omitempty" yaml:"ReferencePM,omitempty"`
	StructurePBO            *StructureVerticalePBOType `xml:"StructurePBO,omitempty" json:"StructurePBO,omitempty" yaml:"StructurePBO,omitempty"`
}

// PositionPmType was auto-generated from WSDL.
type PositionPmType struct {
	InfoFibreModulePM      *string `xml:"InfoFibreModulePM,omitempty" json:"InfoFibreModulePM,omitempty" yaml:"InfoFibreModulePM,omitempty"`
	InfoTubeModulePM       *string `xml:"InfoTubeModulePM,omitempty" json:"InfoTubeModulePM,omitempty" yaml:"InfoTubeModulePM,omitempty"`
	NomModulePM            *string `xml:"NomModulePM,omitempty" json:"NomModulePM,omitempty" yaml:"NomModulePM,omitempty"`
	PositionModulePM       *string `xml:"PositionModulePM,omitempty" json:"PositionModulePM,omitempty" yaml:"PositionModulePM,omitempty"`
	ReferenceCableModulePM *string `xml:"ReferenceCableModulePM,omitempty" json:"ReferenceCableModulePM,omitempty" yaml:"ReferenceCableModulePM,omitempty"`
}

// RecherchePBO was auto-generated from WSDL.
type RecherchePBO struct {
	Demande *RecherchePBODemande `xml:"demande,omitempty" json:"demande,omitempty" yaml:"demande,omitempty"`
}

// RecherchePBODemande was auto-generated from WSDL.
type RecherchePBODemande struct {
	Entete                          *EnteteRequeteType `xml:"Entete,omitempty" json:"Entete,omitempty" yaml:"Entete,omitempty"`
	ReferenceCommandePriseInterneOC *string            `xml:"ReferenceCommandePriseInterneOC,omitempty" json:"ReferenceCommandePriseInterneOC,omitempty" yaml:"ReferenceCommandePriseInterneOC,omitempty"`
	ReferencePrestationPrise        *string            `xml:"ReferencePrestationPrise,omitempty" json:"ReferencePrestationPrise,omitempty" yaml:"ReferencePrestationPrise,omitempty"`
}

// RecherchePBOReponse was auto-generated from WSDL.
type RecherchePBOReponse struct {
	CodeRetour *CodeRetourType    `xml:"CodeRetour,omitempty" json:"CodeRetour,omitempty" yaml:"CodeRetour,omitempty"`
	Entete     *EnteteReponseType `xml:"Entete,omitempty" json:"Entete,omitempty" yaml:"Entete,omitempty"`
	ListePBO   *ListePboType      `xml:"ListePBO,omitempty" json:"ListePBO,omitempty" yaml:"ListePBO,omitempty"`
}

// RecherchePBOResponse was auto-generated from WSDL.
type RecherchePBOResponse struct {
	RecherchePBOResult *RecherchePBOReponse `xml:"RecherchePBOResult,omitempty" json:"RecherchePBOResult,omitempty" yaml:"RecherchePBOResult,omitempty"`
}

// ReferenceAdresseDemandeType was auto-generated from WSDL.
type ReferenceAdresseDemandeType struct {
	IdentifiantImmeuble   *string                       `xml:"IdentifiantImmeuble,omitempty" json:"IdentifiantImmeuble,omitempty" yaml:"IdentifiantImmeuble,omitempty"`
	ReferenceGeographique *CoordonneesGeographiquesType `xml:"ReferenceGeographique,omitempty" json:"ReferenceGeographique,omitempty" yaml:"ReferenceGeographique,omitempty"`
	ReferenceHexacle      *string                       `xml:"ReferenceHexacle,omitempty" json:"ReferenceHexacle,omitempty" yaml:"ReferenceHexacle,omitempty"`
	ReferenceHexacleVoie  *ReferenceHexacleVoieType     `xml:"ReferenceHexacleVoie,omitempty" json:"ReferenceHexacleVoie,omitempty" yaml:"ReferenceHexacleVoie,omitempty"`
	ReferenceRivoli       *ReferenceRivoliType          `xml:"ReferenceRivoli,omitempty" json:"ReferenceRivoli,omitempty" yaml:"ReferenceRivoli,omitempty"`
}

// ReferenceHexacleVoieType was auto-generated from WSDL.
type ReferenceHexacleVoieType struct {
	CodeHexacleVoie      *string `xml:"CodeHexacleVoie,omitempty" json:"CodeHexacleVoie,omitempty" yaml:"CodeHexacleVoie,omitempty"`
	ComplementNumeroVoie *Char   `xml:"ComplementNumeroVoie,omitempty" json:"ComplementNumeroVoie,omitempty" yaml:"ComplementNumeroVoie,omitempty"`
	NumeroVoie           *uint   `xml:"NumeroVoie,omitempty" json:"NumeroVoie,omitempty" yaml:"NumeroVoie,omitempty"`
}

// ReferenceRivoliType was auto-generated from WSDL.
type ReferenceRivoliType struct {
	CodeInsee            *string `xml:"CodeInsee,omitempty" json:"CodeInsee,omitempty" yaml:"CodeInsee,omitempty"`
	CodeRivoli           *string `xml:"CodeRivoli,omitempty" json:"CodeRivoli,omitempty" yaml:"CodeRivoli,omitempty"`
	ComplementNumeroVoie *Char   `xml:"ComplementNumeroVoie,omitempty" json:"ComplementNumeroVoie,omitempty" yaml:"ComplementNumeroVoie,omitempty"`
	NumeroVoie           *uint   `xml:"NumeroVoie,omitempty" json:"NumeroVoie,omitempty" yaml:"NumeroVoie,omitempty"`
}

// RouteOptique was auto-generated from WSDL.
type RouteOptique struct {
	ConnecteurPriseCouleur *string         `xml:"ConnecteurPriseCouleur,omitempty" json:"ConnecteurPriseCouleur,omitempty" yaml:"ConnecteurPriseCouleur,omitempty"`
	ConnecteurPriseNumero  *int            `xml:"ConnecteurPriseNumero,omitempty" json:"ConnecteurPriseNumero,omitempty" yaml:"ConnecteurPriseNumero,omitempty"`
	InformationFibrePBO    *string         `xml:"InformationFibrePBO,omitempty" json:"InformationFibrePBO,omitempty" yaml:"InformationFibrePBO,omitempty"`
	InformationTubePBO     *string         `xml:"InformationTubePBO,omitempty" json:"InformationTubePBO,omitempty" yaml:"InformationTubePBO,omitempty"`
	OC                     *string         `xml:"OC,omitempty" json:"OC,omitempty" yaml:"OC,omitempty"`
	PositionPm             *PositionPmType `xml:"PositionPm,omitempty" json:"PositionPm,omitempty" yaml:"PositionPm,omitempty"`
	ReferenceCablePBO      *string         `xml:"ReferenceCablePBO,omitempty" json:"ReferenceCablePBO,omitempty" yaml:"ReferenceCablePBO,omitempty"`
}

// StructureVerticalePBOType was auto-generated from WSDL.
type StructureVerticalePBOType struct {
	Batiment *string `xml:"Batiment,omitempty" json:"Batiment,omitempty" yaml:"Batiment,omitempty"`
	Escalier *string `xml:"Escalier,omitempty" json:"Escalier,omitempty" yaml:"Escalier,omitempty"`
	Etage    *string `xml:"Etage,omitempty" json:"Etage,omitempty" yaml:"Etage,omitempty"`
}

// Operation wrapper for ConsultationFibres.
// OperationEmutation_ConsultationFibres_InputMessage was auto-generated
// from WSDL.
type OperationEmutation_ConsultationFibres_InputMessage struct {
	Parameters *ConsultationFibres `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for ConsultationFibres.
// OperationEmutation_ConsultationFibres_OutputMessage was auto-generated
// from WSDL.
type OperationEmutation_ConsultationFibres_OutputMessage struct {
	Parameters *ConsultationFibresResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for MiseAJourRouteOptique.
// OperationEmutation_MiseAJourRouteOptique_InputMessage was auto-generated
// from WSDL.
type OperationEmutation_MiseAJourRouteOptique_InputMessage struct {
	Parameters *MiseAJourRouteOptique `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for MiseAJourRouteOptique.
// OperationEmutation_MiseAJourRouteOptique_OutputMessage was auto-generated
// from WSDL.
type OperationEmutation_MiseAJourRouteOptique_OutputMessage struct {
	Parameters *MiseAJourRouteOptiqueResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for RecherchePBO.
// OperationEmutation_RecherchePBO_InputMessage was auto-generated
// from WSDL.
type OperationEmutation_RecherchePBO_InputMessage struct {
	Parameters *RecherchePBO `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation wrapper for RecherchePBO.
// OperationEmutation_RecherchePBO_OutputMessage was auto-generated
// from WSDL.
type OperationEmutation_RecherchePBO_OutputMessage struct {
	Parameters *RecherchePBOResponse `xml:"parameters,omitempty" json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// emutation implements the Emutation interface.
type emutation struct {
	cli *soap.Client
}

// ConsultationFibres was auto-generated from WSDL.
func (p *emutation) ConsultationFibres(parameters *ConsultationFibres) (*ConsultationFibresResponse, error) {
	α := struct {
		OperationEmutation_ConsultationFibres_InputMessage `xml:"tns:ConsultationFibres"`
	}{
		OperationEmutation_ConsultationFibres_InputMessage{
			parameters,
		},
	}

	γ := struct {
		OperationEmutation_ConsultationFibres_OutputMessage `xml:"ConsultationFibresResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://interop-fibre.fr/Emutation/ConsultationFibres", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// MiseAJourRouteOptique was auto-generated from WSDL.
func (p *emutation) MiseAJourRouteOptique(parameters *MiseAJourRouteOptique) (*MiseAJourRouteOptiqueResponse, error) {
	α := struct {
		OperationEmutation_MiseAJourRouteOptique_InputMessage `xml:"tns:MiseAJourRouteOptique"`
	}{
		OperationEmutation_MiseAJourRouteOptique_InputMessage{
			parameters,
		},
	}

	γ := struct {
		OperationEmutation_MiseAJourRouteOptique_OutputMessage `xml:"MiseAJourRouteOptiqueResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://interop-fibre.fr/Emutation/MiseAJourRouteOptique", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}

// RecherchePBO was auto-generated from WSDL.
func (p *emutation) RecherchePBO(parameters *RecherchePBO) (*RecherchePBOResponse, error) {
	α := struct {
		OperationEmutation_RecherchePBO_InputMessage `xml:"tns:RecherchePBO"`
	}{
		OperationEmutation_RecherchePBO_InputMessage{
			parameters,
		},
	}

	γ := struct {
		OperationEmutation_RecherchePBO_OutputMessage `xml:"RecherchePBOResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://interop-fibre.fr/Emutation/RecherchePBO", α, &γ); err != nil {
		return nil, err
	}
	return γ.Parameters, nil
}
