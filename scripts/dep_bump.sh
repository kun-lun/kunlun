#!/bin/bash
projects_array=("artifacts" "built-in-roles" "common" "deployment-producer" "executor" \
"migration-producer" "test-infra" "verification-producer" \
"ashandler" "digester" "infra-producer" "report-producer" \
"tfhandler" "kunlun")
for i in "${projects_array[@]}"
do
  pushd $GOPATH/src/github.com/kun-lun/$i
    $GOPATH/bin/dep ensure
  popd
done
