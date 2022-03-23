set -euo pipefail

docker cp deployment/scripts/switch_to_cosmovisor test_deploy_vn:/var/lib/dcl/
docker cp build/dcld test_deploy_vn:/var/lib/dcl/
docker cp ~/go/bin/cosmovisor test_deploy_vn:/var/lib/dcl/
docker cp deployment/cosmovisor.service test_deploy_vn:/var/lib/dcl/

docker exec -u dcl -i test_deploy_vn ./switch_to_cosmovisor
