#scp -P 22022  -r bin keys config  root@47.56.172.167:/usr/local/server/api

PROJECT_PATH="/usr/local/server/api/bin"
PROJECT_NAME="qb_web_server"
USER_NAME="root"
HOSTS=("你的服务器地址")
PASSWORD="服务器密码"

echo "Please Input the server password: "
#read -s PASSWORD

echo '------------------build------------------'
make web-server
cp ./bin/qb_web_server ./bin/gin_micro_new

echo '-----------------upload-----------------'
# shellcheck disable=SC2068
for host in ${HOSTS[@]}
do
echo "${host}"
./upload.expect ./bin/${PROJECT_NAME}_new ${USER_NAME} "${host}" ${PASSWORD} ${PROJECT_PATH}
# shellcheck disable=SC2181
if [[ "$?" != 0 ]]; then
   exit 2
fi
done

echo '------------------restart-------------------'
# shellcheck disable=SC2068
for host in ${HOSTS[@]}
do
echo "${host}"
./restart.expect ${PROJECT_NAME} ${USER_NAME} "${host}" ${PASSWORD} ${PROJECT_PATH}
done

rm -rf ./bin/gin_micro_new