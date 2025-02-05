package resources

import (
	"fmt"

	"github.com/patrickcping/pingone-go-sdk-v2/risk"
	"github.com/pingidentity/pingcli/internal/connector"
	"github.com/pingidentity/pingcli/internal/connector/common"
	"github.com/pingidentity/pingcli/internal/logger"
)

// Verify that the resource satisfies the exportable resource interface
var (
	_ connector.ExportableResource = &PingOneRiskPredictorResource{}
)

type PingOneRiskPredictorResource struct {
	clientInfo *connector.PingOneClientInfo
}

// Utility method for creating a PingOneRiskPredictorResource
func RiskPredictor(clientInfo *connector.PingOneClientInfo) *PingOneRiskPredictorResource {
	return &PingOneRiskPredictorResource{
		clientInfo: clientInfo,
	}
}

func (r *PingOneRiskPredictorResource) ResourceType() string {
	return "pingone_risk_predictor"
}

func (r *PingOneRiskPredictorResource) ExportAll() (*[]connector.ImportBlock, error) {
	l := logger.Get()
	l.Debug().Msgf("Exporting all '%s' Resources...", r.ResourceType())

	importBlocks := []connector.ImportBlock{}

	riskPredictorData, err := r.getRiskPredictorData()
	if err != nil {
		return nil, err
	}

	for riskPredictorId, riskPredictorInfo := range *riskPredictorData {
		riskPredictorName := riskPredictorInfo[0]
		riskPredictorType := riskPredictorInfo[1]

		commentData := map[string]string{
			"Export Environment ID": r.clientInfo.ExportEnvironmentID,
			"Resource Type":         r.ResourceType(),
			"Risk Predictor ID":     riskPredictorId,
			"Risk Predictor Name":   riskPredictorName,
			"Risk Predictor Type":   riskPredictorType,
		}

		importBlock := connector.ImportBlock{
			ResourceType:       r.ResourceType(),
			ResourceName:       fmt.Sprintf("%s_%s", riskPredictorType, riskPredictorName),
			ResourceID:         fmt.Sprintf("%s/%s", r.clientInfo.ExportEnvironmentID, riskPredictorId),
			CommentInformation: common.GenerateCommentInformation(commentData),
		}

		importBlocks = append(importBlocks, importBlock)
	}

	return &importBlocks, nil
}

func (r *PingOneRiskPredictorResource) getRiskPredictorData() (*map[string][]string, error) {
	riskPredictorData := make(map[string][]string)

	iter := r.clientInfo.ApiClient.RiskAPIClient.RiskAdvancedPredictorsApi.ReadAllRiskPredictors(r.clientInfo.Context, r.clientInfo.ExportEnvironmentID).Execute()

	for cursor, err := range iter {
		ok, err := common.HandleClientResponse(cursor.HTTPResponse, err, "ReadAllRiskPredictors", r.ResourceType())
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, nil
		}

		if cursor.EntityArray == nil {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		embedded, embeddedOk := cursor.EntityArray.GetEmbeddedOk()
		if !embeddedOk {
			return nil, common.DataNilError(r.ResourceType(), cursor.HTTPResponse)
		}

		for _, riskPredictor := range embedded.GetRiskPredictors() {
			var (
				riskPredictorId     *string
				riskPredictorIdOk   bool
				riskPredictorName   *string
				riskPredictorNameOk bool
				riskPredictorType   *risk.EnumPredictorType
				riskPredictorTypeOk bool
			)

			switch {
			case riskPredictor.RiskPredictorAdversaryInTheMiddle != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorAdversaryInTheMiddle.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorAdversaryInTheMiddle.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorAdversaryInTheMiddle.GetTypeOk()
			case riskPredictor.RiskPredictorAnonymousNetwork != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorAnonymousNetwork.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorAnonymousNetwork.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorAnonymousNetwork.GetTypeOk()
			case riskPredictor.RiskPredictorBotDetection != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorBotDetection.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorBotDetection.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorBotDetection.GetTypeOk()
			case riskPredictor.RiskPredictorCommon != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorCommon.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorCommon.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorCommon.GetTypeOk()
			case riskPredictor.RiskPredictorComposite != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorComposite.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorComposite.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorComposite.GetTypeOk()
			case riskPredictor.RiskPredictorCustom != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorCustom.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorCustom.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorCustom.GetTypeOk()
			case riskPredictor.RiskPredictorDevice != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorDevice.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorDevice.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorDevice.GetTypeOk()
			case riskPredictor.RiskPredictorEmailReputation != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorEmailReputation.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorEmailReputation.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorEmailReputation.GetTypeOk()
			case riskPredictor.RiskPredictorGeovelocity != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorGeovelocity.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorGeovelocity.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorGeovelocity.GetTypeOk()
			case riskPredictor.RiskPredictorIPReputation != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorIPReputation.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorIPReputation.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorIPReputation.GetTypeOk()
			case riskPredictor.RiskPredictorUserLocationAnomaly != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorUserLocationAnomaly.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorUserLocationAnomaly.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorUserLocationAnomaly.GetTypeOk()
			case riskPredictor.RiskPredictorUserRiskBehavior != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorUserRiskBehavior.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorUserRiskBehavior.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorUserRiskBehavior.GetTypeOk()
			case riskPredictor.RiskPredictorVelocity != nil:
				riskPredictorId, riskPredictorIdOk = riskPredictor.RiskPredictorVelocity.GetIdOk()
				riskPredictorName, riskPredictorNameOk = riskPredictor.RiskPredictorVelocity.GetNameOk()
				riskPredictorType, riskPredictorTypeOk = riskPredictor.RiskPredictorVelocity.GetTypeOk()
			default:
				continue
			}

			if riskPredictorIdOk && riskPredictorNameOk && riskPredictorTypeOk {
				riskPredictorData[*riskPredictorId] = []string{*riskPredictorName, string(*riskPredictorType)}
			}
		}
	}

	return &riskPredictorData, nil
}
