package catalog

import (
	"github.com/openservicemesh/osm/pkg/constants"
	"github.com/openservicemesh/osm/pkg/service"
)

// GetServiceForServiceAccount returns a service corresponding to a service account
func (mc *MeshCatalog) GetServiceForServiceAccount(sa service.K8sServiceAccount) (service.MeshService, error) {
	for _, provider := range mc.endpointsProviders {
		// TODO (#88) : remove this provider check once we have figured out the service account story for azure vms
		if provider.GetID() != constants.AzureProviderName {
			log.Trace().Msgf("[%s] Looking for Services for Name=%s", provider.GetID(), sa)
			service, err := provider.GetServiceForServiceAccount(sa)
			log.Trace().Msgf("[%s] Found service %s for Name=%s", provider.GetID(), service.String(), sa)
			return service, err
		}
	}
	return service.MeshService{}, errServiceNotFoundForAnyProvider
}
