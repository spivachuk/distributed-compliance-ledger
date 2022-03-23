set -euo pipefail

for node in test_deploy_gvn test_deploy_vn; do
	docker exec -u dcl -i "$node" bash -c "source ~/.profile; dcld config output json"
	docker exec -u dcl -i "$node" bash -c "source ~/.profile; dcld config broadcast-mode block"
	docker exec -u dcl -i "$node" sudo apt update
	docker exec -u dcl -i "$node" sudo apt install jq -y
done

docker exec -u dcl -i test_deploy_gvn mkdir integration_tests
docker cp ./integration_tests/cli test_deploy_gvn:/var/lib/dcl/integration_tests/cli
docker cp ./integration_tests/constants test_deploy_gvn:/var/lib/dcl/integration_tests/constants

docker exec -u dcl -i test_deploy_gvn bash -c "source ~/.profile; integration_tests/cli/model-demo.sh"
docker exec -u dcl -i test_deploy_gvn bash -c "source ~/.profile; integration_tests/cli/modelversion-demo.sh"
docker exec -u dcl -i test_deploy_gvn bash -c "source ~/.profile; integration_tests/cli/auth-demo.sh"
docker exec -u dcl -i test_deploy_gvn bash -c "source ~/.profile; integration_tests/cli/compliance-demo.sh"
docker exec -u dcl -i test_deploy_gvn bash -c "source ~/.profile; integration_tests/cli/pki-demo.sh"
docker exec -u dcl -i test_deploy_gvn bash -c "source ~/.profile; integration_tests/cli/vendorinfo-demo.sh"
