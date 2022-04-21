git pull
tag=$(git describe --abbrev=0)

if [[ tag =~ "fatal" ]]
then
    echo tag="build"
fi
docker build -t search-trace-server:"$tag" .
docker push danbai225/search-trace-server:"$tag"