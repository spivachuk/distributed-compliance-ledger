#!/bin/bash
# Copyright 2020 DSR Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail
source integration_tests/cli/common.sh

random_string user
echo "$user generates keys"
cmd="(echo $passphrase; echo $passphrase) | dcld keys add $user"
result="$(bash -c "$cmd")"

test_divider

echo "Get key info for $user"
result=$(echo $passphrase | dcld keys show $user)
check_response "$result" "\"name\": \"$user\""

test_divider

user_address=$(echo $passphrase | dcld keys show $user -a)
user_pubkey=$(echo $passphrase | dcld keys show $user -p)

jack_address=$(echo $passphrase | dcld keys show jack -a)
alice_address=$(echo $passphrase | dcld keys show alice -a)
bob_address=$(echo $passphrase | dcld keys show bob -a)
anna_address=$(echo $passphrase | dcld keys show anna -a)

vid_in_hex_format=0xA13
pid_in_hex_format=0xA11
vid=2579
pid=2577


echo "Jack proposes account for $user"
result=$(echo $passphrase | dcld tx auth propose-add-account --info="Jack is proposing this account" --address="$user_address" --pubkey="$user_pubkey" --roles="Vendor" --vid="$vid_in_hex_format" --from jack --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get an account for $user is not found"
result=$(dcld query auth account --address=$user_address)
check_response "$result" "Not Found"

echo "Get a proposed account for $user and confirm that the approval contains Jack's address"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "\"address\": \"$user_address\""
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"
response_does_not_contain "$result"  $alice_address "json"

test_divider

echo "Alice approves account for \"$user\""
result=$(echo $passphrase | dcld tx auth approve-add-account --address="$user_address" --info="Alice is approving this account" --from alice --yes)
check_response "$result" "\"code\": 0"

test_divider

echo "Get $user account and confirm that alice and jack are set as approvers"
result=$(dcld query auth account --address=$user_address)
check_response_and_report "$result" $user_address "json"
check_response_and_report "$result"  $jack_address "json"
check_response_and_report "$result"  $alice_address "json"
check_response_and_report "$result"  '"info": "Alice is approving this account"' "json"
check_response_and_report "$result"  '"info": "Jack is proposing this account"' "json"

test_divider

echo "Get $user account"
result=$(dcld query auth account --address=$user_address)
check_response_and_report "$result" "\"address\": \"$user_address\""

test_divider

echo "Get a proposed account for $user is not found"
result=$(dcld query auth proposed-account --address=$user_address)
check_response "$result" "Not Found"

test_divider

productName="Device #1"
echo "$user adds Model with VID: $vid_in_hex_format PID: $pid"
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes)
check_response_and_report "$result" "\"code\": 0"

test_divider

vid_plus_one_in_hex_format=0xA14
vidPlusOne=$((vid_in_hex_format+1))
echo "$user adds Model with a VID: $vid_plus_one_in_hex_format PID: $pid_in_hex_format, This fails with Permission denied as the VID is not associated with this vendor account."
result=$(echo "test1234" | dcld tx model add-model --vid=$vid_plus_one_in_hex_format --pid=$pid_in_hex_format --productName="$productName" --productLabel="Device Description"   --commissioningCustomFlow=0 --deviceTypeID=12 --partNumber=12 --from=$user_address --yes 2>&1) || true
check_response_and_report "$result" "transaction should be signed by a vendor account containing the vendorID $vidPlusOne"

test_divider

echo "$user updates Model with VID: $vid_in_hex_format PID: $pid_in_hex_format"
result=$(echo "test1234" | dcld tx model update-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format --productName="$productName" --productLabel="Device Description" --partNumber=12 --from=$user_address --yes)
check_response_and_report "$result" "\"code\": 0"

test_divider

echo "Get Model with VID: $vid_in_hex_format PID: $pid_in_hex_format"
result=$(dcld query model get-model --vid=$vid_in_hex_format --pid=$pid_in_hex_format)
check_response "$result" "\"vid\": $vid"
check_response "$result" "\"pid\": $pid"
check_response "$result" "\"productName\": \"$productName\""

echo "PASSED"
