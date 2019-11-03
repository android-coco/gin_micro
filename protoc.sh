#proto_path=module/account/proto
#go_out=module/account/proto
#micro_out=module/account/proto
#file_name=module/account/proto/user.proto

echo proto_path
# shellcheck disable=SC2162
read proto_path
echo "${proto_path}"

echo file_name
# shellcheck disable=SC2162
read file_name
echo "${file_name}"
protoc --proto_path="${proto_path}"   --go_out="${proto_path}" --micro_out="${proto_path}" "${proto_path}"/"${file_name}"
