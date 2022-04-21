git pull
tag=$(git describe --abbrev=0)
if [[ $tag == "" ]]
then
    tag="build"
fi
docker build -t -no-cache search-trace-server:"$tag" .
#docker push danbai225/search-trace-server:"$tag"